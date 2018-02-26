package requests

import (
	"github.com/WeberLong/go-zmxy/zmop/interfaces"
)

type ZhimaCreditShareriskGetRequest struct {
	bizApplyNo    string // 业务申请单号
	bizAuthNo     string // 授权合同编号
	bizScene      string // 业务场景[01： 申贷审批； 02： 贷后管理]
	certNo        string // 证件号码
	certType      string // 证件类型[100：身份证]
	name          string // 姓名
	productCode   string // 产品码
	transactionId string // 商户传入的业务流水号。此字段由商户生成，需确保唯一性，用于定位每一次请求，后续按此流水进行对帐。生成规则: 固定30位数字串，前17位为精确到毫秒的时间yyyyMMddhhmmssSSS，后13位为自增数字。

	interfaces.ZhimaRequestParams
}

func (m *ZhimaCreditShareriskGetRequest) InitBizParams(bizApplyNo, bizAuthNo, bizScene, certNo, certType, name, productCode, transactionId string) {
	m.ApiParams = &map[string]string{}
	m.FileParams = &map[string]string{}

	(*m.ApiParams)["biz_apply_no"] = bizApplyNo
	m.bizApplyNo = bizApplyNo

	(*m.ApiParams)["biz_auth_no"] = bizAuthNo
	m.bizAuthNo = bizAuthNo

	(*m.ApiParams)["biz_scene"] = bizScene
	m.bizScene = bizScene

	(*m.ApiParams)["cert_no"] = certNo
	m.certNo = certNo

	(*m.ApiParams)["cert_type"] = certType
	m.certType = certType

	(*m.ApiParams)["name"] = name
	m.name = name

	(*m.ApiParams)["product_code"] = productCode
	m.productCode = productCode

	(*m.ApiParams)["transaction_id"] = transactionId
	m.transactionId = transactionId
}

func (*ZhimaCreditShareriskGetRequest) GetApiMethodName() string {
	return "zhima.credit.sharerisk.get"
}

func (m *ZhimaCreditShareriskGetRequest) GetApiParams() *map[string]string {
	return m.ApiParams
}

func (m *ZhimaCreditShareriskGetRequest) GetFileParams() *map[string]string {
	return m.FileParams
}

func (m *ZhimaCreditShareriskGetRequest) SetApiVersion(ApiVersion string) {
	m.ApiVersion = ApiVersion
}

func (m *ZhimaCreditShareriskGetRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaCreditShareriskGetRequest) SetScene(Scene string) {
	m.Scene = Scene
}

func (m *ZhimaCreditShareriskGetRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaCreditShareriskGetRequest) SetChannel(Channel string) {
	m.Channel = Channel
}

func (m *ZhimaCreditShareriskGetRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaCreditShareriskGetRequest) SetPlatform(Platform string) {
	m.Platform = Platform
}

func (m *ZhimaCreditShareriskGetRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaCreditShareriskGetRequest) SetExtParams(ExtParams string) {
	m.ExtParams = ExtParams
}

func (m *ZhimaCreditShareriskGetRequest) GetExtParams() string {
	return m.ExtParams
}
