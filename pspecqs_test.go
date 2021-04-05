package pspecqs

import (
	"fmt"
	"test.com/a/grequests"
	"testing"
)

func TestCasServer_login(t *testing.T) {
	err:=func() error{
		cli:=grequests.NewSession(&grequests.RequestOptions{
			Headers: map[string]string{
				"Cookie":"JSESSIONID=0001xyMQTc39h6d3U1n8CCRq4Zw:-G00EC;",
			},
		})
		_,_,err:=CasServer_captcha(cli)
		if err != nil {
			return err
		}
		fmt.Println("input ca:",)
		var ca string
		fmt.Scanf("%s",&ca)
		cks,err:=CasServer_login("execution","jinzhaihongzhi","Hz01200017=",ca,cli)
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

