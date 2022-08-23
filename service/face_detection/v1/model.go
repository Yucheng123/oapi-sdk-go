// Package face_detection code generated by oapi sdk gen
/*
 * MIT License
 *
 * Copyright (c) 2022 Lark Technologies Pte. Ltd.
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice, shall be included in all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

package larkface_detection

import (
	"github.com/larksuite/oapi-sdk-go/v3/core"
)

type AttributeItem struct {
	Type        *int     `json:"type,omitempty"`        // 属性
	Probability *float64 `json:"probability,omitempty"` // 识别置信度，[0, 1]，代表判断正确的概率
}

type AttributeItemBuilder struct {
	type_           int // 属性
	typeFlag        bool
	probability     float64 // 识别置信度，[0, 1]，代表判断正确的概率
	probabilityFlag bool
}

func NewAttributeItemBuilder() *AttributeItemBuilder {
	builder := &AttributeItemBuilder{}
	return builder
}

// 属性
// 示例值：0
func (builder *AttributeItemBuilder) Type(type_ int) *AttributeItemBuilder {
	builder.type_ = type_
	builder.typeFlag = true
	return builder
}

// 识别置信度，[0, 1]，代表判断正确的概率
// 示例值：0.95566
func (builder *AttributeItemBuilder) Probability(probability float64) *AttributeItemBuilder {
	builder.probability = probability
	builder.probabilityFlag = true
	return builder
}

func (builder *AttributeItemBuilder) Build() *AttributeItem {
	req := &AttributeItem{}
	if builder.typeFlag {
		req.Type = &builder.type_

	}
	if builder.probabilityFlag {
		req.Probability = &builder.probability

	}
	return req
}

type FaceAttribute struct {
	Gender  *AttributeItem `json:"gender,omitempty"`  // 性别信息：0 男性，1 女性
	Age     *int           `json:"age,omitempty"`     // 年龄大小
	Emotion *AttributeItem `json:"emotion,omitempty"` // 情绪：0 自然， 1 高兴，2 惊讶，3 害怕，4 悲伤，5 生气， 6 厌恶
	Beauty  *int           `json:"beauty,omitempty"`  // 颜值打分：[0, 100]
	Pose    *FacePose      `json:"pose,omitempty"`    // 人脸姿态
	Hat     *AttributeItem `json:"hat,omitempty"`     // 帽子：0 未戴帽子，1 戴帽子
	Glass   *AttributeItem `json:"glass,omitempty"`   // 眼镜：0 未戴眼镜，1 戴眼镜
	Mask    *AttributeItem `json:"mask,omitempty"`    // 口罩：0 未戴口罩，1 戴口罩
}

type FaceAttributeBuilder struct {
	gender      *AttributeItem // 性别信息：0 男性，1 女性
	genderFlag  bool
	age         int // 年龄大小
	ageFlag     bool
	emotion     *AttributeItem // 情绪：0 自然， 1 高兴，2 惊讶，3 害怕，4 悲伤，5 生气， 6 厌恶
	emotionFlag bool
	beauty      int // 颜值打分：[0, 100]
	beautyFlag  bool
	pose        *FacePose // 人脸姿态
	poseFlag    bool
	hat         *AttributeItem // 帽子：0 未戴帽子，1 戴帽子
	hatFlag     bool
	glass       *AttributeItem // 眼镜：0 未戴眼镜，1 戴眼镜
	glassFlag   bool
	mask        *AttributeItem // 口罩：0 未戴口罩，1 戴口罩
	maskFlag    bool
}

func NewFaceAttributeBuilder() *FaceAttributeBuilder {
	builder := &FaceAttributeBuilder{}
	return builder
}

// 性别信息：0 男性，1 女性
// 示例值：
func (builder *FaceAttributeBuilder) Gender(gender *AttributeItem) *FaceAttributeBuilder {
	builder.gender = gender
	builder.genderFlag = true
	return builder
}

// 年龄大小
// 示例值：19
func (builder *FaceAttributeBuilder) Age(age int) *FaceAttributeBuilder {
	builder.age = age
	builder.ageFlag = true
	return builder
}

// 情绪：0 自然， 1 高兴，2 惊讶，3 害怕，4 悲伤，5 生气， 6 厌恶
// 示例值：
func (builder *FaceAttributeBuilder) Emotion(emotion *AttributeItem) *FaceAttributeBuilder {
	builder.emotion = emotion
	builder.emotionFlag = true
	return builder
}

// 颜值打分：[0, 100]
// 示例值：88
func (builder *FaceAttributeBuilder) Beauty(beauty int) *FaceAttributeBuilder {
	builder.beauty = beauty
	builder.beautyFlag = true
	return builder
}

// 人脸姿态
// 示例值：
func (builder *FaceAttributeBuilder) Pose(pose *FacePose) *FaceAttributeBuilder {
	builder.pose = pose
	builder.poseFlag = true
	return builder
}

// 帽子：0 未戴帽子，1 戴帽子
// 示例值：
func (builder *FaceAttributeBuilder) Hat(hat *AttributeItem) *FaceAttributeBuilder {
	builder.hat = hat
	builder.hatFlag = true
	return builder
}

// 眼镜：0 未戴眼镜，1 戴眼镜
// 示例值：
func (builder *FaceAttributeBuilder) Glass(glass *AttributeItem) *FaceAttributeBuilder {
	builder.glass = glass
	builder.glassFlag = true
	return builder
}

// 口罩：0 未戴口罩，1 戴口罩
// 示例值：
func (builder *FaceAttributeBuilder) Mask(mask *AttributeItem) *FaceAttributeBuilder {
	builder.mask = mask
	builder.maskFlag = true
	return builder
}

func (builder *FaceAttributeBuilder) Build() *FaceAttribute {
	req := &FaceAttribute{}
	if builder.genderFlag {
		req.Gender = builder.gender
	}
	if builder.ageFlag {
		req.Age = &builder.age

	}
	if builder.emotionFlag {
		req.Emotion = builder.emotion
	}
	if builder.beautyFlag {
		req.Beauty = &builder.beauty

	}
	if builder.poseFlag {
		req.Pose = builder.pose
	}
	if builder.hatFlag {
		req.Hat = builder.hat
	}
	if builder.glassFlag {
		req.Glass = builder.glass
	}
	if builder.maskFlag {
		req.Mask = builder.mask
	}
	return req
}

type FaceInfo struct {
	Position  *FacePosition  `json:"position,omitempty"`  // 人脸位置信息
	Attribute *FaceAttribute `json:"attribute,omitempty"` // 人脸属性信息
	Quality   *FaceQuality   `json:"quality,omitempty"`   // 人脸质量信息
}

type FaceInfoBuilder struct {
	position      *FacePosition // 人脸位置信息
	positionFlag  bool
	attribute     *FaceAttribute // 人脸属性信息
	attributeFlag bool
	quality       *FaceQuality // 人脸质量信息
	qualityFlag   bool
}

func NewFaceInfoBuilder() *FaceInfoBuilder {
	builder := &FaceInfoBuilder{}
	return builder
}

// 人脸位置信息
// 示例值：
func (builder *FaceInfoBuilder) Position(position *FacePosition) *FaceInfoBuilder {
	builder.position = position
	builder.positionFlag = true
	return builder
}

// 人脸属性信息
// 示例值：
func (builder *FaceInfoBuilder) Attribute(attribute *FaceAttribute) *FaceInfoBuilder {
	builder.attribute = attribute
	builder.attributeFlag = true
	return builder
}

// 人脸质量信息
// 示例值：
func (builder *FaceInfoBuilder) Quality(quality *FaceQuality) *FaceInfoBuilder {
	builder.quality = quality
	builder.qualityFlag = true
	return builder
}

func (builder *FaceInfoBuilder) Build() *FaceInfo {
	req := &FaceInfo{}
	if builder.positionFlag {
		req.Position = builder.position
	}
	if builder.attributeFlag {
		req.Attribute = builder.attribute
	}
	if builder.qualityFlag {
		req.Quality = builder.quality
	}
	return req
}

type FaceOcclude struct {
	Eyebrow  *float64 `json:"eyebrow,omitempty"`   // 眉毛被遮挡情况：[0, 1] 值越大被遮挡的概率越高
	Nose     *float64 `json:"nose,omitempty"`      // 鼻子被遮挡情况：[0, 1] 值越大被遮挡的概率越高
	Cheek    *float64 `json:"cheek,omitempty"`     // 脸颊被遮挡情况：[0, 1] 值越大被遮挡的概率越高
	Mouth    *float64 `json:"mouth,omitempty"`     // 嘴被遮挡情况：[0, 1] 值越大被遮挡的概率越高
	Chin     *float64 `json:"chin,omitempty"`      // 下巴被遮挡情况：[0, 1] 值越大被遮挡的概率越高
	LeftEye  *float64 `json:"left_eye,omitempty"`  // 左眼睛被遮挡情况：[0, 1] 值越大被遮挡的概率越高
	RightEye *float64 `json:"right_eye,omitempty"` // 右眼睛被遮挡情况：[0, 1] 值越大被遮挡的概率越高
}

type FaceOccludeBuilder struct {
	eyebrow      float64 // 眉毛被遮挡情况：[0, 1] 值越大被遮挡的概率越高
	eyebrowFlag  bool
	nose         float64 // 鼻子被遮挡情况：[0, 1] 值越大被遮挡的概率越高
	noseFlag     bool
	cheek        float64 // 脸颊被遮挡情况：[0, 1] 值越大被遮挡的概率越高
	cheekFlag    bool
	mouth        float64 // 嘴被遮挡情况：[0, 1] 值越大被遮挡的概率越高
	mouthFlag    bool
	chin         float64 // 下巴被遮挡情况：[0, 1] 值越大被遮挡的概率越高
	chinFlag     bool
	leftEye      float64 // 左眼睛被遮挡情况：[0, 1] 值越大被遮挡的概率越高
	leftEyeFlag  bool
	rightEye     float64 // 右眼睛被遮挡情况：[0, 1] 值越大被遮挡的概率越高
	rightEyeFlag bool
}

func NewFaceOccludeBuilder() *FaceOccludeBuilder {
	builder := &FaceOccludeBuilder{}
	return builder
}

// 眉毛被遮挡情况：[0, 1] 值越大被遮挡的概率越高
// 示例值：0.02345
func (builder *FaceOccludeBuilder) Eyebrow(eyebrow float64) *FaceOccludeBuilder {
	builder.eyebrow = eyebrow
	builder.eyebrowFlag = true
	return builder
}

// 鼻子被遮挡情况：[0, 1] 值越大被遮挡的概率越高
// 示例值：0.99876
func (builder *FaceOccludeBuilder) Nose(nose float64) *FaceOccludeBuilder {
	builder.nose = nose
	builder.noseFlag = true
	return builder
}

// 脸颊被遮挡情况：[0, 1] 值越大被遮挡的概率越高
// 示例值：0.88767
func (builder *FaceOccludeBuilder) Cheek(cheek float64) *FaceOccludeBuilder {
	builder.cheek = cheek
	builder.cheekFlag = true
	return builder
}

// 嘴被遮挡情况：[0, 1] 值越大被遮挡的概率越高
// 示例值：0.45678
func (builder *FaceOccludeBuilder) Mouth(mouth float64) *FaceOccludeBuilder {
	builder.mouth = mouth
	builder.mouthFlag = true
	return builder
}

// 下巴被遮挡情况：[0, 1] 值越大被遮挡的概率越高
// 示例值：0.66436
func (builder *FaceOccludeBuilder) Chin(chin float64) *FaceOccludeBuilder {
	builder.chin = chin
	builder.chinFlag = true
	return builder
}

// 左眼睛被遮挡情况：[0, 1] 值越大被遮挡的概率越高
// 示例值：0.87699
func (builder *FaceOccludeBuilder) LeftEye(leftEye float64) *FaceOccludeBuilder {
	builder.leftEye = leftEye
	builder.leftEyeFlag = true
	return builder
}

// 右眼睛被遮挡情况：[0, 1] 值越大被遮挡的概率越高
// 示例值：0.78884
func (builder *FaceOccludeBuilder) RightEye(rightEye float64) *FaceOccludeBuilder {
	builder.rightEye = rightEye
	builder.rightEyeFlag = true
	return builder
}

func (builder *FaceOccludeBuilder) Build() *FaceOcclude {
	req := &FaceOcclude{}
	if builder.eyebrowFlag {
		req.Eyebrow = &builder.eyebrow

	}
	if builder.noseFlag {
		req.Nose = &builder.nose

	}
	if builder.cheekFlag {
		req.Cheek = &builder.cheek

	}
	if builder.mouthFlag {
		req.Mouth = &builder.mouth

	}
	if builder.chinFlag {
		req.Chin = &builder.chin

	}
	if builder.leftEyeFlag {
		req.LeftEye = &builder.leftEye

	}
	if builder.rightEyeFlag {
		req.RightEye = &builder.rightEye

	}
	return req
}

type FacePose struct {
	Pitch *int `json:"pitch,omitempty"` // 脸部上下偏移 [-90, 90]
	Yaw   *int `json:"yaw,omitempty"`   // 脸部左右偏移 [-90, 90]
	Roll  *int `json:"roll,omitempty"`  // 平面旋转 [-90, 90]
}

type FacePoseBuilder struct {
	pitch     int // 脸部上下偏移 [-90, 90]
	pitchFlag bool
	yaw       int // 脸部左右偏移 [-90, 90]
	yawFlag   bool
	roll      int // 平面旋转 [-90, 90]
	rollFlag  bool
}

func NewFacePoseBuilder() *FacePoseBuilder {
	builder := &FacePoseBuilder{}
	return builder
}

// 脸部上下偏移 [-90, 90]
// 示例值：-31
func (builder *FacePoseBuilder) Pitch(pitch int) *FacePoseBuilder {
	builder.pitch = pitch
	builder.pitchFlag = true
	return builder
}

// 脸部左右偏移 [-90, 90]
// 示例值：22
func (builder *FacePoseBuilder) Yaw(yaw int) *FacePoseBuilder {
	builder.yaw = yaw
	builder.yawFlag = true
	return builder
}

// 平面旋转 [-90, 90]
// 示例值：-32
func (builder *FacePoseBuilder) Roll(roll int) *FacePoseBuilder {
	builder.roll = roll
	builder.rollFlag = true
	return builder
}

func (builder *FacePoseBuilder) Build() *FacePose {
	req := &FacePose{}
	if builder.pitchFlag {
		req.Pitch = &builder.pitch

	}
	if builder.yawFlag {
		req.Yaw = &builder.yaw

	}
	if builder.rollFlag {
		req.Roll = &builder.roll

	}
	return req
}

type FacePosition struct {
	UpperLeft  *Point `json:"upper_left,omitempty"`  // 人脸框的左上角坐标
	LowerRight *Point `json:"lower_right,omitempty"` // 人脸框的右下角坐标
}

type FacePositionBuilder struct {
	upperLeft      *Point // 人脸框的左上角坐标
	upperLeftFlag  bool
	lowerRight     *Point // 人脸框的右下角坐标
	lowerRightFlag bool
}

func NewFacePositionBuilder() *FacePositionBuilder {
	builder := &FacePositionBuilder{}
	return builder
}

// 人脸框的左上角坐标
// 示例值：
func (builder *FacePositionBuilder) UpperLeft(upperLeft *Point) *FacePositionBuilder {
	builder.upperLeft = upperLeft
	builder.upperLeftFlag = true
	return builder
}

// 人脸框的右下角坐标
// 示例值：
func (builder *FacePositionBuilder) LowerRight(lowerRight *Point) *FacePositionBuilder {
	builder.lowerRight = lowerRight
	builder.lowerRightFlag = true
	return builder
}

func (builder *FacePositionBuilder) Build() *FacePosition {
	req := &FacePosition{}
	if builder.upperLeftFlag {
		req.UpperLeft = builder.upperLeft
	}
	if builder.lowerRightFlag {
		req.LowerRight = builder.lowerRight
	}
	return req
}

type FaceQuality struct {
	Sharpness  *float64     `json:"sharpness,omitempty"`  // 清晰度，值越高越清晰
	Brightness *float64     `json:"brightness,omitempty"` // 亮度
	Occlude    *FaceOcclude `json:"occlude,omitempty"`    // 面部遮挡属性
}

type FaceQualityBuilder struct {
	sharpness      float64 // 清晰度，值越高越清晰
	sharpnessFlag  bool
	brightness     float64 // 亮度
	brightnessFlag bool
	occlude        *FaceOcclude // 面部遮挡属性
	occludeFlag    bool
}

func NewFaceQualityBuilder() *FaceQualityBuilder {
	builder := &FaceQualityBuilder{}
	return builder
}

// 清晰度，值越高越清晰
// 示例值：0.77
func (builder *FaceQualityBuilder) Sharpness(sharpness float64) *FaceQualityBuilder {
	builder.sharpness = sharpness
	builder.sharpnessFlag = true
	return builder
}

// 亮度
// 示例值：0.6
func (builder *FaceQualityBuilder) Brightness(brightness float64) *FaceQualityBuilder {
	builder.brightness = brightness
	builder.brightnessFlag = true
	return builder
}

// 面部遮挡属性
// 示例值：
func (builder *FaceQualityBuilder) Occlude(occlude *FaceOcclude) *FaceQualityBuilder {
	builder.occlude = occlude
	builder.occludeFlag = true
	return builder
}

func (builder *FaceQualityBuilder) Build() *FaceQuality {
	req := &FaceQuality{}
	if builder.sharpnessFlag {
		req.Sharpness = &builder.sharpness

	}
	if builder.brightnessFlag {
		req.Brightness = &builder.brightness

	}
	if builder.occludeFlag {
		req.Occlude = builder.occlude
	}
	return req
}

type Image struct {
	Width  *int `json:"width,omitempty"`  // 图片的宽度
	Height *int `json:"height,omitempty"` // 图片的高度
}

type ImageBuilder struct {
	width      int // 图片的宽度
	widthFlag  bool
	height     int // 图片的高度
	heightFlag bool
}

func NewImageBuilder() *ImageBuilder {
	builder := &ImageBuilder{}
	return builder
}

// 图片的宽度
// 示例值：400
func (builder *ImageBuilder) Width(width int) *ImageBuilder {
	builder.width = width
	builder.widthFlag = true
	return builder
}

// 图片的高度
// 示例值：500
func (builder *ImageBuilder) Height(height int) *ImageBuilder {
	builder.height = height
	builder.heightFlag = true
	return builder
}

func (builder *ImageBuilder) Build() *Image {
	req := &Image{}
	if builder.widthFlag {
		req.Width = &builder.width

	}
	if builder.heightFlag {
		req.Height = &builder.height

	}
	return req
}

type Point struct {
	X *float64 `json:"x,omitempty"` // 横轴坐标
	Y *float64 `json:"y,omitempty"` // 纵轴坐标
}

type PointBuilder struct {
	x     float64 // 横轴坐标
	xFlag bool
	y     float64 // 纵轴坐标
	yFlag bool
}

func NewPointBuilder() *PointBuilder {
	builder := &PointBuilder{}
	return builder
}

// 横轴坐标
// 示例值：200
func (builder *PointBuilder) X(x float64) *PointBuilder {
	builder.x = x
	builder.xFlag = true
	return builder
}

// 纵轴坐标
// 示例值：200
func (builder *PointBuilder) Y(y float64) *PointBuilder {
	builder.y = y
	builder.yFlag = true
	return builder
}

func (builder *PointBuilder) Build() *Point {
	req := &Point{}
	if builder.xFlag {
		req.X = &builder.x

	}
	if builder.yFlag {
		req.Y = &builder.y

	}
	return req
}

type DetectFaceAttributesImageReqBodyBuilder struct {
	image     string // 图片 base64 数据
	imageFlag bool
}

func NewDetectFaceAttributesImageReqBodyBuilder() *DetectFaceAttributesImageReqBodyBuilder {
	builder := &DetectFaceAttributesImageReqBodyBuilder{}
	return builder
}

// 图片 base64 数据
//
//示例值：图片 base64 后的字符串
func (builder *DetectFaceAttributesImageReqBodyBuilder) Image(image string) *DetectFaceAttributesImageReqBodyBuilder {
	builder.image = image
	builder.imageFlag = true
	return builder
}

func (builder *DetectFaceAttributesImageReqBodyBuilder) Build() *DetectFaceAttributesImageReqBody {
	req := &DetectFaceAttributesImageReqBody{}
	if builder.imageFlag {
		req.Image = &builder.image
	}
	return req
}

type DetectFaceAttributesImagePathReqBodyBuilder struct {
	image     string // 图片 base64 数据
	imageFlag bool
}

func NewDetectFaceAttributesImagePathReqBodyBuilder() *DetectFaceAttributesImagePathReqBodyBuilder {
	builder := &DetectFaceAttributesImagePathReqBodyBuilder{}
	return builder
}

// 图片 base64 数据
//
// 示例值：图片 base64 后的字符串
func (builder *DetectFaceAttributesImagePathReqBodyBuilder) Image(image string) *DetectFaceAttributesImagePathReqBodyBuilder {
	builder.image = image
	builder.imageFlag = true
	return builder
}

func (builder *DetectFaceAttributesImagePathReqBodyBuilder) Build() (*DetectFaceAttributesImageReqBody, error) {
	req := &DetectFaceAttributesImageReqBody{}
	if builder.imageFlag {
		req.Image = &builder.image
	}
	return req, nil
}

type DetectFaceAttributesImageReqBuilder struct {
	apiReq *larkcore.ApiReq
	body   *DetectFaceAttributesImageReqBody
}

func NewDetectFaceAttributesImageReqBuilder() *DetectFaceAttributesImageReqBuilder {
	builder := &DetectFaceAttributesImageReqBuilder{}
	builder.apiReq = &larkcore.ApiReq{
		PathParams:  larkcore.PathParams{},
		QueryParams: larkcore.QueryParams{},
	}
	return builder
}

// 检测图片中的人脸属性和质量等信息
func (builder *DetectFaceAttributesImageReqBuilder) Body(body *DetectFaceAttributesImageReqBody) *DetectFaceAttributesImageReqBuilder {
	builder.body = body
	return builder
}

func (builder *DetectFaceAttributesImageReqBuilder) Build() *DetectFaceAttributesImageReq {
	req := &DetectFaceAttributesImageReq{}
	req.apiReq = &larkcore.ApiReq{}
	req.apiReq.Body = builder.body
	return req
}

type DetectFaceAttributesImageReqBody struct {
	Image *string `json:"image,omitempty"` // 图片 base64 数据
}

type DetectFaceAttributesImageReq struct {
	apiReq *larkcore.ApiReq
	Body   *DetectFaceAttributesImageReqBody `body:""`
}

type DetectFaceAttributesImageRespData struct {
	ImageInfo *Image      `json:"image_info,omitempty"` // 图片信息
	FaceInfos []*FaceInfo `json:"face_infos,omitempty"` // 人脸信息列表
}

type DetectFaceAttributesImageResp struct {
	*larkcore.ApiResp `json:"-"`
	larkcore.CodeError
	Data *DetectFaceAttributesImageRespData `json:"data"` // 业务数据
}

func (resp *DetectFaceAttributesImageResp) Success() bool {
	return resp.Code == 0
}
