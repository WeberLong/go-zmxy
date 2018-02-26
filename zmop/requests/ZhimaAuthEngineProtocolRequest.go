package requests

import (
	"github.com/WeberLong/go-zmxy/zmop/interfaces"
)

type ZhimaAuthEngineProtocolRequest struct {
	bizParams     string // 业务扩展入参
	identityParam string // 授权的身份标识参数
	identityType  string // 用户的身份标识类型\n3：身份证+姓名+手机号进行支付宝识别，核身，授权\n4：身份证+姓名+手机号（可选）进行公安网的核身，授权

	interfaces.ZhimaRequestParams
}

func (m *ZhimaAuthEngineProtocolRequest) InitBizParams(bizParams, identityParam, identityType string) {
	m.ApiParams = &map[string]string{}
	m.FileParams = &map[string]string{}

	(*m.ApiParams)["biz_params"] = bizParams
	m.bizParams = bizParams

	(*m.ApiParams)["identity_param"] = identityParam
	m.identityParam = identityParam

	(*m.ApiParams)["identity_type"] = identityType
	m.identityType = identityType
}

func (*ZhimaAuthEngineProtocolRequest) GetApiMethodName() string {
	return "zhima.auth.engine.protocol"
}

func (m *ZhimaAuthEngineProtocolRequest) GetApiParams() *map[string]string {
	return m.ApiParams
}

func (m *ZhimaAuthEngineProtocolRequest) GetFileParams() *map[string]string {
	return m.FileParams
}

func (m *ZhimaAuthEngineProtocolRequest) SetApiVersion(ApiVersion string) {
	m.ApiVersion = ApiVersion
}

func (m *ZhimaAuthEngineProtocolRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaAuthEngineProtocolRequest) SetScene(Scene string) {
	m.Scene = Scene
}

func (m *ZhimaAuthEngineProtocolRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaAuthEngineProtocolRequest) SetChannel(Channel string) {
	m.Channel = Channel
}

func (m *ZhimaAuthEngineProtocolRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaAuthEngineProtocolRequest) SetPlatform(Platform string) {
	m.Platform = Platform
}

func (m *ZhimaAuthEngineProtocolRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaAuthEngineProtocolRequest) SetExtParams(ExtParams string) {
	m.ExtParams = ExtParams
}

func (m *ZhimaAuthEngineProtocolRequest) GetExtParams() string {
	return m.ExtParams
}
