package requests

import (
	"github.com/WeberLong/go-zmxy/zmop/interfaces"
)

type ZhimaCreditEpLawsuitDetailGetRequest struct {
	lawsuitId     string // 涉诉ID
	lawsuitType   string // 涉诉类型。cpws-裁判文书; zxgg-执行公告; sxgg-失信公告; ktgg-开庭公告; fygg-法院公告; ajlc-案件流程信息; bgt-曝光台。
	productCode   string // 产品码
	transactionId string // 商户传入的业务流水号。此字段由商户生成，需确保唯一性，用于定位每一次请求，后续按此流水进行对帐。生成规则: 固定30位数字串，前17位为精确到毫秒的时间yyyyMMddhhmmssSSS，后13位为自增数字。

	interfaces.ZhimaRequestParams
}

func (m *ZhimaCreditEpLawsuitDetailGetRequest) InitBizParams(lawsuitId, lawsuitType, productCode, transactionId string) {
	m.ApiParams = &map[string]string{}
	m.FileParams = &map[string]string{}

	(*m.ApiParams)["lawsuit_id"] = lawsuitId
	m.lawsuitId = lawsuitId

	(*m.ApiParams)["lawsuit_type"] = lawsuitType
	m.lawsuitType = lawsuitType

	(*m.ApiParams)["product_code"] = productCode
	m.productCode = productCode

	(*m.ApiParams)["transaction_id"] = transactionId
	m.transactionId = transactionId
}

func (*ZhimaCreditEpLawsuitDetailGetRequest) GetApiMethodName() string {
	return "zhima.credit.ep.lawsuit.detail.get"
}

func (m *ZhimaCreditEpLawsuitDetailGetRequest) GetApiParams() *map[string]string {
	return m.ApiParams
}

func (m *ZhimaCreditEpLawsuitDetailGetRequest) GetFileParams() *map[string]string {
	return m.FileParams
}

func (m *ZhimaCreditEpLawsuitDetailGetRequest) SetApiVersion(ApiVersion string) {
	m.ApiVersion = ApiVersion
}

func (m *ZhimaCreditEpLawsuitDetailGetRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaCreditEpLawsuitDetailGetRequest) SetScene(Scene string) {
	m.Scene = Scene
}

func (m *ZhimaCreditEpLawsuitDetailGetRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaCreditEpLawsuitDetailGetRequest) SetChannel(Channel string) {
	m.Channel = Channel
}

func (m *ZhimaCreditEpLawsuitDetailGetRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaCreditEpLawsuitDetailGetRequest) SetPlatform(Platform string) {
	m.Platform = Platform
}

func (m *ZhimaCreditEpLawsuitDetailGetRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaCreditEpLawsuitDetailGetRequest) SetExtParams(ExtParams string) {
	m.ExtParams = ExtParams
}

func (m *ZhimaCreditEpLawsuitDetailGetRequest) GetExtParams() string {
	return m.ExtParams
}
