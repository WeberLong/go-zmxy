package requests

import (
	"github.com/WeberLong/go-zmxy/zmop/interfaces"
)

type ZhimaCreditRiskEvaluateRuleQueryRequest struct {
	productCode   string // 签约产品标示，唯一对应一个产品
	ruleCode      string // 可选参数，传值则标示只查询对应规则配置值，不传则输出所有规则配置值
	ruleId        string // 1000806 【规则标识，使用APPID】如果是ISV商户： 传入二级商户APPID如果是普通商户：传入自己调用APPID
	sceneCode     string // 标识对接业务场景，业务场景下商户可做自定义策略配置
	transactionId string // 唯一标示商户每一笔请求

	interfaces.ZhimaRequestParams
}

func (m *ZhimaCreditRiskEvaluateRuleQueryRequest) InitBizParams(productCode, ruleCode, ruleId, sceneCode, transactionId string) {
	m.ApiParams = &map[string]string{}
	m.FileParams = &map[string]string{}

	(*m.ApiParams)["product_code"] = productCode
	m.productCode = productCode

	(*m.ApiParams)["rule_code"] = ruleCode
	m.ruleCode = ruleCode

	(*m.ApiParams)["rule_id"] = ruleId
	m.ruleId = ruleId

	(*m.ApiParams)["scene_code"] = sceneCode
	m.sceneCode = sceneCode

	(*m.ApiParams)["transaction_id"] = transactionId
	m.transactionId = transactionId
}

func (*ZhimaCreditRiskEvaluateRuleQueryRequest) GetApiMethodName() string {
	return "zhima.credit.risk.evaluate.rule.query"
}

func (m *ZhimaCreditRiskEvaluateRuleQueryRequest) GetApiParams() *map[string]string {
	return m.ApiParams
}

func (m *ZhimaCreditRiskEvaluateRuleQueryRequest) GetFileParams() *map[string]string {
	return m.FileParams
}

func (m *ZhimaCreditRiskEvaluateRuleQueryRequest) SetApiVersion(ApiVersion string) {
	m.ApiVersion = ApiVersion
}

func (m *ZhimaCreditRiskEvaluateRuleQueryRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaCreditRiskEvaluateRuleQueryRequest) SetScene(Scene string) {
	m.Scene = Scene
}

func (m *ZhimaCreditRiskEvaluateRuleQueryRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaCreditRiskEvaluateRuleQueryRequest) SetChannel(Channel string) {
	m.Channel = Channel
}

func (m *ZhimaCreditRiskEvaluateRuleQueryRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaCreditRiskEvaluateRuleQueryRequest) SetPlatform(Platform string) {
	m.Platform = Platform
}

func (m *ZhimaCreditRiskEvaluateRuleQueryRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaCreditRiskEvaluateRuleQueryRequest) SetExtParams(ExtParams string) {
	m.ExtParams = ExtParams
}

func (m *ZhimaCreditRiskEvaluateRuleQueryRequest) GetExtParams() string {
	return m.ExtParams
}
