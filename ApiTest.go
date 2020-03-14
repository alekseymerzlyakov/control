package main

import (
	"fmt"
	"github.com/alekseymerzlyakov/control/Requests"
)

const (
	startColumnReq int = 1
	startColumnHeader int = 5
	GenerateTestData int = 7
)

func ApiTest(ApiSheetName string, ApiTestName string) {

	// получаем данные  Metod	Protocol	Domain	Path
	// передаем номер колонки с какой начать
	RequestParameters := GetColumn(startColumnReq, ApiSheetName, ApiTestName) //MetodProtocolDomainPath := GetColumn(4, 1, ap)
	fmt.Println(RequestParameters)

	getHeader := GetRow(0, startColumnHeader, ApiSheetName, ApiTestName) // выбираем список проверок // выборка row с нужной точки и до пустой ячейки
	fmt.Println("test1  --   ",getHeader)

	generateTestData := GetColumn(GenerateTestData, ApiSheetName, ApiTestName) //MetodProtocolDomainPath := GetColumn(4, 1, ap)
	fmt.Println("generateTestData  --------->>   >>>>         ",generateTestData)

	Requests.Request(getHeader, RequestParameters)

	//fmt.Println("GenerateData(generateTestData)  -->>>>>>>>>   ", GenerateData(generateTestData))


	//GenerateData(generateTestData)



	//generateMap := make(map[int][]string)
	//generateMap := strings.Split(generateTestData[0], ";")






}