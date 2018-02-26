package requests

import (
	"github.com/WeberLong/go-zmxy/zmop/interfaces"
)

type ZhimaCreditAntifraudRiskListRequest struct {
	address       string // 地址信息。省+市+区/县+详细地址，长度不超过256，不要包含特殊字符，如","，"\"，"|"，"&"，"^"
	bankCard      string // 银行卡号。中国大陆银行发布的银行卡:借记卡长度19位；信用卡长度16位；各位的取值位[0,9]的整数；不含虚拟卡
	certNo        string // 证件号。证件类型、证件号、姓名三要素均为必填参数
	certType      string // 证件类型。IDENTITY_CARD标识为身份证，目前仅支持身份证类型
	email         string // 电子邮箱。合法email，字母小写，特殊符号以半角形式出现
	imei          string // 国际移动设备标志。15位长度数字
	ip            string // ip地址。以"."分割的四段Ip，如 x.x.x.x，x为[0,255]之间的整数
	mac           string // 物理地址。支持格式如下：xx:xx:xx:xx:xx:xx，xx-xx-xx-xx-xx-xx，xxxxxxxxxxxx，x取值范围[0,9]之间的整数及A，B，C，D，E，F
	mobile        string // 手机号码。中国大陆合法手机号，长度11位，不含国家代码
	name          string // 姓名。长度不超过64，姓名中不要包含特殊字符，如","，"\"，"|"，"&"，"^"
	productCode   string // 产品码。标记商户接入的具体产品；直接使用每个接口入参中示例值即可
	transactionId string // 商户请求的唯一标志，长度64位以内字符串，仅限字母数字下划线组合。该标识作为业务调用的唯一标识，商户要保证其业务唯一性，使用相同transaction_id的查询，芝麻在一段时间内（一般为1天）返回首次查询结果，超过有效期的查询即为无效并返回异常，有效期内的重复查询不重新计费
	wifimac       string // wifi的物理地址。支持格式如下：xx:xx:xx:xx:xx:xx，xx-xx-xx-xx-xx-xx，xxxxxxxxxxxx，x取值范围[0,9]之间的整数及A，B，C，D，E，F

	interfaces.ZhimaRequestParams
}

func (m *ZhimaCreditAntifraudRiskListRequest) InitBizParams(address, bankCard, certNo, certType, email, imei, ip, mac, mobile, name, productCode, transactionId, wifimac string) {
	m.ApiParams = &map[string]string{}
	m.FileParams = &map[string]string{}

	(*m.ApiParams)["address"] = address
	m.address = address

	(*m.ApiParams)["bank_card"] = bankCard
	m.bankCard = bankCard

	(*m.ApiParams)["cert_no"] = certNo
	m.certNo = certNo

	(*m.ApiParams)["cert_type"] = certType
	m.certType = certType

	(*m.ApiParams)["email"] = email
	m.email = email

	(*m.ApiParams)["imei"] = imei
	m.imei = imei

	(*m.ApiParams)["ip"] = ip
	m.ip = ip

	(*m.ApiParams)["mac"] = mac
	m.mac = mac

	(*m.ApiParams)["mobile"] = mobile
	m.mobile = mobile

	(*m.ApiParams)["name"] = name
	m.name = name

	(*m.ApiParams)["product_code"] = productCode
	m.productCode = productCode

	(*m.ApiParams)["transaction_id"] = transactionId
	m.transactionId = transactionId

	(*m.ApiParams)["wifimac"] = wifimac
	m.wifimac = wifimac
}

func (*ZhimaCreditAntifraudRiskListRequest) GetApiMethodName() string {
	return "zhima.credit.antifraud.risk.list"
}

func (m *ZhimaCreditAntifraudRiskListRequest) GetApiParams() *map[string]string {
	return m.ApiParams
}

func (m *ZhimaCreditAntifraudRiskListRequest) GetFileParams() *map[string]string {
	return m.FileParams
}

func (m *ZhimaCreditAntifraudRiskListRequest) SetApiVersion(ApiVersion string) {
	m.ApiVersion = ApiVersion
}

func (m *ZhimaCreditAntifraudRiskListRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaCreditAntifraudRiskListRequest) SetScene(Scene string) {
	m.Scene = Scene
}

func (m *ZhimaCreditAntifraudRiskListRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaCreditAntifraudRiskListRequest) SetChannel(Channel string) {
	m.Channel = Channel
}

func (m *ZhimaCreditAntifraudRiskListRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaCreditAntifraudRiskListRequest) SetPlatform(Platform string) {
	m.Platform = Platform
}

func (m *ZhimaCreditAntifraudRiskListRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaCreditAntifraudRiskListRequest) SetExtParams(ExtParams string) {
	m.ExtParams = ExtParams
}

func (m *ZhimaCreditAntifraudRiskListRequest) GetExtParams() string {
	return m.ExtParams
}
