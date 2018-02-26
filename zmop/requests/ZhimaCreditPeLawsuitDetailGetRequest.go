package requests

import (
	"github.com/WeberLong/go-zmxy/zmop/interfaces"
)

type ZhimaCreditPeLawsuitDetailGetRequest struct {
	lawsuitId     string // 涉诉类型明细ID，对应字段值：裁判文书（partyId)，执行公告(zxggId)，失信记录(shixinId)，曝光台(bgtId)
	lawsuitType   string // 涉诉类型包括：裁判文书（party)，执行公告(zxgg)，失信记录(sxgg)，曝光台(bgt)
	productCode   string // 产品码
	transactionId string // 商户传入的业务流水号。此字段由商户生成，需确保唯一性，用于定位每一次请求，后续按此流水进行对帐。生成规则: 固定30位数字串，前17位为精确到毫秒的时间yyyyMMddhhmmssSSS，后13位为自增数字。

	interfaces.ZhimaRequestParams
}

func (m *ZhimaCreditPeLawsuitDetailGetRequest) InitBizParams(lawsuitId, lawsuitType, productCode, transactionId string) {
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

func (*ZhimaCreditPeLawsuitDetailGetRequest) GetApiMethodName() string {
	return "zhima.credit.pe.lawsuit.detail.get"
}

func (m *ZhimaCreditPeLawsuitDetailGetRequest) GetApiParams() *map[string]string {
	return m.ApiParams
}

func (m *ZhimaCreditPeLawsuitDetailGetRequest) GetFileParams() *map[string]string {
	return m.FileParams
}

func (m *ZhimaCreditPeLawsuitDetailGetRequest) SetApiVersion(ApiVersion string) {
	m.ApiVersion = ApiVersion
}

func (m *ZhimaCreditPeLawsuitDetailGetRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaCreditPeLawsuitDetailGetRequest) SetScene(Scene string) {
	m.Scene = Scene
}

func (m *ZhimaCreditPeLawsuitDetailGetRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaCreditPeLawsuitDetailGetRequest) SetChannel(Channel string) {
	m.Channel = Channel
}

func (m *ZhimaCreditPeLawsuitDetailGetRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaCreditPeLawsuitDetailGetRequest) SetPlatform(Platform string) {
	m.Platform = Platform
}

func (m *ZhimaCreditPeLawsuitDetailGetRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaCreditPeLawsuitDetailGetRequest) SetExtParams(ExtParams string) {
	m.ExtParams = ExtParams
}

func (m *ZhimaCreditPeLawsuitDetailGetRequest) GetExtParams() string {
	return m.ExtParams
}
