package requests

import (
	"github.com/WeberLong/go-zmxy/zmop/interfaces"
)

type ZhimaProfileStatusFeedbackRequest struct {
	bizNo      string // 业务号
	bizType    string // 业务类型
	dataSource string // 当前为BANK或者ACCUFUND
	itemCode   string // 数据抓取code
	memo       string // 更新备注
	orgCode    string // 数据采集机构
	status     string // 数据状态

	interfaces.ZhimaRequestParams
}

func (m *ZhimaProfileStatusFeedbackRequest) InitBizParams(bizNo, bizType, dataSource, itemCode, memo, orgCode, status string) {
	m.ApiParams = &map[string]string{}
	m.FileParams = &map[string]string{}

	(*m.ApiParams)["biz_no"] = bizNo
	m.bizNo = bizNo

	(*m.ApiParams)["biz_type"] = bizType
	m.bizType = bizType

	(*m.ApiParams)["data_source"] = dataSource
	m.dataSource = dataSource

	(*m.ApiParams)["item_code"] = itemCode
	m.itemCode = itemCode

	(*m.ApiParams)["memo"] = memo
	m.memo = memo

	(*m.ApiParams)["org_code"] = orgCode
	m.orgCode = orgCode

	(*m.ApiParams)["status"] = status
	m.status = status
}

func (*ZhimaProfileStatusFeedbackRequest) GetApiMethodName() string {
	return "zhima.profile.status.feedback"
}

func (m *ZhimaProfileStatusFeedbackRequest) GetApiParams() *map[string]string {
	return m.ApiParams
}

func (m *ZhimaProfileStatusFeedbackRequest) GetFileParams() *map[string]string {
	return m.FileParams
}

func (m *ZhimaProfileStatusFeedbackRequest) SetApiVersion(ApiVersion string) {
	m.ApiVersion = ApiVersion
}

func (m *ZhimaProfileStatusFeedbackRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaProfileStatusFeedbackRequest) SetScene(Scene string) {
	m.Scene = Scene
}

func (m *ZhimaProfileStatusFeedbackRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaProfileStatusFeedbackRequest) SetChannel(Channel string) {
	m.Channel = Channel
}

func (m *ZhimaProfileStatusFeedbackRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaProfileStatusFeedbackRequest) SetPlatform(Platform string) {
	m.Platform = Platform
}

func (m *ZhimaProfileStatusFeedbackRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaProfileStatusFeedbackRequest) SetExtParams(ExtParams string) {
	m.ExtParams = ExtParams
}

func (m *ZhimaProfileStatusFeedbackRequest) GetExtParams() string {
	return m.ExtParams
}
