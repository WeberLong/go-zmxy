package requests

import (
	"github.com/WeberLong/go-zmxy/zmop/interfaces"
)

type ZhimaCreditEpWatchlistGetRequest struct {
	dataType      string // 企业关注名单的输入类型：工商注册号（REG_NO）、组织机构代码（ORG_NO）、社会统一代码(SOCIETY_NO)、税务登记号(TAX_NO)、企业名称（COMPANY_NAME）
	dataValue     string // 企业关注名单，对应类型的数据值
	productCode   string // 产品码
	transactionId string // transaction_id是代表一笔请求的唯一标志，该标识作为对账的关键信息，对于用户使用相同transaction_id的查询，芝麻在一天（86400秒）内返回首次查询数据，超过有效期的查询即为无效并返回异常（错误码xxxx），有效期内的重复查询不重新计费。 transaction_id 推荐生成方式是：30位，（其中17位时间值（精确到毫秒）：yyyyMMddHHmmssSSS）加上（13位自增数字：1234567890123）

	interfaces.ZhimaRequestParams
}

func (m *ZhimaCreditEpWatchlistGetRequest) InitBizParams(dataType, dataValue, productCode, transactionId string) {
	m.ApiParams = &map[string]string{}
	m.FileParams = &map[string]string{}

	(*m.ApiParams)["data_type"] = dataType
	m.dataType = dataType

	(*m.ApiParams)["data_value"] = dataValue
	m.dataValue = dataValue

	(*m.ApiParams)["product_code"] = productCode
	m.productCode = productCode

	(*m.ApiParams)["transaction_id"] = transactionId
	m.transactionId = transactionId
}

func (*ZhimaCreditEpWatchlistGetRequest) GetApiMethodName() string {
	return "zhima.credit.ep.watchlist.get"
}

func (m *ZhimaCreditEpWatchlistGetRequest) GetApiParams() *map[string]string {
	return m.ApiParams
}

func (m *ZhimaCreditEpWatchlistGetRequest) GetFileParams() *map[string]string {
	return m.FileParams
}

func (m *ZhimaCreditEpWatchlistGetRequest) SetApiVersion(ApiVersion string) {
	m.ApiVersion = ApiVersion
}

func (m *ZhimaCreditEpWatchlistGetRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaCreditEpWatchlistGetRequest) SetScene(Scene string) {
	m.Scene = Scene
}

func (m *ZhimaCreditEpWatchlistGetRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaCreditEpWatchlistGetRequest) SetChannel(Channel string) {
	m.Channel = Channel
}

func (m *ZhimaCreditEpWatchlistGetRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaCreditEpWatchlistGetRequest) SetPlatform(Platform string) {
	m.Platform = Platform
}

func (m *ZhimaCreditEpWatchlistGetRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaCreditEpWatchlistGetRequest) SetExtParams(ExtParams string) {
	m.ExtParams = ExtParams
}

func (m *ZhimaCreditEpWatchlistGetRequest) GetExtParams() string {
	return m.ExtParams
}
