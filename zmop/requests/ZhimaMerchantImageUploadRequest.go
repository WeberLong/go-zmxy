package requests

import (
	"github.com/WeberLong/go-zmxy/zmop/interfaces"
)

type ZhimaMerchantImageUploadRequest struct {
	imageContent string // 图片的二进制内容,最大支持2M，需要对图片的二进制流做Base64处理转化成字符串
	imageType    string // 芝麻二级商户图标后缀：bmp, jpg, jpeg, png, gif

	interfaces.ZhimaRequestParams
}

func (m *ZhimaMerchantImageUploadRequest) InitBizParams(imageContent, imageType string) {
	m.ApiParams = &map[string]string{}
	m.FileParams = &map[string]string{}

	(*m.ApiParams)["image_content"] = imageContent
	m.imageContent = imageContent

	(*m.ApiParams)["image_type"] = imageType
	m.imageType = imageType
}

func (*ZhimaMerchantImageUploadRequest) GetApiMethodName() string {
	return "zhima.merchant.image.upload"
}

func (m *ZhimaMerchantImageUploadRequest) GetApiParams() *map[string]string {
	return m.ApiParams
}

func (m *ZhimaMerchantImageUploadRequest) GetFileParams() *map[string]string {
	return m.FileParams
}

func (m *ZhimaMerchantImageUploadRequest) SetApiVersion(ApiVersion string) {
	m.ApiVersion = ApiVersion
}

func (m *ZhimaMerchantImageUploadRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaMerchantImageUploadRequest) SetScene(Scene string) {
	m.Scene = Scene
}

func (m *ZhimaMerchantImageUploadRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaMerchantImageUploadRequest) SetChannel(Channel string) {
	m.Channel = Channel
}

func (m *ZhimaMerchantImageUploadRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaMerchantImageUploadRequest) SetPlatform(Platform string) {
	m.Platform = Platform
}

func (m *ZhimaMerchantImageUploadRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaMerchantImageUploadRequest) SetExtParams(ExtParams string) {
	m.ExtParams = ExtParams
}

func (m *ZhimaMerchantImageUploadRequest) GetExtParams() string {
	return m.ExtParams
}
