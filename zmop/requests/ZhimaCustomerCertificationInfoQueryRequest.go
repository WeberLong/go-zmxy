package requests

import (
	"github.com/WeberLong/go-zmxy/zmop/interfaces"
)

type ZhimaCustomerCertificationInfoQueryRequest struct {
	bizNo      string // 一次认证的唯一标识，在商户调用认证初始化接口的时候获取
	merchantId string // 商户id，商户在芝麻的唯一标识

	interfaces.ZhimaRequestParams
}

func (m *ZhimaCustomerCertificationInfoQueryRequest) InitBizParams(bizNo, merchantId string) {
	m.ApiParams = &map[string]string{}
	m.FileParams = &map[string]string{}

	(*m.ApiParams)["biz_no"] = bizNo
	m.bizNo = bizNo

	(*m.ApiParams)["merchant_id"] = merchantId
	m.merchantId = merchantId
}

func (*ZhimaCustomerCertificationInfoQueryRequest) GetApiMethodName() string {
	return "zhima.customer.certification.info.query"
}

func (m *ZhimaCustomerCertificationInfoQueryRequest) GetApiParams() *map[string]string {
	return m.ApiParams
}

func (m *ZhimaCustomerCertificationInfoQueryRequest) GetFileParams() *map[string]string {
	return m.FileParams
}

func (m *ZhimaCustomerCertificationInfoQueryRequest) SetApiVersion(ApiVersion string) {
	m.ApiVersion = ApiVersion
}

func (m *ZhimaCustomerCertificationInfoQueryRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaCustomerCertificationInfoQueryRequest) SetScene(Scene string) {
	m.Scene = Scene
}

func (m *ZhimaCustomerCertificationInfoQueryRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaCustomerCertificationInfoQueryRequest) SetChannel(Channel string) {
	m.Channel = Channel
}

func (m *ZhimaCustomerCertificationInfoQueryRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaCustomerCertificationInfoQueryRequest) SetPlatform(Platform string) {
	m.Platform = Platform
}

func (m *ZhimaCustomerCertificationInfoQueryRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaCustomerCertificationInfoQueryRequest) SetExtParams(ExtParams string) {
	m.ExtParams = ExtParams
}

func (m *ZhimaCustomerCertificationInfoQueryRequest) GetExtParams() string {
	return m.ExtParams
}
