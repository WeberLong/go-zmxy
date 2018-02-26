package requests

import (
	"github.com/WeberLong/go-zmxy/zmop/interfaces"
)

type ZhimaCreditEpCommercialPunishmentQueryRequest struct {
	certNo        string // 查询企业，不需要填写证件号。 查询个人，必须填入证件号。
	certType      string // 查询个人的时候必填，填入IDENTITY_CARD。目前支持内地身份证。
	dataType      string // 枚举值 1:个人 2:企业
	name          string // 企业名称或个人姓名
	productCode   string // 产品码
	transactionId string // transaction_id是代表一笔请求的唯一标志，该标识作为对账的关键信息，对于用户使用相同transaction_id的查询，芝麻在一天（86400秒）内返回首次查询数据，超过有效期的查询即为无效并返回异常（错误码xxxx），有效期内的重复查询不重新计费。 transaction_id 推荐生成方式是：30位，（其中17位时间值（精确到毫秒）：yyyyMMddHHmmssSSS）加上（13位自增数字：1234567890123）

	interfaces.ZhimaRequestParams
}

func (m *ZhimaCreditEpCommercialPunishmentQueryRequest) InitBizParams(certNo, certType, dataType, name, productCode, transactionId string) {
	m.ApiParams = &map[string]string{}
	m.FileParams = &map[string]string{}

	(*m.ApiParams)["cert_no"] = certNo
	m.certNo = certNo

	(*m.ApiParams)["cert_type"] = certType
	m.certType = certType

	(*m.ApiParams)["data_type"] = dataType
	m.dataType = dataType

	(*m.ApiParams)["name"] = name
	m.name = name

	(*m.ApiParams)["product_code"] = productCode
	m.productCode = productCode

	(*m.ApiParams)["transaction_id"] = transactionId
	m.transactionId = transactionId
}

func (*ZhimaCreditEpCommercialPunishmentQueryRequest) GetApiMethodName() string {
	return "zhima.credit.ep.commercial.punishment.query"
}

func (m *ZhimaCreditEpCommercialPunishmentQueryRequest) GetApiParams() *map[string]string {
	return m.ApiParams
}

func (m *ZhimaCreditEpCommercialPunishmentQueryRequest) GetFileParams() *map[string]string {
	return m.FileParams
}

func (m *ZhimaCreditEpCommercialPunishmentQueryRequest) SetApiVersion(ApiVersion string) {
	m.ApiVersion = ApiVersion
}

func (m *ZhimaCreditEpCommercialPunishmentQueryRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaCreditEpCommercialPunishmentQueryRequest) SetScene(Scene string) {
	m.Scene = Scene
}

func (m *ZhimaCreditEpCommercialPunishmentQueryRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaCreditEpCommercialPunishmentQueryRequest) SetChannel(Channel string) {
	m.Channel = Channel
}

func (m *ZhimaCreditEpCommercialPunishmentQueryRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaCreditEpCommercialPunishmentQueryRequest) SetPlatform(Platform string) {
	m.Platform = Platform
}

func (m *ZhimaCreditEpCommercialPunishmentQueryRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaCreditEpCommercialPunishmentQueryRequest) SetExtParams(ExtParams string) {
	m.ExtParams = ExtParams
}

func (m *ZhimaCreditEpCommercialPunishmentQueryRequest) GetExtParams() string {
	return m.ExtParams
}
