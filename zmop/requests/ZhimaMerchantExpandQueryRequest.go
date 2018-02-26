package requests

import (
	"github.com/WeberLong/go-zmxy/zmop/interfaces"
)

type ZhimaMerchantExpandQueryRequest struct {
	bizExtParams string // 业务扩展参数，当前无需设置该参数

	interfaces.ZhimaRequestParams
}

func (m *ZhimaMerchantExpandQueryRequest) InitBizParams(bizExtParams string) {
	m.ApiParams = &map[string]string{}
	m.FileParams = &map[string]string{}

	(*m.ApiParams)["biz_ext_params"] = bizExtParams
	m.bizExtParams = bizExtParams
}

func (*ZhimaMerchantExpandQueryRequest) GetApiMethodName() string {
	return "zhima.merchant.expand.query"
}

func (m *ZhimaMerchantExpandQueryRequest) GetApiParams() *map[string]string {
	return m.ApiParams
}

func (m *ZhimaMerchantExpandQueryRequest) GetFileParams() *map[string]string {
	return m.FileParams
}

func (m *ZhimaMerchantExpandQueryRequest) SetApiVersion(ApiVersion string) {
	m.ApiVersion = ApiVersion
}

func (m *ZhimaMerchantExpandQueryRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaMerchantExpandQueryRequest) SetScene(Scene string) {
	m.Scene = Scene
}

func (m *ZhimaMerchantExpandQueryRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaMerchantExpandQueryRequest) SetChannel(Channel string) {
	m.Channel = Channel
}

func (m *ZhimaMerchantExpandQueryRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaMerchantExpandQueryRequest) SetPlatform(Platform string) {
	m.Platform = Platform
}

func (m *ZhimaMerchantExpandQueryRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaMerchantExpandQueryRequest) SetExtParams(ExtParams string) {
	m.ExtParams = ExtParams
}

func (m *ZhimaMerchantExpandQueryRequest) GetExtParams() string {
	return m.ExtParams
}
