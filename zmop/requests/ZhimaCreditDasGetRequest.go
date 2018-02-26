package requests

import (
	"github.com/WeberLong/go-zmxy/zmop/interfaces"
)

type ZhimaCreditDasGetRequest struct {
	contractFlag  string // 合约外标，服务标识。签约过后将会收到含有该服务标识的邮件，或者向协同您签约的芝麻合作人员索取。
	extParas      string // 扩展参数，JSON字符串格式。当前该字段为保留字段，请忽略该参数，或者赋空值: {}
	openId        string // 芝麻会员在商户端的身份标识
	productCode   string // 云产品产品码
	transactionId string // transaction_id是代表一笔请求的唯一标志，该标识作为对账的关键信息，对于用户使用相同transaction_id的查询，芝麻在一天（86400秒）内返回首次查询数据，超过有效期的查询即为无效并返回异常（错误码xxxx），有效期内的重复查询不重新计费。 transaction_id 推荐生成方式是：30位，（其中17位时间值（精确到毫秒）：yyyyMMddHHmmssSSS）加上（13位自增数字：1234567890123）

	interfaces.ZhimaRequestParams
}

func (m *ZhimaCreditDasGetRequest) InitBizParams(contractFlag, extParas, openId, productCode, transactionId string) {
	m.ApiParams = &map[string]string{}
	m.FileParams = &map[string]string{}

	(*m.ApiParams)["contract_flag"] = contractFlag
	m.contractFlag = contractFlag

	(*m.ApiParams)["ext_paras"] = extParas
	m.extParas = extParas

	(*m.ApiParams)["open_id"] = openId
	m.openId = openId

	(*m.ApiParams)["product_code"] = productCode
	m.productCode = productCode

	(*m.ApiParams)["transaction_id"] = transactionId
	m.transactionId = transactionId
}

func (*ZhimaCreditDasGetRequest) GetApiMethodName() string {
	return "zhima.credit.das.get"
}

func (m *ZhimaCreditDasGetRequest) GetApiParams() *map[string]string {
	return m.ApiParams
}

func (m *ZhimaCreditDasGetRequest) GetFileParams() *map[string]string {
	return m.FileParams
}

func (m *ZhimaCreditDasGetRequest) SetApiVersion(ApiVersion string) {
	m.ApiVersion = ApiVersion
}

func (m *ZhimaCreditDasGetRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaCreditDasGetRequest) SetScene(Scene string) {
	m.Scene = Scene
}

func (m *ZhimaCreditDasGetRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaCreditDasGetRequest) SetChannel(Channel string) {
	m.Channel = Channel
}

func (m *ZhimaCreditDasGetRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaCreditDasGetRequest) SetPlatform(Platform string) {
	m.Platform = Platform
}

func (m *ZhimaCreditDasGetRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaCreditDasGetRequest) SetExtParams(ExtParams string) {
	m.ExtParams = ExtParams
}

func (m *ZhimaCreditDasGetRequest) GetExtParams() string {
	return m.ExtParams
}
