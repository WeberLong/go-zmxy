package requests

import (
	"github.com/WeberLong/go-zmxy/zmop/interfaces"
)

type ZhimaCreditRiskEvaluateQueryRequest struct {
	certNo        string // 证件号。当证件类型为身份证时，cert_no为身份证号
	certType      string // 证件类型 目前支持两种IDENTITY_CARD(身份证),ALIPAY_USER_ID(支付宝uid)
	extendInfo    string // 扩展参数，供提供更多信息给规则引擎做风险判断。以JSON字符串形式配置
	name          string // 姓名，当传入cert_type类型为IDENTITY_CARD时该值为必传项
	productCode   string // 产品码
	ruleId        string // ISV商户传入二级商户APPID普通商户传入自身APPID
	sceneCode     string // 标识对接业务场景，业务场景下商户可做自定义策略配置
	transactionId string // 芝麻业务凭证，详见https://b.zmxy.com.cn/technology/openDoc.htm?id=334

	interfaces.ZhimaRequestParams
}

func (m *ZhimaCreditRiskEvaluateQueryRequest) InitBizParams(certNo, certType, extendInfo, name, productCode, ruleId, sceneCode, transactionId string) {
	m.ApiParams = &map[string]string{}
	m.FileParams = &map[string]string{}

	(*m.ApiParams)["cert_no"] = certNo
	m.certNo = certNo

	(*m.ApiParams)["cert_type"] = certType
	m.certType = certType

	(*m.ApiParams)["extend_info"] = extendInfo
	m.extendInfo = extendInfo

	(*m.ApiParams)["name"] = name
	m.name = name

	(*m.ApiParams)["product_code"] = productCode
	m.productCode = productCode

	(*m.ApiParams)["rule_id"] = ruleId
	m.ruleId = ruleId

	(*m.ApiParams)["scene_code"] = sceneCode
	m.sceneCode = sceneCode

	(*m.ApiParams)["transaction_id"] = transactionId
	m.transactionId = transactionId
}

func (*ZhimaCreditRiskEvaluateQueryRequest) GetApiMethodName() string {
	return "zhima.credit.risk.evaluate.query"
}

func (m *ZhimaCreditRiskEvaluateQueryRequest) GetApiParams() *map[string]string {
	return m.ApiParams
}

func (m *ZhimaCreditRiskEvaluateQueryRequest) GetFileParams() *map[string]string {
	return m.FileParams
}

func (m *ZhimaCreditRiskEvaluateQueryRequest) SetApiVersion(ApiVersion string) {
	m.ApiVersion = ApiVersion
}

func (m *ZhimaCreditRiskEvaluateQueryRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaCreditRiskEvaluateQueryRequest) SetScene(Scene string) {
	m.Scene = Scene
}

func (m *ZhimaCreditRiskEvaluateQueryRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaCreditRiskEvaluateQueryRequest) SetChannel(Channel string) {
	m.Channel = Channel
}

func (m *ZhimaCreditRiskEvaluateQueryRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaCreditRiskEvaluateQueryRequest) SetPlatform(Platform string) {
	m.Platform = Platform
}

func (m *ZhimaCreditRiskEvaluateQueryRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaCreditRiskEvaluateQueryRequest) SetExtParams(ExtParams string) {
	m.ExtParams = ExtParams
}

func (m *ZhimaCreditRiskEvaluateQueryRequest) GetExtParams() string {
	return m.ExtParams
}
