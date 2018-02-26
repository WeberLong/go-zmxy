package requests

import (
	"github.com/WeberLong/go-zmxy/zmop/interfaces"
)

type ZhimaCreditMsxfOnlinejdscoreQueryRequest struct {
	allFddifDivideSixaQdHourbinAmtaorder         string // 特殊订单金额差异占比
	allFddifMinusFiveaRangeHourbinAmtaorder      string // 短时间下单金额的差异系数
	allFddifMinusTwoaSdHourbinAmtaorder          string // 特殊时间下单金额的波动指标
	allFdescMeanPayonlinepaymentAmtorder         string // 特定支付方式金额指标
	allGddescMinusLoantimenowMaxDaydiff          string // 用户购买时间波动系数1
	allGddescMinusLoantimenowMinHourdiff         string // 用户购买时间波动系数2
	allGddifDivCashondeliveryallSumPayAmtorder   string // 用户特殊支付金额占比
	allGddifDivOnlinepaymentallSumPayAmtorder    string // 用户特殊支付金额指标
	allGddifDivSportsoutdoorallNCntprdcategory   string // 用户特殊商品差异性指标
	allGddifDivideFailallNStsCntorder            string // 用户特殊订单占比
	allGddifDivideFiveeightallNHourCntorder      string // 用户特殊时间下单量指标
	allGddifDividePhonedigitalallNCntprdcategory string // 用户特殊商品差异性系数
	allGddifMinusCaMaxProductCntaorder           string // 用户特殊订单量指标
	allGddifMinusCaSumAorderCntproduct           string // 用户特殊订单的差异性指标
	allGddifMinusCsMedianProductCntaorder        string // 用户特殊产品之间差异系数
	allGddifMinusCsSkewAorderAmtaorder           string // 用户特殊订单下单金额异常指标
	allGddifMinusSaEntropyAorderCntproduct       string // 用户购买商品的波动性指标
	allGddifMinusSaSumAorderCntproduct           string // 用户购买商品量
	allGddifMinusSaSumProductCntaorder           string // 用户购买固定商品的稳定性
	allGdescMeanProductCntaorder                 string // 用户购买固定商品的差异
	allGdescMeanSorderAmtaorder                  string // 用户购买特殊订单的数量
	allGdescMinCorderAmtaorder                   string // 用户特殊订单金额指标
	allGdescMinPhoneSumamt                       string // 用户下单稳定性系数
	allGdescMinRecaddrcitySumamt                 string // 用户购买稳定性指标
	allGdescMinRecaddrprovinceAvgamt             string // 用户下单稳定性
	allGdescNormentropyPhoneCntorder             string // 用户购物行为稳定性指标
	allGdescNormentropyProductCntsorder          string // 用户特殊购买行为稳定性指标
	allGdescQlSorderAmtaorder                    string // 用户特殊订单容量指标
	allTsdescAmtorderdiffAmtdiffMedian           string // 用户下单跨度行为指标
	allTsdescAmtorderdiffAmtdiffQu               string // 用户下单跨度行为稳定性
	allTsdescAmtorderdiffAmtdiffSum              string // 用户下单跨度行为波动性
	allTsdescAmtorderdiffTimediffCv              string // 用户下单跨度特殊差异性
	allTsdescAmtorderdiffTimediffQfour           string // 用户下单跨度可靠性
	allTsdescAmtorderdiffTimediffQsix            string // 用户下单金额差异稳定度
	allTsdescAmtorderdiffTimediffQu              string // 用户下单时间稳定度
	allTsdescAmtorderdiffVamtQnine               string // 用户下单行为差异度
	jdauthFddescExistChannelfinanceAuth          string // 用户可信度指标
	jdauthFddescExistLoginnameEqualPhone         string // 用户授信稳定性指标
	jdauthFddescMinusNowauthtimeSeconds          string // 用户信用欺诈指标
	jdbankcardDescDivideNOwnernameReceiver       string // 用户信用稳定性指标
	jdbankcardDescNBankphoneAuthphone            string // 用户可信度差异
	jdbankcardDescNOwnernameReceiver             string // 用户可信度波动系数
	jdbankcardDiffDivideNndBindphone             string // 用户稳定性支付系数
	jdbankcardFdescNBanknameMajorfourbanks       string // 用户主流支付差异
	jdbankcardFdescNBanknameOthers               string // 用户支付信用系数
	jdbankcardFdiffDivideAbcallCntbankname       string // 用户支付差异系数
	jdbankcardFdiffDivideCreditallCntcardtype    string // 用户信用卡稳定性
	jdbankcardFdiffDividePostallCntbankname      string // 用户支付稳定性
	jdbtGddescExtractCreditscoreBt               string // 用户信用指标
	jdbtGddiffMinusOverdraftquotaBtAmt           string // 用户信用稳定系数
	jdoneoneoneonesumGdescAmt                    string // 用户活动金额指标
	jdreceivaddrDescNAddress                     string // 用户居住稳定性指标
	jdreceivaddrDescNNaemail                     string // 用户收货地址差异系数
	jdreceivaddrDescRateNafixphone               string // 用户收货地址稳定性指标
	jdsixoneeightsumGdescAmt                     string // 用户活动金额范围系数
	jduserFddescExistWebloginnameLogname         string // 用户注册差异性指标
	jduserFddescNdCompareThreenames              string // 用户下单时间金额总共的时间精度
	jduserIsbindBothqqwechat                     string // 用户的绑定粘性指标
	oneyFddifDivideSevenaRangeHourbinAmtaorder   string // 1年内短时间金额稳定性指标
	oneyFddifMinusOneaRangeHourbinAmtaorder      string // 1年内短时间金额占比
	oneyFddifMinusSixaRangeHourbinAmtaorder      string // 1年内短时间订单金额稳定性指标
	oneyFdescMeanPaycashondeliveryAmtorder       string // 1年内特殊订单金额平均水平
	oneyFdescSumMeaninvoicecontentAmtorder       string // 1年内特殊订单金额异常指标
	oneyGddifDivOnlinepaymentallSumPayAmtorder   string // 1年内在线支付金额占比
	oneyGddifMinusCaMedianAorderAmtaorder        string // 1年内特殊订单购买能力
	oneyGddifMinusCaSdAmtbinAmtaorder            string // 1年内取消订单订单金额差异性指标
	oneyGddifMinusCaSumAorderCntproduct          string // 1年内订单数量总和差异
	oneyGddifMinusSaEntropyAmtbinAmtaorder       string // 1年内特殊订单金额波动性指标
	oneyGdescCvRecaddrcityAvgamt                 string // 1年内地址差异指标
	oneyGdescNormentropyAmtbinAmtsorder          string // 1年内特殊订单金额分段差异性指标
	oneyTsdescAmtorderdiffTimediffQsix           string // 1年内订单金额特殊时间差异性系数
	oneyTsdescAmtorderdiffVamtRange              string // 1年内下单时间金额波动指标
	openId                                       string // 芝麻会员在商户端的身份标识。
	productCode                                  string // 产品码
	sixmFdescCvHourCntorder                      string // 6月内特殊时间下波动性指标
	sixmGddifDivOnlinepaymentallSumPayAmtorder   string // 6月内在线支付总金额的占比
	sixmGddifDivPhonedigitalallNCntprdcategory   string // 6月内电子产品类目占比
	sixmGddifDivSixmallNHourtwefourteenCntorder  string // 6月内特殊下单量的占比
	sixmGddifDivideFiveeightallNHourCntorder     string // 6月内短时间下单占比
	sixmGddifMinusCaSumAorderCntproduct          string // 6月内异常商品占比
	sixmGdescMinRecaddrcityAvgamt                string // 6月内收货地址平均下单量的差异性指标
	sixmGdescRangeRecaddrprovinceAvgamt          string // 6月内收货地址稳定性指标
	springfestivalGdescQuAamt                    string // 用户活动期间的下单系数
	threemFddifMinusSevenaQdHourbinAmtaorder     string // 3个月内特殊时段购买能力
	threemGddifDivTravelrechargeallNCntprdcateg  string // 3月内特殊用途商品占比
	threemGddifDivideFailallNStsCntorder         string // 3月内异常订单占比
	threemGddifDivideNullallSumPayAmtorder       string // 3月内金额总和异常占比
	threemGdescSumSorderAmtaorder                string // 3月内特殊订单金额指标
	transactionId                                string // 商户传入的业务流水号。此字段由商户生成，需确保唯一性，用于定位每一次请求，后续按此流水进行对帐。生成规则: 固定30位数字串，前17位为精确到毫秒的时间yyyyMMddhhmmssSSS，后13位为自增数字。

	interfaces.ZhimaRequestParams
}

func (m *ZhimaCreditMsxfOnlinejdscoreQueryRequest) InitBizParams(allFddifDivideSixaQdHourbinAmtaorder, allFddifMinusFiveaRangeHourbinAmtaorder, allFddifMinusTwoaSdHourbinAmtaorder, allFdescMeanPayonlinepaymentAmtorder, allGddescMinusLoantimenowMaxDaydiff, allGddescMinusLoantimenowMinHourdiff, allGddifDivCashondeliveryallSumPayAmtorder, allGddifDivOnlinepaymentallSumPayAmtorder, allGddifDivSportsoutdoorallNCntprdcategory, allGddifDivideFailallNStsCntorder, allGddifDivideFiveeightallNHourCntorder, allGddifDividePhonedigitalallNCntprdcategory, allGddifMinusCaMaxProductCntaorder, allGddifMinusCaSumAorderCntproduct, allGddifMinusCsMedianProductCntaorder, allGddifMinusCsSkewAorderAmtaorder, allGddifMinusSaEntropyAorderCntproduct, allGddifMinusSaSumAorderCntproduct, allGddifMinusSaSumProductCntaorder, allGdescMeanProductCntaorder, allGdescMeanSorderAmtaorder, allGdescMinCorderAmtaorder, allGdescMinPhoneSumamt, allGdescMinRecaddrcitySumamt, allGdescMinRecaddrprovinceAvgamt, allGdescNormentropyPhoneCntorder, allGdescNormentropyProductCntsorder, allGdescQlSorderAmtaorder, allTsdescAmtorderdiffAmtdiffMedian, allTsdescAmtorderdiffAmtdiffQu, allTsdescAmtorderdiffAmtdiffSum, allTsdescAmtorderdiffTimediffCv, allTsdescAmtorderdiffTimediffQfour, allTsdescAmtorderdiffTimediffQsix, allTsdescAmtorderdiffTimediffQu, allTsdescAmtorderdiffVamtQnine, jdauthFddescExistChannelfinanceAuth, jdauthFddescExistLoginnameEqualPhone, jdauthFddescMinusNowauthtimeSeconds, jdbankcardDescDivideNOwnernameReceiver, jdbankcardDescNBankphoneAuthphone, jdbankcardDescNOwnernameReceiver, jdbankcardDiffDivideNndBindphone, jdbankcardFdescNBanknameMajorfourbanks, jdbankcardFdescNBanknameOthers, jdbankcardFdiffDivideAbcallCntbankname, jdbankcardFdiffDivideCreditallCntcardtype, jdbankcardFdiffDividePostallCntbankname, jdbtGddescExtractCreditscoreBt, jdbtGddiffMinusOverdraftquotaBtAmt, jdoneoneoneonesumGdescAmt, jdreceivaddrDescNAddress, jdreceivaddrDescNNaemail, jdreceivaddrDescRateNafixphone, jdsixoneeightsumGdescAmt, jduserFddescExistWebloginnameLogname, jduserFddescNdCompareThreenames, jduserIsbindBothqqwechat, oneyFddifDivideSevenaRangeHourbinAmtaorder, oneyFddifMinusOneaRangeHourbinAmtaorder, oneyFddifMinusSixaRangeHourbinAmtaorder, oneyFdescMeanPaycashondeliveryAmtorder, oneyFdescSumMeaninvoicecontentAmtorder, oneyGddifDivOnlinepaymentallSumPayAmtorder, oneyGddifMinusCaMedianAorderAmtaorder, oneyGddifMinusCaSdAmtbinAmtaorder, oneyGddifMinusCaSumAorderCntproduct, oneyGddifMinusSaEntropyAmtbinAmtaorder, oneyGdescCvRecaddrcityAvgamt, oneyGdescNormentropyAmtbinAmtsorder, oneyTsdescAmtorderdiffTimediffQsix, oneyTsdescAmtorderdiffVamtRange, openId, productCode, sixmFdescCvHourCntorder, sixmGddifDivOnlinepaymentallSumPayAmtorder, sixmGddifDivPhonedigitalallNCntprdcategory, sixmGddifDivSixmallNHourtwefourteenCntorder, sixmGddifDivideFiveeightallNHourCntorder, sixmGddifMinusCaSumAorderCntproduct, sixmGdescMinRecaddrcityAvgamt, sixmGdescRangeRecaddrprovinceAvgamt, springfestivalGdescQuAamt, threemFddifMinusSevenaQdHourbinAmtaorder, threemGddifDivTravelrechargeallNCntprdcateg, threemGddifDivideFailallNStsCntorder, threemGddifDivideNullallSumPayAmtorder, threemGdescSumSorderAmtaorder, transactionId string) {
	m.ApiParams = &map[string]string{}
	m.FileParams = &map[string]string{}

	(*m.ApiParams)["all_fddif_divide_sixa_qd_hourbin_amtaorder"] = allFddifDivideSixaQdHourbinAmtaorder
	m.allFddifDivideSixaQdHourbinAmtaorder = allFddifDivideSixaQdHourbinAmtaorder

	(*m.ApiParams)["all_fddif_minus_fivea_range_hourbin_amtaorder"] = allFddifMinusFiveaRangeHourbinAmtaorder
	m.allFddifMinusFiveaRangeHourbinAmtaorder = allFddifMinusFiveaRangeHourbinAmtaorder

	(*m.ApiParams)["all_fddif_minus_twoa_sd_hourbin_amtaorder"] = allFddifMinusTwoaSdHourbinAmtaorder
	m.allFddifMinusTwoaSdHourbinAmtaorder = allFddifMinusTwoaSdHourbinAmtaorder

	(*m.ApiParams)["all_fdesc_mean_payonlinepayment_amtorder"] = allFdescMeanPayonlinepaymentAmtorder
	m.allFdescMeanPayonlinepaymentAmtorder = allFdescMeanPayonlinepaymentAmtorder

	(*m.ApiParams)["all_gddesc_minus_loantimenow_max_daydiff"] = allGddescMinusLoantimenowMaxDaydiff
	m.allGddescMinusLoantimenowMaxDaydiff = allGddescMinusLoantimenowMaxDaydiff

	(*m.ApiParams)["all_gddesc_minus_loantimenow_min_hourdiff"] = allGddescMinusLoantimenowMinHourdiff
	m.allGddescMinusLoantimenowMinHourdiff = allGddescMinusLoantimenowMinHourdiff

	(*m.ApiParams)["all_gddif_div_cashondeliveryall_sum_pay_amtorder"] = allGddifDivCashondeliveryallSumPayAmtorder
	m.allGddifDivCashondeliveryallSumPayAmtorder = allGddifDivCashondeliveryallSumPayAmtorder

	(*m.ApiParams)["all_gddif_div_onlinepaymentall_sum_pay_amtorder"] = allGddifDivOnlinepaymentallSumPayAmtorder
	m.allGddifDivOnlinepaymentallSumPayAmtorder = allGddifDivOnlinepaymentallSumPayAmtorder

	(*m.ApiParams)["all_gddif_div_sportsoutdoorall_n_cntprdcategory"] = allGddifDivSportsoutdoorallNCntprdcategory
	m.allGddifDivSportsoutdoorallNCntprdcategory = allGddifDivSportsoutdoorallNCntprdcategory

	(*m.ApiParams)["all_gddif_divide_failall_n_sts_cntorder"] = allGddifDivideFailallNStsCntorder
	m.allGddifDivideFailallNStsCntorder = allGddifDivideFailallNStsCntorder

	(*m.ApiParams)["all_gddif_divide_fiveeightall_n_hour_cntorder"] = allGddifDivideFiveeightallNHourCntorder
	m.allGddifDivideFiveeightallNHourCntorder = allGddifDivideFiveeightallNHourCntorder

	(*m.ApiParams)["all_gddif_divide_phonedigitalall_n_cntprdcategory"] = allGddifDividePhonedigitalallNCntprdcategory
	m.allGddifDividePhonedigitalallNCntprdcategory = allGddifDividePhonedigitalallNCntprdcategory

	(*m.ApiParams)["all_gddif_minus_ca_max_product_cntaorder"] = allGddifMinusCaMaxProductCntaorder
	m.allGddifMinusCaMaxProductCntaorder = allGddifMinusCaMaxProductCntaorder

	(*m.ApiParams)["all_gddif_minus_ca_sum_aorder_cntproduct"] = allGddifMinusCaSumAorderCntproduct
	m.allGddifMinusCaSumAorderCntproduct = allGddifMinusCaSumAorderCntproduct

	(*m.ApiParams)["all_gddif_minus_cs_median_product_cntaorder"] = allGddifMinusCsMedianProductCntaorder
	m.allGddifMinusCsMedianProductCntaorder = allGddifMinusCsMedianProductCntaorder

	(*m.ApiParams)["all_gddif_minus_cs_skew_aorder_amtaorder"] = allGddifMinusCsSkewAorderAmtaorder
	m.allGddifMinusCsSkewAorderAmtaorder = allGddifMinusCsSkewAorderAmtaorder

	(*m.ApiParams)["all_gddif_minus_sa_entropy_aorder_cntproduct"] = allGddifMinusSaEntropyAorderCntproduct
	m.allGddifMinusSaEntropyAorderCntproduct = allGddifMinusSaEntropyAorderCntproduct

	(*m.ApiParams)["all_gddif_minus_sa_sum_aorder_cntproduct"] = allGddifMinusSaSumAorderCntproduct
	m.allGddifMinusSaSumAorderCntproduct = allGddifMinusSaSumAorderCntproduct

	(*m.ApiParams)["all_gddif_minus_sa_sum_product_cntaorder"] = allGddifMinusSaSumProductCntaorder
	m.allGddifMinusSaSumProductCntaorder = allGddifMinusSaSumProductCntaorder

	(*m.ApiParams)["all_gdesc_mean_product_cntaorder"] = allGdescMeanProductCntaorder
	m.allGdescMeanProductCntaorder = allGdescMeanProductCntaorder

	(*m.ApiParams)["all_gdesc_mean_sorder_amtaorder"] = allGdescMeanSorderAmtaorder
	m.allGdescMeanSorderAmtaorder = allGdescMeanSorderAmtaorder

	(*m.ApiParams)["all_gdesc_min_corder_amtaorder"] = allGdescMinCorderAmtaorder
	m.allGdescMinCorderAmtaorder = allGdescMinCorderAmtaorder

	(*m.ApiParams)["all_gdesc_min_phone_sumamt"] = allGdescMinPhoneSumamt
	m.allGdescMinPhoneSumamt = allGdescMinPhoneSumamt

	(*m.ApiParams)["all_gdesc_min_recaddrcity_sumamt"] = allGdescMinRecaddrcitySumamt
	m.allGdescMinRecaddrcitySumamt = allGdescMinRecaddrcitySumamt

	(*m.ApiParams)["all_gdesc_min_recaddrprovince_avgamt"] = allGdescMinRecaddrprovinceAvgamt
	m.allGdescMinRecaddrprovinceAvgamt = allGdescMinRecaddrprovinceAvgamt

	(*m.ApiParams)["all_gdesc_normentropy_phone_cntorder"] = allGdescNormentropyPhoneCntorder
	m.allGdescNormentropyPhoneCntorder = allGdescNormentropyPhoneCntorder

	(*m.ApiParams)["all_gdesc_normentropy_product_cntsorder"] = allGdescNormentropyProductCntsorder
	m.allGdescNormentropyProductCntsorder = allGdescNormentropyProductCntsorder

	(*m.ApiParams)["all_gdesc_ql_sorder_amtaorder"] = allGdescQlSorderAmtaorder
	m.allGdescQlSorderAmtaorder = allGdescQlSorderAmtaorder

	(*m.ApiParams)["all_tsdesc_amtorderdiff_amtdiff_median"] = allTsdescAmtorderdiffAmtdiffMedian
	m.allTsdescAmtorderdiffAmtdiffMedian = allTsdescAmtorderdiffAmtdiffMedian

	(*m.ApiParams)["all_tsdesc_amtorderdiff_amtdiff_qu"] = allTsdescAmtorderdiffAmtdiffQu
	m.allTsdescAmtorderdiffAmtdiffQu = allTsdescAmtorderdiffAmtdiffQu

	(*m.ApiParams)["all_tsdesc_amtorderdiff_amtdiff_sum"] = allTsdescAmtorderdiffAmtdiffSum
	m.allTsdescAmtorderdiffAmtdiffSum = allTsdescAmtorderdiffAmtdiffSum

	(*m.ApiParams)["all_tsdesc_amtorderdiff_timediff_cv"] = allTsdescAmtorderdiffTimediffCv
	m.allTsdescAmtorderdiffTimediffCv = allTsdescAmtorderdiffTimediffCv

	(*m.ApiParams)["all_tsdesc_amtorderdiff_timediff_qfour"] = allTsdescAmtorderdiffTimediffQfour
	m.allTsdescAmtorderdiffTimediffQfour = allTsdescAmtorderdiffTimediffQfour

	(*m.ApiParams)["all_tsdesc_amtorderdiff_timediff_qsix"] = allTsdescAmtorderdiffTimediffQsix
	m.allTsdescAmtorderdiffTimediffQsix = allTsdescAmtorderdiffTimediffQsix

	(*m.ApiParams)["all_tsdesc_amtorderdiff_timediff_qu"] = allTsdescAmtorderdiffTimediffQu
	m.allTsdescAmtorderdiffTimediffQu = allTsdescAmtorderdiffTimediffQu

	(*m.ApiParams)["all_tsdesc_amtorderdiff_vamt_qnine"] = allTsdescAmtorderdiffVamtQnine
	m.allTsdescAmtorderdiffVamtQnine = allTsdescAmtorderdiffVamtQnine

	(*m.ApiParams)["jdauth_fddesc_exist_channelfinance_auth"] = jdauthFddescExistChannelfinanceAuth
	m.jdauthFddescExistChannelfinanceAuth = jdauthFddescExistChannelfinanceAuth

	(*m.ApiParams)["jdauth_fddesc_exist_loginname_equal_phone"] = jdauthFddescExistLoginnameEqualPhone
	m.jdauthFddescExistLoginnameEqualPhone = jdauthFddescExistLoginnameEqualPhone

	(*m.ApiParams)["jdauth_fddesc_minus_nowauthtime_seconds"] = jdauthFddescMinusNowauthtimeSeconds
	m.jdauthFddescMinusNowauthtimeSeconds = jdauthFddescMinusNowauthtimeSeconds

	(*m.ApiParams)["jdbankcard_desc_divide_n_ownername_receiver"] = jdbankcardDescDivideNOwnernameReceiver
	m.jdbankcardDescDivideNOwnernameReceiver = jdbankcardDescDivideNOwnernameReceiver

	(*m.ApiParams)["jdbankcard_desc_n_bankphone_authphone"] = jdbankcardDescNBankphoneAuthphone
	m.jdbankcardDescNBankphoneAuthphone = jdbankcardDescNBankphoneAuthphone

	(*m.ApiParams)["jdbankcard_desc_n_ownername_receiver"] = jdbankcardDescNOwnernameReceiver
	m.jdbankcardDescNOwnernameReceiver = jdbankcardDescNOwnernameReceiver

	(*m.ApiParams)["jdbankcard_diff_divide_nnd_bindphone"] = jdbankcardDiffDivideNndBindphone
	m.jdbankcardDiffDivideNndBindphone = jdbankcardDiffDivideNndBindphone

	(*m.ApiParams)["jdbankcard_fdesc_n_bankname_majorfourbanks"] = jdbankcardFdescNBanknameMajorfourbanks
	m.jdbankcardFdescNBanknameMajorfourbanks = jdbankcardFdescNBanknameMajorfourbanks

	(*m.ApiParams)["jdbankcard_fdesc_n_bankname_others"] = jdbankcardFdescNBanknameOthers
	m.jdbankcardFdescNBanknameOthers = jdbankcardFdescNBanknameOthers

	(*m.ApiParams)["jdbankcard_fdiff_divide_abcall_cntbankname"] = jdbankcardFdiffDivideAbcallCntbankname
	m.jdbankcardFdiffDivideAbcallCntbankname = jdbankcardFdiffDivideAbcallCntbankname

	(*m.ApiParams)["jdbankcard_fdiff_divide_creditall_cntcardtype"] = jdbankcardFdiffDivideCreditallCntcardtype
	m.jdbankcardFdiffDivideCreditallCntcardtype = jdbankcardFdiffDivideCreditallCntcardtype

	(*m.ApiParams)["jdbankcard_fdiff_divide_postall_cntbankname"] = jdbankcardFdiffDividePostallCntbankname
	m.jdbankcardFdiffDividePostallCntbankname = jdbankcardFdiffDividePostallCntbankname

	(*m.ApiParams)["jdbt_gddesc_extract_creditscore_bt"] = jdbtGddescExtractCreditscoreBt
	m.jdbtGddescExtractCreditscoreBt = jdbtGddescExtractCreditscoreBt

	(*m.ApiParams)["jdbt_gddiff_minus_overdraftquota_bt_amt"] = jdbtGddiffMinusOverdraftquotaBtAmt
	m.jdbtGddiffMinusOverdraftquotaBtAmt = jdbtGddiffMinusOverdraftquotaBtAmt

	(*m.ApiParams)["jdoneoneoneonesum_gdesc_amt"] = jdoneoneoneonesumGdescAmt
	m.jdoneoneoneonesumGdescAmt = jdoneoneoneonesumGdescAmt

	(*m.ApiParams)["jdreceivaddr_desc_n_address"] = jdreceivaddrDescNAddress
	m.jdreceivaddrDescNAddress = jdreceivaddrDescNAddress

	(*m.ApiParams)["jdreceivaddr_desc_n_naemail"] = jdreceivaddrDescNNaemail
	m.jdreceivaddrDescNNaemail = jdreceivaddrDescNNaemail

	(*m.ApiParams)["jdreceivaddr_desc_rate_nafixphone"] = jdreceivaddrDescRateNafixphone
	m.jdreceivaddrDescRateNafixphone = jdreceivaddrDescRateNafixphone

	(*m.ApiParams)["jdsixoneeightsum_gdesc_amt"] = jdsixoneeightsumGdescAmt
	m.jdsixoneeightsumGdescAmt = jdsixoneeightsumGdescAmt

	(*m.ApiParams)["jduser_fddesc_exist_webloginname_logname"] = jduserFddescExistWebloginnameLogname
	m.jduserFddescExistWebloginnameLogname = jduserFddescExistWebloginnameLogname

	(*m.ApiParams)["jduser_fddesc_nd_compare_threenames"] = jduserFddescNdCompareThreenames
	m.jduserFddescNdCompareThreenames = jduserFddescNdCompareThreenames

	(*m.ApiParams)["jduser_isbind_bothqqwechat"] = jduserIsbindBothqqwechat
	m.jduserIsbindBothqqwechat = jduserIsbindBothqqwechat

	(*m.ApiParams)["oney_fddif_divide_sevena_range_hourbin_amtaorder"] = oneyFddifDivideSevenaRangeHourbinAmtaorder
	m.oneyFddifDivideSevenaRangeHourbinAmtaorder = oneyFddifDivideSevenaRangeHourbinAmtaorder

	(*m.ApiParams)["oney_fddif_minus_onea_range_hourbin_amtaorder"] = oneyFddifMinusOneaRangeHourbinAmtaorder
	m.oneyFddifMinusOneaRangeHourbinAmtaorder = oneyFddifMinusOneaRangeHourbinAmtaorder

	(*m.ApiParams)["oney_fddif_minus_sixa_range_hourbin_amtaorder"] = oneyFddifMinusSixaRangeHourbinAmtaorder
	m.oneyFddifMinusSixaRangeHourbinAmtaorder = oneyFddifMinusSixaRangeHourbinAmtaorder

	(*m.ApiParams)["oney_fdesc_mean_paycashondelivery_amtorder"] = oneyFdescMeanPaycashondeliveryAmtorder
	m.oneyFdescMeanPaycashondeliveryAmtorder = oneyFdescMeanPaycashondeliveryAmtorder

	(*m.ApiParams)["oney_fdesc_sum_meaninvoicecontent_amtorder"] = oneyFdescSumMeaninvoicecontentAmtorder
	m.oneyFdescSumMeaninvoicecontentAmtorder = oneyFdescSumMeaninvoicecontentAmtorder

	(*m.ApiParams)["oney_gddif_div_onlinepaymentall_sum_pay_amtorder"] = oneyGddifDivOnlinepaymentallSumPayAmtorder
	m.oneyGddifDivOnlinepaymentallSumPayAmtorder = oneyGddifDivOnlinepaymentallSumPayAmtorder

	(*m.ApiParams)["oney_gddif_minus_ca_median_aorder_amtaorder"] = oneyGddifMinusCaMedianAorderAmtaorder
	m.oneyGddifMinusCaMedianAorderAmtaorder = oneyGddifMinusCaMedianAorderAmtaorder

	(*m.ApiParams)["oney_gddif_minus_ca_sd_amtbin_amtaorder"] = oneyGddifMinusCaSdAmtbinAmtaorder
	m.oneyGddifMinusCaSdAmtbinAmtaorder = oneyGddifMinusCaSdAmtbinAmtaorder

	(*m.ApiParams)["oney_gddif_minus_ca_sum_aorder_cntproduct"] = oneyGddifMinusCaSumAorderCntproduct
	m.oneyGddifMinusCaSumAorderCntproduct = oneyGddifMinusCaSumAorderCntproduct

	(*m.ApiParams)["oney_gddif_minus_sa_entropy_amtbin_amtaorder"] = oneyGddifMinusSaEntropyAmtbinAmtaorder
	m.oneyGddifMinusSaEntropyAmtbinAmtaorder = oneyGddifMinusSaEntropyAmtbinAmtaorder

	(*m.ApiParams)["oney_gdesc_cv_recaddrcity_avgamt"] = oneyGdescCvRecaddrcityAvgamt
	m.oneyGdescCvRecaddrcityAvgamt = oneyGdescCvRecaddrcityAvgamt

	(*m.ApiParams)["oney_gdesc_normentropy_amtbin_amtsorder"] = oneyGdescNormentropyAmtbinAmtsorder
	m.oneyGdescNormentropyAmtbinAmtsorder = oneyGdescNormentropyAmtbinAmtsorder

	(*m.ApiParams)["oney_tsdesc_amtorderdiff_timediff_qsix"] = oneyTsdescAmtorderdiffTimediffQsix
	m.oneyTsdescAmtorderdiffTimediffQsix = oneyTsdescAmtorderdiffTimediffQsix

	(*m.ApiParams)["oney_tsdesc_amtorderdiff_vamt_range"] = oneyTsdescAmtorderdiffVamtRange
	m.oneyTsdescAmtorderdiffVamtRange = oneyTsdescAmtorderdiffVamtRange

	(*m.ApiParams)["open_id"] = openId
	m.openId = openId

	(*m.ApiParams)["product_code"] = productCode
	m.productCode = productCode

	(*m.ApiParams)["sixm_fdesc_cv_hour_cntorder"] = sixmFdescCvHourCntorder
	m.sixmFdescCvHourCntorder = sixmFdescCvHourCntorder

	(*m.ApiParams)["sixm_gddif_div_onlinepaymentall_sum_pay_amtorder"] = sixmGddifDivOnlinepaymentallSumPayAmtorder
	m.sixmGddifDivOnlinepaymentallSumPayAmtorder = sixmGddifDivOnlinepaymentallSumPayAmtorder

	(*m.ApiParams)["sixm_gddif_div_phonedigitalall_n_cntprdcategory"] = sixmGddifDivPhonedigitalallNCntprdcategory
	m.sixmGddifDivPhonedigitalallNCntprdcategory = sixmGddifDivPhonedigitalallNCntprdcategory

	(*m.ApiParams)["sixm_gddif_div_sixmall_n_hourtwefourteen_cntorder"] = sixmGddifDivSixmallNHourtwefourteenCntorder
	m.sixmGddifDivSixmallNHourtwefourteenCntorder = sixmGddifDivSixmallNHourtwefourteenCntorder

	(*m.ApiParams)["sixm_gddif_divide_fiveeightall_n_hour_cntorder"] = sixmGddifDivideFiveeightallNHourCntorder
	m.sixmGddifDivideFiveeightallNHourCntorder = sixmGddifDivideFiveeightallNHourCntorder

	(*m.ApiParams)["sixm_gddif_minus_ca_sum_aorder_cntproduct"] = sixmGddifMinusCaSumAorderCntproduct
	m.sixmGddifMinusCaSumAorderCntproduct = sixmGddifMinusCaSumAorderCntproduct

	(*m.ApiParams)["sixm_gdesc_min_recaddrcity_avgamt"] = sixmGdescMinRecaddrcityAvgamt
	m.sixmGdescMinRecaddrcityAvgamt = sixmGdescMinRecaddrcityAvgamt

	(*m.ApiParams)["sixm_gdesc_range_recaddrprovince_avgamt"] = sixmGdescRangeRecaddrprovinceAvgamt
	m.sixmGdescRangeRecaddrprovinceAvgamt = sixmGdescRangeRecaddrprovinceAvgamt

	(*m.ApiParams)["springfestival_gdesc_qu_aamt"] = springfestivalGdescQuAamt
	m.springfestivalGdescQuAamt = springfestivalGdescQuAamt

	(*m.ApiParams)["threem_fddif_minus_sevena_qd_hourbin_amtaorder"] = threemFddifMinusSevenaQdHourbinAmtaorder
	m.threemFddifMinusSevenaQdHourbinAmtaorder = threemFddifMinusSevenaQdHourbinAmtaorder

	(*m.ApiParams)["threem_gddif_div_travelrechargeall_n_cntprdcateg"] = threemGddifDivTravelrechargeallNCntprdcateg
	m.threemGddifDivTravelrechargeallNCntprdcateg = threemGddifDivTravelrechargeallNCntprdcateg

	(*m.ApiParams)["threem_gddif_divide_failall_n_sts_cntorder"] = threemGddifDivideFailallNStsCntorder
	m.threemGddifDivideFailallNStsCntorder = threemGddifDivideFailallNStsCntorder

	(*m.ApiParams)["threem_gddif_divide_nullall_sum_pay_amtorder"] = threemGddifDivideNullallSumPayAmtorder
	m.threemGddifDivideNullallSumPayAmtorder = threemGddifDivideNullallSumPayAmtorder

	(*m.ApiParams)["threem_gdesc_sum_sorder_amtaorder"] = threemGdescSumSorderAmtaorder
	m.threemGdescSumSorderAmtaorder = threemGdescSumSorderAmtaorder

	(*m.ApiParams)["transaction_id"] = transactionId
	m.transactionId = transactionId
}

func (*ZhimaCreditMsxfOnlinejdscoreQueryRequest) GetApiMethodName() string {
	return "zhima.credit.msxf.onlinejdscore.query"
}

func (m *ZhimaCreditMsxfOnlinejdscoreQueryRequest) GetApiParams() *map[string]string {
	return m.ApiParams
}

func (m *ZhimaCreditMsxfOnlinejdscoreQueryRequest) GetFileParams() *map[string]string {
	return m.FileParams
}

func (m *ZhimaCreditMsxfOnlinejdscoreQueryRequest) SetApiVersion(ApiVersion string) {
	m.ApiVersion = ApiVersion
}

func (m *ZhimaCreditMsxfOnlinejdscoreQueryRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaCreditMsxfOnlinejdscoreQueryRequest) SetScene(Scene string) {
	m.Scene = Scene
}

func (m *ZhimaCreditMsxfOnlinejdscoreQueryRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaCreditMsxfOnlinejdscoreQueryRequest) SetChannel(Channel string) {
	m.Channel = Channel
}

func (m *ZhimaCreditMsxfOnlinejdscoreQueryRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaCreditMsxfOnlinejdscoreQueryRequest) SetPlatform(Platform string) {
	m.Platform = Platform
}

func (m *ZhimaCreditMsxfOnlinejdscoreQueryRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaCreditMsxfOnlinejdscoreQueryRequest) SetExtParams(ExtParams string) {
	m.ExtParams = ExtParams
}

func (m *ZhimaCreditMsxfOnlinejdscoreQueryRequest) GetExtParams() string {
	return m.ExtParams
}
