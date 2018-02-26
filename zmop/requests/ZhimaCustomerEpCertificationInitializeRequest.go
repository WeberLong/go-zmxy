package requests

import (
	"github.com/WeberLong/go-zmxy/zmop/interfaces"
)

type ZhimaCustomerEpCertificationInitializeRequest struct {
	bizCode        string // 认证场景码，支持的场景码有： EP_ALIPAY_ACCOUNT,  签约的协议决定了可以使用的场景
	extBizParam    string // 扩展业务参数，暂时没有用到，接口预留
	identityParam  string // 值为一个json串，无入参时值为"{}"，有入参时必须指定身份类型identity_type，不同的身份类型对应的身份信息不同 当前支持的身份信息为证件信息，identity_type=EP_CERT_INFO  需要输入法人证件三要素，企业证件三要素，如 {"identity_type": "EP_CERT_INFO", "cert_type": "IDENTITY_CARD", "cert_name": "收委", "cert_no":"260104197909275964", "ep_cert_type": "NATIONAL_LEGAL_MERGE", "ep_cert_name": "xxx有限公司", "ep_cert_no":"260104199909275964"}； 特别备注： 上述json串中的 ep_cert_type 属性仅支持2种类型：NATIONAL_LEGAL：工商注册号类型NATIONAL_LEGAL_MERGE ： 社会统一信用代码类型
	merchantConfig string // 认证商户自定义配置，支持一些商户可选的功能,目前为预留的属性值
	productCode    string // 芝麻认证产品码,示例值为真实的产品码
	transactionId  string // 商户请求的唯一标志，32位长度的字母数字下划线组合。该标识作为对账的关键信息，商户要保证其唯一性.建议:前面几位字符是商户自定义的简称,中间可以使用一段日期,结尾可以使用一个序列

	interfaces.ZhimaRequestParams
}

func (m *ZhimaCustomerEpCertificationInitializeRequest) InitBizParams(bizCode, extBizParam, identityParam, merchantConfig, productCode, transactionId string) {
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

func (*ZhimaCustomerEpCertificationInitializeRequest) GetApiMethodName() string {
	return "zhima.customer.ep.certification.initialize"
}

func (m *ZhimaCustomerEpCertificationInitializeRequest) GetApiParams() *map[string]string {
	return m.ApiParams
}

func (m *ZhimaCustomerEpCertificationInitializeRequest) GetFileParams() *map[string]string {
	return m.FileParams
}

func (m *ZhimaCustomerEpCertificationInitializeRequest) SetApiVersion(ApiVersion string) {
	m.ApiVersion = ApiVersion
}

func (m *ZhimaCustomerEpCertificationInitializeRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaCustomerEpCertificationInitializeRequest) SetScene(Scene string) {
	m.Scene = Scene
}

func (m *ZhimaCustomerEpCertificationInitializeRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaCustomerEpCertificationInitializeRequest) SetChannel(Channel string) {
	m.Channel = Channel
}

func (m *ZhimaCustomerEpCertificationInitializeRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaCustomerEpCertificationInitializeRequest) SetPlatform(Platform string) {
	m.Platform = Platform
}

func (m *ZhimaCustomerEpCertificationInitializeRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaCustomerEpCertificationInitializeRequest) SetExtParams(ExtParams string) {
	m.ExtParams = ExtParams
}

func (m *ZhimaCustomerEpCertificationInitializeRequest) GetExtParams() string {
	return m.ExtParams
}
