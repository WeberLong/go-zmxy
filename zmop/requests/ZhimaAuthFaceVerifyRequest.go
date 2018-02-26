package requests

import (
	"github.com/WeberLong/go-zmxy/zmop/interfaces"
)

type ZhimaAuthFaceVerifyRequest struct {
	bizType string // 活体认证类型，目前有认证和授权两种类型；默认为授权类型
	images  string //
	token   string // 标识一次请求流水

	interfaces.ZhimaRequestParams
}

func (m *ZhimaAuthFaceVerifyRequest) InitBizParams(bizType, images, token string) {
	m.ApiParams = &map[string]string{}
	m.FileParams = &map[string]string{}

	(*m.ApiParams)["biz_type"] = bizType
	m.bizType = bizType

	(*m.ApiParams)["images"] = images
	m.images = images

	(*m.ApiParams)["token"] = token
	m.token = token
}

func (*ZhimaAuthFaceVerifyRequest) GetApiMethodName() string {
	return "zhima.auth.face.verify"
}

func (m *ZhimaAuthFaceVerifyRequest) GetApiParams() *map[string]string {
	return m.ApiParams
}

func (m *ZhimaAuthFaceVerifyRequest) GetFileParams() *map[string]string {
	return m.FileParams
}

func (m *ZhimaAuthFaceVerifyRequest) SetApiVersion(ApiVersion string) {
	m.ApiVersion = ApiVersion
}

func (m *ZhimaAuthFaceVerifyRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaAuthFaceVerifyRequest) SetScene(Scene string) {
	m.Scene = Scene
}

func (m *ZhimaAuthFaceVerifyRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaAuthFaceVerifyRequest) SetChannel(Channel string) {
	m.Channel = Channel
}

func (m *ZhimaAuthFaceVerifyRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaAuthFaceVerifyRequest) SetPlatform(Platform string) {
	m.Platform = Platform
}

func (m *ZhimaAuthFaceVerifyRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaAuthFaceVerifyRequest) SetExtParams(ExtParams string) {
	m.ExtParams = ExtParams
}

func (m *ZhimaAuthFaceVerifyRequest) GetExtParams() string {
	return m.ExtParams
}
