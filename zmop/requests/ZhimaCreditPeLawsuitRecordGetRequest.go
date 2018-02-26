package requests

import (
	"github.com/WeberLong/go-zmxy/zmop/interfaces"
)

type ZhimaCreditPeLawsuitRecordGetRequest struct {
	certNo        string // 身份证
	name          string // 姓名
	productCode   string // 产品码
	transactionId string // 商户传入的业务流水号。此字段由商户生成，需确保唯一性，用于定位每一次请求，后续按此流水进行对帐。生成规则: 固定30位数字串，前17位为精确到毫秒的时间yyyyMMddhhmmssSSS，后13位为自增数字。

	interfaces.ZhimaRequestParams
}

func (m *ZhimaCreditPeLawsuitRecordGetRequest) InitBizParams(certNo, name, productCode, transactionId string) {
	m.ApiParams = &map[string]string{}
	m.FileParams = &map[string]string{}

	(*m.ApiParams)["cert_no"] = certNo
	m.certNo = certNo

	(*m.ApiParams)["name"] = name
	m.name = name

	(*m.ApiParams)["product_code"] = productCode
	m.productCode = productCode

	(*m.ApiParams)["transaction_id"] = transactionId
	m.transactionId = transactionId
}

func (*ZhimaCreditPeLawsuitRecordGetRequest) GetApiMethodName() string {
	return "zhima.credit.pe.lawsuit.record.get"
}

func (m *ZhimaCreditPeLawsuitRecordGetRequest) GetApiParams() *map[string]string {
	return m.ApiParams
}

func (m *ZhimaCreditPeLawsuitRecordGetRequest) GetFileParams() *map[string]string {
	return m.FileParams
}

func (m *ZhimaCreditPeLawsuitRecordGetRequest) SetApiVersion(ApiVersion string) {
	m.ApiVersion = ApiVersion
}

func (m *ZhimaCreditPeLawsuitRecordGetRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaCreditPeLawsuitRecordGetRequest) SetScene(Scene string) {
	m.Scene = Scene
}

func (m *ZhimaCreditPeLawsuitRecordGetRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaCreditPeLawsuitRecordGetRequest) SetChannel(Channel string) {
	m.Channel = Channel
}

func (m *ZhimaCreditPeLawsuitRecordGetRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaCreditPeLawsuitRecordGetRequest) SetPlatform(Platform string) {
	m.Platform = Platform
}

func (m *ZhimaCreditPeLawsuitRecordGetRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaCreditPeLawsuitRecordGetRequest) SetExtParams(ExtParams string) {
	m.ExtParams = ExtParams
}

func (m *ZhimaCreditPeLawsuitRecordGetRequest) GetExtParams() string {
	return m.ExtParams
}
