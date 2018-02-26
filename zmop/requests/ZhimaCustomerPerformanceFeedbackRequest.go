package requests

import (
	"github.com/WeberLong/go-zmxy/zmop/interfaces"
)

type ZhimaCustomerPerformanceFeedbackRequest struct {
	certNo     string // 用户证件号码
	certType   string // 证件类型
	name       string // 用户姓名
	repayments string // 商户反馈的某用户还款计划数据，格式：[{"repayment_id": "1234","repayment_desc": "商品名称","installments": [{"installment_id": "1234","installment_amount": "1000","installment_date": "2016-09-11 12:00:00","installment_desc": "已逾期7天","installment_status": "WAITING_REPAY","currency": "CNY"}]}]

	interfaces.ZhimaRequestParams
}

func (m *ZhimaCustomerPerformanceFeedbackRequest) InitBizParams(certNo, certType, name, repayments string) {
	m.ApiParams = &map[string]string{}
	m.FileParams = &map[string]string{}

	(*m.ApiParams)["cert_no"] = certNo
	m.certNo = certNo

	(*m.ApiParams)["cert_type"] = certType
	m.certType = certType

	(*m.ApiParams)["name"] = name
	m.name = name

	(*m.ApiParams)["repayments"] = repayments
	m.repayments = repayments
}

func (*ZhimaCustomerPerformanceFeedbackRequest) GetApiMethodName() string {
	return "zhima.customer.performance.feedback"
}

func (m *ZhimaCustomerPerformanceFeedbackRequest) GetApiParams() *map[string]string {
	return m.ApiParams
}

func (m *ZhimaCustomerPerformanceFeedbackRequest) GetFileParams() *map[string]string {
	return m.FileParams
}

func (m *ZhimaCustomerPerformanceFeedbackRequest) SetApiVersion(ApiVersion string) {
	m.ApiVersion = ApiVersion
}

func (m *ZhimaCustomerPerformanceFeedbackRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaCustomerPerformanceFeedbackRequest) SetScene(Scene string) {
	m.Scene = Scene
}

func (m *ZhimaCustomerPerformanceFeedbackRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaCustomerPerformanceFeedbackRequest) SetChannel(Channel string) {
	m.Channel = Channel
}

func (m *ZhimaCustomerPerformanceFeedbackRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaCustomerPerformanceFeedbackRequest) SetPlatform(Platform string) {
	m.Platform = Platform
}

func (m *ZhimaCustomerPerformanceFeedbackRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaCustomerPerformanceFeedbackRequest) SetExtParams(ExtParams string) {
	m.ExtParams = ExtParams
}

func (m *ZhimaCustomerPerformanceFeedbackRequest) GetExtParams() string {
	return m.ExtParams
}
