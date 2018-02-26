package requests

import (
	"github.com/WeberLong/go-zmxy/zmop/interfaces"
)

type ZhimaCustomerCertificationResultSyncRequest struct {
	bizNo               string // 一次认证的唯一标识,在商户调用认证初始化接口的时候获取
	channelStatuses     string // 各渠道认证状态,失败原因,材料等信息
	errorCode           string // 认证失败的错误码
	errorMessage        string // 认证失败的错误信息
	extBizParam         string // 扩展业务参数
	identifiedPrincipal string // 识别后的主体信息,入参有传就和入参的certify_identity一致
	merchantId          string // 商户id,商户在芝麻的唯一标识
	passed              string // 认证是否通过

	interfaces.ZhimaRequestParams
}

func (m *ZhimaCustomerCertificationResultSyncRequest) InitBizParams(bizNo, channelStatuses, errorCode, errorMessage, extBizParam, identifiedPrincipal, merchantId, passed string) {
	m.ApiParams = &map[string]string{}
	m.FileParams = &map[string]string{}

	(*m.ApiParams)["biz_no"] = bizNo
	m.bizNo = bizNo

	(*m.ApiParams)["channel_statuses"] = channelStatuses
	m.channelStatuses = channelStatuses

	(*m.ApiParams)["error_code"] = errorCode
	m.errorCode = errorCode

	(*m.ApiParams)["error_message"] = errorMessage
	m.errorMessage = errorMessage

	(*m.ApiParams)["ext_biz_param"] = extBizParam
	m.extBizParam = extBizParam

	(*m.ApiParams)["identified_principal"] = identifiedPrincipal
	m.identifiedPrincipal = identifiedPrincipal

	(*m.ApiParams)["merchant_id"] = merchantId
	m.merchantId = merchantId

	(*m.ApiParams)["passed"] = passed
	m.passed = passed
}

func (*ZhimaCustomerCertificationResultSyncRequest) GetApiMethodName() string {
	return "zhima.customer.certification.result.sync"
}

func (m *ZhimaCustomerCertificationResultSyncRequest) GetApiParams() *map[string]string {
	return m.ApiParams
}

func (m *ZhimaCustomerCertificationResultSyncRequest) GetFileParams() *map[string]string {
	return m.FileParams
}

func (m *ZhimaCustomerCertificationResultSyncRequest) SetApiVersion(ApiVersion string) {
	m.ApiVersion = ApiVersion
}

func (m *ZhimaCustomerCertificationResultSyncRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaCustomerCertificationResultSyncRequest) SetScene(Scene string) {
	m.Scene = Scene
}

func (m *ZhimaCustomerCertificationResultSyncRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaCustomerCertificationResultSyncRequest) SetChannel(Channel string) {
	m.Channel = Channel
}

func (m *ZhimaCustomerCertificationResultSyncRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaCustomerCertificationResultSyncRequest) SetPlatform(Platform string) {
	m.Platform = Platform
}

func (m *ZhimaCustomerCertificationResultSyncRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaCustomerCertificationResultSyncRequest) SetExtParams(ExtParams string) {
	m.ExtParams = ExtParams
}

func (m *ZhimaCustomerCertificationResultSyncRequest) GetExtParams() string {
	return m.ExtParams
}
