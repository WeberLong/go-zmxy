package requests

import (
	"github.com/WeberLong/go-zmxy/zmop/interfaces"
)

type ZhimaCreditXuejiVerifyRequest struct {
	collegeName       string // 院校名称
	educationCategory string // 输入为数字：1：表示普通全日制；2：表示硕士或者博士研究生；5：表示成人教育；6：表示高等教育自学考试；7：表示网络教育；8：表示开放教育；9：表示不详
	educationDegree   string // 学历层次：博士、硕士、本科、专科、成人。
	enrollmentYear    string // 入学年份，格式为四位数的年份，如2013
	graduateYear      string // 毕业年份，格式为四位数的年份，如2018
	openId            string // 芝麻会员在商户端的身份标识
	productCode       string // 产品码
	transactionId     string // 商户传入的业务流水号。此字段由商户生成，需确保唯一性，用于定位每一次请求，后续按此流水进行对帐。生成规则: 固定30位数字串，前17位为精确到毫秒的时间yyyyMMddHHmmssSSS，后13位为自增数字。

	interfaces.ZhimaRequestParams
}

func (m *ZhimaCreditXuejiVerifyRequest) InitBizParams(collegeName, educationCategory, educationDegree, enrollmentYear, graduateYear, openId, productCode, transactionId string) {
	m.ApiParams = &map[string]string{}
	m.FileParams = &map[string]string{}

	(*m.ApiParams)["college_name"] = collegeName
	m.collegeName = collegeName

	(*m.ApiParams)["education_category"] = educationCategory
	m.educationCategory = educationCategory

	(*m.ApiParams)["education_degree"] = educationDegree
	m.educationDegree = educationDegree

	(*m.ApiParams)["enrollment_year"] = enrollmentYear
	m.enrollmentYear = enrollmentYear

	(*m.ApiParams)["graduate_year"] = graduateYear
	m.graduateYear = graduateYear

	(*m.ApiParams)["open_id"] = openId
	m.openId = openId

	(*m.ApiParams)["product_code"] = productCode
	m.productCode = productCode

	(*m.ApiParams)["transaction_id"] = transactionId
	m.transactionId = transactionId
}

func (*ZhimaCreditXuejiVerifyRequest) GetApiMethodName() string {
	return "zhima.credit.xueji.verify"
}

func (m *ZhimaCreditXuejiVerifyRequest) GetApiParams() *map[string]string {
	return m.ApiParams
}

func (m *ZhimaCreditXuejiVerifyRequest) GetFileParams() *map[string]string {
	return m.FileParams
}

func (m *ZhimaCreditXuejiVerifyRequest) SetApiVersion(ApiVersion string) {
	m.ApiVersion = ApiVersion
}

func (m *ZhimaCreditXuejiVerifyRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaCreditXuejiVerifyRequest) SetScene(Scene string) {
	m.Scene = Scene
}

func (m *ZhimaCreditXuejiVerifyRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaCreditXuejiVerifyRequest) SetChannel(Channel string) {
	m.Channel = Channel
}

func (m *ZhimaCreditXuejiVerifyRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaCreditXuejiVerifyRequest) SetPlatform(Platform string) {
	m.Platform = Platform
}

func (m *ZhimaCreditXuejiVerifyRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaCreditXuejiVerifyRequest) SetExtParams(ExtParams string) {
	m.ExtParams = ExtParams
}

func (m *ZhimaCreditXuejiVerifyRequest) GetExtParams() string {
	return m.ExtParams
}
