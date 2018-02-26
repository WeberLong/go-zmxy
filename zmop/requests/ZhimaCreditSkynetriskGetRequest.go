package requests

import (
	"github.com/WeberLong/go-zmxy/zmop/interfaces"
)

type ZhimaCreditSkynetriskGetRequest struct {
	alipayLogonId string // 支付宝登陆号
	certNo        string // 身份证号码
	contractFlag  string // 合约外标，服务标识。签约过后将会收到含有该服务标识的邮件，或者向协同您签约的芝麻合作人员索取。
	mobile        string // 手机号码
	name          string // 姓名
	principalType string // 主体类型和对应参数使用身份证信息查询填cert，对应cert_no和name参数必填；使用支付宝登陆账号查询填alipayLogonId，对应alipay_logon_id参数必填；使用支付宝用户ID查询填userId，对应user_id参数必填；使用手机号查询填mobile，对应mobile参数必填
	productCode   string // 产品码，固定为w1010100000000001000
	transactionId string // transaction_id是代表一笔请求的唯一标志，该标识作为对账的关键信息，对于用户使用相同transaction_id的查询，芝麻在一天（86400秒）内返回首次查询数据，超过有效期的查询即为无效并返回异常（错误码xxxx），有效期内的重复查询不重新计费。 transaction_id 推荐生成方式是：30位，（其中17位时间值（精确到毫秒）：yyyyMMddHHmmssSSS）加上（13位自增数字：1234567890123）
	userId        string // 支付宝账号ID

	interfaces.ZhimaRequestParams
}

func (m *ZhimaCreditSkynetriskGetRequest) InitBizParams(alipayLogonId, certNo, contractFlag, mobile, name, principalType, productCode, transactionId, userId string) {
	m.ApiParams = &map[string]string{}
	m.FileParams = &map[string]string{}

	(*m.ApiParams)["alipay_logon_id"] = alipayLogonId
	m.alipayLogonId = alipayLogonId

	(*m.ApiParams)["cert_no"] = certNo
	m.certNo = certNo

	(*m.ApiParams)["contract_flag"] = contractFlag
	m.contractFlag = contractFlag

	(*m.ApiParams)["mobile"] = mobile
	m.mobile = mobile

	(*m.ApiParams)["name"] = name
	m.name = name

	(*m.ApiParams)["principal_type"] = principalType
	m.principalType = principalType

	(*m.ApiParams)["product_code"] = productCode
	m.productCode = productCode

	(*m.ApiParams)["transaction_id"] = transactionId
	m.transactionId = transactionId

	(*m.ApiParams)["user_id"] = userId
	m.userId = userId
}

func (*ZhimaCreditSkynetriskGetRequest) GetApiMethodName() string {
	return "zhima.credit.skynetrisk.get"
}

func (m *ZhimaCreditSkynetriskGetRequest) GetApiParams() *map[string]string {
	return m.ApiParams
}

func (m *ZhimaCreditSkynetriskGetRequest) GetFileParams() *map[string]string {
	return m.FileParams
}

func (m *ZhimaCreditSkynetriskGetRequest) SetApiVersion(ApiVersion string) {
	m.ApiVersion = ApiVersion
}

func (m *ZhimaCreditSkynetriskGetRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaCreditSkynetriskGetRequest) SetScene(Scene string) {
	m.Scene = Scene
}

func (m *ZhimaCreditSkynetriskGetRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaCreditSkynetriskGetRequest) SetChannel(Channel string) {
	m.Channel = Channel
}

func (m *ZhimaCreditSkynetriskGetRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaCreditSkynetriskGetRequest) SetPlatform(Platform string) {
	m.Platform = Platform
}

func (m *ZhimaCreditSkynetriskGetRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaCreditSkynetriskGetRequest) SetExtParams(ExtParams string) {
	m.ExtParams = ExtParams
}

func (m *ZhimaCreditSkynetriskGetRequest) GetExtParams() string {
	return m.ExtParams
}
