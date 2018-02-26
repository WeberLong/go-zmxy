package requests

import (
	"github.com/WeberLong/go-zmxy/zmop/interfaces"
)

type ZhimaCreditXueliVerifyRequest struct {
	certNo            string // 证件号码，暂时只支持身份证号
	certType          string // 证件类型，暂时支持身份证
	collegeName       string // 院校名称
	educationCategory string // 学历类别：普通，研究生，成人，高等教育自学考试，网络教育，开放教育，不详
	educationDegree   string // 学历层次：博士研究生，硕士研究生，研究生班，第二学士学位，本科，高升本，专升本，第二本科专科，专科(高职)，第二专科，夜大电大函大普通班，大学
	graduateYear      string // 毕业年份，格式为四位数的年份，如2018
	name              string // 姓名
	productCode       string // 产品码
	transactionId     string // 商户传入的业务流水号。此字段由商户生成，需确保唯一性，用于定位每一次请求，后续按此流水进行对帐。生成规则: 固定30位数字串，前17位为精确到毫秒的时间yyyyMMddHHmmssSSS，后13位为自增数字。

	interfaces.ZhimaRequestParams
}

func (m *ZhimaCreditXueliVerifyRequest) InitBizParams(certNo, certType, collegeName, educationCategory, educationDegree, graduateYear, name, productCode, transactionId string) {
	m.ApiParams = &map[string]string{}
	m.FileParams = &map[string]string{}

	(*m.ApiParams)["cert_no"] = certNo
	m.certNo = certNo

	(*m.ApiParams)["cert_type"] = certType
	m.certType = certType

	(*m.ApiParams)["college_name"] = collegeName
	m.collegeName = collegeName

	(*m.ApiParams)["education_category"] = educationCategory
	m.educationCategory = educationCategory

	(*m.ApiParams)["education_degree"] = educationDegree
	m.educationDegree = educationDegree

	(*m.ApiParams)["graduate_year"] = graduateYear
	m.graduateYear = graduateYear

	(*m.ApiParams)["name"] = name
	m.name = name

	(*m.ApiParams)["product_code"] = productCode
	m.productCode = productCode

	(*m.ApiParams)["transaction_id"] = transactionId
	m.transactionId = transactionId
}

func (*ZhimaCreditXueliVerifyRequest) GetApiMethodName() string {
	return "zhima.credit.xueli.verify"
}

func (m *ZhimaCreditXueliVerifyRequest) GetApiParams() *map[string]string {
	return m.ApiParams
}

func (m *ZhimaCreditXueliVerifyRequest) GetFileParams() *map[string]string {
	return m.FileParams
}

func (m *ZhimaCreditXueliVerifyRequest) SetApiVersion(ApiVersion string) {
	m.ApiVersion = ApiVersion
}

func (m *ZhimaCreditXueliVerifyRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaCreditXueliVerifyRequest) SetScene(Scene string) {
	m.Scene = Scene
}

func (m *ZhimaCreditXueliVerifyRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaCreditXueliVerifyRequest) SetChannel(Channel string) {
	m.Channel = Channel
}

func (m *ZhimaCreditXueliVerifyRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaCreditXueliVerifyRequest) SetPlatform(Platform string) {
	m.Platform = Platform
}

func (m *ZhimaCreditXueliVerifyRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaCreditXueliVerifyRequest) SetExtParams(ExtParams string) {
	m.ExtParams = ExtParams
}

func (m *ZhimaCreditXueliVerifyRequest) GetExtParams() string {
	return m.ExtParams
}
