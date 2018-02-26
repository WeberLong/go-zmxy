package requests

import (
	"github.com/WeberLong/go-zmxy/zmop/interfaces"
)

type ZhimaCustomerCertificationQueryRequest struct {
	bizNo string // 一次认证的唯一标识，在商户调用认证初始化接口的时候获取

	interfaces.ZhimaRequestParams
}

func (m *ZhimaCustomerCertificationQueryRequest) InitBizParams(bizNo string) {
	m.ApiParams = &map[string]string{}
	m.FileParams = &map[string]string{}

	(*m.ApiParams)["biz_no"] = bizNo
	m.bizNo = bizNo
}

func (*ZhimaCustomerCertificationQueryRequest) GetApiMethodName() string {
	return "zhima.customer.certification.query"
}

func (m *ZhimaCustomerCertificationQueryRequest) GetApiParams() *map[string]string {
	return m.ApiParams
}

func (m *ZhimaCustomerCertificationQueryRequest) GetFileParams() *map[string]string {
	return m.FileParams
}

func (m *ZhimaCustomerCertificationQueryRequest) SetApiVersion(ApiVersion string) {
	m.ApiVersion = ApiVersion
}

func (m *ZhimaCustomerCertificationQueryRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaCustomerCertificationQueryRequest) SetScene(Scene string) {
	m.Scene = Scene
}

func (m *ZhimaCustomerCertificationQueryRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaCustomerCertificationQueryRequest) SetChannel(Channel string) {
	m.Channel = Channel
}

func (m *ZhimaCustomerCertificationQueryRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaCustomerCertificationQueryRequest) SetPlatform(Platform string) {
	m.Platform = Platform
}

func (m *ZhimaCustomerCertificationQueryRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaCustomerCertificationQueryRequest) SetExtParams(ExtParams string) {
	m.ExtParams = ExtParams
}

func (m *ZhimaCustomerCertificationQueryRequest) GetExtParams() string {
	return m.ExtParams
}
