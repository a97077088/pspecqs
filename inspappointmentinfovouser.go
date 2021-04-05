package pspecqs

import "test.com/a/app"
const(
	//已处理
	ProcessingState_YES=iota
	//未处理
	ProcessingState_NO
)
type InspAppointmentInfoVoUser struct {
	app.AppUser
	ClassName string
	//状态
	ProcessingState string
	Uuid string
	//预约编号
	AppointmentId string
	//预约时间?
	ApplicationDate string
	//不知道是什么
	ReservationState string
	//联系电话
	ReservationsTel	string
	//使用单位代码
	UseCompanyCode	string
	//使用单位名称
	UseCompanyName string
	//item数量
	MeterTotal string
	//预约人
	Reservations string
	//批次
	BatchNumber string
	//预约单位地址
	CompanyAddr	string
	//操作时间?
	OperactionTime string
	//发送状态，不知道是什么
	SendStatus string
}


func NewInspAppointmentInfoVoUser()*InspAppointmentInfoVoUser{
	u:=&InspAppointmentInfoVoUser{}
	return u
}
func InspAppointmentInfoVoUserFromInterface(unique interface{})*InspAppointmentInfoVoUser {
	return unique.(*InspAppointmentInfoVoUser)
}