package requests

import (
	"github.com/WeberLong/go-zmxy/zmop/interfaces"
)

type ZhimaCreditContactAnalyzeQueryRequest struct {
	productCode   string // 芝麻开放平台信用产品码， 唯一标示云产品
	transactionId string // 商户请求的唯一标志，长度64位以内字符串，仅限字母数字下划线组合。 该标识作为业务调用的唯一标识，商户要保证其业务唯一性，使用相同transaction_id的查询， 芝麻在一段时间内（一般为1天）返回首次查询结果， 超过有效期的查询即为无效并返回异常，有效期内的重复查询不重新计费。
	userIds       string // 支付宝用户id的列表，String类型，多个uid之间用逗号【,】分隔

	interfaces.ZhimaRequestParams
}

func (m *ZhimaCreditContactAnalyzeQueryRequest) InitBizParams(productCode, transactionId, userIds string) {
	m.ApiParams = &map[string]string{}
	m.FileParams = &map[string]string{}

	(*m.ApiParams)["product_code"] = productCode
	m.productCode = productCode

	(*m.ApiParams)["transaction_id"] = transactionId
	m.transactionId = transactionId

	(*m.ApiParams)["user_ids"] = userIds
	m.userIds = userIds
}

func (*ZhimaCreditContactAnalyzeQueryRequest) GetApiMethodName() string {
	return "zhima.credit.contact.analyze.query"
}

func (m *ZhimaCreditContactAnalyzeQueryRequest) GetApiParams() *map[string]string {
	return m.ApiParams
}

func (m *ZhimaCreditContactAnalyzeQueryRequest) GetFileParams() *map[string]string {
	return m.FileParams
}

func (m *ZhimaCreditContactAnalyzeQueryRequest) SetApiVersion(ApiVersion string) {
	m.ApiVersion = ApiVersion
}

func (m *ZhimaCreditContactAnalyzeQueryRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaCreditContactAnalyzeQueryRequest) SetScene(Scene string) {
	m.Scene = Scene
}

func (m *ZhimaCreditContactAnalyzeQueryRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaCreditContactAnalyzeQueryRequest) SetChannel(Channel string) {
	m.Channel = Channel
}

func (m *ZhimaCreditContactAnalyzeQueryRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaCreditContactAnalyzeQueryRequest) SetPlatform(Platform string) {
	m.Platform = Platform
}

func (m *ZhimaCreditContactAnalyzeQueryRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaCreditContactAnalyzeQueryRequest) SetExtParams(ExtParams string) {
	m.ExtParams = ExtParams
}

func (m *ZhimaCreditContactAnalyzeQueryRequest) GetExtParams() string {
	return m.ExtParams
}
