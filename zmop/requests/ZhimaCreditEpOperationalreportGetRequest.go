package requests

import (
	"github.com/WeberLong/go-zmxy/zmop/interfaces"
)

type ZhimaCreditEpOperationalreportGetRequest struct {
	openId        string // 用户授权后的openId
	productCode   string // 产品码
	transactionId string // 商户传入的业务流水号。此字段由商户生成，需确保唯一性，用于定位每一次请求，后续按此流水进行对帐。生成规则: 固定30位数字串，前17位为精确到毫秒的时间yyyyMMddHHmmssSSS，后13位为自增数字。

	interfaces.ZhimaRequestParams
}

func (m *ZhimaCreditEpOperationalreportGetRequest) InitBizParams(openId, productCode, transactionId string) {
	m.ApiParams = &map[string]string{}
	m.FileParams = &map[string]string{}

	(*m.ApiParams)["open_id"] = openId
	m.openId = openId

	(*m.ApiParams)["product_code"] = productCode
	m.productCode = productCode

	(*m.ApiParams)["transaction_id"] = transactionId
	m.transactionId = transactionId
}

func (*ZhimaCreditEpOperationalreportGetRequest) GetApiMethodName() string {
	return "zhima.credit.ep.operationalreport.get"
}

func (m *ZhimaCreditEpOperationalreportGetRequest) GetApiParams() *map[string]string {
	return m.ApiParams
}

func (m *ZhimaCreditEpOperationalreportGetRequest) GetFileParams() *map[string]string {
	return m.FileParams
}

func (m *ZhimaCreditEpOperationalreportGetRequest) SetApiVersion(ApiVersion string) {
	m.ApiVersion = ApiVersion
}

func (m *ZhimaCreditEpOperationalreportGetRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaCreditEpOperationalreportGetRequest) SetScene(Scene string) {
	m.Scene = Scene
}

func (m *ZhimaCreditEpOperationalreportGetRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaCreditEpOperationalreportGetRequest) SetChannel(Channel string) {
	m.Channel = Channel
}

func (m *ZhimaCreditEpOperationalreportGetRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaCreditEpOperationalreportGetRequest) SetPlatform(Platform string) {
	m.Platform = Platform
}

func (m *ZhimaCreditEpOperationalreportGetRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaCreditEpOperationalreportGetRequest) SetExtParams(ExtParams string) {
	m.ExtParams = ExtParams
}

func (m *ZhimaCreditEpOperationalreportGetRequest) GetExtParams() string {
	return m.ExtParams
}
