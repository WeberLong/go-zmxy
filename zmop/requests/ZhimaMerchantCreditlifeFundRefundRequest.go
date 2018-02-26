package requests

import (
	"github.com/WeberLong/go-zmxy/zmop/interfaces"
)

type ZhimaMerchantCreditlifeFundRefundRequest struct {
	bizProduct string //
	outOrderNo string // 商户发起扣款时的订单号
	payAmount  string // 退款金额
	remark     string // 交易信息说明(退款原因)

	interfaces.ZhimaRequestParams
}

func (m *ZhimaMerchantCreditlifeFundRefundRequest) InitBizParams(bizProduct, outOrderNo, payAmount, remark string) {
	m.ApiParams = &map[string]string{}
	m.FileParams = &map[string]string{}

	(*m.ApiParams)["biz_product"] = bizProduct
	m.bizProduct = bizProduct

	(*m.ApiParams)["out_order_no"] = outOrderNo
	m.outOrderNo = outOrderNo

	(*m.ApiParams)["pay_amount"] = payAmount
	m.payAmount = payAmount

	(*m.ApiParams)["remark"] = remark
	m.remark = remark
}

func (*ZhimaMerchantCreditlifeFundRefundRequest) GetApiMethodName() string {
	return "zhima.merchant.creditlife.fund.refund"
}

func (m *ZhimaMerchantCreditlifeFundRefundRequest) GetApiParams() *map[string]string {
	return m.ApiParams
}

func (m *ZhimaMerchantCreditlifeFundRefundRequest) GetFileParams() *map[string]string {
	return m.FileParams
}

func (m *ZhimaMerchantCreditlifeFundRefundRequest) SetApiVersion(ApiVersion string) {
	m.ApiVersion = ApiVersion
}

func (m *ZhimaMerchantCreditlifeFundRefundRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaMerchantCreditlifeFundRefundRequest) SetScene(Scene string) {
	m.Scene = Scene
}

func (m *ZhimaMerchantCreditlifeFundRefundRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaMerchantCreditlifeFundRefundRequest) SetChannel(Channel string) {
	m.Channel = Channel
}

func (m *ZhimaMerchantCreditlifeFundRefundRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaMerchantCreditlifeFundRefundRequest) SetPlatform(Platform string) {
	m.Platform = Platform
}

func (m *ZhimaMerchantCreditlifeFundRefundRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaMerchantCreditlifeFundRefundRequest) SetExtParams(ExtParams string) {
	m.ExtParams = ExtParams
}

func (m *ZhimaMerchantCreditlifeFundRefundRequest) GetExtParams() string {
	return m.ExtParams
}
