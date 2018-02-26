package requests

import (
	"github.com/WeberLong/go-zmxy/zmop/interfaces"
)

type ZhimaCreditEpInfoGetRequest struct {
	dataType      string // 查询类型。1-社会信用代码；2-企业名称；3-企业注册号；4-组织机构代码。
	dataValue     string // 数据类型的对应值。
	productCode   string // 产品码
	transactionId string // 商户传入的业务流水号。此字段由商户生成，需确保唯一性，用于定位每一次请求，后续按此流水进行对帐。生成规则: 固定30位数字串，前17位为精确到毫秒的时间yyyyMMddhhmmssSSS，后13位为自增数字。

	interfaces.ZhimaRequestParams
}

func (m *ZhimaCreditEpInfoGetRequest) InitBizParams(dataType, dataValue, productCode, transactionId string) {
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

func (*ZhimaCreditEpInfoGetRequest) GetApiMethodName() string {
	return "zhima.credit.ep.info.get"
}

func (m *ZhimaCreditEpInfoGetRequest) GetApiParams() *map[string]string {
	return m.ApiParams
}

func (m *ZhimaCreditEpInfoGetRequest) GetFileParams() *map[string]string {
	return m.FileParams
}

func (m *ZhimaCreditEpInfoGetRequest) SetApiVersion(ApiVersion string) {
	m.ApiVersion = ApiVersion
}

func (m *ZhimaCreditEpInfoGetRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaCreditEpInfoGetRequest) SetScene(Scene string) {
	m.Scene = Scene
}

func (m *ZhimaCreditEpInfoGetRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaCreditEpInfoGetRequest) SetChannel(Channel string) {
	m.Channel = Channel
}

func (m *ZhimaCreditEpInfoGetRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaCreditEpInfoGetRequest) SetPlatform(Platform string) {
	m.Platform = Platform
}

func (m *ZhimaCreditEpInfoGetRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaCreditEpInfoGetRequest) SetExtParams(ExtParams string) {
	m.ExtParams = ExtParams
}

func (m *ZhimaCreditEpInfoGetRequest) GetExtParams() string {
	return m.ExtParams
}
