package requests

import (
	"github.com/WeberLong/go-zmxy/zmop/interfaces"
)

type ZhimaMerchantOrderConfirmRequest struct {
	orderNo       string //
	transactionId string // transaction_id是代表一笔请求的唯一标志，该标识作为对账的关键信息，对于用户使用相同transaction_id的查询，芝麻在一天（86400秒）内返回首次查询数据，超过有效期的查询即为无效并返回异常，有效期内的反复查询不重新计费。transaction_id 推荐生成方式是：30位，（其中17位时间值（精确到毫秒）：yyyyMMddHHmmssSSS）加上（13位自增数字：1234567890123）

	interfaces.ZhimaRequestParams
}

func (m *ZhimaMerchantOrderConfirmRequest) InitBizParams(orderNo, transactionId string) {
	m.ApiParams = &map[string]string{}
	m.FileParams = &map[string]string{}

	(*m.ApiParams)["order_no"] = orderNo
	m.orderNo = orderNo

	(*m.ApiParams)["transaction_id"] = transactionId
	m.transactionId = transactionId
}

func (*ZhimaMerchantOrderConfirmRequest) GetApiMethodName() string {
	return "zhima.merchant.order.confirm"
}

func (m *ZhimaMerchantOrderConfirmRequest) GetApiParams() *map[string]string {
	return m.ApiParams
}

func (m *ZhimaMerchantOrderConfirmRequest) GetFileParams() *map[string]string {
	return m.FileParams
}

func (m *ZhimaMerchantOrderConfirmRequest) SetApiVersion(ApiVersion string) {
	m.ApiVersion = ApiVersion
}

func (m *ZhimaMerchantOrderConfirmRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaMerchantOrderConfirmRequest) SetScene(Scene string) {
	m.Scene = Scene
}

func (m *ZhimaMerchantOrderConfirmRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaMerchantOrderConfirmRequest) SetChannel(Channel string) {
	m.Channel = Channel
}

func (m *ZhimaMerchantOrderConfirmRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaMerchantOrderConfirmRequest) SetPlatform(Platform string) {
	m.Platform = Platform
}

func (m *ZhimaMerchantOrderConfirmRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaMerchantOrderConfirmRequest) SetExtParams(ExtParams string) {
	m.ExtParams = ExtParams
}

func (m *ZhimaMerchantOrderConfirmRequest) GetExtParams() string {
	return m.ExtParams
}
