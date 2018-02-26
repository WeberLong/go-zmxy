package requests

import (
	"github.com/WeberLong/go-zmxy/zmop/interfaces"
)

type ZhimaCustomerCertificationMaterialCertifyRequest struct {
	bizCode        string // 认证场景码，支持的场景码有： FACE_API：通过接口进行人脸识别签约的协议决定了可以使用的场景
	extBizParam    string // 扩展业务参数，暂时没有用到，接口预留
	identityParam  string // 值为一个json串，必须指定身份类型identity_type，不同的身份类型对应的身份信息不同当前支持：身份信息为证件信息，identity_type值为CERT_INFO:证件类型为身份证cert_type值为IDENTITY_CARD,必要信息cert_type，cert_name和cert_no；身份信息为个人正面照片，identity_type值为FACIAL_PICTURE_FRONT:必要信息cert_type，cert_name和cert_no和FACIAL_PICTURE_FRONT示例：{"identity_type": "FACIAL_PICTURE_FRONT", "cert_type": "IDENTITY_CARD", "cert_name": "收委", "cert_no": "260104197909275964", "FACIAL_PICTURE_FRONT": "/9j/4AAQSkZJR"}
	materials      string // 认证过程中需要的认证材料，不同认证场景需要的材料不同biz_code值为FACE_API时需要材料FACIAL_PICTURE_FRONT
	merchantConfig string // 认证商户自定义配置，支持一些商户可选的功能
	productCode    string // 芝麻认证产品码,示例值为真实的产品码
	transactionId  string // 商户请求的唯一标志，32位长度的字母数字下划线组合。该标识作为对账的关键信息，商户要保证其唯一性.建议:前面几位字符是商户自定义的简称,中间可以使用一段日期,结尾可以使用一个序列

	interfaces.ZhimaRequestParams
}

func (m *ZhimaCustomerCertificationMaterialCertifyRequest) InitBizParams(bizCode, extBizParam, identityParam, materials, merchantConfig, productCode, transactionId string) {
	m.ApiParams = &map[string]string{}
	m.FileParams = &map[string]string{}

	(*m.ApiParams)["biz_code"] = bizCode
	m.bizCode = bizCode

	(*m.ApiParams)["ext_biz_param"] = extBizParam
	m.extBizParam = extBizParam

	(*m.ApiParams)["identity_param"] = identityParam
	m.identityParam = identityParam

	(*m.ApiParams)["materials"] = materials
	m.materials = materials

	(*m.ApiParams)["merchant_config"] = merchantConfig
	m.merchantConfig = merchantConfig

	(*m.ApiParams)["product_code"] = productCode
	m.productCode = productCode

	(*m.ApiParams)["transaction_id"] = transactionId
	m.transactionId = transactionId
}

func (*ZhimaCustomerCertificationMaterialCertifyRequest) GetApiMethodName() string {
	return "zhima.customer.certification.material.certify"
}

func (m *ZhimaCustomerCertificationMaterialCertifyRequest) GetApiParams() *map[string]string {
	return m.ApiParams
}

func (m *ZhimaCustomerCertificationMaterialCertifyRequest) GetFileParams() *map[string]string {
	return m.FileParams
}

func (m *ZhimaCustomerCertificationMaterialCertifyRequest) SetApiVersion(ApiVersion string) {
	m.ApiVersion = ApiVersion
}

func (m *ZhimaCustomerCertificationMaterialCertifyRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaCustomerCertificationMaterialCertifyRequest) SetScene(Scene string) {
	m.Scene = Scene
}

func (m *ZhimaCustomerCertificationMaterialCertifyRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaCustomerCertificationMaterialCertifyRequest) SetChannel(Channel string) {
	m.Channel = Channel
}

func (m *ZhimaCustomerCertificationMaterialCertifyRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaCustomerCertificationMaterialCertifyRequest) SetPlatform(Platform string) {
	m.Platform = Platform
}

func (m *ZhimaCustomerCertificationMaterialCertifyRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaCustomerCertificationMaterialCertifyRequest) SetExtParams(ExtParams string) {
	m.ExtParams = ExtParams
}

func (m *ZhimaCustomerCertificationMaterialCertifyRequest) GetExtParams() string {
	return m.ExtParams
}
