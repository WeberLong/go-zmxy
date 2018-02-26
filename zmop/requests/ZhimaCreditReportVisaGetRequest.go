package requests

import (
	"github.com/WeberLong/go-zmxy/zmop/interfaces"
)

type ZhimaCreditReportVisaGetRequest struct {
	openId        string // 芝麻会员在商户端的身份标识
	productCode   string // 产品码
	transactionId string // 商户传入的业务流水号

	interfaces.ZhimaRequestParams
}

func (m *ZhimaCreditReportVisaGetRequest) InitBizParams(openId, productCode, transactionId string) {
	m.ApiParams = &map[string]string{}
	m.FileParams = &map[string]string{}

	(*m.ApiParams)["open_id"] = openId
	m.openId = openId

	(*m.ApiParams)["product_code"] = productCode
	m.productCode = productCode

	(*m.ApiParams)["transaction_id"] = transactionId
	m.transactionId = transactionId
}

func (*ZhimaCreditReportVisaGetRequest) GetApiMethodName() string {
	return "zhima.credit.report.visa.get"
}

func (m *ZhimaCreditReportVisaGetRequest) GetApiParams() *map[string]string {
	return m.ApiParams
}

func (m *ZhimaCreditReportVisaGetRequest) GetFileParams() *map[string]string {
	return m.FileParams
}

func (m *ZhimaCreditReportVisaGetRequest) SetApiVersion(ApiVersion string) {
	m.ApiVersion = ApiVersion
}

func (m *ZhimaCreditReportVisaGetRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaCreditReportVisaGetRequest) SetScene(Scene string) {
	m.Scene = Scene
}

func (m *ZhimaCreditReportVisaGetRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaCreditReportVisaGetRequest) SetChannel(Channel string) {
	m.Channel = Channel
}

func (m *ZhimaCreditReportVisaGetRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaCreditReportVisaGetRequest) SetPlatform(Platform string) {
	m.Platform = Platform
}

func (m *ZhimaCreditReportVisaGetRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaCreditReportVisaGetRequest) SetExtParams(ExtParams string) {
	m.ExtParams = ExtParams
}

func (m *ZhimaCreditReportVisaGetRequest) GetExtParams() string {
	return m.ExtParams
}
