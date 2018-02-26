package requests

import (
	"github.com/WeberLong/go-zmxy/zmop/interfaces"
)

type ZhimaCreditScoreBriefGetRequest struct {
	admittanceScore string // 350～950之间 业务判断的准入标准 建议业务确定一个稳定的判断标准 频繁的变更该标准可能导致接口被停用
	certNo          string // 证件号
	certType        string // 证件类型 目前支持三种IDENTITY_CARD(身份证),PASSPORT(护照),ALIPAY_USER_ID(支付宝uid)
	name            string // 姓名
	productCode     string // 产品码，直接使用［示例］给出的值
	transactionId   string // transaction_id是代表一笔请求的唯一标志，该标识作为对账的关键信息，对于用户使用相同transaction_id的查询，芝麻在一天（86400秒）内返回首次查询数据，超过有效期的查询即为无效并返回异常，有效期内的反复查询不重新计费。transaction_id 推荐生成方式是：30位，（其中17位时间值（精确到毫秒）：yyyyMMddHHmmssSSS）加上（13位自增数字：1234567890123）

	interfaces.ZhimaRequestParams
}

func (m *ZhimaCreditScoreBriefGetRequest) InitBizParams(admittanceScore, certNo, certType, name, productCode, transactionId string) {
	m.ApiParams = &map[string]string{}
	m.FileParams = &map[string]string{}

	(*m.ApiParams)["admittance_score"] = admittanceScore
	m.admittanceScore = admittanceScore

	(*m.ApiParams)["cert_no"] = certNo
	m.certNo = certNo

	(*m.ApiParams)["cert_type"] = certType
	m.certType = certType

	(*m.ApiParams)["name"] = name
	m.name = name

	(*m.ApiParams)["product_code"] = productCode
	m.productCode = productCode

	(*m.ApiParams)["transaction_id"] = transactionId
	m.transactionId = transactionId
}

func (*ZhimaCreditScoreBriefGetRequest) GetApiMethodName() string {
	return "zhima.credit.score.brief.get"
}

func (m *ZhimaCreditScoreBriefGetRequest) GetApiParams() *map[string]string {
	return m.ApiParams
}

func (m *ZhimaCreditScoreBriefGetRequest) GetFileParams() *map[string]string {
	return m.FileParams
}

func (m *ZhimaCreditScoreBriefGetRequest) SetApiVersion(ApiVersion string) {
	m.ApiVersion = ApiVersion
}

func (m *ZhimaCreditScoreBriefGetRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaCreditScoreBriefGetRequest) SetScene(Scene string) {
	m.Scene = Scene
}

func (m *ZhimaCreditScoreBriefGetRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaCreditScoreBriefGetRequest) SetChannel(Channel string) {
	m.Channel = Channel
}

func (m *ZhimaCreditScoreBriefGetRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaCreditScoreBriefGetRequest) SetPlatform(Platform string) {
	m.Platform = Platform
}

func (m *ZhimaCreditScoreBriefGetRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaCreditScoreBriefGetRequest) SetExtParams(ExtParams string) {
	m.ExtParams = ExtParams
}

func (m *ZhimaCreditScoreBriefGetRequest) GetExtParams() string {
	return m.ExtParams
}
