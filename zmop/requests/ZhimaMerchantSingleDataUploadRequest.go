package requests

import (
	"github.com/WeberLong/go-zmxy/zmop/interfaces"
)

type ZhimaMerchantSingleDataUploadRequest struct {
	bizExtParams string // 公用回传参数，这个字段由商户传入，系统透传给商户，便于商户做逻辑关联，请使用json格式。
	data         string // 传入的json数据，商户通过json格式将数据传给芝麻 ， json中的字段可以通过如下步骤获取：首先调用zhima.merchant.data.upload.initialize接口获取数据模板，该接口会返回一个数据模板文件的url地址，如：http://zmxymerchant-prod.oss-cn-shenzhen.zmxy.com.cn/openApi/openDoc/信用护航-负面记录和信用足迹商户数据模板V1.0.xlsx，该数据模板文件详细列出了需要传入的字段，及各字段的要求，data中的各字段就是该文件中列出的字段编码。
	primaryKeys  string // 主键列使用传入字段进行组合，也可以使用传入的某个单字段（确保主键稳定，而且可以很好的区分不同的数据）。例如order_no,pay_month或者order_no,bill_month组合，对于一个order_no只会有一条数据的情况，直接使用order_no作为主键列
	sceneCode    string // 场景码，每个场景码对应的数据模板不一样，请使用zhima.merchant.data.upload.initialize接口获取场景码对应的数据模板。

	interfaces.ZhimaRequestParams
}

func (m *ZhimaMerchantSingleDataUploadRequest) InitBizParams(bizExtParams, data, primaryKeys, sceneCode string) {
	m.ApiParams = &map[string]string{}
	m.FileParams = &map[string]string{}

	(*m.ApiParams)["biz_ext_params"] = bizExtParams
	m.bizExtParams = bizExtParams

	(*m.ApiParams)["data"] = data
	m.data = data

	(*m.ApiParams)["primary_keys"] = primaryKeys
	m.primaryKeys = primaryKeys

	(*m.ApiParams)["scene_code"] = sceneCode
	m.sceneCode = sceneCode
}

func (*ZhimaMerchantSingleDataUploadRequest) GetApiMethodName() string {
	return "zhima.merchant.single.data.upload"
}

func (m *ZhimaMerchantSingleDataUploadRequest) GetApiParams() *map[string]string {
	return m.ApiParams
}

func (m *ZhimaMerchantSingleDataUploadRequest) GetFileParams() *map[string]string {
	return m.FileParams
}

func (m *ZhimaMerchantSingleDataUploadRequest) SetApiVersion(ApiVersion string) {
	m.ApiVersion = ApiVersion
}

func (m *ZhimaMerchantSingleDataUploadRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaMerchantSingleDataUploadRequest) SetScene(Scene string) {
	m.Scene = Scene
}

func (m *ZhimaMerchantSingleDataUploadRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaMerchantSingleDataUploadRequest) SetChannel(Channel string) {
	m.Channel = Channel
}

func (m *ZhimaMerchantSingleDataUploadRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaMerchantSingleDataUploadRequest) SetPlatform(Platform string) {
	m.Platform = Platform
}

func (m *ZhimaMerchantSingleDataUploadRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaMerchantSingleDataUploadRequest) SetExtParams(ExtParams string) {
	m.ExtParams = ExtParams
}

func (m *ZhimaMerchantSingleDataUploadRequest) GetExtParams() string {
	return m.ExtParams
}
