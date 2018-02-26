package requests

import (
	"github.com/WeberLong/go-zmxy/zmop/interfaces"
)

type ZhimaCreditEpRoleGetRequest struct {
	certNo        string // 身份证号
	productCode   string // 产品码
	transactionId string // 商户传入的业务流水号。此字段由商户生成，需确保唯一性，用于定位每一次请求，后续按此流水进行对帐。生成规则: 固定30位数字串，前17位为精确到毫秒的时间yyyyMMddhhmmssSSS，后13位为自增数字。

	interfaces.ZhimaRequestParams
}

func (m *ZhimaCreditEpRoleGetRequest) InitBizParams(certNo, productCode, transactionId string) {
	m.ApiParams = &map[string]string{}
	m.FileParams = &map[string]string{}

	(*m.ApiParams)["cert_no"] = certNo
	m.certNo = certNo

	(*m.ApiParams)["product_code"] = productCode
	m.productCode = productCode

	(*m.ApiParams)["transaction_id"] = transactionId
	m.transactionId = transactionId
}

func (*ZhimaCreditEpRoleGetRequest) GetApiMethodName() string {
	return "zhima.credit.ep.role.get"
}

func (m *ZhimaCreditEpRoleGetRequest) GetApiParams() *map[string]string {
	return m.ApiParams
}

func (m *ZhimaCreditEpRoleGetRequest) GetFileParams() *map[string]string {
	return m.FileParams
}

func (m *ZhimaCreditEpRoleGetRequest) SetApiVersion(ApiVersion string) {
	m.ApiVersion = ApiVersion
}

func (m *ZhimaCreditEpRoleGetRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaCreditEpRoleGetRequest) SetScene(Scene string) {
	m.Scene = Scene
}

func (m *ZhimaCreditEpRoleGetRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaCreditEpRoleGetRequest) SetChannel(Channel string) {
	m.Channel = Channel
}

func (m *ZhimaCreditEpRoleGetRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaCreditEpRoleGetRequest) SetPlatform(Platform string) {
	m.Platform = Platform
}

func (m *ZhimaCreditEpRoleGetRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaCreditEpRoleGetRequest) SetExtParams(ExtParams string) {
	m.ExtParams = ExtParams
}

func (m *ZhimaCreditEpRoleGetRequest) GetExtParams() string {
	return m.ExtParams
}
