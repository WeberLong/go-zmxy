package requests

import (
	"github.com/WeberLong/go-zmxy/zmop/interfaces"
)

type ZhimaCreditKkcreditAbcscoreQueryRequest struct {
	age                     string // 年龄
	crdAgeUnclsAvg          string // 未销信用卡开户距今月数的平均数
	gender                  string // 性别，男=1，女=0
	lnizedLnitCttPpl        string // 近90天深夜联系人
	normCdtBalUsedPctAvg    string // 当前正常的信用卡账户已用额度与可用额度之比的均值
	openId                  string // 芝麻会员在商户端的身份标识。
	phoneUseMth             string // 手机注册时长
	productCode             string // 产品码
	smsLonfizedSendPpl      string // 近150天短信发送联系人
	transactionId           string // 商户传入的业务流水号。此字段由商户生成，需确保唯一性，用于定位每一次请求，后续按此流水进行对帐。生成规则: 固定30位数字串，前17位为精确到毫秒的时间yyyyMMddhhmmssSSS，后13位为自增数字。
	trcLsmfiAvgPlanTotalPct string // 近5月套餐金额占比

	interfaces.ZhimaRequestParams
}

func (m *ZhimaCreditKkcreditAbcscoreQueryRequest) InitBizParams(age, crdAgeUnclsAvg, gender, lnizedLnitCttPpl, normCdtBalUsedPctAvg, openId, phoneUseMth, productCode, smsLonfizedSendPpl, transactionId, trcLsmfiAvgPlanTotalPct string) {
	m.ApiParams = &map[string]string{}
	m.FileParams = &map[string]string{}

	(*m.ApiParams)["age"] = age
	m.age = age

	(*m.ApiParams)["crd_age_uncls_avg"] = crdAgeUnclsAvg
	m.crdAgeUnclsAvg = crdAgeUnclsAvg

	(*m.ApiParams)["gender"] = gender
	m.gender = gender

	(*m.ApiParams)["lnized_lnit_ctt_ppl"] = lnizedLnitCttPpl
	m.lnizedLnitCttPpl = lnizedLnitCttPpl

	(*m.ApiParams)["norm_cdt_bal_used_pct_avg"] = normCdtBalUsedPctAvg
	m.normCdtBalUsedPctAvg = normCdtBalUsedPctAvg

	(*m.ApiParams)["open_id"] = openId
	m.openId = openId

	(*m.ApiParams)["phone_use_mth"] = phoneUseMth
	m.phoneUseMth = phoneUseMth

	(*m.ApiParams)["product_code"] = productCode
	m.productCode = productCode

	(*m.ApiParams)["sms_lonfized_send_ppl"] = smsLonfizedSendPpl
	m.smsLonfizedSendPpl = smsLonfizedSendPpl

	(*m.ApiParams)["transaction_id"] = transactionId
	m.transactionId = transactionId

	(*m.ApiParams)["trc_lsmfi_avg_plan_total_pct"] = trcLsmfiAvgPlanTotalPct
	m.trcLsmfiAvgPlanTotalPct = trcLsmfiAvgPlanTotalPct
}

func (*ZhimaCreditKkcreditAbcscoreQueryRequest) GetApiMethodName() string {
	return "zhima.credit.kkcredit.abcscore.query"
}

func (m *ZhimaCreditKkcreditAbcscoreQueryRequest) GetApiParams() *map[string]string {
	return m.ApiParams
}

func (m *ZhimaCreditKkcreditAbcscoreQueryRequest) GetFileParams() *map[string]string {
	return m.FileParams
}

func (m *ZhimaCreditKkcreditAbcscoreQueryRequest) SetApiVersion(ApiVersion string) {
	m.ApiVersion = ApiVersion
}

func (m *ZhimaCreditKkcreditAbcscoreQueryRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaCreditKkcreditAbcscoreQueryRequest) SetScene(Scene string) {
	m.Scene = Scene
}

func (m *ZhimaCreditKkcreditAbcscoreQueryRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaCreditKkcreditAbcscoreQueryRequest) SetChannel(Channel string) {
	m.Channel = Channel
}

func (m *ZhimaCreditKkcreditAbcscoreQueryRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaCreditKkcreditAbcscoreQueryRequest) SetPlatform(Platform string) {
	m.Platform = Platform
}

func (m *ZhimaCreditKkcreditAbcscoreQueryRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaCreditKkcreditAbcscoreQueryRequest) SetExtParams(ExtParams string) {
	m.ExtParams = ExtParams
}

func (m *ZhimaCreditKkcreditAbcscoreQueryRequest) GetExtParams() string {
	return m.ExtParams
}
