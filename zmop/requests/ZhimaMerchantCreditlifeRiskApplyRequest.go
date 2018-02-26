package requests

import (
	"github.com/WeberLong/go-zmxy/zmop/interfaces"
)

type ZhimaMerchantCreditlifeRiskApplyRequest struct {
	address       string //
	certNo        string //
	itemId        string //
	mobile        string //
	name          string //
	transactionId string //

	interfaces.ZhimaRequestParams
}

func (m *ZhimaMerchantCreditlifeRiskApplyRequest) InitBizParams(address, certNo, itemId, mobile, name, transactionId string) {
	m.ApiParams = &map[string]string{}
	m.FileParams = &map[string]string{}

	(*m.ApiParams)["address"] = address
	m.address = address

	(*m.ApiParams)["cert_no"] = certNo
	m.certNo = certNo

	(*m.ApiParams)["item_id"] = itemId
	m.itemId = itemId

	(*m.ApiParams)["mobile"] = mobile
	m.mobile = mobile

	(*m.ApiParams)["name"] = name
	m.name = name

	(*m.ApiParams)["transaction_id"] = transactionId
	m.transactionId = transactionId
}

func (*ZhimaMerchantCreditlifeRiskApplyRequest) GetApiMethodName() string {
	return "zhima.merchant.creditlife.risk.apply"
}

func (m *ZhimaMerchantCreditlifeRiskApplyRequest) GetApiParams() *map[string]string {
	return m.ApiParams
}

func (m *ZhimaMerchantCreditlifeRiskApplyRequest) GetFileParams() *map[string]string {
	return m.FileParams
}

func (m *ZhimaMerchantCreditlifeRiskApplyRequest) SetApiVersion(ApiVersion string) {
	m.ApiVersion = ApiVersion
}

func (m *ZhimaMerchantCreditlifeRiskApplyRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaMerchantCreditlifeRiskApplyRequest) SetScene(Scene string) {
	m.Scene = Scene
}

func (m *ZhimaMerchantCreditlifeRiskApplyRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaMerchantCreditlifeRiskApplyRequest) SetChannel(Channel string) {
	m.Channel = Channel
}

func (m *ZhimaMerchantCreditlifeRiskApplyRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaMerchantCreditlifeRiskApplyRequest) SetPlatform(Platform string) {
	m.Platform = Platform
}

func (m *ZhimaMerchantCreditlifeRiskApplyRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaMerchantCreditlifeRiskApplyRequest) SetExtParams(ExtParams string) {
	m.ExtParams = ExtParams
}

func (m *ZhimaMerchantCreditlifeRiskApplyRequest) GetExtParams() string {
	return m.ExtParams
}
