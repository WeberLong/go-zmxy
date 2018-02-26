package requests

import (
	"github.com/WeberLong/go-zmxy/zmop/interfaces"
)

type ZhimaCreditReportFinanceGetRequest struct {
	bizScene      string // 调用次产品的使用场景：SN001      贷前审批SN002      贷后管理SN003      风控模型开发SN004      风险定价SN000      其他
	openId        string // 芝麻会员在商户端的身份标识
	productCode   string // 产品码
	transactionId string // 商户传入的业务流水号。此字段由商户生成，需确保唯一性，用于定位每一次请求，后续按此流水进行对帐。生成规则: 固定30位数字串，前17位为精确到毫秒的时间yyyyMMddHHmmssSSS，后13位为自增数字

	interfaces.ZhimaRequestParams
}

func (m *ZhimaCreditReportFinanceGetRequest) InitBizParams(bizScene, openId, productCode, transactionId string) {
	m.ApiParams = &map[string]string{}
	m.FileParams = &map[string]string{}

	(*m.ApiParams)["biz_scene"] = bizScene
	m.bizScene = bizScene

	(*m.ApiParams)["open_id"] = openId
	m.openId = openId

	(*m.ApiParams)["product_code"] = productCode
	m.productCode = productCode

	(*m.ApiParams)["transaction_id"] = transactionId
	m.transactionId = transactionId
}

func (*ZhimaCreditReportFinanceGetRequest) GetApiMethodName() string {
	return "zhima.credit.report.finance.get"
}

func (m *ZhimaCreditReportFinanceGetRequest) GetApiParams() *map[string]string {
	return m.ApiParams
}

func (m *ZhimaCreditReportFinanceGetRequest) GetFileParams() *map[string]string {
	return m.FileParams
}

func (m *ZhimaCreditReportFinanceGetRequest) SetApiVersion(ApiVersion string) {
	m.ApiVersion = ApiVersion
}

func (m *ZhimaCreditReportFinanceGetRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaCreditReportFinanceGetRequest) SetScene(Scene string) {
	m.Scene = Scene
}

func (m *ZhimaCreditReportFinanceGetRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaCreditReportFinanceGetRequest) SetChannel(Channel string) {
	m.Channel = Channel
}

func (m *ZhimaCreditReportFinanceGetRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaCreditReportFinanceGetRequest) SetPlatform(Platform string) {
	m.Platform = Platform
}

func (m *ZhimaCreditReportFinanceGetRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaCreditReportFinanceGetRequest) SetExtParams(ExtParams string) {
	m.ExtParams = ExtParams
}

func (m *ZhimaCreditReportFinanceGetRequest) GetExtParams() string {
	return m.ExtParams
}
