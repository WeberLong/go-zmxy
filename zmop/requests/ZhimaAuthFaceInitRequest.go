package requests

import (
	"github.com/WeberLong/go-zmxy/zmop/interfaces"
)

type ZhimaAuthFaceInitRequest struct {
	apiKey   string // 应用的标识
	authMsg  string // 参数的加密串
	bizType  string // 用于标识使用人脸业务的类型是授权或者认证，默认为授权类型
	bundleId string // 不同商户的bundle_id,用来做缓存
	token    string // 请求的sessionId

	interfaces.ZhimaRequestParams
}

func (m *ZhimaAuthFaceInitRequest) InitBizParams(apiKey, authMsg, bizType, bundleId, token string) {
	m.ApiParams = &map[string]string{}
	m.FileParams = &map[string]string{}

	(*m.ApiParams)["api_key"] = apiKey
	m.apiKey = apiKey

	(*m.ApiParams)["auth_msg"] = authMsg
	m.authMsg = authMsg

	(*m.ApiParams)["biz_type"] = bizType
	m.bizType = bizType

	(*m.ApiParams)["bundle_id"] = bundleId
	m.bundleId = bundleId

	(*m.ApiParams)["token"] = token
	m.token = token
}

func (*ZhimaAuthFaceInitRequest) GetApiMethodName() string {
	return "zhima.auth.face.init"
}

func (m *ZhimaAuthFaceInitRequest) GetApiParams() *map[string]string {
	return m.ApiParams
}

func (m *ZhimaAuthFaceInitRequest) GetFileParams() *map[string]string {
	return m.FileParams
}

func (m *ZhimaAuthFaceInitRequest) SetApiVersion(ApiVersion string) {
	m.ApiVersion = ApiVersion
}

func (m *ZhimaAuthFaceInitRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaAuthFaceInitRequest) SetScene(Scene string) {
	m.Scene = Scene
}

func (m *ZhimaAuthFaceInitRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaAuthFaceInitRequest) SetChannel(Channel string) {
	m.Channel = Channel
}

func (m *ZhimaAuthFaceInitRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaAuthFaceInitRequest) SetPlatform(Platform string) {
	m.Platform = Platform
}

func (m *ZhimaAuthFaceInitRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaAuthFaceInitRequest) SetExtParams(ExtParams string) {
	m.ExtParams = ExtParams
}

func (m *ZhimaAuthFaceInitRequest) GetExtParams() string {
	return m.ExtParams
}
