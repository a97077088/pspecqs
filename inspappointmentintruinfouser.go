package pspecqs

import "test.com/a/app"

type InspAppointmentIntruInfoUser struct {
	app.AppUser
	ClassName string
	StrongInspUuid string
	VerificationState string
	MeterStatus string
	Uuid string
	SampleNo string
	//器具名称
	MeterIntruName string
	//规格型号
	ModelSpec string
	//测量范围
	MeasureRang string
	//准确度等级
	AccuracyLevel string
	//出厂编号
	SerialNum string
	//安装/使用地点
	SetupPlace string
	//生产厂家
	PorgaoName string
	//有效期至
	ExpiryDate int64
	//检验单位
	ApplyOpinion string
	//器具状态
	InstrPurposeCode string
	//计量器具用途
	InstrPurposeName string
	MeterType string
	MeterCategoryCode string
	//器具一级目录
	MeterCategoryName string
	MeterClassCode string
	//器具二级目录
	MeterClassName string
	MeterNumber	int
	ProdAddrTypeClass string
	ProdAddrTypeName string
	UseAddrCode	string
	UseAddrName	string
	ApplyDate int64
	ForceInsp string
	//检验公司代码
	ApplyCompanyCode string
	//检验公司名称
	ApplyCompanyName string
	//许可证
	LicenceNo string
	SendStatus string
	ReservationState string
	InspSendType string
	//预约编号
	AppointmentId string
	ReservationId string
}
func (U *InspAppointmentIntruInfoUser) Unique() interface{} {
	return U.AppUser.Unique()
}

func NewInspAppointmentIntruInfoUser()*InspAppointmentIntruInfoUser{
	u:=&InspAppointmentIntruInfoUser{}
	return u
}
func InspAppointmentIntruInfoUserFromInterface(unique interface{})*InspAppointmentIntruInfoUser {
	return unique.(*InspAppointmentIntruInfoUser)
}