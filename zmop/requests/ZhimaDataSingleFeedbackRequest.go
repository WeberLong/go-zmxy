package requests

import (
	"github.com/WeberLong/go-zmxy/zmop/interfaces"
)

type ZhimaDataSingleFeedbackRequest struct {
	bizExtParams string // 扩展参数
	data         string // 反馈的json格式的数据内容
	identity     string // 主键使用反馈字段进行组合，也可以使用反馈的某个单字段（确保主键稳定，而且可以很好的区分不同的数据）。例如order_no
	typeId       string // 芝麻系统中配置的值，由芝麻信用提供，需要匹配，测试反馈和正式反馈使用不同的TYPE_ID

	interfaces.ZhimaRequestParams
}

func (m *ZhimaDataSingleFeedbackRequest) InitBizParams(bizExtParams, data, identity, typeId string) {
	m.ApiParams = &map[string]string{}
	m.FileParams = &map[string]string{}

	(*m.ApiParams)["biz_ext_params"] = bizExtParams
	m.bizExtParams = bizExtParams

	(*m.ApiParams)["data"] = data
	m.data = data

	(*m.ApiParams)["identity"] = identity
	m.identity = identity

	(*m.ApiParams)["type_id"] = typeId
	m.typeId = typeId
}

func (*ZhimaDataSingleFeedbackRequest) GetApiMethodName() string {
	return "zhima.data.single.feedback"
}

func (m *ZhimaDataSingleFeedbackRequest) GetApiParams() *map[string]string {
	return m.ApiParams
}

func (m *ZhimaDataSingleFeedbackRequest) GetFileParams() *map[string]string {
	return m.FileParams
}

func (m *ZhimaDataSingleFeedbackRequest) SetApiVersion(ApiVersion string) {
	m.ApiVersion = ApiVersion
}

func (m *ZhimaDataSingleFeedbackRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaDataSingleFeedbackRequest) SetScene(Scene string) {
	m.Scene = Scene
}

func (m *ZhimaDataSingleFeedbackRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaDataSingleFeedbackRequest) SetChannel(Channel string) {
	m.Channel = Channel
}

func (m *ZhimaDataSingleFeedbackRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaDataSingleFeedbackRequest) SetPlatform(Platform string) {
	m.Platform = Platform
}

func (m *ZhimaDataSingleFeedbackRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaDataSingleFeedbackRequest) SetExtParams(ExtParams string) {
	m.ExtParams = ExtParams
}

func (m *ZhimaDataSingleFeedbackRequest) GetExtParams() string {
	return m.ExtParams
}
