package requests

import (
	"github.com/WeberLong/go-zmxy/zmop/interfaces"
)

type ZhimaMerchantCreditlifePreauthCancelRequest struct {
	outOrderNo string // 待解冻预授权冻结资金订单号，或解冻请求流水号
	preAuthNo  string // 预授权号
	remark     string // 取消预授权冻结资金原因

	interfaces.ZhimaRequestParams
}

func (m *ZhimaMerchantCreditlifePreauthCancelRequest) InitBizParams(outOrderNo, preAuthNo, remark string) {
	m.ApiParams = &map[string]string{}
	m.FileParams = &map[string]string{}

	(*m.ApiParams)["out_order_no"] = outOrderNo
	m.outOrderNo = outOrderNo

	(*m.ApiParams)["pre_auth_no"] = preAuthNo
	m.preAuthNo = preAuthNo

	(*m.ApiParams)["remark"] = remark
	m.remark = remark
}

func (*ZhimaMerchantCreditlifePreauthCancelRequest) GetApiMethodName() string {
	return "zhima.merchant.creditlife.preauth.cancel"
}

func (m *ZhimaMerchantCreditlifePreauthCancelRequest) GetApiParams() *map[string]string {
	return m.ApiParams
}

func (m *ZhimaMerchantCreditlifePreauthCancelRequest) GetFileParams() *map[string]string {
	return m.FileParams
}

func (m *ZhimaMerchantCreditlifePreauthCancelRequest) SetApiVersion(ApiVersion string) {
	m.ApiVersion = ApiVersion
}

func (m *ZhimaMerchantCreditlifePreauthCancelRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaMerchantCreditlifePreauthCancelRequest) SetScene(Scene string) {
	m.Scene = Scene
}

func (m *ZhimaMerchantCreditlifePreauthCancelRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaMerchantCreditlifePreauthCancelRequest) SetChannel(Channel string) {
	m.Channel = Channel
}

func (m *ZhimaMerchantCreditlifePreauthCancelRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaMerchantCreditlifePreauthCancelRequest) SetPlatform(Platform string) {
	m.Platform = Platform
}

func (m *ZhimaMerchantCreditlifePreauthCancelRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaMerchantCreditlifePreauthCancelRequest) SetExtParams(ExtParams string) {
	m.ExtParams = ExtParams
}

func (m *ZhimaMerchantCreditlifePreauthCancelRequest) GetExtParams() string {
	return m.ExtParams
}
