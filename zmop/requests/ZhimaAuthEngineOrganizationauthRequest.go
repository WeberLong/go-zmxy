package requests

import (
	"github.com/WeberLong/go-zmxy/zmop/interfaces"
)

type ZhimaAuthEngineOrganizationauthRequest struct {
	bizParams     string // 授权码入参，1). 若identity_Type=2时，{"auth_code":"M_P_COMPANY_CERT"} 2). 若identity_Type=5时，{"auth_code":"M_P_COMPANY_UID"}
	identityParam string // 证件号目前支持2种：a. 工商注册号：NATIONAL_LEGAL  b. 社会统一信用代码 : NATIONAL_LEGAL_MERGE1). 若identity_type=2时：identity_param={\"certNo\":\"440000400004160\",\"certType\":\"NATIONAL_LEGAL\",\"name\":\"中国移动通信集团广东有限公司\"}"2). 若identity_type=5时：identity_param={\"userId\":\"2088xxxxxx\"}"
	identityType  string // 用户身份标识类型：5： userId入参的类型标识；2： 证件号和姓名的入参的类型标识

	interfaces.ZhimaRequestParams
}

func (m *ZhimaAuthEngineOrganizationauthRequest) InitBizParams(bizParams, identityParam, identityType string) {
	m.ApiParams = &map[string]string{}
	m.FileParams = &map[string]string{}

	(*m.ApiParams)["biz_params"] = bizParams
	m.bizParams = bizParams

	(*m.ApiParams)["identity_param"] = identityParam
	m.identityParam = identityParam

	(*m.ApiParams)["identity_type"] = identityType
	m.identityType = identityType
}

func (*ZhimaAuthEngineOrganizationauthRequest) GetApiMethodName() string {
	return "zhima.auth.engine.organizationauth"
}

func (m *ZhimaAuthEngineOrganizationauthRequest) GetApiParams() *map[string]string {
	return m.ApiParams
}

func (m *ZhimaAuthEngineOrganizationauthRequest) GetFileParams() *map[string]string {
	return m.FileParams
}

func (m *ZhimaAuthEngineOrganizationauthRequest) SetApiVersion(ApiVersion string) {
	m.ApiVersion = ApiVersion
}

func (m *ZhimaAuthEngineOrganizationauthRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaAuthEngineOrganizationauthRequest) SetScene(Scene string) {
	m.Scene = Scene
}

func (m *ZhimaAuthEngineOrganizationauthRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaAuthEngineOrganizationauthRequest) SetChannel(Channel string) {
	m.Channel = Channel
}

func (m *ZhimaAuthEngineOrganizationauthRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaAuthEngineOrganizationauthRequest) SetPlatform(Platform string) {
	m.Platform = Platform
}

func (m *ZhimaAuthEngineOrganizationauthRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaAuthEngineOrganizationauthRequest) SetExtParams(ExtParams string) {
	m.ExtParams = ExtParams
}

func (m *ZhimaAuthEngineOrganizationauthRequest) GetExtParams() string {
	return m.ExtParams
}
