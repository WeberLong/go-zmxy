package requests

import (
	"github.com/WeberLong/go-zmxy/zmop/interfaces"
)

type ZhimaCustomerCertificationCertifyRequest struct {
	bizNo     string // 一次认证的唯一标识，在完成芝麻认证初始化后可以获取
	returnUrl string // 商户回调地址，在用户完成认证后会调转回商户地址

	interfaces.ZhimaRequestParams
}

func (m *ZhimaCustomerCertificationCertifyRequest) InitBizParams(bizNo, returnUrl string) {
	m.ApiParams = &map[string]string{}
	m.FileParams = &map[string]string{}

	(*m.ApiParams)["biz_no"] = bizNo
	m.bizNo = bizNo

	(*m.ApiParams)["return_url"] = returnUrl
	m.returnUrl = returnUrl
}

func (*ZhimaCustomerCertificationCertifyRequest) GetApiMethodName() string {
	return "zhima.customer.certification.certify"
}

func (m *ZhimaCustomerCertificationCertifyRequest) GetApiParams() *map[string]string {
	return m.ApiParams
}

func (m *ZhimaCustomerCertificationCertifyRequest) GetFileParams() *map[string]string {
	return m.FileParams
}

func (m *ZhimaCustomerCertificationCertifyRequest) SetApiVersion(ApiVersion string) {
	m.ApiVersion = ApiVersion
}

func (m *ZhimaCustomerCertificationCertifyRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaCustomerCertificationCertifyRequest) SetScene(Scene string) {
	m.Scene = Scene
}

func (m *ZhimaCustomerCertificationCertifyRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaCustomerCertificationCertifyRequest) SetChannel(Channel string) {
	m.Channel = Channel
}

func (m *ZhimaCustomerCertificationCertifyRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaCustomerCertificationCertifyRequest) SetPlatform(Platform string) {
	m.Platform = Platform
}

func (m *ZhimaCustomerCertificationCertifyRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaCustomerCertificationCertifyRequest) SetExtParams(ExtParams string) {
	m.ExtParams = ExtParams
}

func (m *ZhimaCustomerCertificationCertifyRequest) GetExtParams() string {
	return m.ExtParams
}
