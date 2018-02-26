package requests

import (
	"github.com/WeberLong/go-zmxy/zmop/interfaces"
)

type ZhimaMerchantCloseloopDataUploadRequest struct {
	bizExtParams      string // 扩展参数
	columns           string // 单条数据的数据列，多个列以逗号隔开。
	file              string // 传入的json格式的文件，其中{"records":  是每个文件的固定开头。
	fileCharset       string // 是传入文件的数据编码，如果文件格式是UTF-8，则填写UTF-8，如果文件格式是GBK，则填写GBK
	primaryKeyColumns string // 主键列使用反馈字段进行组合，也可以使用反馈的某个单字段（确保主键稳定，而且可以很好的区分不同的数据）。例如order_no,pay_month或者order_no,bill_month组合，对于一个order_no只会有一条数据的情况，直接使用order_no作为主键列
	records           string // 文件数据记录条数
	sceneCode         string // 场景码，用于标识上传的文件的用途，不同的场景码，file中的json格式不一样

	interfaces.ZhimaRequestParams
}

func (m *ZhimaMerchantCloseloopDataUploadRequest) InitBizParams(bizExtParams, columns, file, fileCharset, primaryKeyColumns, records, sceneCode string) {
	m.ApiParams = &map[string]string{}
	m.FileParams = &map[string]string{}

	(*m.ApiParams)["biz_ext_params"] = bizExtParams
	m.bizExtParams = bizExtParams

	(*m.ApiParams)["columns"] = columns
	m.columns = columns

	(*m.FileParams)["file"] = file
	m.file = file

	(*m.ApiParams)["file_charset"] = fileCharset
	m.fileCharset = fileCharset

	(*m.ApiParams)["primary_key_columns"] = primaryKeyColumns
	m.primaryKeyColumns = primaryKeyColumns

	(*m.ApiParams)["records"] = records
	m.records = records

	(*m.ApiParams)["scene_code"] = sceneCode
	m.sceneCode = sceneCode
}

func (*ZhimaMerchantCloseloopDataUploadRequest) GetApiMethodName() string {
	return "zhima.merchant.closeloop.data.upload"
}

func (m *ZhimaMerchantCloseloopDataUploadRequest) GetApiParams() *map[string]string {
	return m.ApiParams
}

func (m *ZhimaMerchantCloseloopDataUploadRequest) GetFileParams() *map[string]string {
	return m.FileParams
}

func (m *ZhimaMerchantCloseloopDataUploadRequest) SetApiVersion(ApiVersion string) {
	m.ApiVersion = ApiVersion
}

func (m *ZhimaMerchantCloseloopDataUploadRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaMerchantCloseloopDataUploadRequest) SetScene(Scene string) {
	m.Scene = Scene
}

func (m *ZhimaMerchantCloseloopDataUploadRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaMerchantCloseloopDataUploadRequest) SetChannel(Channel string) {
	m.Channel = Channel
}

func (m *ZhimaMerchantCloseloopDataUploadRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaMerchantCloseloopDataUploadRequest) SetPlatform(Platform string) {
	m.Platform = Platform
}

func (m *ZhimaMerchantCloseloopDataUploadRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaMerchantCloseloopDataUploadRequest) SetExtParams(ExtParams string) {
	m.ExtParams = ExtParams
}

func (m *ZhimaMerchantCloseloopDataUploadRequest) GetExtParams() string {
	return m.ExtParams
}
