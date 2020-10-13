package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"gopkg.in/h2non/gentleman.v2"
	"gopkg.in/h2non/gentleman.v2/plugins/body"
	"os"
)

type Response2 struct {
	ResponseCode2 int
	ResponseBody2 string
	URL2          string
	Header        string
}

func (b Data) Request2(requestBody2 string, header2 map[int][]string, RequestParameters2 map[int][]string) Response2 {
	// Create a new client
	var respons2 = Response2{}
	cli2 := gentleman.New()

	var UR2 string
	if RequestParameters2[3][0] != "nil" {
		UR2 = RequestParameters2[1][0] + "://" + RequestParameters2[2][0] + RequestParameters2[3][0]
	} else {
		UR2 = RequestParameters2[2][0]
	}
	fmt.Println("QQQQQQ   RequestParameters2[3][0]:         ", UR2)
	fmt.Println("QQQQQQ   METHOD:         ", RequestParameters2[0][0])
	// Define base URL

	cli2.URL(UR2)
	cli2.Method(RequestParameters2[0][0])
	cli2.Use(body.JSON(requestBody2))
	fmt.Println("QQQQQQ   requestBody2:         ", requestBody2)

	req2 := cli2.Request()
	fmt.Println("QQQQQQ   req2:         ", req2.BodyString)
	//ur := RequestParameters[3][0]

	// Выполнение теста - перебираем список Requests из вкладки APIlist
	for rIdx := 0; rIdx < len(header2); rIdx++ {
		req2.SetHeader(header2[rIdx][0], header2[rIdx][1])
		fmt.Println("QQQQQQ   SetHeader:         ", header2[rIdx][0], header2[rIdx][1])
	}

	// Perform the request
	res2, err := req2.Send()

	if err != nil {
		fmt.Printf("Request error: %s\n", err)
		msg := "Request error: ->    " + err.Error()
		msg = msg + "\nВыполнение теста остановленно"
		msg = msg + "\n ApiSheetName2 ->" + RendomData.ApiSheetName
		msg = msg + "\n ApiTestName2 ->" + RendomData.ApiTestName
		msg = msg + "\nURL -> " + UR2
		msg = msg + "\nRequest Body -> " + "\n" + requestBody2
		telegram(msg)
		os.Exit(0)

	}
	//if !res.Ok {
	//	fmt.Printf("Invalid server response: %d\n", res.StatusCode)
	//	//return
	//}

	//respons2.ResponseBody2 = res2.String()
	//respons2.ResponseCode2 = res2.StatusCode
	//respons2.URL2 = res2.RawRequest.URL.String()
	//respons2.Header = res2.Header.Get("authorization")
	//// Reads the whole body and returns it as string
	//
	//
	//respons2.URL2 = res2.RawRequest.URL.String()
	//fmt.Println("\n---------------------------Request URL-----------------------------")
	//log.Info("\n---------------------------Request URL-----------------------------")
	//fmt.Println("Request2 URL         -->      ", respons2.URL2)
	//log.Info("Request2 URL         -->      ", respons2.URL2)
	//
	//fmt.Println("\n---------------------------Request Body-----------------------------")
	//log.Info("\n---------------------------Request Body-----------------------------")
	//fmt.Println("Request2 Body         -->      ", requestBody2)
	//log.Info("Request Body         -->      ", requestBody2)
	//
	////fmt.Println("len(cook):", len(cook))
	//fmt.Println("\n\n---------------------------Response Cookies-----------------------------")
	//if len(res2.Cookies) > 0 {
	//	cookieFull := ""
	//	for _, cookie := range res2.Cookies {
	//		cookieFull = cookieFull + cookie.Name + "=" + cookie.Value + "; "
	//
	//		fmt.Println("cookie:", cookie)
	//		fmt.Println("Cookie.name:", cookie.Name)
	//		fmt.Println("Cookie.Value", cookie.Value)
	//		SetPropValue("testdata.cookie."+cookie.Name, cookie.Name+"="+cookie.Value)
	//	}
	//
	//	SetPropValue("testdata.cookiefull", cookieFull)
	//} else {
	//	fmt.Println("count cookie = 0")
	//
	//}
	//fmt.Println("\n---------------------------Response Headers-----------------------------")
	//fmt.Println("Response Headers:", res2.Header)
	////for _, header := range res.Header {
	////	fmt.Println("Response Headers:", res.Header)
	////}
	//respons2.ResponseBody2 = ""
	//respons2.ResponseBody2 = res2.String()
	//fmt.Println("\n---------------------------Response Body-----------------------------")
	//log.Info("\n---------------------------Response Body-----------------------------")
	//
	//fmt.Println("Response Body       -->     ", respons2.ResponseBody2)
	//log.Info("Response Body       -->     ", respons2.ResponseBody2)
	//
	//respons2.ResponseCode2 = res2.StatusCode
	//
	//
	//
	//
	////fmt.Printf("Body: %s", res.String())
	//return respons2
	respons2.URL2 = res2.RawRequest.URL.String()
	fmt.Println("\n---------------------------Request2 URL-----------------------------")
	log.Info("\n---------------------------Request2 URL-----------------------------")
	fmt.Println("Request URL         -->      ", respons2.URL2)
	log.Info("Request URL         -->      ", respons2.URL2)

	fmt.Println("\n---------------------------Request Body-----------------------------")
	log.Info("\n---------------------------Request Body-----------------------------")
	fmt.Println("Request Body         -->      ", requestBody2)
	log.Info("Request Body         -->      ", requestBody2)

	//fmt.Println("len(cook):", len(cook))
	fmt.Println("\n\n---------------------------Response Cookies-----------------------------")
	if len(res2.Cookies) > 0 {
		cookieFull := ""
		for _, cookie := range res2.Cookies {
			cookieFull = cookieFull + cookie.Name + "=" + cookie.Value + "; "

			fmt.Println("cookie:", cookie)
			fmt.Println("Cookie.name:", cookie.Name)
			fmt.Println("Cookie.Value", cookie.Value)
			SetPropValue("testdata.cookie."+cookie.Name, cookie.Name+"="+cookie.Value)
		}

		SetPropValue("testdata.cookiefull", cookieFull)
	} else {
		fmt.Println("count cookie = 0")

	}
	fmt.Println("\n---------------------------Response Headers-----------------------------")
	fmt.Println("Response Headers:", res2.Header)
	//for _, header := range res.Header {
	//	fmt.Println("Response Headers:", res.Header)
	//}
	respons2.ResponseBody2 = ""
	respons2.ResponseBody2 = res2.String()
	fmt.Println("\n---------------------------Response Body-----------------------------")
	log.Info("\n---------------------------Response Body-----------------------------")

	fmt.Println("Response Body       -->     ", respons2.ResponseBody2)
	log.Info("Response Body       -->     ", respons2.ResponseBody2)

	respons2.ResponseCode2 = res2.StatusCode

	//respons.Header = res.Header.Get("authorization")
	// Reads the whole body and returns it as string
	//fmt.Printf("Body: %s", res.String())
	return respons2
}
