package requests

import (
	"github.com/WeberLong/go-zmxy/zmop/interfaces"
)

type ZhimaCustomerBorrowScanRequest struct {
	goodsId     string //
	productCode string //
	scenceCode  string //
	shopCode    string //

	interfaces.ZhimaRequestParams
}

func (m *ZhimaCustomerBorrowScanRequest) InitBizParams(goodsId, productCode, scenceCode, shopCode string) {
	m.ApiParams = &map[string]string{}
	m.FileParams = &map[string]string{}

	(*m.ApiParams)["goods_id"] = goodsId
	m.goodsId = goodsId

	(*m.ApiParams)["product_code"] = productCode
	m.productCode = productCode

	(*m.ApiParams)["scence_code"] = scenceCode
	m.scenceCode = scenceCode

	(*m.ApiParams)["shop_code"] = shopCode
	m.shopCode = shopCode
}

func (*ZhimaCustomerBorrowScanRequest) GetApiMethodName() string {
	return "zhima.customer.borrow.scan"
}

func (m *ZhimaCustomerBorrowScanRequest) GetApiParams() *map[string]string {
	return m.ApiParams
}

func (m *ZhimaCustomerBorrowScanRequest) GetFileParams() *map[string]string {
	return m.FileParams
}

func (m *ZhimaCustomerBorrowScanRequest) SetApiVersion(ApiVersion string) {
	m.ApiVersion = ApiVersion
}

func (m *ZhimaCustomerBorrowScanRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaCustomerBorrowScanRequest) SetScene(Scene string) {
	m.Scene = Scene
}

func (m *ZhimaCustomerBorrowScanRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaCustomerBorrowScanRequest) SetChannel(Channel string) {
	m.Channel = Channel
}

func (m *ZhimaCustomerBorrowScanRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaCustomerBorrowScanRequest) SetPlatform(Platform string) {
	m.Platform = Platform
}

func (m *ZhimaCustomerBorrowScanRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaCustomerBorrowScanRequest) SetExtParams(ExtParams string) {
	m.ExtParams = ExtParams
}

func (m *ZhimaCustomerBorrowScanRequest) GetExtParams() string {
	return m.ExtParams
}
