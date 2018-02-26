package requests

import (
	"github.com/WeberLong/go-zmxy/zmop/interfaces"
)

type ZhimaMerchantEnterpriseQueryRequest struct {
	bizExtParams string // 业务扩展参数，当前无需设置该参数

	interfaces.ZhimaRequestParams
}

func (m *ZhimaMerchantEnterpriseQueryRequest) InitBizParams(bizExtParams string) {
	m.ApiParams = &map[string]string{}
	m.FileParams = &map[string]string{}

	(*m.ApiParams)["biz_ext_params"] = bizExtParams
	m.bizExtParams = bizExtParams
}

func (*ZhimaMerchantEnterpriseQueryRequest) GetApiMethodName() string {
	return "zhima.merchant.enterprise.query"
}

func (m *ZhimaMerchantEnterpriseQueryRequest) GetApiParams() *map[string]string {
	return m.ApiParams
}

func (m *ZhimaMerchantEnterpriseQueryRequest) GetFileParams() *map[string]string {
	return m.FileParams
}

func (m *ZhimaMerchantEnterpriseQueryRequest) SetApiVersion(ApiVersion string) {
	m.ApiVersion = ApiVersion
}

func (m *ZhimaMerchantEnterpriseQueryRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaMerchantEnterpriseQueryRequest) SetScene(Scene string) {
	m.Scene = Scene
}

func (m *ZhimaMerchantEnterpriseQueryRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaMerchantEnterpriseQueryRequest) SetChannel(Channel string) {
	m.Channel = Channel
}

func (m *ZhimaMerchantEnterpriseQueryRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaMerchantEnterpriseQueryRequest) SetPlatform(Platform string) {
	m.Platform = Platform
}

func (m *ZhimaMerchantEnterpriseQueryRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaMerchantEnterpriseQueryRequest) SetExtParams(ExtParams string) {
	m.ExtParams = ExtParams
}

func (m *ZhimaMerchantEnterpriseQueryRequest) GetExtParams() string {
	return m.ExtParams
}
