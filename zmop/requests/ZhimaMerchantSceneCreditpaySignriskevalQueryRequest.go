package requests

import (
	"github.com/WeberLong/go-zmxy/zmop/interfaces"
)

type ZhimaMerchantSceneCreditpaySignriskevalQueryRequest struct {
	linkedAppId      string // 二级商户应用id
	linkedMerchantId string // 二级商户merchantId
	openId           string // openId
	ruleId           string // 规则id
	sceneCode        string // 风险评估场景码
	userId           string // 支付宝uid

	interfaces.ZhimaRequestParams
}

func (m *ZhimaMerchantSceneCreditpaySignriskevalQueryRequest) InitBizParams(linkedAppId, linkedMerchantId, openId, ruleId, sceneCode, userId string) {
	m.ApiParams = &map[string]string{}
	m.FileParams = &map[string]string{}

	(*m.ApiParams)["linked_app_id"] = linkedAppId
	m.linkedAppId = linkedAppId

	(*m.ApiParams)["linked_merchant_id"] = linkedMerchantId
	m.linkedMerchantId = linkedMerchantId

	(*m.ApiParams)["open_id"] = openId
	m.openId = openId

	(*m.ApiParams)["rule_id"] = ruleId
	m.ruleId = ruleId

	(*m.ApiParams)["scene_code"] = sceneCode
	m.sceneCode = sceneCode

	(*m.ApiParams)["user_id"] = userId
	m.userId = userId
}

func (*ZhimaMerchantSceneCreditpaySignriskevalQueryRequest) GetApiMethodName() string {
	return "zhima.merchant.scene.creditpay.signriskeval.query"
}

func (m *ZhimaMerchantSceneCreditpaySignriskevalQueryRequest) GetApiParams() *map[string]string {
	return m.ApiParams
}

func (m *ZhimaMerchantSceneCreditpaySignriskevalQueryRequest) GetFileParams() *map[string]string {
	return m.FileParams
}

func (m *ZhimaMerchantSceneCreditpaySignriskevalQueryRequest) SetApiVersion(ApiVersion string) {
	m.ApiVersion = ApiVersion
}

func (m *ZhimaMerchantSceneCreditpaySignriskevalQueryRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaMerchantSceneCreditpaySignriskevalQueryRequest) SetScene(Scene string) {
	m.Scene = Scene
}

func (m *ZhimaMerchantSceneCreditpaySignriskevalQueryRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaMerchantSceneCreditpaySignriskevalQueryRequest) SetChannel(Channel string) {
	m.Channel = Channel
}

func (m *ZhimaMerchantSceneCreditpaySignriskevalQueryRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaMerchantSceneCreditpaySignriskevalQueryRequest) SetPlatform(Platform string) {
	m.Platform = Platform
}

func (m *ZhimaMerchantSceneCreditpaySignriskevalQueryRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaMerchantSceneCreditpaySignriskevalQueryRequest) SetExtParams(ExtParams string) {
	m.ExtParams = ExtParams
}

func (m *ZhimaMerchantSceneCreditpaySignriskevalQueryRequest) GetExtParams() string {
	return m.ExtParams
}
