package requests

import (
	"github.com/WeberLong/go-zmxy/zmop/interfaces"
)

type ZhimaCreditScoreCobuildGetRequest struct {
	applyDate     string // 申请日期。可空
	bizParams     string // 业务扩展参数。
	mobile        string // 手机号。可空
	openId        string // 芝麻会员在商户端的身份标识
	productCode   string // 产品码
	transactionId string // 商户传入的业务流水号。此字段由商户生成，需确保唯一性，用于定位每一次请求，后续按此流水进行对帐。生成规则: 固定30位数字串，前17位为精确到毫秒的时间yyyyMMddhhmmssSSS，后13位为自增数字。

	interfaces.ZhimaRequestParams
}

func (m *ZhimaCreditScoreCobuildGetRequest) InitBizParams(applyDate, bizParams, mobile, openId, productCode, transactionId string) {
	m.ApiParams = &map[string]string{}
	m.FileParams = &map[string]string{}

	(*m.ApiParams)["apply_date"] = applyDate
	m.applyDate = applyDate

	(*m.ApiParams)["biz_params"] = bizParams
	m.bizParams = bizParams

	(*m.ApiParams)["mobile"] = mobile
	m.mobile = mobile

	(*m.ApiParams)["open_id"] = openId
	m.openId = openId

	(*m.ApiParams)["product_code"] = productCode
	m.productCode = productCode

	(*m.ApiParams)["transaction_id"] = transactionId
	m.transactionId = transactionId
}

func (*ZhimaCreditScoreCobuildGetRequest) GetApiMethodName() string {
	return "zhima.credit.score.cobuild.get"
}

func (m *ZhimaCreditScoreCobuildGetRequest) GetApiParams() *map[string]string {
	return m.ApiParams
}

func (m *ZhimaCreditScoreCobuildGetRequest) GetFileParams() *map[string]string {
	return m.FileParams
}

func (m *ZhimaCreditScoreCobuildGetRequest) SetApiVersion(ApiVersion string) {
	m.ApiVersion = ApiVersion
}

func (m *ZhimaCreditScoreCobuildGetRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaCreditScoreCobuildGetRequest) SetScene(Scene string) {
	m.Scene = Scene
}

func (m *ZhimaCreditScoreCobuildGetRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaCreditScoreCobuildGetRequest) SetChannel(Channel string) {
	m.Channel = Channel
}

func (m *ZhimaCreditScoreCobuildGetRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaCreditScoreCobuildGetRequest) SetPlatform(Platform string) {
	m.Platform = Platform
}

func (m *ZhimaCreditScoreCobuildGetRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaCreditScoreCobuildGetRequest) SetExtParams(ExtParams string) {
	m.ExtParams = ExtParams
}

func (m *ZhimaCreditScoreCobuildGetRequest) GetExtParams() string {
	return m.ExtParams
}
