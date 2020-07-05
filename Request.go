package main

import (
	"fmt"
	"gopkg.in/h2non/gentleman.v2"
	"gopkg.in/h2non/gentleman.v2/plugins/body"
)

type Response struct {
	ResponseCode int
	ResponseBody string
	URL string
	Header string
}

func Request(requestBody string, header map[int][]string, RequestParameters map[int][]string) Response {
	// Create a new client
	var respons = Response{}
	cli := gentleman.New()

	var UR string
	if RequestParameters[3][0] != "nil" {
		UR = RequestParameters[1][0] + "://" + RequestParameters[2][0] +  RequestParameters[3][0]
	} else {
		UR = RequestParameters[2][0]
	}


	// Define base URL
	fmt.Println("UR    -->     ", UR)
	cli.URL(UR)
	cli.Method(RequestParameters[0][0])
	cli.Use(body.JSON(requestBody))
	req := cli.Request()

	//ur := RequestParameters[3][0]
	fmt.Println("RequestParameters[3][0]: %s\n         ", RequestParameters[3][0])

	// Выполнение теста - перебираем список Requests из вкладки APIlist
	for rIdx := 0; rIdx < len(header); rIdx++ {
		req.SetHeader(header[rIdx][0], header[rIdx][1])
		fmt.Println("Header: ->        ", header[rIdx][0],header[rIdx][1])
	}

	// Perform the request
	res, err := req.Send()
	if err != nil {
		fmt.Printf("Request error: %s\n", err)
		//fmt.Println("Ошибка. В вкладке APIList есть ссылка на не существующий тест")
		//os.Exit(0)
		//return
	}
	//if !res.Ok {
	//	fmt.Printf("Invalid server response: %d\n", res.StatusCode)
	//	//return
	//}

	respons.ResponseBody = res.String()
	respons.ResponseCode = res.StatusCode
	respons.URL = res.RawRequest.URL.String()
	//respons.Header = res.Header.Get("authorization")
	// Reads the whole body and returns it as string
	//fmt.Printf("Body: %s", res.String())
	return respons
}