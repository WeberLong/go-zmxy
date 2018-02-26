package requests

import (
	"github.com/WeberLong/go-zmxy/zmop/interfaces"
)

type ZhimaCreditWatchlistBriefGetRequest struct {
	certNo        string // 证件类型对应的证件号码， 如：身份证号， 护照号，userId
	certType      string // 当前支持3种类型的输入：IDENTITY_CARD (身份证)PASSPORT (护照)ALIPAY_USER_ID (支付宝uid)
	name          string // 当cert_type 为ALIPAY_USER_ID时证件名称可为空
	productCode   string // 芝麻开放平台信用产品码， 唯一标示云产品
	transactionId string // 商户请求的唯一标志，长度64位以内字符串，仅限字母数字下划线组合。该标识作为业务调用的唯一标识，商户要保证其业务唯一性，使用相同transaction_id的查询，芝麻在一段时间内（一般为1天）返回首次查询结果，超过有效期的查询即为无效并返回异常，有效期内的重复查询不重新计费。

	interfaces.ZhimaRequestParams
}

func (m *ZhimaCreditWatchlistBriefGetRequest) InitBizParams(certNo, certType, name, productCode, transactionId string) {
	m.ApiParams = &map[string]string{}
	m.FileParams = &map[string]string{}

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

func (*ZhimaCreditWatchlistBriefGetRequest) GetApiMethodName() string {
	return "zhima.credit.watchlist.brief.get"
}

func (m *ZhimaCreditWatchlistBriefGetRequest) GetApiParams() *map[string]string {
	return m.ApiParams
}

func (m *ZhimaCreditWatchlistBriefGetRequest) GetFileParams() *map[string]string {
	return m.FileParams
}

func (m *ZhimaCreditWatchlistBriefGetRequest) SetApiVersion(ApiVersion string) {
	m.ApiVersion = ApiVersion
}

func (m *ZhimaCreditWatchlistBriefGetRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaCreditWatchlistBriefGetRequest) SetScene(Scene string) {
	m.Scene = Scene
}

func (m *ZhimaCreditWatchlistBriefGetRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaCreditWatchlistBriefGetRequest) SetChannel(Channel string) {
	m.Channel = Channel
}

func (m *ZhimaCreditWatchlistBriefGetRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaCreditWatchlistBriefGetRequest) SetPlatform(Platform string) {
	m.Platform = Platform
}

func (m *ZhimaCreditWatchlistBriefGetRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaCreditWatchlistBriefGetRequest) SetExtParams(ExtParams string) {
	m.ExtParams = ExtParams
}

func (m *ZhimaCreditWatchlistBriefGetRequest) GetExtParams() string {
	return m.ExtParams
}
