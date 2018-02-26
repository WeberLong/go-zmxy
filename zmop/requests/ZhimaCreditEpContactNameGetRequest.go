package requests

import (
	"github.com/WeberLong/go-zmxy/zmop/interfaces"
)

type ZhimaCreditEpContactNameGetRequest struct {
	companyName   string // 公司名称
	productCode   string // 产品码
	transactionId string // 商户传入的业务流水号。此字段由商户生成，需确保唯一性，用于定位每一次请求，后续按此流水进行对帐。生成规则: 固定30位数字串，前17位为精确到毫秒的时间yyyyMMddhhmmssSSS，后13位为自增数字。

	interfaces.ZhimaRequestParams
}

func (m *ZhimaCreditEpContactNameGetRequest) InitBizParams(companyName, productCode, transactionId string) {
	m.ApiParams = &map[string]string{}
	m.FileParams = &map[string]string{}

	(*m.ApiParams)["company_name"] = companyName
	m.companyName = companyName

	(*m.ApiParams)["product_code"] = productCode
	m.productCode = productCode

	(*m.ApiParams)["transaction_id"] = transactionId
	m.transactionId = transactionId
}

func (*ZhimaCreditEpContactNameGetRequest) GetApiMethodName() string {
	return "zhima.credit.ep.contact.name.get"
}

func (m *ZhimaCreditEpContactNameGetRequest) GetApiParams() *map[string]string {
	return m.ApiParams
}

func (m *ZhimaCreditEpContactNameGetRequest) GetFileParams() *map[string]string {
	return m.FileParams
}

func (m *ZhimaCreditEpContactNameGetRequest) SetApiVersion(ApiVersion string) {
	m.ApiVersion = ApiVersion
}

func (m *ZhimaCreditEpContactNameGetRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaCreditEpContactNameGetRequest) SetScene(Scene string) {
	m.Scene = Scene
}

func (m *ZhimaCreditEpContactNameGetRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaCreditEpContactNameGetRequest) SetChannel(Channel string) {
	m.Channel = Channel
}

func (m *ZhimaCreditEpContactNameGetRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaCreditEpContactNameGetRequest) SetPlatform(Platform string) {
	m.Platform = Platform
}

func (m *ZhimaCreditEpContactNameGetRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaCreditEpContactNameGetRequest) SetExtParams(ExtParams string) {
	m.ExtParams = ExtParams
}

func (m *ZhimaCreditEpContactNameGetRequest) GetExtParams() string {
	return m.ExtParams
}
