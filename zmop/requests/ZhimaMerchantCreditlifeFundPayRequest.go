package requests

import (
	"github.com/WeberLong/go-zmxy/zmop/interfaces"
)

type ZhimaMerchantCreditlifeFundPayRequest struct {
	agreementNo   string // 代扣协议号(代扣扣款时必须提供)
	fundPayType   string // 扣款类型(withholding_pay:代扣扣款,preauth_pay:预授权转支付)
	goodsTitle    string //
	goodsType     string // 商品类型(0:虚拟物品,1:实物)
	outOrderNo    string // 商户订单号
	payAmount     string // 支付金额
	preAuthNo     string // 预授权号(付款方式为预授权转支付时必须提供)
	roleId        string // 芝麻用户id
	sellerId      string // 收款方支付宝id
	transactionId string //
	userId        string // 支付宝用户id（付款方id）

	interfaces.ZhimaRequestParams
}

func (m *ZhimaMerchantCreditlifeFundPayRequest) InitBizParams(agreementNo, fundPayType, goodsTitle, goodsType, outOrderNo, payAmount, preAuthNo, roleId, sellerId, transactionId, userId string) {
	m.ApiParams = &map[string]string{}
	m.FileParams = &map[string]string{}

	(*m.ApiParams)["agreement_no"] = agreementNo
	m.agreementNo = agreementNo

	(*m.ApiParams)["fund_pay_type"] = fundPayType
	m.fundPayType = fundPayType

	(*m.ApiParams)["goods_title"] = goodsTitle
	m.goodsTitle = goodsTitle

	(*m.ApiParams)["goods_type"] = goodsType
	m.goodsType = goodsType

	(*m.ApiParams)["out_order_no"] = outOrderNo
	m.outOrderNo = outOrderNo

	(*m.ApiParams)["pay_amount"] = payAmount
	m.payAmount = payAmount

	(*m.ApiParams)["pre_auth_no"] = preAuthNo
	m.preAuthNo = preAuthNo

	(*m.ApiParams)["role_id"] = roleId
	m.roleId = roleId

	(*m.ApiParams)["seller_id"] = sellerId
	m.sellerId = sellerId

	(*m.ApiParams)["transaction_id"] = transactionId
	m.transactionId = transactionId

	(*m.ApiParams)["user_id"] = userId
	m.userId = userId
}

func (*ZhimaMerchantCreditlifeFundPayRequest) GetApiMethodName() string {
	return "zhima.merchant.creditlife.fund.pay"
}

func (m *ZhimaMerchantCreditlifeFundPayRequest) GetApiParams() *map[string]string {
	return m.ApiParams
}

func (m *ZhimaMerchantCreditlifeFundPayRequest) GetFileParams() *map[string]string {
	return m.FileParams
}

func (m *ZhimaMerchantCreditlifeFundPayRequest) SetApiVersion(ApiVersion string) {
	m.ApiVersion = ApiVersion
}

func (m *ZhimaMerchantCreditlifeFundPayRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaMerchantCreditlifeFundPayRequest) SetScene(Scene string) {
	m.Scene = Scene
}

func (m *ZhimaMerchantCreditlifeFundPayRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaMerchantCreditlifeFundPayRequest) SetChannel(Channel string) {
	m.Channel = Channel
}

func (m *ZhimaMerchantCreditlifeFundPayRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaMerchantCreditlifeFundPayRequest) SetPlatform(Platform string) {
	m.Platform = Platform
}

func (m *ZhimaMerchantCreditlifeFundPayRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaMerchantCreditlifeFundPayRequest) SetExtParams(ExtParams string) {
	m.ExtParams = ExtParams
}

func (m *ZhimaMerchantCreditlifeFundPayRequest) GetExtParams() string {
	return m.ExtParams
}
