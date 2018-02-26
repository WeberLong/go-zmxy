package requests

import (
	"github.com/WeberLong/go-zmxy/zmop/interfaces"
)

type ZhimaCreditContactGetRequest struct {
	address       string // 用户的地址：最多传入三组，用|分隔。地址中不能有如下特殊字符&,^,\
	isOverdue     string // 是否逾期。T表示逾期，F表示非逾期
	mobile        string // 用户手机号，最多传入三个，每个手机号之间以“|"分隔。
	openId        string // 芝麻会员在商户端的身份标识。
	overdueDays   string // 逾期天数。请输入数字。
	productCode   string // 产品码
	transactionId string // 商户传入的业务流水号。此字段由商户生成，需确保唯一性，用于定位每一次请求，后续按此流水进行对帐。生成规则: 固定30位数字串，前17位为精确到毫秒的时间yyyyMMddhhmmssSSS，后13位为自增数字。

	interfaces.ZhimaRequestParams
}

func (m *ZhimaCreditContactGetRequest) InitBizParams(address, isOverdue, mobile, openId, overdueDays, productCode, transactionId string) {
	m.ApiParams = &map[string]string{}
	m.FileParams = &map[string]string{}

	(*m.ApiParams)["address"] = address
	m.address = address

	(*m.ApiParams)["is_overdue"] = isOverdue
	m.isOverdue = isOverdue

	(*m.ApiParams)["mobile"] = mobile
	m.mobile = mobile

	(*m.ApiParams)["open_id"] = openId
	m.openId = openId

	(*m.ApiParams)["overdue_days"] = overdueDays
	m.overdueDays = overdueDays

	(*m.ApiParams)["product_code"] = productCode
	m.productCode = productCode

	(*m.ApiParams)["transaction_id"] = transactionId
	m.transactionId = transactionId
}

func (*ZhimaCreditContactGetRequest) GetApiMethodName() string {
	return "zhima.credit.contact.get"
}

func (m *ZhimaCreditContactGetRequest) GetApiParams() *map[string]string {
	return m.ApiParams
}

func (m *ZhimaCreditContactGetRequest) GetFileParams() *map[string]string {
	return m.FileParams
}

func (m *ZhimaCreditContactGetRequest) SetApiVersion(ApiVersion string) {
	m.ApiVersion = ApiVersion
}

func (m *ZhimaCreditContactGetRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaCreditContactGetRequest) SetScene(Scene string) {
	m.Scene = Scene
}

func (m *ZhimaCreditContactGetRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaCreditContactGetRequest) SetChannel(Channel string) {
	m.Channel = Channel
}

func (m *ZhimaCreditContactGetRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaCreditContactGetRequest) SetPlatform(Platform string) {
	m.Platform = Platform
}

func (m *ZhimaCreditContactGetRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaCreditContactGetRequest) SetExtParams(ExtParams string) {
	m.ExtParams = ExtParams
}

func (m *ZhimaCreditContactGetRequest) GetExtParams() string {
	return m.ExtParams
}
