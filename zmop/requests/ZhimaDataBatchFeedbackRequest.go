package requests

import (
	"github.com/WeberLong/go-zmxy/zmop/interfaces"
)

type ZhimaDataBatchFeedbackRequest struct {
	bizExtParams      string // 扩展参数
	columns           string // 单条数据的数据列，多个列以逗号隔开
	file              string // 反馈的json格式的文件，其中{"records":  是每个文件的固定开头。
	fileCharset       string // 是反馈文件的数据编码，如果文件格式是UTF-8，则填写UTF-8，如果文件格式是GBK，则填写GBK
	fileDescription   string // 文件描述信息
	fileType          string // 反馈的数据类型，目前只支持json_data
	primaryKeyColumns string // 主键列使用反馈字段进行组合，也可以使用反馈的某个单字段（确保主键稳定，而且可以很好的区分不同的数据）。例如order_no,pay_month或者order_no,bill_month组合，对于一个order_no只会有一条数据的情况，直接使用order_no作为主键列
	records           string // 文件数据记录条数
	typeId            string // 芝麻系统中配置的值，由芝麻信用提供，需要匹配，测试反馈和正式反馈使用不同的type_id。其中测试type_id与反馈字段模板会通过邮件统一提供给合作伙伴，在测试反馈通过之后，再通过邮件提供正式反馈type_id给合作伙伴。

	interfaces.ZhimaRequestParams
}

func (m *ZhimaDataBatchFeedbackRequest) InitBizParams(bizExtParams, columns, file, fileCharset, fileDescription, fileType, primaryKeyColumns, records, typeId string) {
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

	(*m.ApiParams)["file_description"] = fileDescription
	m.fileDescription = fileDescription

	(*m.ApiParams)["file_type"] = fileType
	m.fileType = fileType

	(*m.ApiParams)["primary_key_columns"] = primaryKeyColumns
	m.primaryKeyColumns = primaryKeyColumns

	(*m.ApiParams)["records"] = records
	m.records = records

	(*m.ApiParams)["type_id"] = typeId
	m.typeId = typeId
}

func (*ZhimaDataBatchFeedbackRequest) GetApiMethodName() string {
	return "zhima.data.batch.feedback"
}

func (m *ZhimaDataBatchFeedbackRequest) GetApiParams() *map[string]string {
	return m.ApiParams
}

func (m *ZhimaDataBatchFeedbackRequest) GetFileParams() *map[string]string {
	return m.FileParams
}

func (m *ZhimaDataBatchFeedbackRequest) SetApiVersion(ApiVersion string) {
	m.ApiVersion = ApiVersion
}

func (m *ZhimaDataBatchFeedbackRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaDataBatchFeedbackRequest) SetScene(Scene string) {
	m.Scene = Scene
}

func (m *ZhimaDataBatchFeedbackRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaDataBatchFeedbackRequest) SetChannel(Channel string) {
	m.Channel = Channel
}

func (m *ZhimaDataBatchFeedbackRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaDataBatchFeedbackRequest) SetPlatform(Platform string) {
	m.Platform = Platform
}

func (m *ZhimaDataBatchFeedbackRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaDataBatchFeedbackRequest) SetExtParams(ExtParams string) {
	m.ExtParams = ExtParams
}

func (m *ZhimaDataBatchFeedbackRequest) GetExtParams() string {
	return m.ExtParams
}
