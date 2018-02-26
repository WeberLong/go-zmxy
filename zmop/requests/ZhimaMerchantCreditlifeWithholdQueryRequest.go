package requests

import (
	"github.com/WeberLong/go-zmxy/zmop/interfaces"
)

type ZhimaMerchantCreditlifeWithholdQueryRequest struct {
	userId string // 支付宝用户id

	interfaces.ZhimaRequestParams
}

func (m *ZhimaMerchantCreditlifeWithholdQueryRequest) InitBizParams(userId string) {
	m.ApiParams = &map[string]string{}
	m.FileParams = &map[string]string{}

	(*m.ApiParams)["user_id"] = userId
	m.userId = userId
}

func (*ZhimaMerchantCreditlifeWithholdQueryRequest) GetApiMethodName() string {
	return "zhima.merchant.creditlife.withhold.query"
}

func (m *ZhimaMerchantCreditlifeWithholdQueryRequest) GetApiParams() *map[string]string {
	return m.ApiParams
}

func (m *ZhimaMerchantCreditlifeWithholdQueryRequest) GetFileParams() *map[string]string {
	return m.FileParams
}

func (m *ZhimaMerchantCreditlifeWithholdQueryRequest) SetApiVersion(ApiVersion string) {
	m.ApiVersion = ApiVersion
}

func (m *ZhimaMerchantCreditlifeWithholdQueryRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaMerchantCreditlifeWithholdQueryRequest) SetScene(Scene string) {
	m.Scene = Scene
}

func (m *ZhimaMerchantCreditlifeWithholdQueryRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaMerchantCreditlifeWithholdQueryRequest) SetChannel(Channel string) {
	m.Channel = Channel
}

func (m *ZhimaMerchantCreditlifeWithholdQueryRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaMerchantCreditlifeWithholdQueryRequest) SetPlatform(Platform string) {
	m.Platform = Platform
}

func (m *ZhimaMerchantCreditlifeWithholdQueryRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaMerchantCreditlifeWithholdQueryRequest) SetExtParams(ExtParams string) {
	m.ExtParams = ExtParams
}

func (m *ZhimaMerchantCreditlifeWithholdQueryRequest) GetExtParams() string {
	return m.ExtParams
}
