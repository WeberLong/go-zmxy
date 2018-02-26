package requests

import (
	"github.com/WeberLong/go-zmxy/zmop/interfaces"
)

type ZhimaMerchantSceneCreditpayRiskevalQueryRequest struct {
	needAuth string // 是否需要授权
	userId   string // 支付宝userid

	interfaces.ZhimaRequestParams
}

func (m *ZhimaMerchantSceneCreditpayRiskevalQueryRequest) InitBizParams(needAuth, userId string) {
	m.ApiParams = &map[string]string{}
	m.FileParams = &map[string]string{}

	(*m.ApiParams)["need_auth"] = needAuth
	m.needAuth = needAuth

	(*m.ApiParams)["user_id"] = userId
	m.userId = userId
}

func (*ZhimaMerchantSceneCreditpayRiskevalQueryRequest) GetApiMethodName() string {
	return "zhima.merchant.scene.creditpay.riskeval.query"
}

func (m *ZhimaMerchantSceneCreditpayRiskevalQueryRequest) GetApiParams() *map[string]string {
	return m.ApiParams
}

func (m *ZhimaMerchantSceneCreditpayRiskevalQueryRequest) GetFileParams() *map[string]string {
	return m.FileParams
}

func (m *ZhimaMerchantSceneCreditpayRiskevalQueryRequest) SetApiVersion(ApiVersion string) {
	m.ApiVersion = ApiVersion
}

func (m *ZhimaMerchantSceneCreditpayRiskevalQueryRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaMerchantSceneCreditpayRiskevalQueryRequest) SetScene(Scene string) {
	m.Scene = Scene
}

func (m *ZhimaMerchantSceneCreditpayRiskevalQueryRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaMerchantSceneCreditpayRiskevalQueryRequest) SetChannel(Channel string) {
	m.Channel = Channel
}

func (m *ZhimaMerchantSceneCreditpayRiskevalQueryRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaMerchantSceneCreditpayRiskevalQueryRequest) SetPlatform(Platform string) {
	m.Platform = Platform
}

func (m *ZhimaMerchantSceneCreditpayRiskevalQueryRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaMerchantSceneCreditpayRiskevalQueryRequest) SetExtParams(ExtParams string) {
	m.ExtParams = ExtParams
}

func (m *ZhimaMerchantSceneCreditpayRiskevalQueryRequest) GetExtParams() string {
	return m.ExtParams
}
