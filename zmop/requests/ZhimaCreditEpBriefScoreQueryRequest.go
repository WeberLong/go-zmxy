package requests

import (
	"github.com/WeberLong/go-zmxy/zmop/interfaces"
)

type ZhimaCreditEpBriefScoreQueryRequest struct {
	name          string // 企业名称
	productCode   string // 产品码
	transactionId string // transaction_id是代表一笔请求的唯一标志，该标识作为对账的关键信息，对于用户使用相同transaction_id的查询，芝麻在一天（86400秒）内返回首次查询数据，超过有效期的查询即为无效并返回异常（错误码xxxx），有效期内的重复查询不重新计费。 transaction_id 推荐生成方式是：30位，（其中17位时间值（精确到毫秒）：yyyyMMddHHmmssSSS）加上（13位自增数字：1234567890123）

	interfaces.ZhimaRequestParams
}

func (m *ZhimaCreditEpBriefScoreQueryRequest) InitBizParams(name, productCode, transactionId string) {
	m.ApiParams = &map[string]string{}
	m.FileParams = &map[string]string{}

	(*m.ApiParams)["name"] = name
	m.name = name

	(*m.ApiParams)["product_code"] = productCode
	m.productCode = productCode

	(*m.ApiParams)["transaction_id"] = transactionId
	m.transactionId = transactionId
}

func (*ZhimaCreditEpBriefScoreQueryRequest) GetApiMethodName() string {
	return "zhima.credit.ep.brief.score.query"
}

func (m *ZhimaCreditEpBriefScoreQueryRequest) GetApiParams() *map[string]string {
	return m.ApiParams
}

func (m *ZhimaCreditEpBriefScoreQueryRequest) GetFileParams() *map[string]string {
	return m.FileParams
}

func (m *ZhimaCreditEpBriefScoreQueryRequest) SetApiVersion(ApiVersion string) {
	m.ApiVersion = ApiVersion
}

func (m *ZhimaCreditEpBriefScoreQueryRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaCreditEpBriefScoreQueryRequest) SetScene(Scene string) {
	m.Scene = Scene
}

func (m *ZhimaCreditEpBriefScoreQueryRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaCreditEpBriefScoreQueryRequest) SetChannel(Channel string) {
	m.Channel = Channel
}

func (m *ZhimaCreditEpBriefScoreQueryRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaCreditEpBriefScoreQueryRequest) SetPlatform(Platform string) {
	m.Platform = Platform
}

func (m *ZhimaCreditEpBriefScoreQueryRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaCreditEpBriefScoreQueryRequest) SetExtParams(ExtParams string) {
	m.ExtParams = ExtParams
}

func (m *ZhimaCreditEpBriefScoreQueryRequest) GetExtParams() string {
	return m.ExtParams
}
