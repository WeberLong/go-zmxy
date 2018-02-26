package requests

import (
	"github.com/WeberLong/go-zmxy/zmop/interfaces"
)

type ZhimaCreditDriverVerifyRequest struct {
	archiveNo     string // 驾驶证档案编号
	issueDate     string // 初次领证日期，格式为yyyyMMdd
	licenseNo     string // 驾驶证号码
	name          string // 驾驶证上的姓名
	productCode   string // 产品码
	transactionId string // 芝麻业务凭证，详见https://b.zmxy.com.cn/technology/openDoc.htm?id=334
	vehicleClass  string // 准驾车型，多个车型之间以 || 隔开，如C1 || C2 || B2

	interfaces.ZhimaRequestParams
}

func (m *ZhimaCreditDriverVerifyRequest) InitBizParams(archiveNo, issueDate, licenseNo, name, productCode, transactionId, vehicleClass string) {
	m.ApiParams = &map[string]string{}
	m.FileParams = &map[string]string{}

	(*m.ApiParams)["archive_no"] = archiveNo
	m.archiveNo = archiveNo

	(*m.ApiParams)["issue_date"] = issueDate
	m.issueDate = issueDate

	(*m.ApiParams)["license_no"] = licenseNo
	m.licenseNo = licenseNo

	(*m.ApiParams)["name"] = name
	m.name = name

	(*m.ApiParams)["product_code"] = productCode
	m.productCode = productCode

	(*m.ApiParams)["transaction_id"] = transactionId
	m.transactionId = transactionId

	(*m.ApiParams)["vehicle_class"] = vehicleClass
	m.vehicleClass = vehicleClass
}

func (*ZhimaCreditDriverVerifyRequest) GetApiMethodName() string {
	return "zhima.credit.driver.verify"
}

func (m *ZhimaCreditDriverVerifyRequest) GetApiParams() *map[string]string {
	return m.ApiParams
}

func (m *ZhimaCreditDriverVerifyRequest) GetFileParams() *map[string]string {
	return m.FileParams
}

func (m *ZhimaCreditDriverVerifyRequest) SetApiVersion(ApiVersion string) {
	m.ApiVersion = ApiVersion
}

func (m *ZhimaCreditDriverVerifyRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaCreditDriverVerifyRequest) SetScene(Scene string) {
	m.Scene = Scene
}

func (m *ZhimaCreditDriverVerifyRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaCreditDriverVerifyRequest) SetChannel(Channel string) {
	m.Channel = Channel
}

func (m *ZhimaCreditDriverVerifyRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaCreditDriverVerifyRequest) SetPlatform(Platform string) {
	m.Platform = Platform
}

func (m *ZhimaCreditDriverVerifyRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaCreditDriverVerifyRequest) SetExtParams(ExtParams string) {
	m.ExtParams = ExtParams
}

func (m *ZhimaCreditDriverVerifyRequest) GetExtParams() string {
	return m.ExtParams
}
