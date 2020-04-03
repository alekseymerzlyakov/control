package main

import (
	"fmt"
	"github.com/tidwall/gjson"
	"regexp"
	"strconv"
	"strings"
)

func otp(otp string) {

	for oIdx := 0; oIdx < 10; oIdx++ {
		path := "ops.#.list." + strconv.Itoa(oIdx) + ".data.phone"
		text := "ops.#.list." + strconv.Itoa(oIdx) + ".data.text"
		fmt.Println("OTP -->>>  ", path)
		phone := gjson.Get(otp,path).String()
		sms :=   gjson.Get(otp,text).String()

			if strings.Contains(phone, GetPropValue("phone")) {
				fmt.Println("phone -->>>  ", phone)
				fmt.Println("text -->>>  ", sms)
				re := regexp.MustCompile("[0-9]+")
				s := re.FindAllString(sms, -1)
				element := strings.Trim(s[0], "[")
				element = strings.Trim(element, "]")
				SetPropValue("otp",element)
				fmt.Println("text -->>>  ", element)
				break

			}



	}



}
