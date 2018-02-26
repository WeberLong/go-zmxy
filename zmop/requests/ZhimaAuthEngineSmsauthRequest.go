package requests

import (
	"github.com/WeberLong/go-zmxy/zmop/interfaces"
)

type ZhimaAuthEngineSmsauthRequest struct {
	bizParams     string // 业务扩展字段,页面授权接口需要传入auth_code,channelType,stateauth_code授权码,短信唤起支付宝客户端接口的值为M_SMSchannelType渠道类型,每个授权码支持不同的渠道类型appsdk:sdk接入apppc:商户pc页面接入api:后台api接入windows:支付宝服务窗接入app:商户app接入state是商户自定义的数据,页面授权接口会原样把这个数据返回个商户
	identityParam string // 身份参数,短信唤起支付宝客户端授权需要传入姓名+证件类型+证件号码+手机号
	identityType  string // 身份类型,短信唤起支付宝客户端接口的身份识别类型为2:按照姓名+证件类型+证件号码+手机号进行授权

	interfaces.ZhimaRequestParams
}

func (m *ZhimaAuthEngineSmsauthRequest) InitBizParams(bizParams, identityParam, identityType string) {
	m.ApiParams = &map[string]string{}
	m.FileParams = &map[string]string{}

	(*m.ApiParams)["biz_params"] = bizParams
	m.bizParams = bizParams

	(*m.ApiParams)["identity_param"] = identityParam
	m.identityParam = identityParam

	(*m.ApiParams)["identity_type"] = identityType
	m.identityType = identityType
}

func (*ZhimaAuthEngineSmsauthRequest) GetApiMethodName() string {
	return "zhima.auth.engine.smsauth"
}

func (m *ZhimaAuthEngineSmsauthRequest) GetApiParams() *map[string]string {
	return m.ApiParams
}

func (m *ZhimaAuthEngineSmsauthRequest) GetFileParams() *map[string]string {
	return m.FileParams
}

func (m *ZhimaAuthEngineSmsauthRequest) SetApiVersion(ApiVersion string) {
	m.ApiVersion = ApiVersion
}

func (m *ZhimaAuthEngineSmsauthRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaAuthEngineSmsauthRequest) SetScene(Scene string) {
	m.Scene = Scene
}

func (m *ZhimaAuthEngineSmsauthRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaAuthEngineSmsauthRequest) SetChannel(Channel string) {
	m.Channel = Channel
}

func (m *ZhimaAuthEngineSmsauthRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaAuthEngineSmsauthRequest) SetPlatform(Platform string) {
	m.Platform = Platform
}

func (m *ZhimaAuthEngineSmsauthRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaAuthEngineSmsauthRequest) SetExtParams(ExtParams string) {
	m.ExtParams = ExtParams
}

func (m *ZhimaAuthEngineSmsauthRequest) GetExtParams() string {
	return m.ExtParams
}
