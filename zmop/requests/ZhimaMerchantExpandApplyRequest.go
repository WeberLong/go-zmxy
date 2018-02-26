package requests

import (
	"github.com/WeberLong/go-zmxy/zmop/interfaces"
)

type ZhimaMerchantExpandApplyRequest struct {
	aliasName                  string // 浙江飞猪网络技术有限公司，企业别称请填写【飞猪】
	alipayWindowName           string // 为用户提供服务的支付宝服务窗名称，如有请填写支付宝服务窗，微信公众号，网站地址，应用名称必须是4选1
	appName                    string // 为用户提供服务的手机应用，如有请填写支付宝服务窗，微信公众号，网站地址，应用名称必须是4选1
	applyMemo                  string // 用于记录本次提交信息中，那些字段有所调整
	bizScene                   string // 签约芝麻信用产品的用途，最多5个场景，ISV可以自定义
	dataFeedbackContractEmail  string // 数据反馈联系人邮箱地址(使用芝麻信用评分、行业关注名单时，联动解决数据问题的接口人。)
	dataFeedbackContractMobile string // 数据反馈联系人手机号码(使用芝麻信用评分、行业关注名单时，联动解决数据问题的接口人。)
	dataFeedbackContractName   string // 数据反馈联系人姓名(使用芝麻信用评分、行业关注名单时，联动解决数据问题的接口人。)
	logoImage                  string // 商户logo的图片内容，把图片的二进制转化成String传递过来，最大2M
	logoImageType              string // 商户图标类型，只支持如下格式：bmp, jpg, jpeg, png, gif
	majorContractEmail         string // 主要联系人邮箱地址(联系人将用于接收重要通知，如签约进度、技术集成、合同到期等)
	majorContractMobile        string // 主要联系人手机号码(联系人将用于接收重要通知，如签约进度、技术集成、合同到期等)
	majorContractName          string // 主要联系人姓名(联系人将用于接收重要通知，如签约进度、技术集成、合同到期等)
	objectionContractEmail     string // 异议处理联系人邮箱地址(用户对所披露的负面信息存在异议时，联动核查的接口人)
	objectionContractMobile    string // 异议处理联系人手机号码(用户对所披露的负面信息存在异议时，联动核查的接口人)
	objectionContractName      string // 异议处理联系人姓名(用户对所披露的负面信息存在异议时，联动核查的接口人)
	oneLevelMcc                string // 芝麻特定的商户一级行业归属，比如生活行业，金融行业，政府行业
	qualificationImage         string // 商户业务许可证图片内容，把图片的二进制转化成String传递过来，最大2M
	qualificationImageType     string // 商户业务许可证图片类型，只支持如下格式：bmp, jpg, jpeg, png, gif
	serviceContractEmail       string // 服务联动联系人邮箱地址(联动解决用户服务相关问题的接口人。比如用户投诉或出现重大服务事件时，可以协调解决问题的联系人)
	serviceContractMobile      string // 服务联动联系人手机号码(联动解决用户服务相关问题的接口人。比如用户投诉或出现重大服务事件时，可以协调解决问题的联系人)
	serviceContractName        string // 服务联动联系人姓名(联动解决用户服务相关问题的接口人。比如用户投诉或出现重大服务事件时，可以协调解决问题的联系人)
	twoLevelMcc                string // 芝麻特有的商户二级行业归属，比如汽车服务
	webchatAmount              string // 用户提供服务的微信公众号，如有请填写支付宝服务窗，微信公众号，网站地址，应用名称必须是4选1
	websitUrl                  string // 为用户提供服务的网站，如有请填写支付宝服务窗，微信公众号，网站地址，应用名称必须是4选1

	interfaces.ZhimaRequestParams
}

func (m *ZhimaMerchantExpandApplyRequest) InitBizParams(aliasName, alipayWindowName, appName, applyMemo, bizScene, dataFeedbackContractEmail, dataFeedbackContractMobile, dataFeedbackContractName, logoImage, logoImageType, majorContractEmail, majorContractMobile, majorContractName, objectionContractEmail, objectionContractMobile, objectionContractName, oneLevelMcc, qualificationImage, qualificationImageType, serviceContractEmail, serviceContractMobile, serviceContractName, twoLevelMcc, webchatAmount, websitUrl string) {
	m.ApiParams = &map[string]string{}
	m.FileParams = &map[string]string{}

	(*m.ApiParams)["alias_name"] = aliasName
	m.aliasName = aliasName

	(*m.ApiParams)["alipay_window_name"] = alipayWindowName
	m.alipayWindowName = alipayWindowName

	(*m.ApiParams)["app_name"] = appName
	m.appName = appName

	(*m.ApiParams)["apply_memo"] = applyMemo
	m.applyMemo = applyMemo

	(*m.ApiParams)["biz_scene"] = bizScene
	m.bizScene = bizScene

	(*m.ApiParams)["data_feedback_contract_email"] = dataFeedbackContractEmail
	m.dataFeedbackContractEmail = dataFeedbackContractEmail

	(*m.ApiParams)["data_feedback_contract_mobile"] = dataFeedbackContractMobile
	m.dataFeedbackContractMobile = dataFeedbackContractMobile

	(*m.ApiParams)["data_feedback_contract_name"] = dataFeedbackContractName
	m.dataFeedbackContractName = dataFeedbackContractName

	(*m.ApiParams)["logo_image"] = logoImage
	m.logoImage = logoImage

	(*m.ApiParams)["logo_image_type"] = logoImageType
	m.logoImageType = logoImageType

	(*m.ApiParams)["major_contract_email"] = majorContractEmail
	m.majorContractEmail = majorContractEmail

	(*m.ApiParams)["major_contract_mobile"] = majorContractMobile
	m.majorContractMobile = majorContractMobile

	(*m.ApiParams)["major_contract_name"] = majorContractName
	m.majorContractName = majorContractName

	(*m.ApiParams)["objection_contract_email"] = objectionContractEmail
	m.objectionContractEmail = objectionContractEmail

	(*m.ApiParams)["objection_contract_mobile"] = objectionContractMobile
	m.objectionContractMobile = objectionContractMobile

	(*m.ApiParams)["objection_contract_name"] = objectionContractName
	m.objectionContractName = objectionContractName

	(*m.ApiParams)["one_level_mcc"] = oneLevelMcc
	m.oneLevelMcc = oneLevelMcc

	(*m.ApiParams)["qualification_image"] = qualificationImage
	m.qualificationImage = qualificationImage

	(*m.ApiParams)["qualification_image_type"] = qualificationImageType
	m.qualificationImageType = qualificationImageType

	(*m.ApiParams)["service_contract_email"] = serviceContractEmail
	m.serviceContractEmail = serviceContractEmail

	(*m.ApiParams)["service_contract_mobile"] = serviceContractMobile
	m.serviceContractMobile = serviceContractMobile

	(*m.ApiParams)["service_contract_name"] = serviceContractName
	m.serviceContractName = serviceContractName

	(*m.ApiParams)["two_level_mcc"] = twoLevelMcc
	m.twoLevelMcc = twoLevelMcc

	(*m.ApiParams)["webchat_amount"] = webchatAmount
	m.webchatAmount = webchatAmount

	(*m.ApiParams)["websit_url"] = websitUrl
	m.websitUrl = websitUrl
}

func (*ZhimaMerchantExpandApplyRequest) GetApiMethodName() string {
	return "zhima.merchant.expand.apply"
}

func (m *ZhimaMerchantExpandApplyRequest) GetApiParams() *map[string]string {
	return m.ApiParams
}

func (m *ZhimaMerchantExpandApplyRequest) GetFileParams() *map[string]string {
	return m.FileParams
}

func (m *ZhimaMerchantExpandApplyRequest) SetApiVersion(ApiVersion string) {
	m.ApiVersion = ApiVersion
}

func (m *ZhimaMerchantExpandApplyRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaMerchantExpandApplyRequest) SetScene(Scene string) {
	m.Scene = Scene
}

func (m *ZhimaMerchantExpandApplyRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaMerchantExpandApplyRequest) SetChannel(Channel string) {
	m.Channel = Channel
}

func (m *ZhimaMerchantExpandApplyRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaMerchantExpandApplyRequest) SetPlatform(Platform string) {
	m.Platform = Platform
}

func (m *ZhimaMerchantExpandApplyRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaMerchantExpandApplyRequest) SetExtParams(ExtParams string) {
	m.ExtParams = ExtParams
}

func (m *ZhimaMerchantExpandApplyRequest) GetExtParams() string {
	return m.ExtParams
}
