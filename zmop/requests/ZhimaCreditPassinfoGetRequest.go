package requests

import (
	"github.com/WeberLong/go-zmxy/zmop/interfaces"
)

type ZhimaCreditPassinfoGetRequest struct {
	contractFlag  string // 合约外标，服务标识。签约过后将会收到含有该服务标识的邮件，或者向协同您签约的芝麻合作人员索取。
	openId        string // 用户授权芝麻后的身份标识
	productCode   string // 云产品码
	transactionId string // 业务流水号，标示每一次调用

	interfaces.ZhimaRequestParams
}

func (m *ZhimaCreditPassinfoGetRequest) InitBizParams(contractFlag, openId, productCode, transactionId string) {
	m.ApiParams = &map[string]string{}
	m.FileParams = &map[string]string{}

	(*m.ApiParams)["contract_flag"] = contractFlag
	m.contractFlag = contractFlag

	(*m.ApiParams)["open_id"] = openId
	m.openId = openId

	(*m.ApiParams)["product_code"] = productCode
	m.productCode = productCode

	(*m.ApiParams)["transaction_id"] = transactionId
	m.transactionId = transactionId
}

func (*ZhimaCreditPassinfoGetRequest) GetApiMethodName() string {
	return "zhima.credit.passinfo.get"
}

func (m *ZhimaCreditPassinfoGetRequest) GetApiParams() *map[string]string {
	return m.ApiParams
}

func (m *ZhimaCreditPassinfoGetRequest) GetFileParams() *map[string]string {
	return m.FileParams
}

func (m *ZhimaCreditPassinfoGetRequest) SetApiVersion(ApiVersion string) {
	m.ApiVersion = ApiVersion
}

func (m *ZhimaCreditPassinfoGetRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaCreditPassinfoGetRequest) SetScene(Scene string) {
	m.Scene = Scene
}

func (m *ZhimaCreditPassinfoGetRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaCreditPassinfoGetRequest) SetChannel(Channel string) {
	m.Channel = Channel
}

func (m *ZhimaCreditPassinfoGetRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaCreditPassinfoGetRequest) SetPlatform(Platform string) {
	m.Platform = Platform
}

func (m *ZhimaCreditPassinfoGetRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaCreditPassinfoGetRequest) SetExtParams(ExtParams string) {
	m.ExtParams = ExtParams
}

func (m *ZhimaCreditPassinfoGetRequest) GetExtParams() string {
	return m.ExtParams
}
