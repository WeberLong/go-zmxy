package requests

import (
	"github.com/WeberLong/go-zmxy/zmop/interfaces"
)

type ZhimaAuthEngineMultiauthRequest struct {
	authAppId      string // 外部商户应用id
	authMerchantId string // 外部商户id
	userId         string // 支付宝用户id

	interfaces.ZhimaRequestParams
}

func (m *ZhimaAuthEngineMultiauthRequest) InitBizParams(authAppId, authMerchantId, userId string) {
	m.ApiParams = &map[string]string{}
	m.FileParams = &map[string]string{}

	(*m.ApiParams)["auth_app_id"] = authAppId
	m.authAppId = authAppId

	(*m.ApiParams)["auth_merchant_id"] = authMerchantId
	m.authMerchantId = authMerchantId

	(*m.ApiParams)["user_id"] = userId
	m.userId = userId
}

func (*ZhimaAuthEngineMultiauthRequest) GetApiMethodName() string {
	return "zhima.auth.engine.multiauth"
}

func (m *ZhimaAuthEngineMultiauthRequest) GetApiParams() *map[string]string {
	return m.ApiParams
}

func (m *ZhimaAuthEngineMultiauthRequest) GetFileParams() *map[string]string {
	return m.FileParams
}

func (m *ZhimaAuthEngineMultiauthRequest) SetApiVersion(ApiVersion string) {
	m.ApiVersion = ApiVersion
}

func (m *ZhimaAuthEngineMultiauthRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaAuthEngineMultiauthRequest) SetScene(Scene string) {
	m.Scene = Scene
}

func (m *ZhimaAuthEngineMultiauthRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaAuthEngineMultiauthRequest) SetChannel(Channel string) {
	m.Channel = Channel
}

func (m *ZhimaAuthEngineMultiauthRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaAuthEngineMultiauthRequest) SetPlatform(Platform string) {
	m.Platform = Platform
}

func (m *ZhimaAuthEngineMultiauthRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaAuthEngineMultiauthRequest) SetExtParams(ExtParams string) {
	m.ExtParams = ExtParams
}

func (m *ZhimaAuthEngineMultiauthRequest) GetExtParams() string {
	return m.ExtParams
}
