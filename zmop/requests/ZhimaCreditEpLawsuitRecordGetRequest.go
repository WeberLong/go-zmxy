package requests

import (
	"github.com/WeberLong/go-zmxy/zmop/interfaces"
)

type ZhimaCreditEpLawsuitRecordGetRequest struct {
	companyName   string // 企业名称和组织机构号，2个参数必须要输入一个
	lawsuitType   string // 涉诉类型。cpws-裁判文书; zxgg-执行公告; sxgg-失信公告; ktgg-开庭公告; fygg-法院公告; ajlc-案件流程信息; bgt-曝光台。
	orgNo         string // 企业组织机构代码和企业名称2个参数，最少有一个
	productCode   string // 产品码
	transactionId string // 商户传入的业务流水号。此字段由商户生成，需确保唯一性，用于定位每一次请求，后续按此流水进行对帐。生成规则: 固定30位数字串，前17位为精确到毫秒的时间yyyyMMddhhmmssSSS，后13位为自增数字。

	interfaces.ZhimaRequestParams
}

func (m *ZhimaCreditEpLawsuitRecordGetRequest) InitBizParams(companyName, lawsuitType, orgNo, productCode, transactionId string) {
	m.ApiParams = &map[string]string{}
	m.FileParams = &map[string]string{}

	(*m.ApiParams)["company_name"] = companyName
	m.companyName = companyName

	(*m.ApiParams)["lawsuit_type"] = lawsuitType
	m.lawsuitType = lawsuitType

	(*m.ApiParams)["org_no"] = orgNo
	m.orgNo = orgNo

	(*m.ApiParams)["product_code"] = productCode
	m.productCode = productCode

	(*m.ApiParams)["transaction_id"] = transactionId
	m.transactionId = transactionId
}

func (*ZhimaCreditEpLawsuitRecordGetRequest) GetApiMethodName() string {
	return "zhima.credit.ep.lawsuit.record.get"
}

func (m *ZhimaCreditEpLawsuitRecordGetRequest) GetApiParams() *map[string]string {
	return m.ApiParams
}

func (m *ZhimaCreditEpLawsuitRecordGetRequest) GetFileParams() *map[string]string {
	return m.FileParams
}

func (m *ZhimaCreditEpLawsuitRecordGetRequest) SetApiVersion(ApiVersion string) {
	m.ApiVersion = ApiVersion
}

func (m *ZhimaCreditEpLawsuitRecordGetRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaCreditEpLawsuitRecordGetRequest) SetScene(Scene string) {
	m.Scene = Scene
}

func (m *ZhimaCreditEpLawsuitRecordGetRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaCreditEpLawsuitRecordGetRequest) SetChannel(Channel string) {
	m.Channel = Channel
}

func (m *ZhimaCreditEpLawsuitRecordGetRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaCreditEpLawsuitRecordGetRequest) SetPlatform(Platform string) {
	m.Platform = Platform
}

func (m *ZhimaCreditEpLawsuitRecordGetRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaCreditEpLawsuitRecordGetRequest) SetExtParams(ExtParams string) {
	m.ExtParams = ExtParams
}

func (m *ZhimaCreditEpLawsuitRecordGetRequest) GetExtParams() string {
	return m.ExtParams
}
