package requests

import (
	"github.com/WeberLong/go-zmxy/zmop/interfaces"
)

type ZhimaCustomerCertifyApplyRequest struct {
	token string // 芝麻认证每一次请求返回的令牌，发起页面认证请求和认证请求结果查询接口都需要使用到该返回值作为入参

	interfaces.ZhimaRequestParams
}

func (m *ZhimaCustomerCertifyApplyRequest) InitBizParams(token string) {
	m.ApiParams = &map[string]string{}
	m.FileParams = &map[string]string{}

	(*m.ApiParams)["token"] = token
	m.token = token
}

func (*ZhimaCustomerCertifyApplyRequest) GetApiMethodName() string {
	return "zhima.customer.certify.apply"
}

func (m *ZhimaCustomerCertifyApplyRequest) GetApiParams() *map[string]string {
	return m.ApiParams
}

func (m *ZhimaCustomerCertifyApplyRequest) GetFileParams() *map[string]string {
	return m.FileParams
}

func (m *ZhimaCustomerCertifyApplyRequest) SetApiVersion(ApiVersion string) {
	m.ApiVersion = ApiVersion
}

func (m *ZhimaCustomerCertifyApplyRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaCustomerCertifyApplyRequest) SetScene(Scene string) {
	m.Scene = Scene
}

func (m *ZhimaCustomerCertifyApplyRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaCustomerCertifyApplyRequest) SetChannel(Channel string) {
	m.Channel = Channel
}

func (m *ZhimaCustomerCertifyApplyRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaCustomerCertifyApplyRequest) SetPlatform(Platform string) {
	m.Platform = Platform
}

func (m *ZhimaCustomerCertifyApplyRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaCustomerCertifyApplyRequest) SetExtParams(ExtParams string) {
	m.ExtParams = ExtParams
}

func (m *ZhimaCustomerCertifyApplyRequest) GetExtParams() string {
	return m.ExtParams
}
