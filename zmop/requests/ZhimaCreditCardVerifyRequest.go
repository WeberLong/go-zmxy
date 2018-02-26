package requests

import (
	"github.com/WeberLong/go-zmxy/zmop/interfaces"
)

type ZhimaCreditCardVerifyRequest struct {
	address       string // 用户地址.1. 如传入，则进行有效性校验，输出地址有效性等级；2. 如不传，则不进行地址有效性校验，输出地址有效性等级为未知；
	instId        string // 机构ID
	openId        string // 芝麻用户的身份标志，openid
	orderId       string // 订单号
	productCode   string // 云产品ID
	transactionId string // 商户传入的业务流水号。此字段由商户生成，需确保唯一性，用于定位每一次请求，后续按此流水进行对帐。生成规则: 固定30位数字串，前17位为精确到毫秒的时间yyyyMMddhhmmssSSS，后13位为自增数字。
	userId        string // 支付宝用户ID

	interfaces.ZhimaRequestParams
}

func (m *ZhimaCreditCardVerifyRequest) InitBizParams(address, instId, openId, orderId, productCode, transactionId, userId string) {
	m.ApiParams = &map[string]string{}
	m.FileParams = &map[string]string{}

	(*m.ApiParams)["address"] = address
	m.address = address

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

func (*ZhimaCreditCardVerifyRequest) GetApiMethodName() string {
	return "zhima.credit.card.verify"
}

func (m *ZhimaCreditCardVerifyRequest) GetApiParams() *map[string]string {
	return m.ApiParams
}

func (m *ZhimaCreditCardVerifyRequest) GetFileParams() *map[string]string {
	return m.FileParams
}

func (m *ZhimaCreditCardVerifyRequest) SetApiVersion(ApiVersion string) {
	m.ApiVersion = ApiVersion
}

func (m *ZhimaCreditCardVerifyRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaCreditCardVerifyRequest) SetScene(Scene string) {
	m.Scene = Scene
}

func (m *ZhimaCreditCardVerifyRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaCreditCardVerifyRequest) SetChannel(Channel string) {
	m.Channel = Channel
}

func (m *ZhimaCreditCardVerifyRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaCreditCardVerifyRequest) SetPlatform(Platform string) {
	m.Platform = Platform
}

func (m *ZhimaCreditCardVerifyRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaCreditCardVerifyRequest) SetExtParams(ExtParams string) {
	m.ExtParams = ExtParams
}

func (m *ZhimaCreditCardVerifyRequest) GetExtParams() string {
	return m.ExtParams
}
