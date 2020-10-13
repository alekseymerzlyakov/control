package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"gopkg.in/h2non/gentleman.v2"
	"gopkg.in/h2non/gentleman.v2/plugins/body"
	"os"
)

type Response struct {
	ResponseCode int
	ResponseBody string
	URL          string
	Header       string
}

func (a Data) Request(requestBody string, header map[int][]string, RequestParameters map[int][]string) Response {
	// Create a new client
	var respons = Response{}
	cli := gentleman.New()

	var UR string
	if RequestParameters[3][0] != "" {
		UR = RequestParameters[1][0] + "://" + RequestParameters[2][0] + RequestParameters[3][0]
	} else {
		UR = RequestParameters[2][0]
	}

	UR = a.Replace(UR)

	cli.URL(UR)
	cli.Method(RequestParameters[0][0])
	cli.Use(body.JSON(requestBody))

	req := cli.Request()

	// Выполнение теста - перебираем список Requests из вкладки APIlist
	fmt.Println("\n\n---------------------------Request Header -----------------------------\n")
	for rIdx := 0; rIdx < len(header); rIdx++ {
		req.SetHeader(header[rIdx][0], header[rIdx][1])
		fmt.Println("Header:    -->        ", header[rIdx][0], header[rIdx][1])
	}

	// Perform the request
	res, err := req.Send()

	if err != nil {
		fmt.Printf("Request error: %s\n", err)
		msg := "Request error: ->    " + err.Error()
		msg = msg + "\nВыполнение теста остановленно"
		msg = msg + "\n ApiSheetName ->" + RendomData.ApiSheetName
		msg = msg + "\n ApiTestName ->" + RendomData.ApiTestName
		msg = msg + "\nURL -> " + UR
		msg = msg + "\nRequest Body -> " + "\n" + requestBody
		telegram(msg)
		os.Exit(0)

	}

	respons.URL = res.RawRequest.URL.String()
	fmt.Println("\n---------------------------Request URL-----------------------------")
	log.Info("\n---------------------------Request URL-----------------------------")
	fmt.Println("Request URL         -->      ", respons.URL)
	log.Info("Request URL         -->      ", respons.URL)

	fmt.Println("\n---------------------------Request Body-----------------------------")
	log.Info("\n---------------------------Request Body-----------------------------")
	fmt.Println("Request Body         -->      ", requestBody)
	log.Info("Request Body         -->      ", requestBody)

	//fmt.Println("len(cook):", len(cook))
	fmt.Println("\n\n---------------------------Response Cookies-----------------------------")
	if len(res.Cookies) > 0 {
		cookieFull := ""
		for _, cookie := range res.Cookies {
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
	fmt.Println("Response Headers:", res.Header)
	//for _, header := range res.Header {
	//	fmt.Println("Response Headers:", res.Header)
	//}
	respons.ResponseBody = ""
	respons.ResponseBody = res.String()
	fmt.Println("\n---------------------------Response Body-----------------------------")
	log.Info("\n---------------------------Response Body-----------------------------")

	fmt.Println("Response Body       -->     ", respons.ResponseBody)
	log.Info("Response Body       -->     ", respons.ResponseBody)

	respons.ResponseCode = res.StatusCode

	//respons.Header = res.Header.Get("authorization")
	// Reads the whole body and returns it as string
	//fmt.Printf("Body: %s", res.String())
	return respons
}
