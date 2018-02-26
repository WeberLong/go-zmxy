package requests

import (
	"github.com/WeberLong/go-zmxy/zmop/interfaces"
)

type ZhimaCreditCardLimitGetRequest struct {
	instId        string // 机构ID
	openId        string // 芝麻开放平台OPENID
	orderId       string // 订单号
	productCode   string // 云产品id
	transactionId string // test
	userId        string // 支付宝用户ID

	interfaces.ZhimaRequestParams
}

func (m *ZhimaCreditCardLimitGetRequest) InitBizParams(instId, openId, orderId, productCode, transactionId, userId string) {
	m.ApiParams = &map[string]string{}
	m.FileParams = &map[string]string{}

	(*m.ApiParams)["inst_id"] = instId
	m.instId = instId

	(*m.ApiParams)["open_id"] = openId
	m.openId = openId

	(*m.ApiParams)["order_id"] = orderId
	m.orderId = orderId

	(*m.ApiParams)["product_code"] = productCode
	m.productCode = productCode

	(*m.ApiParams)["transaction_id"] = transactionId
	m.transactionId = transactionId

	(*m.ApiParams)["user_id"] = userId
	m.userId = userId
}

func (*ZhimaCreditCardLimitGetRequest) GetApiMethodName() string {
	return "zhima.credit.card.limit.get"
}

func (m *ZhimaCreditCardLimitGetRequest) GetApiParams() *map[string]string {
	return m.ApiParams
}

func (m *ZhimaCreditCardLimitGetRequest) GetFileParams() *map[string]string {
	return m.FileParams
}

func (m *ZhimaCreditCardLimitGetRequest) SetApiVersion(ApiVersion string) {
	m.ApiVersion = ApiVersion
}

func (m *ZhimaCreditCardLimitGetRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaCreditCardLimitGetRequest) SetScene(Scene string) {
	m.Scene = Scene
}

func (m *ZhimaCreditCardLimitGetRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaCreditCardLimitGetRequest) SetChannel(Channel string) {
	m.Channel = Channel
}

func (m *ZhimaCreditCardLimitGetRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaCreditCardLimitGetRequest) SetPlatform(Platform string) {
	m.Platform = Platform
}

func (m *ZhimaCreditCardLimitGetRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaCreditCardLimitGetRequest) SetExtParams(ExtParams string) {
	m.ExtParams = ExtParams
}

func (m *ZhimaCreditCardLimitGetRequest) GetExtParams() string {
	return m.ExtParams
}
