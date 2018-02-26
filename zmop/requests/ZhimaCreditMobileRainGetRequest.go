package requests

import (
	"github.com/WeberLong/go-zmxy/zmop/interfaces"
)

type ZhimaCreditMobileRainGetRequest struct {
	mobile        string // 手机号，11位长度的国内手机号
	productCode   string // 产品码
	transactionId string // transaction_id是代表一笔请求的唯一标志，该标识作为对账的关键信息，对于用户使用相同transaction_id的查询，芝麻在一天（86400秒）内返回首次查询数据，超过有效期的查询即为无效并返回异常（错误码xxxx），有效期内的重复查询不重新计费。 transaction_id 推荐生成方式是：30位，（其中17位时间值（精确到毫秒）：yyyyMMddHHmmssSSS）加上（13位自增数字：1234567890123）

	interfaces.ZhimaRequestParams
}

func (m *ZhimaCreditMobileRainGetRequest) InitBizParams(mobile, productCode, transactionId string) {
	m.ApiParams = &map[string]string{}
	m.FileParams = &map[string]string{}

	(*m.ApiParams)["mobile"] = mobile
	m.mobile = mobile

	(*m.ApiParams)["product_code"] = productCode
	m.productCode = productCode

	(*m.ApiParams)["transaction_id"] = transactionId
	m.transactionId = transactionId
}

func (*ZhimaCreditMobileRainGetRequest) GetApiMethodName() string {
	return "zhima.credit.mobile.rain.get"
}

func (m *ZhimaCreditMobileRainGetRequest) GetApiParams() *map[string]string {
	return m.ApiParams
}

func (m *ZhimaCreditMobileRainGetRequest) GetFileParams() *map[string]string {
	return m.FileParams
}

func (m *ZhimaCreditMobileRainGetRequest) SetApiVersion(ApiVersion string) {
	m.ApiVersion = ApiVersion
}

func (m *ZhimaCreditMobileRainGetRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaCreditMobileRainGetRequest) SetScene(Scene string) {
	m.Scene = Scene
}

func (m *ZhimaCreditMobileRainGetRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaCreditMobileRainGetRequest) SetChannel(Channel string) {
	m.Channel = Channel
}

func (m *ZhimaCreditMobileRainGetRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaCreditMobileRainGetRequest) SetPlatform(Platform string) {
	m.Platform = Platform
}

func (m *ZhimaCreditMobileRainGetRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaCreditMobileRainGetRequest) SetExtParams(ExtParams string) {
	m.ExtParams = ExtParams
}

func (m *ZhimaCreditMobileRainGetRequest) GetExtParams() string {
	return m.ExtParams
}
