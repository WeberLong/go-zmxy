package requests

import (
	"github.com/WeberLong/go-zmxy/zmop/interfaces"
)

type ZhimaCustomerCertificationInitializeRequest struct {
	bizCode        string // 认证场景码，支持的场景码有：FACE：多因子活体人脸认证，SMART_FACE：多因子快捷活体人脸认证，FACE_SDK：SDK活体人脸认证签约的协议决定了可以使用的场景
	extBizParam    string // 扩展业务参数，暂时没有用到，接口预留
	identityParam  string // 值为一个json串，无入参时值为"{}"，有入参时必须指定身份类型identity_type，不同的身份类型对应的身份信息不同当前支持：身份信息为证件信息，identity_type值为CERT_INFO：证件类型为身份证cert_type值为IDENTITY_CARD，必要信息cert_name和cert_no；身份信息为短信手机号，适用于短信认证，identity_type值为SMS_MOBILE_NO：证件类型可以为空，当证件类型为身份证cert_type值为IDENTITY_CARD，必要信息cert_name和cert_no，mobile_no可以为空，以上信息没有传入的时候会上用户录入；身份信息为支付宝UID，identity_type值为USER_ID:必要信息user_id示例：{"identity_type": "USER_ID", "user_id": "2088172637486509"}
	merchantConfig string // 认证商户自定义配置，支持一些商户可选的功能need_user_authorization： 值为true或者false当值为true时，在认证用户引导页会展示用户授权协议，在认证通过后会进行授权，但是授权是否成功不影响认证结果
	productCode    string // 芝麻认证产品码,示例值为真实的产品码
	transactionId  string // 商户请求的唯一标志，32位长度的字母数字下划线组合。该标识作为对账的关键信息，商户要保证其唯一性.建议:前面几位字符是商户自定义的简称,中间可以使用一段日期,结尾可以使用一个序列

	interfaces.ZhimaRequestParams
}

func (m *ZhimaCustomerCertificationInitializeRequest) InitBizParams(bizCode, extBizParam, identityParam, merchantConfig, productCode, transactionId string) {
	m.ApiParams = &map[string]string{}
	m.FileParams = &map[string]string{}

	(*m.ApiParams)["biz_code"] = bizCode
	m.bizCode = bizCode

	(*m.ApiParams)["ext_biz_param"] = extBizParam
	m.extBizParam = extBizParam

	(*m.ApiParams)["identity_param"] = identityParam
	m.identityParam = identityParam

	(*m.ApiParams)["merchant_config"] = merchantConfig
	m.merchantConfig = merchantConfig

	(*m.ApiParams)["product_code"] = productCode
	m.productCode = productCode

	(*m.ApiParams)["transaction_id"] = transactionId
	m.transactionId = transactionId
}

func (*ZhimaCustomerCertificationInitializeRequest) GetApiMethodName() string {
	return "zhima.customer.certification.initialize"
}

func (m *ZhimaCustomerCertificationInitializeRequest) GetApiParams() *map[string]string {
	return m.ApiParams
}

func (m *ZhimaCustomerCertificationInitializeRequest) GetFileParams() *map[string]string {
	return m.FileParams
}

func (m *ZhimaCustomerCertificationInitializeRequest) SetApiVersion(ApiVersion string) {
	m.ApiVersion = ApiVersion
}

func (m *ZhimaCustomerCertificationInitializeRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaCustomerCertificationInitializeRequest) SetScene(Scene string) {
	m.Scene = Scene
}

func (m *ZhimaCustomerCertificationInitializeRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaCustomerCertificationInitializeRequest) SetChannel(Channel string) {
	m.Channel = Channel
}

func (m *ZhimaCustomerCertificationInitializeRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaCustomerCertificationInitializeRequest) SetPlatform(Platform string) {
	m.Platform = Platform
}

func (m *ZhimaCustomerCertificationInitializeRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaCustomerCertificationInitializeRequest) SetExtParams(ExtParams string) {
	m.ExtParams = ExtParams
}

func (m *ZhimaCustomerCertificationInitializeRequest) GetExtParams() string {
	return m.ExtParams
}
