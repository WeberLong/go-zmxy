package requests

import (
	"github.com/WeberLong/go-zmxy/zmop/interfaces"
)

type ZhimaCustomerBorrowRestoreRequest struct {
	goodsId     string //
	productCode string //
	scenceCode  string //
	shopCode    string //

	interfaces.ZhimaRequestParams
}

func (m *ZhimaCustomerBorrowRestoreRequest) InitBizParams(goodsId, productCode, scenceCode, shopCode string) {
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

func (*ZhimaCustomerBorrowRestoreRequest) GetApiMethodName() string {
	return "zhima.customer.borrow.restore"
}

func (m *ZhimaCustomerBorrowRestoreRequest) GetApiParams() *map[string]string {
	return m.ApiParams
}

func (m *ZhimaCustomerBorrowRestoreRequest) GetFileParams() *map[string]string {
	return m.FileParams
}

func (m *ZhimaCustomerBorrowRestoreRequest) SetApiVersion(ApiVersion string) {
	m.ApiVersion = ApiVersion
}

func (m *ZhimaCustomerBorrowRestoreRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaCustomerBorrowRestoreRequest) SetScene(Scene string) {
	m.Scene = Scene
}

func (m *ZhimaCustomerBorrowRestoreRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaCustomerBorrowRestoreRequest) SetChannel(Channel string) {
	m.Channel = Channel
}

func (m *ZhimaCustomerBorrowRestoreRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaCustomerBorrowRestoreRequest) SetPlatform(Platform string) {
	m.Platform = Platform
}

func (m *ZhimaCustomerBorrowRestoreRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaCustomerBorrowRestoreRequest) SetExtParams(ExtParams string) {
	m.ExtParams = ExtParams
}

func (m *ZhimaCustomerBorrowRestoreRequest) GetExtParams() string {
	return m.ExtParams
}
