package requests

import (
	"github.com/WeberLong/go-zmxy/zmop/interfaces"
)

type ZhimaCustomerCertifyCheckRequest struct {
	token string // 芝麻认证每一次请求返回的令牌，发起页面认证请求和认证请求结果查询接口都需要使用到该返回值作为入参

	interfaces.ZhimaRequestParams
}

func (m *ZhimaCustomerCertifyCheckRequest) InitBizParams(token string) {
	m.ApiParams = &map[string]string{}
	m.FileParams = &map[string]string{}

	(*m.ApiParams)["token"] = token
	m.token = token
}

func (*ZhimaCustomerCertifyCheckRequest) GetApiMethodName() string {
	return "zhima.customer.certify.check"
}

func (m *ZhimaCustomerCertifyCheckRequest) GetApiParams() *map[string]string {
	return m.ApiParams
}

func (m *ZhimaCustomerCertifyCheckRequest) GetFileParams() *map[string]string {
	return m.FileParams
}

func (m *ZhimaCustomerCertifyCheckRequest) SetApiVersion(ApiVersion string) {
	m.ApiVersion = ApiVersion
}

func (m *ZhimaCustomerCertifyCheckRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaCustomerCertifyCheckRequest) SetScene(Scene string) {
	m.Scene = Scene
}

func (m *ZhimaCustomerCertifyCheckRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaCustomerCertifyCheckRequest) SetChannel(Channel string) {
	m.Channel = Channel
}

func (m *ZhimaCustomerCertifyCheckRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaCustomerCertifyCheckRequest) SetPlatform(Platform string) {
	m.Platform = Platform
}

func (m *ZhimaCustomerCertifyCheckRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaCustomerCertifyCheckRequest) SetExtParams(ExtParams string) {
	m.ExtParams = ExtParams
}

func (m *ZhimaCustomerCertifyCheckRequest) GetExtParams() string {
	return m.ExtParams
}
