package requests

import (
	"github.com/WeberLong/go-zmxy/zmop/interfaces"
)

type ZhimaCreditIvsGetRequest struct {
	address       string // 用户地址，最多输入三个，多个地址之间用|分隔，如 杭州市西湖区天目山路266号|杭州市西湖区万塘路999号。备注：证件号、姓名、手机号、地址、银行卡、电子邮箱至少传其中两项
	bankCard      string // 银行卡号，最多输入两个，多个银行卡号之间用|分隔，如602436748024138|622536748024139。备注：证件号、姓名、手机号、地址、银行卡、电子邮箱至少传其中两项
	certNo        string // 证件号。备注：证件号、姓名、手机号、地址、银行卡、电子邮箱至少传其中两项
	certType      string // 证件类型。备注：证件号、姓名、手机号、地址、银行卡、电子邮箱至少传其中两项
	email         string // 电子邮箱，最多传入两个，多个邮箱之间用|分隔，如jnlxhy@alitest.com|john.sss@alitest.com。备注：证件号、姓名、手机号、地址、银行卡、电子邮箱至少传其中两项
	imei          string // 国际移动设备标志
	imsi          string // 国际移动用户识别码
	ip            string // ip地址
	mac           string // 物理地址
	mobile        string // 手机号，最多传入三个，多个手机号之间用|分隔，如15256797367|18669152789。备注：证件号、姓名、手机号、地址、银行卡、电子邮箱至少传其中两项
	name          string // 姓名。备注：证件号、姓名、手机号、地址、银行卡、电子邮箱至少传其中两项
	productCode   string // 产品码
	transactionId string // 商户传入的业务流水号。此字段由商户生成，需确保唯一性，用于定位每一次请求，后续按此流水进行对帐。参考生成规则: 固定30位数字串，前17位为精确到毫秒的时间yyyyMMddHHmmssSSS，后13位为自增数字。
	wifimac       string // wifi的物理地址

	interfaces.ZhimaRequestParams
}

func (m *ZhimaCreditIvsGetRequest) InitBizParams(address, bankCard, certNo, certType, email, imei, imsi, ip, mac, mobile, name, productCode, transactionId, wifimac string) {
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

	(*m.ApiParams)["imsi"] = imsi
	m.imsi = imsi

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

func (*ZhimaCreditIvsGetRequest) GetApiMethodName() string {
	return "zhima.credit.ivs.get"
}

func (m *ZhimaCreditIvsGetRequest) GetApiParams() *map[string]string {
	return m.ApiParams
}

func (m *ZhimaCreditIvsGetRequest) GetFileParams() *map[string]string {
	return m.FileParams
}

func (m *ZhimaCreditIvsGetRequest) SetApiVersion(ApiVersion string) {
	m.ApiVersion = ApiVersion
}

func (m *ZhimaCreditIvsGetRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaCreditIvsGetRequest) SetScene(Scene string) {
	m.Scene = Scene
}

func (m *ZhimaCreditIvsGetRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaCreditIvsGetRequest) SetChannel(Channel string) {
	m.Channel = Channel
}

func (m *ZhimaCreditIvsGetRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaCreditIvsGetRequest) SetPlatform(Platform string) {
	m.Platform = Platform
}

func (m *ZhimaCreditIvsGetRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaCreditIvsGetRequest) SetExtParams(ExtParams string) {
	m.ExtParams = ExtParams
}

func (m *ZhimaCreditIvsGetRequest) GetExtParams() string {
	return m.ExtParams
}
