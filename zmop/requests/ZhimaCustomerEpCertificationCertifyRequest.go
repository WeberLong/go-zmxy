package requests

import (
	"github.com/WeberLong/go-zmxy/zmop/interfaces"
)

type ZhimaCustomerEpCertificationCertifyRequest struct {
	bizNo     string // 一次认证的唯一标识，在完成芝麻认证初始化后可以获取
	returnUrl string // 商户回调地址，在用户完成认证后会调转回商户地址

	interfaces.ZhimaRequestParams
}

func (m *ZhimaCustomerEpCertificationCertifyRequest) InitBizParams(bizNo, returnUrl string) {
	m.ApiParams = &map[string]string{}
	m.FileParams = &map[string]string{}

	(*m.ApiParams)["biz_no"] = bizNo
	m.bizNo = bizNo

	(*m.ApiParams)["return_url"] = returnUrl
	m.returnUrl = returnUrl
}

func (*ZhimaCustomerEpCertificationCertifyRequest) GetApiMethodName() string {
	return "zhima.customer.ep.certification.certify"
}

func (m *ZhimaCustomerEpCertificationCertifyRequest) GetApiParams() *map[string]string {
	return m.ApiParams
}

func (m *ZhimaCustomerEpCertificationCertifyRequest) GetFileParams() *map[string]string {
	return m.FileParams
}

func (m *ZhimaCustomerEpCertificationCertifyRequest) SetApiVersion(ApiVersion string) {
	m.ApiVersion = ApiVersion
}

func (m *ZhimaCustomerEpCertificationCertifyRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaCustomerEpCertificationCertifyRequest) SetScene(Scene string) {
	m.Scene = Scene
}

func (m *ZhimaCustomerEpCertificationCertifyRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaCustomerEpCertificationCertifyRequest) SetChannel(Channel string) {
	m.Channel = Channel
}

func (m *ZhimaCustomerEpCertificationCertifyRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaCustomerEpCertificationCertifyRequest) SetPlatform(Platform string) {
	m.Platform = Platform
}

func (m *ZhimaCustomerEpCertificationCertifyRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaCustomerEpCertificationCertifyRequest) SetExtParams(ExtParams string) {
	m.ExtParams = ExtParams
}

func (m *ZhimaCustomerEpCertificationCertifyRequest) GetExtParams() string {
	return m.ExtParams
}
