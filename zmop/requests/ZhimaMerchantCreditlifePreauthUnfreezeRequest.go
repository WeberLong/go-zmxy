package requests

import (
	"github.com/WeberLong/go-zmxy/zmop/interfaces"
)

type ZhimaMerchantCreditlifePreauthUnfreezeRequest struct {
	payAmount     string // 待解冻资金(元)
	preAuthNo     string // 预授权后产生的预授权号
	remark        string // 发起资金解冻原因
	transactionId string // 交易流水号

	interfaces.ZhimaRequestParams
}

func (m *ZhimaMerchantCreditlifePreauthUnfreezeRequest) InitBizParams(payAmount, preAuthNo, remark, transactionId string) {
	m.ApiParams = &map[string]string{}
	m.FileParams = &map[string]string{}

	(*m.ApiParams)["pay_amount"] = payAmount
	m.payAmount = payAmount

	(*m.ApiParams)["pre_auth_no"] = preAuthNo
	m.preAuthNo = preAuthNo

	(*m.ApiParams)["remark"] = remark
	m.remark = remark

	(*m.ApiParams)["transaction_id"] = transactionId
	m.transactionId = transactionId
}

func (*ZhimaMerchantCreditlifePreauthUnfreezeRequest) GetApiMethodName() string {
	return "zhima.merchant.creditlife.preauth.unfreeze"
}

func (m *ZhimaMerchantCreditlifePreauthUnfreezeRequest) GetApiParams() *map[string]string {
	return m.ApiParams
}

func (m *ZhimaMerchantCreditlifePreauthUnfreezeRequest) GetFileParams() *map[string]string {
	return m.FileParams
}

func (m *ZhimaMerchantCreditlifePreauthUnfreezeRequest) SetApiVersion(ApiVersion string) {
	m.ApiVersion = ApiVersion
}

func (m *ZhimaMerchantCreditlifePreauthUnfreezeRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaMerchantCreditlifePreauthUnfreezeRequest) SetScene(Scene string) {
	m.Scene = Scene
}

func (m *ZhimaMerchantCreditlifePreauthUnfreezeRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaMerchantCreditlifePreauthUnfreezeRequest) SetChannel(Channel string) {
	m.Channel = Channel
}

func (m *ZhimaMerchantCreditlifePreauthUnfreezeRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaMerchantCreditlifePreauthUnfreezeRequest) SetPlatform(Platform string) {
	m.Platform = Platform
}

func (m *ZhimaMerchantCreditlifePreauthUnfreezeRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaMerchantCreditlifePreauthUnfreezeRequest) SetExtParams(ExtParams string) {
	m.ExtParams = ExtParams
}

func (m *ZhimaMerchantCreditlifePreauthUnfreezeRequest) GetExtParams() string {
	return m.ExtParams
}
