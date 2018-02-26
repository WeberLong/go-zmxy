package requests

import (
	"github.com/WeberLong/go-zmxy/zmop/interfaces"
)

type ZhimaMerchantSceneCreditpayBizriskevalQueryRequest struct {
	linkedAppId      string // 二级商户应用id
	linkedMerchantId string // 二级商户merchantId
	openId           string // openId
	ruleId           string // 规则id
	scenceCode       string // 风险评估场景码
	userId           string // 支付宝uid

	interfaces.ZhimaRequestParams
}

func (m *ZhimaMerchantSceneCreditpayBizriskevalQueryRequest) InitBizParams(linkedAppId, linkedMerchantId, openId, ruleId, scenceCode, userId string) {
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

	(*m.ApiParams)["scence_code"] = scenceCode
	m.scenceCode = scenceCode

	(*m.ApiParams)["user_id"] = userId
	m.userId = userId
}

func (*ZhimaMerchantSceneCreditpayBizriskevalQueryRequest) GetApiMethodName() string {
	return "zhima.merchant.scene.creditpay.bizriskeval.query"
}

func (m *ZhimaMerchantSceneCreditpayBizriskevalQueryRequest) GetApiParams() *map[string]string {
	return m.ApiParams
}

func (m *ZhimaMerchantSceneCreditpayBizriskevalQueryRequest) GetFileParams() *map[string]string {
	return m.FileParams
}

func (m *ZhimaMerchantSceneCreditpayBizriskevalQueryRequest) SetApiVersion(ApiVersion string) {
	m.ApiVersion = ApiVersion
}

func (m *ZhimaMerchantSceneCreditpayBizriskevalQueryRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaMerchantSceneCreditpayBizriskevalQueryRequest) SetScene(Scene string) {
	m.Scene = Scene
}

func (m *ZhimaMerchantSceneCreditpayBizriskevalQueryRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaMerchantSceneCreditpayBizriskevalQueryRequest) SetChannel(Channel string) {
	m.Channel = Channel
}

func (m *ZhimaMerchantSceneCreditpayBizriskevalQueryRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaMerchantSceneCreditpayBizriskevalQueryRequest) SetPlatform(Platform string) {
	m.Platform = Platform
}

func (m *ZhimaMerchantSceneCreditpayBizriskevalQueryRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaMerchantSceneCreditpayBizriskevalQueryRequest) SetExtParams(ExtParams string) {
	m.ExtParams = ExtParams
}

func (m *ZhimaMerchantSceneCreditpayBizriskevalQueryRequest) GetExtParams() string {
	return m.ExtParams
}
