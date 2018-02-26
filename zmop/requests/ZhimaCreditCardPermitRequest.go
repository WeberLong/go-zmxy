package requests

import (
	"github.com/WeberLong/go-zmxy/zmop/interfaces"
)

type ZhimaCreditCardPermitRequest struct {
	instId        string // 机构ID
	orderId       string // 订单编号
	productCode   string // 云产品id
	transactionId string // 商户传入的业务流水号。此字段由商户生成，需确保唯一性，用于定位每一次请求，后续按此流水进行对帐。生成规则: 固定30位数字串，前17位为精确到毫秒的时间yyyyMMddhhmmssSSS，后13位为自增数字。
	userId        string // 支付宝用户ID

	interfaces.ZhimaRequestParams
}

func (m *ZhimaCreditCardPermitRequest) InitBizParams(instId, orderId, productCode, transactionId, userId string) {
	m.ApiParams = &map[string]string{}
	m.FileParams = &map[string]string{}

	(*m.ApiParams)["inst_id"] = instId
	m.instId = instId

	(*m.ApiParams)["order_id"] = orderId
	m.orderId = orderId

	(*m.ApiParams)["product_code"] = productCode
	m.productCode = productCode

	(*m.ApiParams)["transaction_id"] = transactionId
	m.transactionId = transactionId

	(*m.ApiParams)["user_id"] = userId
	m.userId = userId
}

func (*ZhimaCreditCardPermitRequest) GetApiMethodName() string {
	return "zhima.credit.card.permit"
}

func (m *ZhimaCreditCardPermitRequest) GetApiParams() *map[string]string {
	return m.ApiParams
}

func (m *ZhimaCreditCardPermitRequest) GetFileParams() *map[string]string {
	return m.FileParams
}

func (m *ZhimaCreditCardPermitRequest) SetApiVersion(ApiVersion string) {
	m.ApiVersion = ApiVersion
}

func (m *ZhimaCreditCardPermitRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaCreditCardPermitRequest) SetScene(Scene string) {
	m.Scene = Scene
}

func (m *ZhimaCreditCardPermitRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaCreditCardPermitRequest) SetChannel(Channel string) {
	m.Channel = Channel
}

func (m *ZhimaCreditCardPermitRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaCreditCardPermitRequest) SetPlatform(Platform string) {
	m.Platform = Platform
}

func (m *ZhimaCreditCardPermitRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaCreditCardPermitRequest) SetExtParams(ExtParams string) {
	m.ExtParams = ExtParams
}

func (m *ZhimaCreditCardPermitRequest) GetExtParams() string {
	return m.ExtParams
}
