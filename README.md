# sf6go

## 项目介绍

这是一个对[seetaface6](https://github.com/SeetaFace6Open/index)项目的一个go语言的封装，方便go语言调用这个优秀的人脸识别库

> 目前仅在ubuntu20.04系统上编译了seetaface6，windows平台的dll还要等一等

## 使用说明

### 准备
> 环境要求
> 
> ubuntu 20.04
> 
> `go >= 1.17`
> 
> | CPU指令集 | 性能 |
> |-----------|------|
> | AVX2+FMA  | 高   |
> | AVX2      | 中   |
> | SSE2      | 低   |
> 
> CPU最低要求需要支持SSE2指令集，否则将无法运行

1. 下载[模型和动态链接库](https://github.com/jyk1987/sf6data)
2. 配置动态链接库的环境变量，将lib下面相应系统动态连接库文件夹路径配置到"LD_LIBRARY_PATH"中
```
export LD_LIBRARY_PATH=$LD_LIBRARY_PATH:路径
```
3. `sudo apt update && sudo apt install cmake g++ gcc -y`
4. `go get github.com/jyk1987/sf6go`

### 使用

#### 加载模型

> 在使用前需要先加载模型目录，模型目录只需要加载一次

```
sf6go.InitModelPath("模型存放的目录")
```

#### 准备数据

seetaface6使用的是BGR色彩空间数据，我们通常情况下用到的数据都是RGB色彩空间，所以图像（图片）数据是需要先进行转换才能够传递给seetaface6进行使用的
我再sf6go中已经实现了几种常用的转换方法

```
    // 通过图片文件创建SeetaImageData
    sf6go.NewSeetaImageDataFromFile("image path")
    
    // 通过Image创建SeetaImageData
    sf6go.NewSeetaImageDataFromImage(img)
    
    // 通过base64字符串创建SeetaImageData
    sf6go.NewSeetaImageDataFromBase64("base64 string")
    
    /** 
    如果使用opencv(govc)获取图像数据，不能直接转换成SeetaImageData
    需要先创建一个SeetaImageData然后再通过设置数据的方式来使用
    这是为了让sf6go不依赖gocv方便服务器端部署使用
    **/
    // 创建SeetaImageData, channels一般情况下都是3（彩色数据）
    imageData := sf6go.NewSeetaImageData(width,height,channels)
    imageData.SetUint8(mat.ToBytes()) // mat 为 gocv.Mat
```

#### 人脸检测

```
    // 人脸检测器
	fd := sf6go.NewFaceDetector()
	defer fd.Close()
	postions := fd.Detect(imageData)// postions为图像中所检测到的人脸的位置和大小信息
	for i := 0; i < len(postions); i++ {
	   postion := postions[i].Postion
	   log.Println(postion)
	}
```
> func (s *FaceDetector) Detect(img *SeetaImageData) []*SeetaFaceInfo
> 
> 返回数据SeetaFaceInfo结构：

```
    type SeetaFaceInfo struct {
    	Postion *SeetaRect // 人脸位置信息
    	Score   float32    // 结果置信度
    }
```

#### 人脸特征点定位

> 使用5点信息模型,5点模型定位的特征是后续操作的必要数据
> 
> 用于后续的人脸对其裁剪+特征提取操作等多种操作
> 
> 不能使用其他模型进行定位
> 
> 可选用的模型:
> 
> 1. 5点模型sf6go.ModelType_light，用于后续的前置必要数据
> 2. 5点遮挡信息模型sf6go.ModelType_mask，包含5个特征点（左眉心、右眉心、鼻尖、左嘴角、右嘴角）的遮挡信息，进用于遮挡判断，不能用于后续
```
    fl := sf6go.NewFaceLandmarker(sf6go.ModelType_light)
    defer fl.Close()
    // 定位5点特征
    pointInfo := fl.Mark(imageData, postion)
    log.Println(pointInfo)
```

> func (s *FaceLandmarker) Mark(img *SeetaImageData, postion *SeetaRect) *SeetaPointInfo
> 
> 返回数据SeetaPointInfo结构：

```
    // SeetaPointInfo 人脸特征点信息
    type SeetaPointInfo struct {
    	PointCount int                     // 特征点数
    	Points     []C.struct_SeetaPointF  // 特征点坐标
    	Masks      []int                   // 特征点这档信息，需要使用sf6go.ModelType_mask模型
    }
```
#### 特征提取

> 特征提取使用的特征定位信息必须是sf6go.ModelType_light模型提取的，参看"人脸特征点定位"章节
> 特征提取可用模型:
> 1. 68点高精度特征模型 sf6go.ModelType_default，用于需要高精度特征的场景
> 
> 2. 5点轻量特征模型 sf6go.ModelType_light，用于大部分的通用场景
> 
> 3. 5点口罩特征模型 sf6go.ModelType_mask，用于对佩戴口罩人脸进行特征提取
> 
> **注意:** 通常情况5点特征就够用了，68点高精度特征并不一定就会比5点特征能提升巨大的准确度（seetaface6官方文档中是这样说的），同时68点特征进行定位和特征对比时时间都会比较长。如果需要在实时场景（基于视频的人脸识别系统）中使用68点特征，那将对开发者和硬件都提出很高的要求。
```
    // 5点轻量特征提取器
    fr := sf6go.NewFaceRecognizer(sf6go.ModelType_light)
    defer fr.Close()
```
1. 人脸裁剪对齐+特征提取，一气呵成
```
    // 组合方法特征提取
    success, features := fr.Extract(imageData, pointInfo)
```
2. 人脸裁剪对齐+特征提取，拆分操作
```
    // 裁剪并对齐人脸
    face := fr.CropFaceV2(imageData, pointInfo)
    // 通过裁剪的人脸获取特征
    success, features := fr.ExtractCroppedFace(face)
```
> 特征提取方法返回的数据为：bool, []float32
> 
> 第一个参数代表提取是否成功，第二个参数为特征数据
> 
> **特征长度:** 5点轻量模型和5点口罩模型提取的特征长度512，68点高精度模型提取的特征是1024长度

#### 特征相似度计算

1. 使用seetaface6远程的c++代码进行相似度计算
> func (s *FaceRecognizer) CalculateSimilarity(features1, features2 []float32) float32 
>
> 不推荐这样进行特征相似度对比，因为在大的底库中进行对比时，需要反复通过cgo中转调用调用C，然后再转到C++中，这个中转过程相对还是比较耗时的。大一些的底库会浪费很长时间。

2. 直接使用go代码进行特征相似度计算，**推荐**
> 我在sf6go中没有加入这个代码,下面是对比代码
```
    // CompareFeatures 对比特征相似度
    func CompareFeatures(features1 []float32, features2 []float32) float32 {
    	var sum float32 = 0.0
    	fcount := len(features1)
    	for i := 0; i < fcount; i++ {
    		sum += features1[i] * features2[i]
    	}
    	return sum
    }
```

> **相似度阈值:**
>
> | 模型                    | 特征长度 | 一般阈值 |
> |-------------------------|----------|---------|
> | sf6go.ModelType_default | 1024     | 0.62    |
> | sf6go.ModelType_light   | 512      | 0.55    |
> | sf6go.ModelType_mask    | 512      | 0.48    |

### 事例

详细查看 [./cmd/main.go](https://github.com/jyk1987/sf6go/blob/master/cmd/main.go)