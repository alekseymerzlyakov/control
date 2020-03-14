package Requests

import (
	"fmt"
	"gopkg.in/h2non/gentleman.v2"
)

func Request(header map[int][]string, RequestParameters map[int][]string) {
	// Create a new client

	cli := gentleman.New()
	UR := RequestParameters[1][0] + "://" + RequestParameters[2][0]


	// Define base URL
	cli.URL(UR)
	cli.Method(RequestParameters[0][0])

	// Create a new request based on the current client
	req := cli.Request()

	// Define the URL path at request level
	req.Path(RequestParameters[3][0])
	// Set a new header field
	//req.SetHeader("accept", "application/json]")



	// Выполнение теста - перебираем список Requests из вкладки APIlist
	for rIdx := 0; rIdx < len(header); rIdx++ {
		//fmt.Println(header[rIdx][0],header[rIdx][1])
		req.SetHeader(header[rIdx][0], header[rIdx][1])
		fmt.Println("header  - >    ", header[rIdx][0], header[rIdx][1])
	}



	// Perform the request
	res, err := req.Send()
	if err != nil {
		fmt.Printf("Request error: %s\n", err)
		return
	}
	if !res.Ok {
		fmt.Printf("Invalid server response: %d\n", res.StatusCode)
		return
	}

	// Reads the whole body and returns it as string
	fmt.Printf("Body: %s", res.String())

}