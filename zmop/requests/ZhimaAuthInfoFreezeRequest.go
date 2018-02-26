package requests

import (
	"github.com/WeberLong/go-zmxy/zmop/interfaces"
)

type ZhimaAuthInfoFreezeRequest struct {
	bizNo     string // 用户在商户处申请业务的唯一识别码;\n每个申请仅对应一条冻结记录，若存在相同流水号的冻结周期将覆盖
	bizParams string // 扩展字段，json字符串的key-value格式
	bizType   string // 申请原因001: 贷款申请中, 002:信用卡申请中, 003:租车申请中, 004:信贷服务期内, 005:信贷逾期中
	endDate   string // 冻结结束时间，若该时间减去冻结开始时间的差值大于冻结周期，则该时间将赋值冻结开始时间+冻结周期
	openId    string // 用户在商端的身份标识
	startDate string // 冻结开始时间，若该时间早于系统当前时间，则会取当前时间作为冻结开始时间

	interfaces.ZhimaRequestParams
}

func (m *ZhimaAuthInfoFreezeRequest) InitBizParams(bizNo, bizParams, bizType, endDate, openId, startDate string) {
	m.ApiParams = &map[string]string{}
	m.FileParams = &map[string]string{}

	(*m.ApiParams)["biz_no"] = bizNo
	m.bizNo = bizNo

	(*m.ApiParams)["biz_params"] = bizParams
	m.bizParams = bizParams

	(*m.ApiParams)["biz_type"] = bizType
	m.bizType = bizType

	(*m.ApiParams)["end_date"] = endDate
	m.endDate = endDate

	(*m.ApiParams)["open_id"] = openId
	m.openId = openId

	(*m.ApiParams)["start_date"] = startDate
	m.startDate = startDate
}

func (*ZhimaAuthInfoFreezeRequest) GetApiMethodName() string {
	return "zhima.auth.info.freeze"
}

func (m *ZhimaAuthInfoFreezeRequest) GetApiParams() *map[string]string {
	return m.ApiParams
}

func (m *ZhimaAuthInfoFreezeRequest) GetFileParams() *map[string]string {
	return m.FileParams
}

func (m *ZhimaAuthInfoFreezeRequest) SetApiVersion(ApiVersion string) {
	m.ApiVersion = ApiVersion
}

func (m *ZhimaAuthInfoFreezeRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaAuthInfoFreezeRequest) SetScene(Scene string) {
	m.Scene = Scene
}

func (m *ZhimaAuthInfoFreezeRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaAuthInfoFreezeRequest) SetChannel(Channel string) {
	m.Channel = Channel
}

func (m *ZhimaAuthInfoFreezeRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaAuthInfoFreezeRequest) SetPlatform(Platform string) {
	m.Platform = Platform
}

func (m *ZhimaAuthInfoFreezeRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaAuthInfoFreezeRequest) SetExtParams(ExtParams string) {
	m.ExtParams = ExtParams
}

func (m *ZhimaAuthInfoFreezeRequest) GetExtParams() string {
	return m.ExtParams
}
