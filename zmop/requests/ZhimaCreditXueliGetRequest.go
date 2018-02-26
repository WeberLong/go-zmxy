package requests

import (
	"github.com/WeberLong/go-zmxy/zmop/interfaces"
)

type ZhimaCreditXueliGetRequest struct {
	openId        string // 芝麻会员在商户端的身份标识
	productCode   string // 产品码
	transactionId string // 商户传入的业务流水号。此字段由商户生成，需确保唯一性，用于定位每一次请求，后续按此流水进行对帐。生成规则: 固定30位数字串，前17位为精确到毫秒的时间yyyyMMddHHmmssSSS，后13位为自增数字。

	interfaces.ZhimaRequestParams
}

func (m *ZhimaCreditXueliGetRequest) InitBizParams(openId, productCode, transactionId string) {
	m.ApiParams = &map[string]string{}
	m.FileParams = &map[string]string{}

	(*m.ApiParams)["open_id"] = openId
	m.openId = openId

	(*m.ApiParams)["product_code"] = productCode
	m.productCode = productCode

	(*m.ApiParams)["transaction_id"] = transactionId
	m.transactionId = transactionId
}

func (*ZhimaCreditXueliGetRequest) GetApiMethodName() string {
	return "zhima.credit.xueli.get"
}

func (m *ZhimaCreditXueliGetRequest) GetApiParams() *map[string]string {
	return m.ApiParams
}

func (m *ZhimaCreditXueliGetRequest) GetFileParams() *map[string]string {
	return m.FileParams
}

func (m *ZhimaCreditXueliGetRequest) SetApiVersion(ApiVersion string) {
	m.ApiVersion = ApiVersion
}

func (m *ZhimaCreditXueliGetRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaCreditXueliGetRequest) SetScene(Scene string) {
	m.Scene = Scene
}

func (m *ZhimaCreditXueliGetRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaCreditXueliGetRequest) SetChannel(Channel string) {
	m.Channel = Channel
}

func (m *ZhimaCreditXueliGetRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaCreditXueliGetRequest) SetPlatform(Platform string) {
	m.Platform = Platform
}

func (m *ZhimaCreditXueliGetRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaCreditXueliGetRequest) SetExtParams(ExtParams string) {
	m.ExtParams = ExtParams
}

func (m *ZhimaCreditXueliGetRequest) GetExtParams() string {
	return m.ExtParams
}
