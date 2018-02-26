package requests

import (
	"github.com/WeberLong/go-zmxy/zmop/interfaces"
)

type ZhimaMerchantEnterpriseApplyRequest struct {
	aliasName                  string // 商户别名
	alipayWindowName           string // 为用户提供服务的支付宝服务窗名称，如有请填写 支付宝服务窗，微信公众号，网站地址，应用名称必须是4选1
	appName                    string // 为用户提供服务的手机应用，如有请填写 支付宝服务窗，微信公众号，网站地址，应用名称必须是4选1
	applyMemo                  string // 用于记录本次提交信息中，那些字段有所调整
	authLetterUrl              string // 业务授权书url，请通过zhima.merchant.image.upload上传文件；在商户类型为政府机构或者事业单位时，与证照图片二选一；在商户类型为企业时，与法人信息二选一
	bizScene                   string // 签约芝麻信用产品的用途，最多5个场景，ISV可以自定义
	companyAddress             string // 企业地址
	companyName                string // 企业名称
	dataFeedbackContractEmail  string // 数据反馈联系人邮箱地址(使用芝麻信用评分、行业关注名单时，联动解决数据问题的接口人。)
	dataFeedbackContractMobile string // 数据反馈联系人手机号码(使用芝麻信用评分、行业关注名单时，联动解决数据问题的接口人。)
	dataFeedbackContractName   string // 数据反馈联系人姓名(使用芝麻信用评分、行业关注名单时，联动解决数据问题的接口人。)
	endBusinessDate            string // 证照结束日期9999-12-31   代表长期
	legalCertNo                string // 法人身份证号码在商户类型为企业时，与业务授权书二选一
	legalCertNoBackImageUrl    string // 法人身份证反面图片url请通过zhima.merchant.image.upload上传文件;在商户类型为企业时，与业务授权书二选一
	legalCertNoFrontImageUrl   string // 法人身份证正面图片url请通过zhima.merchant.image.upload上传文件;在商户类型为企业时，与业务授权书二选一
	legalCertValidEndDate      string // 法人身份证有效期结束日期;在商户类型为企业时，与业务授权书二选一
	legalCertValidStartDate    string // 法人身份证有效期开始日期;在商户类型为企业时，与业务授权书二选一
	legalName                  string // 法人姓名在商户类型为企业时，与业务授权书二选一
	licenseImageUrl            string // 证照图片url，请通过zhima.merchant.image.upload上传对应证照类型的图片，在商户类型为政府机构或者事业单位时，与业务授权书二选一
	licenseNo                  string // 证照号码
	licenseType                string // 证照类型：1:普通营业执照2:多证合一3:组织结构代码证4:其他证照company_type=1时，证照类型只能选择1和2；company_type=2,3时，证照类型只能选择3和4
	logoImage                  string // 芝麻二级商户图标的二进制流,最大100K(80*80),需要对图片的二进制流做Base64处理转化成字符串
	logoImageType              string // 芝麻二级商户图标后缀：bmp, jpg, jpeg, png, gif
	majorContractEmail         string // 主要联系人邮箱地址(联系人将用于接收重要通知，如签约进度、技术集成、合同到期等)
	majorContractMobile        string // 主要联系人手机号码(联系人将用于接收重要通知，如签约进度、技术集成、合同到期等)
	majorContractName          string // 主要联系人姓名(联系人将用于接收重要通知，如签约进度、技术集成、合同到期等)
	merchantType               string // 商户类型1:企业2:政府机构3:事业单位
	objectionContractEmail     string // 异议处理联系人邮箱地址(用户对所披露的负面信息存在异议时，联动核查的接口人)
	objectionContractMobile    string // 异议处理联系人手机号码(用户对所披露的负面信息存在异议时，联动核查的接口人)
	objectionContractName      string // 异议处理联系人姓名(用户对所披露的负面信息存在异议时，联动核查的接口人)
	oneLevelMcc                string // 芝麻商户一级行业归属，芝麻提供
	organizationImageUrl       string // 组织结构代码证图片url,营业执照时普通类型时(废弃)
	organizationNo             string // 组织机构代码，营业执照时普通类型时(废弃)
	outOrderNo                 string // 外部订单号
	proxyCertNo                string // 代理人身份证号码(废弃)
	proxyCertNoBackImageUrl    string // 代理人身份证反面图片url(废弃)
	proxyCertNoFrontImageUrl   string // 代理人身份证正面图片url(废弃)
	proxyCertValidEndDate      string // 代理人身份证有效期结束日期(废弃)
	proxyCertValidStartDate    string // 代理人身份证有效期开始日期(废弃)
	proxyMandateUrl            string // 代理人委托书url地址(废弃)
	proxyName                  string // 代理人姓名(废弃)
	qualificationImageUrl      string // 业务许可证图片url请通过zhima.merchant.image.upload上传文件
	serviceContractEmail       string // 服务联动联系人邮箱地址(联动解决用户服务相关问题的接口人。比如用户投诉或出现重大服务事件时，可以协调解决问题的联系人)
	serviceContractMobile      string // 服务联动联系人手机号码(联动解决用户服务相关问题的接口人。比如用户投诉或出现重大服务事件时，可以协调解决问题的联系人)
	serviceContractName        string // 服务联动联系人姓名(联动解决用户服务相关问题的接口人。比如用户投诉或出现重大服务事件时，可以协调解决问题的联系人)
	startBusinessDate          string // 证照开始日期
	twoLevelMcc                string // 芝麻商户二级行业归属，芝麻提供
	webchatAmount              string // 用户提供服务的微信公众号，如有请填写 支付宝服务窗，微信公众号，网站地址，应用名称必须是4选1
	websitUrl                  string // 为用户提供服务的网站，如有请填写 支付宝服务窗，微信公众号，网站地址，应用名称必须是4选1

	interfaces.ZhimaRequestParams
}

func (m *ZhimaMerchantEnterpriseApplyRequest) InitBizParams(aliasName, alipayWindowName, appName, applyMemo, authLetterUrl, bizScene, companyAddress, companyName, dataFeedbackContractEmail, dataFeedbackContractMobile, dataFeedbackContractName, endBusinessDate, legalCertNo, legalCertNoBackImageUrl, legalCertNoFrontImageUrl, legalCertValidEndDate, legalCertValidStartDate, legalName, licenseImageUrl, licenseNo, licenseType, logoImage, logoImageType, majorContractEmail, majorContractMobile, majorContractName, merchantType, objectionContractEmail, objectionContractMobile, objectionContractName, oneLevelMcc, organizationImageUrl, organizationNo, outOrderNo, proxyCertNo, proxyCertNoBackImageUrl, proxyCertNoFrontImageUrl, proxyCertValidEndDate, proxyCertValidStartDate, proxyMandateUrl, proxyName, qualificationImageUrl, serviceContractEmail, serviceContractMobile, serviceContractName, startBusinessDate, twoLevelMcc, webchatAmount, websitUrl string) {
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

	(*m.ApiParams)["auth_letter_url"] = authLetterUrl
	m.authLetterUrl = authLetterUrl

	(*m.ApiParams)["biz_scene"] = bizScene
	m.bizScene = bizScene

	(*m.ApiParams)["company_address"] = companyAddress
	m.companyAddress = companyAddress

	(*m.ApiParams)["company_name"] = companyName
	m.companyName = companyName

	(*m.ApiParams)["data_feedback_contract_email"] = dataFeedbackContractEmail
	m.dataFeedbackContractEmail = dataFeedbackContractEmail

	(*m.ApiParams)["data_feedback_contract_mobile"] = dataFeedbackContractMobile
	m.dataFeedbackContractMobile = dataFeedbackContractMobile

	(*m.ApiParams)["data_feedback_contract_name"] = dataFeedbackContractName
	m.dataFeedbackContractName = dataFeedbackContractName

	(*m.ApiParams)["end_business_date"] = endBusinessDate
	m.endBusinessDate = endBusinessDate

	(*m.ApiParams)["legal_cert_no"] = legalCertNo
	m.legalCertNo = legalCertNo

	(*m.ApiParams)["legal_cert_no_back_image_url"] = legalCertNoBackImageUrl
	m.legalCertNoBackImageUrl = legalCertNoBackImageUrl

	(*m.ApiParams)["legal_cert_no_front_image_url"] = legalCertNoFrontImageUrl
	m.legalCertNoFrontImageUrl = legalCertNoFrontImageUrl

	(*m.ApiParams)["legal_cert_valid_end_date"] = legalCertValidEndDate
	m.legalCertValidEndDate = legalCertValidEndDate

	(*m.ApiParams)["legal_cert_valid_start_date"] = legalCertValidStartDate
	m.legalCertValidStartDate = legalCertValidStartDate

	(*m.ApiParams)["legal_name"] = legalName
	m.legalName = legalName

	(*m.ApiParams)["license_image_url"] = licenseImageUrl
	m.licenseImageUrl = licenseImageUrl

	(*m.ApiParams)["license_no"] = licenseNo
	m.licenseNo = licenseNo

	(*m.ApiParams)["license_type"] = licenseType
	m.licenseType = licenseType

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

	(*m.ApiParams)["merchant_type"] = merchantType
	m.merchantType = merchantType

	(*m.ApiParams)["objection_contract_email"] = objectionContractEmail
	m.objectionContractEmail = objectionContractEmail

	(*m.ApiParams)["objection_contract_mobile"] = objectionContractMobile
	m.objectionContractMobile = objectionContractMobile

	(*m.ApiParams)["objection_contract_name"] = objectionContractName
	m.objectionContractName = objectionContractName

	(*m.ApiParams)["one_level_mcc"] = oneLevelMcc
	m.oneLevelMcc = oneLevelMcc

	(*m.ApiParams)["organization_image_url"] = organizationImageUrl
	m.organizationImageUrl = organizationImageUrl

	(*m.ApiParams)["organization_no"] = organizationNo
	m.organizationNo = organizationNo

	(*m.ApiParams)["out_order_no"] = outOrderNo
	m.outOrderNo = outOrderNo

	(*m.ApiParams)["proxy_cert_no"] = proxyCertNo
	m.proxyCertNo = proxyCertNo

	(*m.ApiParams)["proxy_cert_no_back_image_url"] = proxyCertNoBackImageUrl
	m.proxyCertNoBackImageUrl = proxyCertNoBackImageUrl

	(*m.ApiParams)["proxy_cert_no_front_image_url"] = proxyCertNoFrontImageUrl
	m.proxyCertNoFrontImageUrl = proxyCertNoFrontImageUrl

	(*m.ApiParams)["proxy_cert_valid_end_date"] = proxyCertValidEndDate
	m.proxyCertValidEndDate = proxyCertValidEndDate

	(*m.ApiParams)["proxy_cert_valid_start_date"] = proxyCertValidStartDate
	m.proxyCertValidStartDate = proxyCertValidStartDate

	(*m.ApiParams)["proxy_mandate_url"] = proxyMandateUrl
	m.proxyMandateUrl = proxyMandateUrl

	(*m.ApiParams)["proxy_name"] = proxyName
	m.proxyName = proxyName

	(*m.ApiParams)["qualification_image_url"] = qualificationImageUrl
	m.qualificationImageUrl = qualificationImageUrl

	(*m.ApiParams)["service_contract_email"] = serviceContractEmail
	m.serviceContractEmail = serviceContractEmail

	(*m.ApiParams)["service_contract_mobile"] = serviceContractMobile
	m.serviceContractMobile = serviceContractMobile

	(*m.ApiParams)["service_contract_name"] = serviceContractName
	m.serviceContractName = serviceContractName

	(*m.ApiParams)["start_business_date"] = startBusinessDate
	m.startBusinessDate = startBusinessDate

	(*m.ApiParams)["two_level_mcc"] = twoLevelMcc
	m.twoLevelMcc = twoLevelMcc

	(*m.ApiParams)["webchat_amount"] = webchatAmount
	m.webchatAmount = webchatAmount

	(*m.ApiParams)["websit_url"] = websitUrl
	m.websitUrl = websitUrl
}

func (*ZhimaMerchantEnterpriseApplyRequest) GetApiMethodName() string {
	return "zhima.merchant.enterprise.apply"
}

func (m *ZhimaMerchantEnterpriseApplyRequest) GetApiParams() *map[string]string {
	return m.ApiParams
}

func (m *ZhimaMerchantEnterpriseApplyRequest) GetFileParams() *map[string]string {
	return m.FileParams
}

func (m *ZhimaMerchantEnterpriseApplyRequest) SetApiVersion(ApiVersion string) {
	m.ApiVersion = ApiVersion
}

func (m *ZhimaMerchantEnterpriseApplyRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaMerchantEnterpriseApplyRequest) SetScene(Scene string) {
	m.Scene = Scene
}

func (m *ZhimaMerchantEnterpriseApplyRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaMerchantEnterpriseApplyRequest) SetChannel(Channel string) {
	m.Channel = Channel
}

func (m *ZhimaMerchantEnterpriseApplyRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaMerchantEnterpriseApplyRequest) SetPlatform(Platform string) {
	m.Platform = Platform
}

func (m *ZhimaMerchantEnterpriseApplyRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaMerchantEnterpriseApplyRequest) SetExtParams(ExtParams string) {
	m.ExtParams = ExtParams
}

func (m *ZhimaMerchantEnterpriseApplyRequest) GetExtParams() string {
	return m.ExtParams
}
