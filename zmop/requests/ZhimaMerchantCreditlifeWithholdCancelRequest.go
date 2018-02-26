package requests

import (
	"github.com/WeberLong/go-zmxy/zmop/interfaces"
)

type ZhimaMerchantCreditlifeWithholdCancelRequest struct {
	userId string // 用户支付宝id

	interfaces.ZhimaRequestParams
}

func (m *ZhimaMerchantCreditlifeWithholdCancelRequest) InitBizParams(userId string) {
	m.ApiParams = &map[string]string{}
	m.FileParams = &map[string]string{}

	(*m.ApiParams)["user_id"] = userId
	m.userId = userId
}

func (*ZhimaMerchantCreditlifeWithholdCancelRequest) GetApiMethodName() string {
	return "zhima.merchant.creditlife.withhold.cancel"
}

func (m *ZhimaMerchantCreditlifeWithholdCancelRequest) GetApiParams() *map[string]string {
	return m.ApiParams
}

func (m *ZhimaMerchantCreditlifeWithholdCancelRequest) GetFileParams() *map[string]string {
	return m.FileParams
}

func (m *ZhimaMerchantCreditlifeWithholdCancelRequest) SetApiVersion(ApiVersion string) {
	m.ApiVersion = ApiVersion
}

func (m *ZhimaMerchantCreditlifeWithholdCancelRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaMerchantCreditlifeWithholdCancelRequest) SetScene(Scene string) {
	m.Scene = Scene
}

func (m *ZhimaMerchantCreditlifeWithholdCancelRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaMerchantCreditlifeWithholdCancelRequest) SetChannel(Channel string) {
	m.Channel = Channel
}

func (m *ZhimaMerchantCreditlifeWithholdCancelRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaMerchantCreditlifeWithholdCancelRequest) SetPlatform(Platform string) {
	m.Platform = Platform
}

func (m *ZhimaMerchantCreditlifeWithholdCancelRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaMerchantCreditlifeWithholdCancelRequest) SetExtParams(ExtParams string) {
	m.ExtParams = ExtParams
}

func (m *ZhimaMerchantCreditlifeWithholdCancelRequest) GetExtParams() string {
	return m.ExtParams
}
