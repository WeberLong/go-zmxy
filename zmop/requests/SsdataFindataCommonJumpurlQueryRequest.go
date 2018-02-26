package requests

import (
	"github.com/WeberLong/go-zmxy/zmop/interfaces"
)

type SsdataFindataCommonJumpurlQueryRequest struct {
	bizExtParams string // 业务扩展参数
	bizNo        string // 业务流水号

	interfaces.ZhimaRequestParams
}

func (m *SsdataFindataCommonJumpurlQueryRequest) InitBizParams(bizExtParams, bizNo string) {
	m.ApiParams = &map[string]string{}
	m.FileParams = &map[string]string{}

	(*m.ApiParams)["biz_ext_params"] = bizExtParams
	m.bizExtParams = bizExtParams

	(*m.ApiParams)["biz_no"] = bizNo
	m.bizNo = bizNo
}

func (*SsdataFindataCommonJumpurlQueryRequest) GetApiMethodName() string {
	return "ssdata.findata.common.jumpurl.query"
}

func (m *SsdataFindataCommonJumpurlQueryRequest) GetApiParams() *map[string]string {
	return m.ApiParams
}

func (m *SsdataFindataCommonJumpurlQueryRequest) GetFileParams() *map[string]string {
	return m.FileParams
}

func (m *SsdataFindataCommonJumpurlQueryRequest) SetApiVersion(ApiVersion string) {
	m.ApiVersion = ApiVersion
}

func (m *SsdataFindataCommonJumpurlQueryRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *SsdataFindataCommonJumpurlQueryRequest) SetScene(Scene string) {
	m.Scene = Scene
}

func (m *SsdataFindataCommonJumpurlQueryRequest) GetScene() string {
	return m.Scene
}

func (m *SsdataFindataCommonJumpurlQueryRequest) SetChannel(Channel string) {
	m.Channel = Channel
}

func (m *SsdataFindataCommonJumpurlQueryRequest) GetChannel() string {
	return m.Channel
}

func (m *SsdataFindataCommonJumpurlQueryRequest) SetPlatform(Platform string) {
	m.Platform = Platform
}

func (m *SsdataFindataCommonJumpurlQueryRequest) GetPlatform() string {
	return m.Platform
}

func (m *SsdataFindataCommonJumpurlQueryRequest) SetExtParams(ExtParams string) {
	m.ExtParams = ExtParams
}

func (m *SsdataFindataCommonJumpurlQueryRequest) GetExtParams() string {
	return m.ExtParams
}
