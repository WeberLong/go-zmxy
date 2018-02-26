package requests

import (
	"github.com/WeberLong/go-zmxy/zmop/interfaces"
)

type ZhimaCreditTrueidentityGetRequest struct {
	dataType      string // 数据类型。1-支付宝账号；2-手机号；3-支付宝登录账号
	dataValue     string // 数据类型对应的数据值。
	productCode   string // 产品码
	transactionId string // 商户传入的业务流水号。此字段由商户生成，需确保唯一性，用于定位每一次请求，后续按此流水进行对帐。生成规则: 固定30位数字串，前17位为精确到毫秒的时间yyyyMMddhhmmssSSS，后13位为自增数字。

	interfaces.ZhimaRequestParams
}

func (m *ZhimaCreditTrueidentityGetRequest) InitBizParams(dataType, dataValue, productCode, transactionId string) {
	m.ApiParams = &map[string]string{}
	m.FileParams = &map[string]string{}

	(*m.ApiParams)["data_type"] = dataType
	m.dataType = dataType

	(*m.ApiParams)["data_value"] = dataValue
	m.dataValue = dataValue

	(*m.ApiParams)["product_code"] = productCode
	m.productCode = productCode

	(*m.ApiParams)["transaction_id"] = transactionId
	m.transactionId = transactionId
}

func (*ZhimaCreditTrueidentityGetRequest) GetApiMethodName() string {
	return "zhima.credit.trueidentity.get"
}

func (m *ZhimaCreditTrueidentityGetRequest) GetApiParams() *map[string]string {
	return m.ApiParams
}

func (m *ZhimaCreditTrueidentityGetRequest) GetFileParams() *map[string]string {
	return m.FileParams
}

func (m *ZhimaCreditTrueidentityGetRequest) SetApiVersion(ApiVersion string) {
	m.ApiVersion = ApiVersion
}

func (m *ZhimaCreditTrueidentityGetRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaCreditTrueidentityGetRequest) SetScene(Scene string) {
	m.Scene = Scene
}

func (m *ZhimaCreditTrueidentityGetRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaCreditTrueidentityGetRequest) SetChannel(Channel string) {
	m.Channel = Channel
}

func (m *ZhimaCreditTrueidentityGetRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaCreditTrueidentityGetRequest) SetPlatform(Platform string) {
	m.Platform = Platform
}

func (m *ZhimaCreditTrueidentityGetRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaCreditTrueidentityGetRequest) SetExtParams(ExtParams string) {
	m.ExtParams = ExtParams
}

func (m *ZhimaCreditTrueidentityGetRequest) GetExtParams() string {
	return m.ExtParams
}
