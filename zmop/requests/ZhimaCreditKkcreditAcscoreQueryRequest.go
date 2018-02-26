package requests

import (
	"github.com/WeberLong/go-zmxy/zmop/interfaces"
)

type ZhimaCreditKkcreditAcscoreQueryRequest struct {
	lnizedLnitCttPpl        string // 近90天深夜联系人
	lonfizedAnsCttDay       string // 近150天被叫通话天数
	lonfizedRgCttTm         string // 近150天固话通话时长
	lontwzedWeekCttPplPct   string // 近120天工作日通话联系人占比
	openId                  string // 芝麻会员在商户端的身份标识。
	phoneUseMth             string // 手机注册时长
	productCode             string // 产品码
	smsLonfizedSendPpl      string // 近150天短信发送联系人
	transactionId           string // 商户传入的业务流水号。此字段由商户生成，需确保唯一性，用于定位每一次请求，后续按此流水进行对帐。生成规则: 固定30位数字串，前17位为精确到毫秒的时间yyyyMMddhhmmssSSS，后13位为自增数字。
	trcLsmfiAvgPlanTotalPct string // 近5月内平均套餐金额占比

	interfaces.ZhimaRequestParams
}

func (m *ZhimaCreditKkcreditAcscoreQueryRequest) InitBizParams(lnizedLnitCttPpl, lonfizedAnsCttDay, lonfizedRgCttTm, lontwzedWeekCttPplPct, openId, phoneUseMth, productCode, smsLonfizedSendPpl, transactionId, trcLsmfiAvgPlanTotalPct string) {
	m.ApiParams = &map[string]string{}
	m.FileParams = &map[string]string{}

	(*m.ApiParams)["lnized_lnit_ctt_ppl"] = lnizedLnitCttPpl
	m.lnizedLnitCttPpl = lnizedLnitCttPpl

	(*m.ApiParams)["lonfized_ans_ctt_day"] = lonfizedAnsCttDay
	m.lonfizedAnsCttDay = lonfizedAnsCttDay

	(*m.ApiParams)["lonfized_rg_ctt_tm"] = lonfizedRgCttTm
	m.lonfizedRgCttTm = lonfizedRgCttTm

	(*m.ApiParams)["lontwzed_week_ctt_ppl_pct"] = lontwzedWeekCttPplPct
	m.lontwzedWeekCttPplPct = lontwzedWeekCttPplPct

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

func (*ZhimaCreditKkcreditAcscoreQueryRequest) GetApiMethodName() string {
	return "zhima.credit.kkcredit.acscore.query"
}

func (m *ZhimaCreditKkcreditAcscoreQueryRequest) GetApiParams() *map[string]string {
	return m.ApiParams
}

func (m *ZhimaCreditKkcreditAcscoreQueryRequest) GetFileParams() *map[string]string {
	return m.FileParams
}

func (m *ZhimaCreditKkcreditAcscoreQueryRequest) SetApiVersion(ApiVersion string) {
	m.ApiVersion = ApiVersion
}

func (m *ZhimaCreditKkcreditAcscoreQueryRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaCreditKkcreditAcscoreQueryRequest) SetScene(Scene string) {
	m.Scene = Scene
}

func (m *ZhimaCreditKkcreditAcscoreQueryRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaCreditKkcreditAcscoreQueryRequest) SetChannel(Channel string) {
	m.Channel = Channel
}

func (m *ZhimaCreditKkcreditAcscoreQueryRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaCreditKkcreditAcscoreQueryRequest) SetPlatform(Platform string) {
	m.Platform = Platform
}

func (m *ZhimaCreditKkcreditAcscoreQueryRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaCreditKkcreditAcscoreQueryRequest) SetExtParams(ExtParams string) {
	m.ExtParams = ExtParams
}

func (m *ZhimaCreditKkcreditAcscoreQueryRequest) GetExtParams() string {
	return m.ExtParams
}
