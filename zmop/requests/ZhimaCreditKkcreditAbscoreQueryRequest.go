package requests

import (
	"github.com/WeberLong/go-zmxy/zmop/interfaces"
)

type ZhimaCreditKkcreditAbscoreQueryRequest struct {
	age                  string // 年龄
	crdAgeUnclsAvg       string // 未销信用卡开户距今月数的平均数
	gender               string // 性别，男=1，女=0
	normCdtBalUsedPctAvg string // 当前正常的信用卡账户已用额度与可用额度之比的均值
	openId               string // 芝麻会员在商户端的身份标识。
	productCode          string // 产品码
	transactionId        string // 商户传入的业务流水号。此字段由商户生成，需确保唯一性，用于定位每一次请求，后续按此流水进行对帐。生成规则: 固定30位数字串，前17位为精确到毫秒的时间yyyyMMddhhmmssSSS，后13位为自增数字。

	interfaces.ZhimaRequestParams
}

func (m *ZhimaCreditKkcreditAbscoreQueryRequest) InitBizParams(age, crdAgeUnclsAvg, gender, normCdtBalUsedPctAvg, openId, productCode, transactionId string) {
	m.ApiParams = &map[string]string{}
	m.FileParams = &map[string]string{}

	(*m.ApiParams)["age"] = age
	m.age = age

	(*m.ApiParams)["crd_age_uncls_avg"] = crdAgeUnclsAvg
	m.crdAgeUnclsAvg = crdAgeUnclsAvg

	(*m.ApiParams)["gender"] = gender
	m.gender = gender

	(*m.ApiParams)["norm_cdt_bal_used_pct_avg"] = normCdtBalUsedPctAvg
	m.normCdtBalUsedPctAvg = normCdtBalUsedPctAvg

	(*m.ApiParams)["open_id"] = openId
	m.openId = openId

	(*m.ApiParams)["product_code"] = productCode
	m.productCode = productCode

	(*m.ApiParams)["transaction_id"] = transactionId
	m.transactionId = transactionId
}

func (*ZhimaCreditKkcreditAbscoreQueryRequest) GetApiMethodName() string {
	return "zhima.credit.kkcredit.abscore.query"
}

func (m *ZhimaCreditKkcreditAbscoreQueryRequest) GetApiParams() *map[string]string {
	return m.ApiParams
}

func (m *ZhimaCreditKkcreditAbscoreQueryRequest) GetFileParams() *map[string]string {
	return m.FileParams
}

func (m *ZhimaCreditKkcreditAbscoreQueryRequest) SetApiVersion(ApiVersion string) {
	m.ApiVersion = ApiVersion
}

func (m *ZhimaCreditKkcreditAbscoreQueryRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaCreditKkcreditAbscoreQueryRequest) SetScene(Scene string) {
	m.Scene = Scene
}

func (m *ZhimaCreditKkcreditAbscoreQueryRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaCreditKkcreditAbscoreQueryRequest) SetChannel(Channel string) {
	m.Channel = Channel
}

func (m *ZhimaCreditKkcreditAbscoreQueryRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaCreditKkcreditAbscoreQueryRequest) SetPlatform(Platform string) {
	m.Platform = Platform
}

func (m *ZhimaCreditKkcreditAbscoreQueryRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaCreditKkcreditAbscoreQueryRequest) SetExtParams(ExtParams string) {
	m.ExtParams = ExtParams
}

func (m *ZhimaCreditKkcreditAbscoreQueryRequest) GetExtParams() string {
	return m.ExtParams
}
