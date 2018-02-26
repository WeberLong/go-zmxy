package requests

import (
	"github.com/WeberLong/go-zmxy/zmop/interfaces"
)

type ZhimaCustomerCreditrecordSummaryQueryRequest struct {
	alipayUserId string // 用户支付宝ID
	bizScene     string // 足迹业务场景

	interfaces.ZhimaRequestParams
}

func (m *ZhimaCustomerCreditrecordSummaryQueryRequest) InitBizParams(alipayUserId, bizScene string) {
	m.ApiParams = &map[string]string{}
	m.FileParams = &map[string]string{}

	(*m.ApiParams)["alipay_user_id"] = alipayUserId
	m.alipayUserId = alipayUserId

	(*m.ApiParams)["biz_scene"] = bizScene
	m.bizScene = bizScene
}

func (*ZhimaCustomerCreditrecordSummaryQueryRequest) GetApiMethodName() string {
	return "zhima.customer.creditrecord.summary.query"
}

func (m *ZhimaCustomerCreditrecordSummaryQueryRequest) GetApiParams() *map[string]string {
	return m.ApiParams
}

func (m *ZhimaCustomerCreditrecordSummaryQueryRequest) GetFileParams() *map[string]string {
	return m.FileParams
}

func (m *ZhimaCustomerCreditrecordSummaryQueryRequest) SetApiVersion(ApiVersion string) {
	m.ApiVersion = ApiVersion
}

func (m *ZhimaCustomerCreditrecordSummaryQueryRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaCustomerCreditrecordSummaryQueryRequest) SetScene(Scene string) {
	m.Scene = Scene
}

func (m *ZhimaCustomerCreditrecordSummaryQueryRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaCustomerCreditrecordSummaryQueryRequest) SetChannel(Channel string) {
	m.Channel = Channel
}

func (m *ZhimaCustomerCreditrecordSummaryQueryRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaCustomerCreditrecordSummaryQueryRequest) SetPlatform(Platform string) {
	m.Platform = Platform
}

func (m *ZhimaCustomerCreditrecordSummaryQueryRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaCustomerCreditrecordSummaryQueryRequest) SetExtParams(ExtParams string) {
	m.ExtParams = ExtParams
}

func (m *ZhimaCustomerCreditrecordSummaryQueryRequest) GetExtParams() string {
	return m.ExtParams
}
