package requests

import (
	"github.com/WeberLong/go-zmxy/zmop/interfaces"
)

type ZhimaDataFeedbackurlQueryRequest struct {
	merchantId string //

	interfaces.ZhimaRequestParams
}

func (m *ZhimaDataFeedbackurlQueryRequest) InitBizParams(merchantId string) {
	m.ApiParams = &map[string]string{}
	m.FileParams = &map[string]string{}

	(*m.ApiParams)["merchant_id"] = merchantId
	m.merchantId = merchantId
}

func (*ZhimaDataFeedbackurlQueryRequest) GetApiMethodName() string {
	return "zhima.data.feedbackurl.query"
}

func (m *ZhimaDataFeedbackurlQueryRequest) GetApiParams() *map[string]string {
	return m.ApiParams
}

func (m *ZhimaDataFeedbackurlQueryRequest) GetFileParams() *map[string]string {
	return m.FileParams
}

func (m *ZhimaDataFeedbackurlQueryRequest) SetApiVersion(ApiVersion string) {
	m.ApiVersion = ApiVersion
}

func (m *ZhimaDataFeedbackurlQueryRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaDataFeedbackurlQueryRequest) SetScene(Scene string) {
	m.Scene = Scene
}

func (m *ZhimaDataFeedbackurlQueryRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaDataFeedbackurlQueryRequest) SetChannel(Channel string) {
	m.Channel = Channel
}

func (m *ZhimaDataFeedbackurlQueryRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaDataFeedbackurlQueryRequest) SetPlatform(Platform string) {
	m.Platform = Platform
}

func (m *ZhimaDataFeedbackurlQueryRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaDataFeedbackurlQueryRequest) SetExtParams(ExtParams string) {
	m.ExtParams = ExtParams
}

func (m *ZhimaDataFeedbackurlQueryRequest) GetExtParams() string {
	return m.ExtParams
}
