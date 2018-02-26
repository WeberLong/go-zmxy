package requests

import (
	"github.com/WeberLong/go-zmxy/zmop/interfaces"
)

type ZhimaCreditPeLawsuitDetailQueryRequest struct {
	lawsuitId     string // 涉诉类型明细ID，对应字段值：裁判文书（partyId)，执行公告(zxggId)，失信记录(shixinId)，曝光台(bgtId)
	lawsuitType   string // 涉诉类型包括：裁判文书（party)，执行公告(zxgg)，失信记录(sxgg)，曝光台(bgt)
	productCode   string // 产品码
	transactionId string // 商户传入的业务流水号。此字段由商户生成，需确保唯一性，用于定位每一次请求，后续按此流水进行对帐。生成规则: 固定30位数字串，前17位为精确到毫秒的时间yyyyMMddhhmmssSSS，后13位为自增数字。

	interfaces.ZhimaRequestParams
}

func (m *ZhimaCreditPeLawsuitDetailQueryRequest) InitBizParams(lawsuitId, lawsuitType, productCode, transactionId string) {
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

func (*ZhimaCreditPeLawsuitDetailQueryRequest) GetApiMethodName() string {
	return "zhima.credit.pe.lawsuit.detail.query"
}

func (m *ZhimaCreditPeLawsuitDetailQueryRequest) GetApiParams() *map[string]string {
	return m.ApiParams
}

func (m *ZhimaCreditPeLawsuitDetailQueryRequest) GetFileParams() *map[string]string {
	return m.FileParams
}

func (m *ZhimaCreditPeLawsuitDetailQueryRequest) SetApiVersion(ApiVersion string) {
	m.ApiVersion = ApiVersion
}

func (m *ZhimaCreditPeLawsuitDetailQueryRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaCreditPeLawsuitDetailQueryRequest) SetScene(Scene string) {
	m.Scene = Scene
}

func (m *ZhimaCreditPeLawsuitDetailQueryRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaCreditPeLawsuitDetailQueryRequest) SetChannel(Channel string) {
	m.Channel = Channel
}

func (m *ZhimaCreditPeLawsuitDetailQueryRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaCreditPeLawsuitDetailQueryRequest) SetPlatform(Platform string) {
	m.Platform = Platform
}

func (m *ZhimaCreditPeLawsuitDetailQueryRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaCreditPeLawsuitDetailQueryRequest) SetExtParams(ExtParams string) {
	m.ExtParams = ExtParams
}

func (m *ZhimaCreditPeLawsuitDetailQueryRequest) GetExtParams() string {
	return m.ExtParams
}
