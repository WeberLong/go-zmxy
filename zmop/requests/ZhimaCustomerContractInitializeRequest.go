package requests

import (
	"github.com/WeberLong/go-zmxy/zmop/interfaces"
)

type ZhimaCustomerContractInitializeRequest struct {
	contractFile string // 合约内容，PDF文件流，BASE64编码
	contractName string // 合约名称，展示给签约方
	productCode  string // 芝麻认证产品码,示例值为真实的产品码

	interfaces.ZhimaRequestParams
}

func (m *ZhimaCustomerContractInitializeRequest) InitBizParams(contractFile, contractName, productCode string) {
	m.ApiParams = &map[string]string{}
	m.FileParams = &map[string]string{}

	(*m.ApiParams)["contract_file"] = contractFile
	m.contractFile = contractFile

	(*m.ApiParams)["contract_name"] = contractName
	m.contractName = contractName

	(*m.ApiParams)["product_code"] = productCode
	m.productCode = productCode
}

func (*ZhimaCustomerContractInitializeRequest) GetApiMethodName() string {
	return "zhima.customer.contract.initialize"
}

func (m *ZhimaCustomerContractInitializeRequest) GetApiParams() *map[string]string {
	return m.ApiParams
}

func (m *ZhimaCustomerContractInitializeRequest) GetFileParams() *map[string]string {
	return m.FileParams
}

func (m *ZhimaCustomerContractInitializeRequest) SetApiVersion(ApiVersion string) {
	m.ApiVersion = ApiVersion
}

func (m *ZhimaCustomerContractInitializeRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaCustomerContractInitializeRequest) SetScene(Scene string) {
	m.Scene = Scene
}

func (m *ZhimaCustomerContractInitializeRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaCustomerContractInitializeRequest) SetChannel(Channel string) {
	m.Channel = Channel
}

func (m *ZhimaCustomerContractInitializeRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaCustomerContractInitializeRequest) SetPlatform(Platform string) {
	m.Platform = Platform
}

func (m *ZhimaCustomerContractInitializeRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaCustomerContractInitializeRequest) SetExtParams(ExtParams string) {
	m.ExtParams = ExtParams
}

func (m *ZhimaCustomerContractInitializeRequest) GetExtParams() string {
	return m.ExtParams
}
