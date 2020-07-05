package main

import (
	"fmt"
	"strings"
)

const (
	startColumnReq2 int = 1
	startColumnHeader2 int = 5
	GenerateTestData2 int = 7
)

func (b Data) ApiTest2 (ApiSheetName string, numTestLine2 int) {
	startRow2 := numTestLine2

	// получаем данные  Metod	Protocol	Domain	Path
	RequestParameters2 := GetColumn2(startColumnReq2, startRow2, b) //MetodProtocolDomainPath := GetColumn(4, 1, ap)
	fmt.Println(RequestParameters2)

	// получаем Header
	getHeader2 := GetRow3(0, startColumnHeader2, startRow2, b) // выбираем список проверок // выборка row с нужной точки и до пустой ячейки
	fmt.Println("test1  --   ",getHeader2)

	//Generate test data
	generateTestData2 := GetColumn2(GenerateTestData2, startRow2, b) //
	b.GenerateData(generateTestData2) //генерация тестовых данных

	b.DataForRequest2(numTestLine2, getHeader2, RequestParameters2)  // создание запроса

	// Если в PATH есть }/reg? то генерируем токен
	if strings.Contains(RequestParameters2[3][0],"/reg?") {Token()}

}