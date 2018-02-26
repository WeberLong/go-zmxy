package requests

import (
	"github.com/WeberLong/go-zmxy/zmop/interfaces"
)

type ZhimaCreditVehicleVerifyRequest struct {
	engineNo      string // 发动机号码。vin与engine_no至少包含一项
	owner         string // 所有人，支持个人和机构
	plateNo       string // 车牌号
	productCode   string // 产品码
	registerDate  string // 注册日期，格式yyyyMMdd
	transactionId string // 商户传入的业务流水号。此字段由商户生成，需确保唯一性，用于定位每一次请求，后续按此流水进行对帐。生成规则: 固定30位数字串，前17位为精确到毫秒的时间yyyyMMddHHmmssSSS，后13位为自增数字。
	vehicleBrand  string // 车辆品牌（行驶证中中文部分）
	vehicleSeries string // 车辆型号（行驶证中英文部分）
	vin           string // 车辆识别代号。vin与engine_no至少包含一项

	interfaces.ZhimaRequestParams
}

func (m *ZhimaCreditVehicleVerifyRequest) InitBizParams(engineNo, owner, plateNo, productCode, registerDate, transactionId, vehicleBrand, vehicleSeries, vin string) {
	m.ApiParams = &map[string]string{}
	m.FileParams = &map[string]string{}

	(*m.ApiParams)["engine_no"] = engineNo
	m.engineNo = engineNo

	(*m.ApiParams)["owner"] = owner
	m.owner = owner

	(*m.ApiParams)["plate_no"] = plateNo
	m.plateNo = plateNo

	(*m.ApiParams)["product_code"] = productCode
	m.productCode = productCode

	(*m.ApiParams)["register_date"] = registerDate
	m.registerDate = registerDate

	(*m.ApiParams)["transaction_id"] = transactionId
	m.transactionId = transactionId

	(*m.ApiParams)["vehicle_brand"] = vehicleBrand
	m.vehicleBrand = vehicleBrand

	(*m.ApiParams)["vehicle_series"] = vehicleSeries
	m.vehicleSeries = vehicleSeries

	(*m.ApiParams)["vin"] = vin
	m.vin = vin
}

func (*ZhimaCreditVehicleVerifyRequest) GetApiMethodName() string {
	return "zhima.credit.vehicle.verify"
}

func (m *ZhimaCreditVehicleVerifyRequest) GetApiParams() *map[string]string {
	return m.ApiParams
}

func (m *ZhimaCreditVehicleVerifyRequest) GetFileParams() *map[string]string {
	return m.FileParams
}

func (m *ZhimaCreditVehicleVerifyRequest) SetApiVersion(ApiVersion string) {
	m.ApiVersion = ApiVersion
}

func (m *ZhimaCreditVehicleVerifyRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaCreditVehicleVerifyRequest) SetScene(Scene string) {
	m.Scene = Scene
}

func (m *ZhimaCreditVehicleVerifyRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaCreditVehicleVerifyRequest) SetChannel(Channel string) {
	m.Channel = Channel
}

func (m *ZhimaCreditVehicleVerifyRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaCreditVehicleVerifyRequest) SetPlatform(Platform string) {
	m.Platform = Platform
}

func (m *ZhimaCreditVehicleVerifyRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaCreditVehicleVerifyRequest) SetExtParams(ExtParams string) {
	m.ExtParams = ExtParams
}

func (m *ZhimaCreditVehicleVerifyRequest) GetExtParams() string {
	return m.ExtParams
}
