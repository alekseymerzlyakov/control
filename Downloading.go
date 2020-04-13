package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"os"
)


func Downloading() {

	url := GetPropValue("urlxlsxfile")
	fileName := "APP/" + GetPropValue("countryname") + ".xlsx"
	fmt.Println("Downloading file...")

	output, err := os.Create(fileName)
	defer output.Close()

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println("req", req)
			fmt.Println("via", via)
			fmt.Println("headers", req.Header)
			return nil
		},
	}
	jar, err := cookiejar.New(nil)
	if err != nil {
		fmt.Println("Error while creating cookie jar", url, "-", err)
		return
	}
	client.Jar = jar

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("GET failed", err)
		return
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux i686; rv:17.0) Gecko/20100101 Firefox/17.0")
	response, err := client.Do(req)
	//response, err := client.Get(url)

	if err != nil {
		fmt.Println("Error while downloading", url, "-", err)
		return
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		fmt.Println(response.Status)
		return
	}

	n, err := io.Copy(output, response.Body)
	if err != nil {
		fmt.Println("Error while downloading", url, "-", err)
		return
	}

	fmt.Println(n, "bytes downloaded")
	fmt.Println(jar)
}

