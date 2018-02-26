package requests

import (
	"github.com/WeberLong/go-zmxy/zmop/interfaces"
)

type ZhimaCreditHetroneDasscoreQueryRequest struct {
	amtBankcardTransacThreeMonths  string // 近3月交易总金额
	cntBankcardTransacTwelveMonths string // 近十二个月有交易的月数
	cntMobileOnline                string // 手机在网时长
	contactScore                   string // 通讯录分数
	existsBankcardTransacOversea   string // 最近有无境外消费
	gender                         string // 性别
	openId                         string // 用户在商端的身份标识
	productCode                    string // 信用产品码，对应云产品的标识
	transactionId                  string // 代表一笔请求的唯一标志，该标识作为对账的关键信息，对于用户使用相同transaction_id的查询，芝麻在一天（86400秒）内返回首次查询数据，超过有效期的查询即为无效并返回异常，有效期内的反复查询不重新计费。transaction_id 推荐生成方式是：30位，（其中17位时间值（精确到毫秒）：yyyyMMddHHmmssSSS）加上（13位自增数字：1234567890123）

	interfaces.ZhimaRequestParams
}

func (m *ZhimaCreditHetroneDasscoreQueryRequest) InitBizParams(amtBankcardTransacThreeMonths, cntBankcardTransacTwelveMonths, cntMobileOnline, contactScore, existsBankcardTransacOversea, gender, openId, productCode, transactionId string) {
	m.ApiParams = &map[string]string{}
	m.FileParams = &map[string]string{}

	(*m.ApiParams)["amt_bankcard_transac_three_months"] = amtBankcardTransacThreeMonths
	m.amtBankcardTransacThreeMonths = amtBankcardTransacThreeMonths

	(*m.ApiParams)["cnt_bankcard_transac_twelve_months"] = cntBankcardTransacTwelveMonths
	m.cntBankcardTransacTwelveMonths = cntBankcardTransacTwelveMonths

	(*m.ApiParams)["cnt_mobile_online"] = cntMobileOnline
	m.cntMobileOnline = cntMobileOnline

	(*m.ApiParams)["contact_score"] = contactScore
	m.contactScore = contactScore

	(*m.ApiParams)["exists_bankcard_transac_oversea"] = existsBankcardTransacOversea
	m.existsBankcardTransacOversea = existsBankcardTransacOversea

	(*m.ApiParams)["gender"] = gender
	m.gender = gender

	(*m.ApiParams)["open_id"] = openId
	m.openId = openId

	(*m.ApiParams)["product_code"] = productCode
	m.productCode = productCode

	(*m.ApiParams)["transaction_id"] = transactionId
	m.transactionId = transactionId
}

func (*ZhimaCreditHetroneDasscoreQueryRequest) GetApiMethodName() string {
	return "zhima.credit.hetrone.dasscore.query"
}

func (m *ZhimaCreditHetroneDasscoreQueryRequest) GetApiParams() *map[string]string {
	return m.ApiParams
}

func (m *ZhimaCreditHetroneDasscoreQueryRequest) GetFileParams() *map[string]string {
	return m.FileParams
}

func (m *ZhimaCreditHetroneDasscoreQueryRequest) SetApiVersion(ApiVersion string) {
	m.ApiVersion = ApiVersion
}

func (m *ZhimaCreditHetroneDasscoreQueryRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaCreditHetroneDasscoreQueryRequest) SetScene(Scene string) {
	m.Scene = Scene
}

func (m *ZhimaCreditHetroneDasscoreQueryRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaCreditHetroneDasscoreQueryRequest) SetChannel(Channel string) {
	m.Channel = Channel
}

func (m *ZhimaCreditHetroneDasscoreQueryRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaCreditHetroneDasscoreQueryRequest) SetPlatform(Platform string) {
	m.Platform = Platform
}

func (m *ZhimaCreditHetroneDasscoreQueryRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaCreditHetroneDasscoreQueryRequest) SetExtParams(ExtParams string) {
	m.ExtParams = ExtParams
}

func (m *ZhimaCreditHetroneDasscoreQueryRequest) GetExtParams() string {
	return m.ExtParams
}
