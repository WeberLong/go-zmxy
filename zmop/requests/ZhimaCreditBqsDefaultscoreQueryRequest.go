package requests

import (
	"github.com/WeberLong/go-zmxy/zmop/interfaces"
)

type ZhimaCreditBqsDefaultscoreQueryRequest struct {
	acceptPercentApply    string // 申请事件通过比率
	age                   string // 年龄
	applyHour             string // 申请时间（小时，0-23）
	applyPartnerTypeCount string // 多头申请商户类型数量
	blackCount            string // 黑名单命中个数
	callActiveArea        string // 本人主要通话活动区域在几线城市
	contactExcludedCount  string // 排除被叫电话很短的联系人个数
	contactsActiveArea    string // 朋友圈活动区域在几线城市
	deviceCount           string // 关联设备数量
	gender                string // 性别
	gpsCityCount          string // GPS城市数量
	inactiveDays          string // 全天未使用通话和短信功能天数
	ipCityCount           string // IP城市数量
	loanAppCount          string // 设备中借贷app数量
	mobile                string // 手机号
	multiapplyCount       string // 多头申请商户数量
	nightCalls            string // 夜间通话次数
	noneMobileCount       string // 联系人中非手机个数
	onlyTerminCount       string // 仅有被叫联系人个数
	openDays              string // 入网时长
	openId                string // 用户在商端的身份标识
	phoneDays             string // 该用户第一次事件距今时间
	productCode           string // 信用产品码，对应云产品的标识
	provinceId            string // 省份代码
	rejectPercentApply    string // 申请事件拒绝比率
	sumInfoCostMoney      string // 话费消费总金额
	topContact            string // 最常用联系人，多个用逗号分隔
	transactionId         string // 代表一笔请求的唯一标志，该标识作为对账的关键信息，对于用户使用相同transaction_id的查询，芝麻在一天（86400秒）内返回首次查询数据，超过有效期的查询即为无效并返回异常，有效期内的反复查询不重新计费。transaction_id 推荐生成方式是：30位，（其中17位时间值（精确到毫秒）：yyyyMMddHHmmssSSS）加上（13位自增数字：1234567890123）
	whiteGrade            string // 白名单等级
	workCityCount         string // 上班时间手机号关联城市数量

	interfaces.ZhimaRequestParams
}

func (m *ZhimaCreditBqsDefaultscoreQueryRequest) InitBizParams(acceptPercentApply, age, applyHour, applyPartnerTypeCount, blackCount, callActiveArea, contactExcludedCount, contactsActiveArea, deviceCount, gender, gpsCityCount, inactiveDays, ipCityCount, loanAppCount, mobile, multiapplyCount, nightCalls, noneMobileCount, onlyTerminCount, openDays, openId, phoneDays, productCode, provinceId, rejectPercentApply, sumInfoCostMoney, topContact, transactionId, whiteGrade, workCityCount string) {
	m.ApiParams = &map[string]string{}
	m.FileParams = &map[string]string{}

	(*m.ApiParams)["accept_percent_apply"] = acceptPercentApply
	m.acceptPercentApply = acceptPercentApply

	(*m.ApiParams)["age"] = age
	m.age = age

	(*m.ApiParams)["apply_hour"] = applyHour
	m.applyHour = applyHour

	(*m.ApiParams)["apply_partner_type_count"] = applyPartnerTypeCount
	m.applyPartnerTypeCount = applyPartnerTypeCount

	(*m.ApiParams)["black_count"] = blackCount
	m.blackCount = blackCount

	(*m.ApiParams)["call_active_area"] = callActiveArea
	m.callActiveArea = callActiveArea

	(*m.ApiParams)["contact_excluded_count"] = contactExcludedCount
	m.contactExcludedCount = contactExcludedCount

	(*m.ApiParams)["contacts_active_area"] = contactsActiveArea
	m.contactsActiveArea = contactsActiveArea

	(*m.ApiParams)["device_count"] = deviceCount
	m.deviceCount = deviceCount

	(*m.ApiParams)["gender"] = gender
	m.gender = gender

	(*m.ApiParams)["gps_city_count"] = gpsCityCount
	m.gpsCityCount = gpsCityCount

	(*m.ApiParams)["inactive_days"] = inactiveDays
	m.inactiveDays = inactiveDays

	(*m.ApiParams)["ip_city_count"] = ipCityCount
	m.ipCityCount = ipCityCount

	(*m.ApiParams)["loan_app_count"] = loanAppCount
	m.loanAppCount = loanAppCount

	(*m.ApiParams)["mobile"] = mobile
	m.mobile = mobile

	(*m.ApiParams)["multiapply_count"] = multiapplyCount
	m.multiapplyCount = multiapplyCount

	(*m.ApiParams)["night_calls"] = nightCalls
	m.nightCalls = nightCalls

	(*m.ApiParams)["none_mobile_count"] = noneMobileCount
	m.noneMobileCount = noneMobileCount

	(*m.ApiParams)["only_termin_count"] = onlyTerminCount
	m.onlyTerminCount = onlyTerminCount

	(*m.ApiParams)["open_days"] = openDays
	m.openDays = openDays

	(*m.ApiParams)["open_id"] = openId
	m.openId = openId

	(*m.ApiParams)["phone_days"] = phoneDays
	m.phoneDays = phoneDays

	(*m.ApiParams)["product_code"] = productCode
	m.productCode = productCode

	(*m.ApiParams)["province_id"] = provinceId
	m.provinceId = provinceId

	(*m.ApiParams)["reject_percent_apply"] = rejectPercentApply
	m.rejectPercentApply = rejectPercentApply

	(*m.ApiParams)["sum_info_cost_money"] = sumInfoCostMoney
	m.sumInfoCostMoney = sumInfoCostMoney

	(*m.ApiParams)["top_contact"] = topContact
	m.topContact = topContact

	(*m.ApiParams)["transaction_id"] = transactionId
	m.transactionId = transactionId

	(*m.ApiParams)["white_grade"] = whiteGrade
	m.whiteGrade = whiteGrade

	(*m.ApiParams)["work_city_count"] = workCityCount
	m.workCityCount = workCityCount
}

func (*ZhimaCreditBqsDefaultscoreQueryRequest) GetApiMethodName() string {
	return "zhima.credit.bqs.defaultscore.query"
}

func (m *ZhimaCreditBqsDefaultscoreQueryRequest) GetApiParams() *map[string]string {
	return m.ApiParams
}

func (m *ZhimaCreditBqsDefaultscoreQueryRequest) GetFileParams() *map[string]string {
	return m.FileParams
}

func (m *ZhimaCreditBqsDefaultscoreQueryRequest) SetApiVersion(ApiVersion string) {
	m.ApiVersion = ApiVersion
}

func (m *ZhimaCreditBqsDefaultscoreQueryRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaCreditBqsDefaultscoreQueryRequest) SetScene(Scene string) {
	m.Scene = Scene
}

func (m *ZhimaCreditBqsDefaultscoreQueryRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaCreditBqsDefaultscoreQueryRequest) SetChannel(Channel string) {
	m.Channel = Channel
}

func (m *ZhimaCreditBqsDefaultscoreQueryRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaCreditBqsDefaultscoreQueryRequest) SetPlatform(Platform string) {
	m.Platform = Platform
}

func (m *ZhimaCreditBqsDefaultscoreQueryRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaCreditBqsDefaultscoreQueryRequest) SetExtParams(ExtParams string) {
	m.ExtParams = ExtParams
}

func (m *ZhimaCreditBqsDefaultscoreQueryRequest) GetExtParams() string {
	return m.ExtParams
}
