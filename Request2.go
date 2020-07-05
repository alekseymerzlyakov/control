package main

import (
	"fmt"
	"gopkg.in/h2non/gentleman.v2"
	"gopkg.in/h2non/gentleman.v2/plugins/body"
)

type Response2 struct {
	ResponseCode2 int
	ResponseBody2 string
	URL2 string
	Header string
}

func Request2(requestBody2 string, header2 map[int][]string, RequestParameters2 map[int][]string) Response2 {
	// Create a new client
	var respons2 = Response2{}
	cli2 := gentleman.New()

	var UR2 string
	if RequestParameters2[3][0] != "nil" {
		UR2 = RequestParameters2[1][0] + "://" + RequestParameters2[2][0] +  RequestParameters2[3][0]
	} else {
		UR2 = RequestParameters2[2][0]
	}


	// Define base URL
	fmt.Println("UR    -->     ", UR2)
	cli2.URL(UR2)
	cli2.Method(RequestParameters2[0][0])
	cli2.Use(body.JSON(requestBody2))
	req2 := cli2.Request()

	//ur := RequestParameters[3][0]
	fmt.Println("RequestParameters2[3][0]: %s\n         ", RequestParameters2[3][0])

	// Выполнение теста - перебираем список Requests из вкладки APIlist
	for rIdx := 0; rIdx < len(header2); rIdx++ {
		req2.SetHeader(header2[rIdx][0], header2[rIdx][1])
	}

	// Perform the request
	res2, err := req2.Send()
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

	respons2.ResponseBody2 = res2.String()
	respons2.ResponseCode2 = res2.StatusCode
	respons2.URL2 = res2.RawRequest.URL.String()
	respons2.Header = res2.Header.Get("authorization")
	// Reads the whole body and returns it as string
	//fmt.Printf("Body: %s", res.String())
	return respons2
}