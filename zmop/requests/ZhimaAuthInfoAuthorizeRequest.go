package requests

import (
	"fmt"

	"github.com/WeberLong/go-zmxy/zmop/interfaces"
)

type ZhimaAuthInfoAuthorizeRequest struct {
	bizParams     string // 业务扩展字段,页面授权接口需要传入auth_code,channelType,stateauth_code授权码,值取决于授权方式和身份类型PC方式,身份类型identity_type=1:{"auth_code":"M_MOBILE_APPPC"}PC方式,身份类型identity_type=2:{"auth_code":"M_APPPC_CERT"}H5方式(身份类型identity_type为任何值):{"auth_code":"M_H5"}SDK方式(身份类型identity_type为任何值):{"auth_code":"M_APPSDK"}channelType渠道类型,每个授权码支持不同的渠道类型appsdk:sdk接入apppc:商户pc页面接入api:后台api接入windows:支付宝服务窗接入app:商户app接入state是商户自定义的数据,页面授权接口会原样把这个数据返回个商户
	identityParam string // 不同身份类型传入的参数列表,json字符串的key-value格式身份类型identityType=1:{"mobileNo":"15158657683"}身份类型identityType=2:{"certNo":"330100xxxxxxxxxxxx","name":"张三","certType":"IDENTITY_CARD"}
	identityType  string // 身份标识类型<p>1:按照手机号进行授权</p>2:按照身份证+姓名进行授权

	interfaces.ZhimaRequestParams
}

func (m *ZhimaAuthInfoAuthorizeRequest) InitBizParams(bizParams, identityParam, identityType string) {
	m.ApiParams = &map[string]string{}
	m.FileParams = &map[string]string{}

	(*m.ApiParams)["biz_params"] = bizParams
	m.bizParams = bizParams

	(*m.ApiParams)["identity_param"] = identityParam
	m.identityParam = identityParam

	(*m.ApiParams)["identity_type"] = identityType
	m.identityType = identityType

	fmt.Println((*m.ApiParams)["biz_params"])
	fmt.Println((*m.ApiParams)["identity_param"])
	fmt.Println((*m.ApiParams)["identity_type"])
}

func (*ZhimaAuthInfoAuthorizeRequest) GetApiMethodName() string {
	return "zhima.auth.info.authorize"
}

func (m *ZhimaAuthInfoAuthorizeRequest) GetApiParams() *map[string]string {
	return m.ApiParams
}

func (m *ZhimaAuthInfoAuthorizeRequest) GetFileParams() *map[string]string {
	return m.FileParams
}

func (m *ZhimaAuthInfoAuthorizeRequest) SetApiVersion(ApiVersion string) {
	m.ApiVersion = ApiVersion
}

func (m *ZhimaAuthInfoAuthorizeRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaAuthInfoAuthorizeRequest) SetScene(Scene string) {
	m.Scene = Scene
}

func (m *ZhimaAuthInfoAuthorizeRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaAuthInfoAuthorizeRequest) SetChannel(Channel string) {
	m.Channel = Channel
}

func (m *ZhimaAuthInfoAuthorizeRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaAuthInfoAuthorizeRequest) SetPlatform(Platform string) {
	m.Platform = Platform
}

func (m *ZhimaAuthInfoAuthorizeRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaAuthInfoAuthorizeRequest) SetExtParams(ExtParams string) {
	m.ExtParams = ExtParams
}

func (m *ZhimaAuthInfoAuthorizeRequest) GetExtParams() string {
	return m.ExtParams
}
