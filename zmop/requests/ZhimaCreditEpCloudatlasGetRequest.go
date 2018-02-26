package requests

import (
	"github.com/WeberLong/go-zmxy/zmop/interfaces"
)

type ZhimaCreditEpCloudatlasGetRequest struct {
	cloudatlasCode string // 请输入调用云图产品申请接口返回的云图编码
	productCode    string // 产品码
	transactionId  string // 此transaction_id请传云图查询请求接口的transaction_id

	interfaces.ZhimaRequestParams
}

func (m *ZhimaCreditEpCloudatlasGetRequest) InitBizParams(cloudatlasCode, productCode, transactionId string) {
	m.ApiParams = &map[string]string{}
	m.FileParams = &map[string]string{}

	(*m.ApiParams)["cloudatlas_code"] = cloudatlasCode
	m.cloudatlasCode = cloudatlasCode

	(*m.ApiParams)["product_code"] = productCode
	m.productCode = productCode

	(*m.ApiParams)["transaction_id"] = transactionId
	m.transactionId = transactionId
}

func (*ZhimaCreditEpCloudatlasGetRequest) GetApiMethodName() string {
	return "zhima.credit.ep.cloudatlas.get"
}

func (m *ZhimaCreditEpCloudatlasGetRequest) GetApiParams() *map[string]string {
	return m.ApiParams
}

func (m *ZhimaCreditEpCloudatlasGetRequest) GetFileParams() *map[string]string {
	return m.FileParams
}

func (m *ZhimaCreditEpCloudatlasGetRequest) SetApiVersion(ApiVersion string) {
	m.ApiVersion = ApiVersion
}

func (m *ZhimaCreditEpCloudatlasGetRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaCreditEpCloudatlasGetRequest) SetScene(Scene string) {
	m.Scene = Scene
}

func (m *ZhimaCreditEpCloudatlasGetRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaCreditEpCloudatlasGetRequest) SetChannel(Channel string) {
	m.Channel = Channel
}

func (m *ZhimaCreditEpCloudatlasGetRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaCreditEpCloudatlasGetRequest) SetPlatform(Platform string) {
	m.Platform = Platform
}

func (m *ZhimaCreditEpCloudatlasGetRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaCreditEpCloudatlasGetRequest) SetExtParams(ExtParams string) {
	m.ExtParams = ExtParams
}

func (m *ZhimaCreditEpCloudatlasGetRequest) GetExtParams() string {
	return m.ExtParams
}
