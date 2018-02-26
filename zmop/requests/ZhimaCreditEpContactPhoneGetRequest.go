package requests

import (
	"github.com/WeberLong/go-zmxy/zmop/interfaces"
)

type ZhimaCreditEpContactPhoneGetRequest struct {
	phone         string // 电话号码
	productCode   string // 产品码
	transactionId string // 商户传入的业务流水号。此字段由商户生成，需确保唯一性，用于定位每一次请求，后续按此流水进行对帐。生成规则: 固定30位数字串，前17位为精确到毫秒的时间yyyyMMddhhmmssSSS，后13位为自增数字。

	interfaces.ZhimaRequestParams
}

func (m *ZhimaCreditEpContactPhoneGetRequest) InitBizParams(phone, productCode, transactionId string) {
	m.ApiParams = &map[string]string{}
	m.FileParams = &map[string]string{}

	(*m.ApiParams)["phone"] = phone
	m.phone = phone

	(*m.ApiParams)["product_code"] = productCode
	m.productCode = productCode

	(*m.ApiParams)["transaction_id"] = transactionId
	m.transactionId = transactionId
}

func (*ZhimaCreditEpContactPhoneGetRequest) GetApiMethodName() string {
	return "zhima.credit.ep.contact.phone.get"
}

func (m *ZhimaCreditEpContactPhoneGetRequest) GetApiParams() *map[string]string {
	return m.ApiParams
}

func (m *ZhimaCreditEpContactPhoneGetRequest) GetFileParams() *map[string]string {
	return m.FileParams
}

func (m *ZhimaCreditEpContactPhoneGetRequest) SetApiVersion(ApiVersion string) {
	m.ApiVersion = ApiVersion
}

func (m *ZhimaCreditEpContactPhoneGetRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaCreditEpContactPhoneGetRequest) SetScene(Scene string) {
	m.Scene = Scene
}

func (m *ZhimaCreditEpContactPhoneGetRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaCreditEpContactPhoneGetRequest) SetChannel(Channel string) {
	m.Channel = Channel
}

func (m *ZhimaCreditEpContactPhoneGetRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaCreditEpContactPhoneGetRequest) SetPlatform(Platform string) {
	m.Platform = Platform
}

func (m *ZhimaCreditEpContactPhoneGetRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaCreditEpContactPhoneGetRequest) SetExtParams(ExtParams string) {
	m.ExtParams = ExtParams
}

func (m *ZhimaCreditEpContactPhoneGetRequest) GetExtParams() string {
	return m.ExtParams
}
