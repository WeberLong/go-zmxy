package requests

import (
	"github.com/WeberLong/go-zmxy/zmop/interfaces"
)

type ZhimaOpenLogFeedbackRequest struct {
	logInfo string //

	interfaces.ZhimaRequestParams
}

func (m *ZhimaOpenLogFeedbackRequest) InitBizParams(logInfo string) {
	m.ApiParams = &map[string]string{}
	m.FileParams = &map[string]string{}

	(*m.ApiParams)["log_info"] = logInfo
	m.logInfo = logInfo
}

func (*ZhimaOpenLogFeedbackRequest) GetApiMethodName() string {
	return "zhima.open.log.feedback"
}

func (m *ZhimaOpenLogFeedbackRequest) GetApiParams() *map[string]string {
	return m.ApiParams
}

func (m *ZhimaOpenLogFeedbackRequest) GetFileParams() *map[string]string {
	return m.FileParams
}

func (m *ZhimaOpenLogFeedbackRequest) SetApiVersion(ApiVersion string) {
	m.ApiVersion = ApiVersion
}

func (m *ZhimaOpenLogFeedbackRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaOpenLogFeedbackRequest) SetScene(Scene string) {
	m.Scene = Scene
}

func (m *ZhimaOpenLogFeedbackRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaOpenLogFeedbackRequest) SetChannel(Channel string) {
	m.Channel = Channel
}

func (m *ZhimaOpenLogFeedbackRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaOpenLogFeedbackRequest) SetPlatform(Platform string) {
	m.Platform = Platform
}

func (m *ZhimaOpenLogFeedbackRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaOpenLogFeedbackRequest) SetExtParams(ExtParams string) {
	m.ExtParams = ExtParams
}

func (m *ZhimaOpenLogFeedbackRequest) GetExtParams() string {
	return m.ExtParams
}
