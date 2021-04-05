package main

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	jsoniter "github.com/json-iterator/go"
	"net/url"
	"strings"
	"test.com/a/grequests"
)

func CasServer_captcha(session *grequests.Session)error{
	cli:=grequests.NewCli(session)
	r,err:=cli.Get("http://psp.e-cqs.cn/casServer/captcha.jpg?0.5466787858743882",nil)
	if err != nil {
		return err
	}
	r.DownloadToFile("test.png")
	return nil
}
func CasServer_login(username,password,authcode string,session *grequests.Session)(string,error){
	cli:=grequests.NewCli(session)
	r,err:=cli.Get("http://psp.e-cqs.cn/casServer/login?service=http://psp.e-cqs.cn/InspSystemV11/reservation/reservationQuery.jsp&code=90000&tacsaccess=1&flag=false",&grequests.RequestOptions{})
	if err != nil {
		return "", err
	}
	gq,err:=goquery.NewDocumentFromReader(bytes.NewReader(r.Bytes()))
	if err != nil {
		return "",err
	}
	execution,_:=gq.Find(".loginBtn input[name=execution]").Attr("value")
	reqdata:=map[string]string{
		"usertype":"1",
		"strongPassword":"1",
		"username":EncodeUserName(username),
		"password":EncodePass(password,authcode),
		"authcode":authcode,
		"execution":execution,
		"_eventId":"submit",
		"submit":"",
	}
	r,err=cli.Post("http://psp.e-cqs.cn/casServer/login?service=http://psp.e-cqs.cn/InspSystemV11/reservation/reservationQuery.jsp&code=90000&tacsaccess=1&flag=false",&grequests.RequestOptions{
		Data: reqdata,
	})
	if err != nil {
		return "", err
	}
	gq,err=goquery.NewDocumentFromReader(bytes.NewReader(r.Bytes()))
	if err != nil {
		return "",err
	}
	if gq.Find("html head title").Text()!="检定预约受理"{
		errorsmsg:=gq.Find(".errors#msg").Text()
		if errorsmsg!=""{
			return "",errors.New(errorsmsg)
		}
		return "",errors.New("登录失败")
	}
	ul,_:=url.Parse("http://psp.e-cqs.cn/casServer/")
	cks:=cli.HTTPClient.Jar.Cookies(ul)
	cookies:=strings.Builder{}
	for idx, it := range cks  {
		if idx!=0{
			cookies.WriteString(";")
		}
		cookies.WriteString(fmt.Sprintf("%s=%s",it.Name,it.Value))
	}
	return cookies.String(),nil
}
func GetIntruInfoByReservationId(params string,filters *FilterSets,session *grequests.Session)([]*InspAppointmentIntruInfoUser,error){
	cli:=grequests.NewCli(session)
	ft:=NewFilterSets()
	if filters!=nil{
		ft=filters
	}
	reqquery:=map[string]string{
		"encoding":"true",
		"jsonValue":jsoniter.Wrap(map[string]interface{}{
			"serviceClassName":"com.itown.inspSystem.reservation.service.ReservationService",
			"methodName":"getIntruInfoByReservationId",
			"serviceObject":nil,
			"type":nil,
			"params":[]interface{}{params},
		}).ToString(),
	}
	reqdata:=map[string]string{
		"page":"1",
		"rows":"100",
		"__filterSet":jsoniter.Wrap(ft.Maps()).ToString(),
	}
	r,err:=cli.Post("http://psp.e-cqs.cn/InspSystemV11/jsonClient.action",&grequests.RequestOptions{
		Params: reqquery,
		Data: reqdata,
		Headers: map[string]string{
			"Content-Type":"application/x-www-form-urlencoded; charset=UTF-8",
			"X-Requested-With":"XMLHttpRequest",
		},
		RedirectLimit: -1,
	})
	if err != nil {
		return nil, err
	}
	err=r.HttpStatusOK()
	if err != nil {
		if r.StatusCode==302{
			return nil,errors.New("登录状态错误")
		}
		return nil,err
	}
	rjs,err:=r.ToJson()
	if err != nil {
		return nil,err
	}
	classname:=rjs.Get("__className").ToString()
	if classname=="com.itown.rcp.proxy.impl.http.json.ExceptionWrapper"{
		return nil,errors.New(rjs.Get("message").ToString())
	}
	totalcount:=rjs.Get("transferableProperties","fspParameter","pagination","totalCount").ToInt()
	_=totalcount
	users:=make([]*InspAppointmentIntruInfoUser,0)
	vals:=rjs.Get("returnValue","value")
	for i := 0; i < vals.Size(); i++ {
		it:=vals.Get(i)
		val:=InspAppointmentIntruInfoUser{
			ClassName: it.Get("__className").ToString(),
			StrongInspUuid: it.Get("strongInspUuid").ToString(),
			VerificationState: it.Get("verificationState").ToString(),
			MeterStatus: it.Get("meterStatus").ToString(),
			Uuid: it.Get("uuid").ToString(),
			SampleNo: it.Get("sampleNo").ToString(),
			MeterIntruName: it.Get("meterIntruName").ToString(),
			ModelSpec: it.Get("modelSpec").ToString(),
			MeasureRang: it.Get("measureRang").ToString(),
			AccuracyLevel: it.Get("accuracyLevel").ToString(),
			SerialNum: it.Get("serialNum").ToString(),
			SetupPlace: it.Get("setupPlace").ToString(),
			PorgaoName: it.Get("porgaoName").ToString(),
			ExpiryDate: it.Get("expiryDate","value").ToInt64(),
			ApplyOpinion: it.Get("applyOpinion").ToString(),
			InstrPurposeCode: it.Get("instrPurposeCode").ToString(),
			InstrPurposeName: it.Get("instrPurposeName").ToString(),
			MeterType: it.Get("meterType").ToString(),
			MeterCategoryCode: it.Get("meterCategoryCode").ToString(),
			MeterCategoryName: it.Get("meterCategoryName").ToString(),
			MeterClassCode: it.Get("meterClassCode").ToString(),
			MeterClassName: it.Get("meterClassName").ToString(),
			MeterNumber: it.Get("meterNumber","value").ToInt(),
			ProdAddrTypeClass: it.Get("prodAddrTypeClass").ToString(),
			ProdAddrTypeName: it.Get("prodAddrTypeName").ToString(),
			UseAddrCode: it.Get("useAddrCode").ToString(),
			UseAddrName: it.Get("useAddrName").ToString(),
			ApplyDate: it.Get("applyDate","value").ToInt64(),
			ForceInsp: it.Get("forceInsp").ToString(),
			ApplyCompanyCode: it.Get("applyCompanyCode").ToString(),
			ApplyCompanyName: it.Get("applyCompanyName").ToString(),
			LicenceNo: it.Get("licenceNo").ToString(),
			SendStatus: it.Get("sendStatus").ToString(),
			ReservationState: it.Get("reservationState").ToString(),
			InspSendType: it.Get("inspSendType").ToString(),
			AppointmentId: it.Get("appointmentId").ToString(),
			ReservationId: it.Get("reservationId").ToString(),
		}
		users=append(users,&val)
	}
	return users,nil
}
func GetAppointmentInfoCount(filters *FilterSets,session *grequests.Session)(int,error){
	cli:=grequests.NewCli(session)
	ft:=NewFilterSets()
	if filters!=nil{
		ft=filters
	}
	reqquery:=map[string]string{
		"encoding":"true",
		"jsonValue":jsoniter.Wrap(map[string]interface{}{
			"serviceClassName":"com.itown.inspSystem.reservation.service.ReservationService",
			"methodName":"getAppointmentInfo",
			"serviceObject":nil,
			"type":nil,
			"params":[]interface{}{},
		}).ToString(),
	}
	reqdata:=map[string]string{
		"page":"1",
		"rows":"1",
		"__filterSet":jsoniter.Wrap(ft.Maps()).ToString(),
	}
	r,err:=cli.Post("http://psp.e-cqs.cn/InspSystemV11/jsonClient.action",&grequests.RequestOptions{
		Params: reqquery,
		Data: reqdata,
		Headers: map[string]string{
			"Content-Type":"application/x-www-form-urlencoded; charset=UTF-8",
			"X-Requested-With":"XMLHttpRequest",
		},
		RedirectLimit: -1,
	})
	if err != nil {
		if r.StatusCode==302{
			return 0,errors.New("登录状态错误")
		}
		return 0,err
	}
	rjs,err:=r.ToJson()
	if err != nil {
		return 0,err
	}
	classname:=rjs.Get("__className").ToString()
	if classname=="com.itown.rcp.proxy.impl.http.json.ExceptionWrapper"{
		return 0,errors.New(rjs.Get("message").ToString())
	}
	totalcount:=rjs.Get("transferableProperties","fspParameter","pagination","totalCount").ToInt()
	return totalcount,nil
}
func GetAppointmentInfo(rows int,filters *FilterSets,session *grequests.Session)([]*InspAppointmentInfoVoUser,error){
	cli:=grequests.NewCli(session)
	ft:=NewFilterSets()
	if filters!=nil{
		ft=filters
	}
	reqquery:=map[string]string{
		"encoding":"true",
		"jsonValue":jsoniter.Wrap(map[string]interface{}{
			"serviceClassName":"com.itown.inspSystem.reservation.service.ReservationService",
			"methodName":"getAppointmentInfo",
			"serviceObject":nil,
			"type":nil,
			"params":[]interface{}{},
		}).ToString(),
	}
	reqdata:=map[string]string{
		"page":"1",
		"rows":fmt.Sprintf("%d",rows),
		"__filterSet":jsoniter.Wrap(ft.Maps()).ToString(),
	}
	r,err:=cli.Post("http://psp.e-cqs.cn/InspSystemV11/jsonClient.action",&grequests.RequestOptions{
		Params: reqquery,
		Data: reqdata,
		Headers: map[string]string{
			"Content-Type":"application/x-www-form-urlencoded; charset=UTF-8",
			"X-Requested-With":"XMLHttpRequest",
		},
		RedirectLimit: -1,
	})
	if err != nil {
		if r.StatusCode==302{
			return nil,errors.New("登录状态错误")
		}
		return nil,err
	}
	rjs,err:=r.ToJson()
	if err != nil {
		return nil,err
	}
	classname:=rjs.Get("__className").ToString()
	if classname=="com.itown.rcp.proxy.impl.http.json.ExceptionWrapper"{
		return nil,errors.New(rjs.Get("message").ToString())
	}
	totalcount:=rjs.Get("transferableProperties","fspParameter","pagination","totalCount").ToInt()
	_=totalcount
	users:=make([]*InspAppointmentInfoVoUser,0)
	vals:=rjs.Get("returnValue","value")
	for i := 0; i < vals.Size(); i++ {
		it:=vals.Get(i)
		val:=InspAppointmentInfoVoUser{
			ClassName: it.Get("__className").ToString(),
			ProcessingState: it.Get("processingState").ToString(),
			Uuid: it.Get("uuid").ToString(),
			AppointmentId: it.Get("appointmentId").ToString(),
			ApplicationDate: it.Get("applicationDate").Get("value").ToString(),
			ReservationState: it.Get("reservationState").ToString(),
			ReservationsTel: it.Get("reservationsTel").ToString(),
			UseCompanyCode: it.Get("useCompanyCode").ToString(),
			UseCompanyName: it.Get("useCompanyName").ToString(),
			MeterTotal: it.Get("meterTotal","value").ToString(),
			Reservations: it.Get("reservations").ToString(),
			BatchNumber: it.Get("batchNumber","value").ToString(),
			CompanyAddr: it.Get("companyAddr").ToString(),
			OperactionTime: it.Get("operationTime","value").ToString(),
			SendStatus: it.Get("sendStatus").ToString(),
		}
		users=append(users,&val)
	}
	return users,nil
}

func main() {
	err:=func() error{
		cli:=grequests.NewSession(&grequests.RequestOptions{
			Headers: map[string]string{
				"Cookie":"JSESSIONID=0001xyMQTc39h6d3U1n8CCRq4Zw:-G00EC;",
			},
		})
		err:=CasServer_captcha(cli)
		if err != nil {
			return err
		}
		fmt.Println("input ca:",)
		var ca string
		fmt.Scanf("%s",&ca)
		cks,err:=CasServer_login("jinzhaihongzhi","Hz01200017=",ca,cli)
		if err != nil {
			return err
		}
		_=cks
		users,err:=GetIntruInfoByReservationId("7BB017210AE6491417B2861C81A25459",nil,cli)
		if err != nil {
			return err
		}
		for _, it := range users {
			fmt.Println(it)
		}
	    return nil
	}()
	if err!=nil{
	    fmt.Println(err)
	}
}
