package requests

import (
	"github.com/WeberLong/go-zmxy/zmop/interfaces"
)

type ZhimaCustomerCertifyInitialRequest struct {
	bizParams     string // 业务扩展参数入参，传递方式例如{"xx":"xxxxx"};针对KBA的认证方式需要关注，biz_params中需要传入入参：{"verifyScene":"向芝麻申请获得的scene值"}
	contractFlag  string // 与芝麻信用签订的合约外标，即使合约改签或续签该值不会发生变化。请联系技术支持
	identityParam string // 不同身份类型的参数列表，json字符串的key-value格式：如：当identity_type= "BY_CERTNO_AND_NAME";时identity_param={"certNo":"xxx","name":"张三","certType":"IDENTITY_CARD","mobileNo":"13901234567"};或者当identity_type= "BY_MOBILE_NO";时identity_param={"mobileNo":"13901234567"};或者当identify_type="BY_CERT_IMAGE"identity_param={"frontCertImage":"oioiweroeworewoiho2323","backCertImage":"dsrrwerewew"}
	identityType  string // 身份标识类型（后续可以扩展）：BY_CERTNO_AND_NAME:按照身份证+姓名（手机号可选）进行身份识别BY_MOBILE_NO:按照手机号进行身份识别BY_CERT_IMAGE: 根据证件图片识别
	pageUrl       string // 商户页面回调地址，芝麻认证完成后通过此url地址回传给商户认证的结果；SDK模式接入的场景为非必填项，其他渠道类型的必填项；
	productCode   string // 当前使用的产品码
	schemaUrl     string // 商户App的回调地址，通过商户App发起人脸核身的芝麻认证时必传；其他场景为非必填项；
	sourceType    string // 请求来源类型，为比填项， 例如h5, pc , app, sdk,window；1.h5 ：商户H5端接入芝麻应用的场景；2.pc：商户pc端接入芝麻认证的场景；3.app：商户app应用接入芝麻认证的场景；4.sdk：商户调用芝麻的sdk进行芝麻认证的场景:5.window：服务窗进行芝麻认证的场景；
	state         string // 芝麻认证过程中的冗余字段，在认证申请时传入，在结果页面回调中原样透传给商户端。支持json格式。【建议使用方式】用于商户端唯一标记发起认证的用户信息，在接收到芝麻信用认证结果回调后确定用户
	transactionId string // 商户传入的业务流水号。此字段由商户生成，需确保唯一性，用于定位每一次请求，后续按此流水进行对帐。生成规则: 固定30位数字串，前17位为精确到毫秒的时间yyyyMMddhhmmssSSS，后13位为自增数字。

	interfaces.ZhimaRequestParams
}

func (m *ZhimaCustomerCertifyInitialRequest) InitBizParams(bizParams, contractFlag, identityParam, identityType, pageUrl, productCode, schemaUrl, sourceType, state, transactionId string) {
	m.ApiParams = &map[string]string{}
	m.FileParams = &map[string]string{}

	(*m.ApiParams)["biz_params"] = bizParams
	m.bizParams = bizParams

	(*m.ApiParams)["contract_flag"] = contractFlag
	m.contractFlag = contractFlag

	(*m.ApiParams)["identity_param"] = identityParam
	m.identityParam = identityParam

	(*m.ApiParams)["identity_type"] = identityType
	m.identityType = identityType

	(*m.ApiParams)["page_url"] = pageUrl
	m.pageUrl = pageUrl

	(*m.ApiParams)["product_code"] = productCode
	m.productCode = productCode

	(*m.ApiParams)["schema_url"] = schemaUrl
	m.schemaUrl = schemaUrl

	(*m.ApiParams)["source_type"] = sourceType
	m.sourceType = sourceType

	(*m.ApiParams)["state"] = state
	m.state = state

	(*m.ApiParams)["transaction_id"] = transactionId
	m.transactionId = transactionId
}

func (*ZhimaCustomerCertifyInitialRequest) GetApiMethodName() string {
	return "zhima.customer.certify.initial"
}

func (m *ZhimaCustomerCertifyInitialRequest) GetApiParams() *map[string]string {
	return m.ApiParams
}

func (m *ZhimaCustomerCertifyInitialRequest) GetFileParams() *map[string]string {
	return m.FileParams
}

func (m *ZhimaCustomerCertifyInitialRequest) SetApiVersion(ApiVersion string) {
	m.ApiVersion = ApiVersion
}

func (m *ZhimaCustomerCertifyInitialRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaCustomerCertifyInitialRequest) SetScene(Scene string) {
	m.Scene = Scene
}

func (m *ZhimaCustomerCertifyInitialRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaCustomerCertifyInitialRequest) SetChannel(Channel string) {
	m.Channel = Channel
}

func (m *ZhimaCustomerCertifyInitialRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaCustomerCertifyInitialRequest) SetPlatform(Platform string) {
	m.Platform = Platform
}

func (m *ZhimaCustomerCertifyInitialRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaCustomerCertifyInitialRequest) SetExtParams(ExtParams string) {
	m.ExtParams = ExtParams
}

func (m *ZhimaCustomerCertifyInitialRequest) GetExtParams() string {
	return m.ExtParams
}
