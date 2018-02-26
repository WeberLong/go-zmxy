package requests

import (
	"github.com/WeberLong/go-zmxy/zmop/interfaces"
)

type ZhimaCustomerEpCertificationQueryRequest struct {
	bizNo string // 一次认证的唯一标识，在商户调用认证初始化接口的时候获取

	interfaces.ZhimaRequestParams
}

func (m *ZhimaCustomerEpCertificationQueryRequest) InitBizParams(bizNo string) {
	m.ApiParams = &map[string]string{}
	m.FileParams = &map[string]string{}

	(*m.ApiParams)["biz_no"] = bizNo
	m.bizNo = bizNo
}

func (*ZhimaCustomerEpCertificationQueryRequest) GetApiMethodName() string {
	return "zhima.customer.ep.certification.query"
}

func (m *ZhimaCustomerEpCertificationQueryRequest) GetApiParams() *map[string]string {
	return m.ApiParams
}

func (m *ZhimaCustomerEpCertificationQueryRequest) GetFileParams() *map[string]string {
	return m.FileParams
}

func (m *ZhimaCustomerEpCertificationQueryRequest) SetApiVersion(ApiVersion string) {
	m.ApiVersion = ApiVersion
}

func (m *ZhimaCustomerEpCertificationQueryRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaCustomerEpCertificationQueryRequest) SetScene(Scene string) {
	m.Scene = Scene
}

func (m *ZhimaCustomerEpCertificationQueryRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaCustomerEpCertificationQueryRequest) SetChannel(Channel string) {
	m.Channel = Channel
}

func (m *ZhimaCustomerEpCertificationQueryRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaCustomerEpCertificationQueryRequest) SetPlatform(Platform string) {
	m.Platform = Platform
}

func (m *ZhimaCustomerEpCertificationQueryRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaCustomerEpCertificationQueryRequest) SetExtParams(ExtParams string) {
	m.ExtParams = ExtParams
}

func (m *ZhimaCustomerEpCertificationQueryRequest) GetExtParams() string {
	return m.ExtParams
}
