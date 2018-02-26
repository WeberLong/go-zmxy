package requests

import (
	"github.com/WeberLong/go-zmxy/zmop/interfaces"
)

type ZhimaAuthInfoAuthqueryRequest struct {
	authCategory  string // 授权类型，用于识别当前查询是否授权的分流；可传参数“B2B”或“C2B”，自助商户请填写“C2B”。
	identityParam string // 不同身份类型传入的参数列表,json字符串的key-value格式身份类型identityType=0:{"openId":"268801234567890123456"}身份类型identityType=2:{"certNo":"330100xxxxxxxxxxxx","name":"张三","certType":"IDENTITY_CARD"}
	identityType  string // 身份标识类型0:按照openId查询2:按照身份证+姓名查询

	interfaces.ZhimaRequestParams
}

func (m *ZhimaAuthInfoAuthqueryRequest) InitBizParams(authCategory, identityParam, identityType string) {
	m.ApiParams = &map[string]string{}
	m.FileParams = &map[string]string{}

	(*m.ApiParams)["auth_category"] = authCategory
	m.authCategory = authCategory

	(*m.ApiParams)["identity_param"] = identityParam
	m.identityParam = identityParam

	(*m.ApiParams)["identity_type"] = identityType
	m.identityType = identityType
}

func (*ZhimaAuthInfoAuthqueryRequest) GetApiMethodName() string {
	return "zhima.auth.info.authquery"
}

func (m *ZhimaAuthInfoAuthqueryRequest) GetApiParams() *map[string]string {
	return m.ApiParams
}

func (m *ZhimaAuthInfoAuthqueryRequest) GetFileParams() *map[string]string {
	return m.FileParams
}

func (m *ZhimaAuthInfoAuthqueryRequest) SetApiVersion(ApiVersion string) {
	m.ApiVersion = ApiVersion
}

func (m *ZhimaAuthInfoAuthqueryRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaAuthInfoAuthqueryRequest) SetScene(Scene string) {
	m.Scene = Scene
}

func (m *ZhimaAuthInfoAuthqueryRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaAuthInfoAuthqueryRequest) SetChannel(Channel string) {
	m.Channel = Channel
}

func (m *ZhimaAuthInfoAuthqueryRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaAuthInfoAuthqueryRequest) SetPlatform(Platform string) {
	m.Platform = Platform
}

func (m *ZhimaAuthInfoAuthqueryRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaAuthInfoAuthqueryRequest) SetExtParams(ExtParams string) {
	m.ExtParams = ExtParams
}

func (m *ZhimaAuthInfoAuthqueryRequest) GetExtParams() string {
	return m.ExtParams
}
