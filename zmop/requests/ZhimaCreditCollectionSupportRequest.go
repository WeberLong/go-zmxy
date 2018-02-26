package requests

import (
	"github.com/WeberLong/go-zmxy/zmop/interfaces"
)

type ZhimaCreditCollectionSupportRequest struct {
	certNo        string // 证件号码。当前仅支持身份证。请输入身份证号码
	certType      string // 证件类型。当前只支持身份证：IDENTITY_CARD
	name          string // 要查询的用户姓名
	numberStatus  string // 使用过的无效联系号码及状态列表。每一项格式为：${号码}|${状态};${号码}|${状态}。例如：18604020393|A;18604020394|B
	productCode   string // 云产品id
	transactionId string // 商户传入的业务流水号。此字段由商户生成，需确保唯一性，用于定位每一次请求，后续按此流水进行对帐。生成规则: 固定30位数字串，前17位为精确到毫秒的时间yyyyMMddhhmmssSSS，后13位为自增数字。

	interfaces.ZhimaRequestParams
}

func (m *ZhimaCreditCollectionSupportRequest) InitBizParams(certNo, certType, name, numberStatus, productCode, transactionId string) {
	m.ApiParams = &map[string]string{}
	m.FileParams = &map[string]string{}

	(*m.ApiParams)["cert_no"] = certNo
	m.certNo = certNo

	(*m.ApiParams)["cert_type"] = certType
	m.certType = certType

	(*m.ApiParams)["name"] = name
	m.name = name

	(*m.ApiParams)["number_status"] = numberStatus
	m.numberStatus = numberStatus

	(*m.ApiParams)["product_code"] = productCode
	m.productCode = productCode

	(*m.ApiParams)["transaction_id"] = transactionId
	m.transactionId = transactionId
}

func (*ZhimaCreditCollectionSupportRequest) GetApiMethodName() string {
	return "zhima.credit.collection.support"
}

func (m *ZhimaCreditCollectionSupportRequest) GetApiParams() *map[string]string {
	return m.ApiParams
}

func (m *ZhimaCreditCollectionSupportRequest) GetFileParams() *map[string]string {
	return m.FileParams
}

func (m *ZhimaCreditCollectionSupportRequest) SetApiVersion(ApiVersion string) {
	m.ApiVersion = ApiVersion
}

func (m *ZhimaCreditCollectionSupportRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaCreditCollectionSupportRequest) SetScene(Scene string) {
	m.Scene = Scene
}

func (m *ZhimaCreditCollectionSupportRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaCreditCollectionSupportRequest) SetChannel(Channel string) {
	m.Channel = Channel
}

func (m *ZhimaCreditCollectionSupportRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaCreditCollectionSupportRequest) SetPlatform(Platform string) {
	m.Platform = Platform
}

func (m *ZhimaCreditCollectionSupportRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaCreditCollectionSupportRequest) SetExtParams(ExtParams string) {
	m.ExtParams = ExtParams
}

func (m *ZhimaCreditCollectionSupportRequest) GetExtParams() string {
	return m.ExtParams
}
