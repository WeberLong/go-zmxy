package requests

import (
	"github.com/WeberLong/go-zmxy/zmop/interfaces"
)

type ZhimaMerchantDataUploadInitializeRequest struct {
	sceneCode string //

	interfaces.ZhimaRequestParams
}

func (m *ZhimaMerchantDataUploadInitializeRequest) InitBizParams(sceneCode string) {
	m.ApiParams = &map[string]string{}
	m.FileParams = &map[string]string{}

	(*m.ApiParams)["scene_code"] = sceneCode
	m.sceneCode = sceneCode
}

func (*ZhimaMerchantDataUploadInitializeRequest) GetApiMethodName() string {
	return "zhima.merchant.data.upload.initialize"
}

func (m *ZhimaMerchantDataUploadInitializeRequest) GetApiParams() *map[string]string {
	return m.ApiParams
}

func (m *ZhimaMerchantDataUploadInitializeRequest) GetFileParams() *map[string]string {
	return m.FileParams
}

func (m *ZhimaMerchantDataUploadInitializeRequest) SetApiVersion(ApiVersion string) {
	m.ApiVersion = ApiVersion
}

func (m *ZhimaMerchantDataUploadInitializeRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaMerchantDataUploadInitializeRequest) SetScene(Scene string) {
	m.Scene = Scene
}

func (m *ZhimaMerchantDataUploadInitializeRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaMerchantDataUploadInitializeRequest) SetChannel(Channel string) {
	m.Channel = Channel
}

func (m *ZhimaMerchantDataUploadInitializeRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaMerchantDataUploadInitializeRequest) SetPlatform(Platform string) {
	m.Platform = Platform
}

func (m *ZhimaMerchantDataUploadInitializeRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaMerchantDataUploadInitializeRequest) SetExtParams(ExtParams string) {
	m.ExtParams = ExtParams
}

func (m *ZhimaMerchantDataUploadInitializeRequest) GetExtParams() string {
	return m.ExtParams
}
