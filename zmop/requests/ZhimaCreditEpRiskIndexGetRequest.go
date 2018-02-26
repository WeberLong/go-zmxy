package requests

import (
	"github.com/WeberLong/go-zmxy/zmop/interfaces"
)

type ZhimaCreditEpRiskIndexGetRequest struct {
	productCode   string // 产品码
	transactionId string // transaction_id是代表一笔请求的唯一标志，该标识作为对账的关键信息，对于用户使用相同transaction_id的查询，芝麻在一天（86400秒）内返回首次查询数据，超过有效期的查询即为无效并返回异常，有效期内的反复查询不重新计费。 transaction_id 推荐生成方式是：30位，（其中17位时间值（精确到毫秒）：yyyyMMddHHmmssSSS）加上（13位自增数字：1234567890123）
	opType        string // -1： 所有风险指数

	interfaces.ZhimaRequestParams
}

func (m *ZhimaCreditEpRiskIndexGetRequest) InitBizParams(productCode, transactionId, opType string) {
	m.ApiParams = &map[string]string{}
	m.FileParams = &map[string]string{}

	(*m.ApiParams)["product_code"] = productCode
	m.productCode = productCode

	(*m.ApiParams)["transaction_id"] = transactionId
	m.transactionId = transactionId

	(*m.ApiParams)["op_type"] = opType
	m.opType = opType
}

func (*ZhimaCreditEpRiskIndexGetRequest) GetApiMethodName() string {
	return "zhima.credit.ep.risk.index.get"
}

func (m *ZhimaCreditEpRiskIndexGetRequest) GetApiParams() *map[string]string {
	return m.ApiParams
}

func (m *ZhimaCreditEpRiskIndexGetRequest) GetFileParams() *map[string]string {
	return m.FileParams
}

func (m *ZhimaCreditEpRiskIndexGetRequest) SetApiVersion(ApiVersion string) {
	m.ApiVersion = ApiVersion
}

func (m *ZhimaCreditEpRiskIndexGetRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaCreditEpRiskIndexGetRequest) SetScene(Scene string) {
	m.Scene = Scene
}

func (m *ZhimaCreditEpRiskIndexGetRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaCreditEpRiskIndexGetRequest) SetChannel(Channel string) {
	m.Channel = Channel
}

func (m *ZhimaCreditEpRiskIndexGetRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaCreditEpRiskIndexGetRequest) SetPlatform(Platform string) {
	m.Platform = Platform
}

func (m *ZhimaCreditEpRiskIndexGetRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaCreditEpRiskIndexGetRequest) SetExtParams(ExtParams string) {
	m.ExtParams = ExtParams
}

func (m *ZhimaCreditEpRiskIndexGetRequest) GetExtParams() string {
	return m.ExtParams
}
