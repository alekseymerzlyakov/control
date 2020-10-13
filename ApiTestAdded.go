package main

import (
	"fmt"
	"strings"
)

const (
	startColumnReq2    int = 1
	startColumnHeader2 int = 5
	GenerateTestData2  int = 7
)

func (a Data) ApiTest2(ApiSheetName string, numTestLine2 int) {
	startRow2 := numTestLine2

	// получаем данные  Metod	Protocol	Domain	Path
	RequestParameters2 := GetColumn2(startColumnReq2, startRow2, a) //MetodProtocolDomainPath := GetColumn(4, 1, ap)
	//fmt.Println(RequestParameters2)

	// получаем Header

	getHeader2 := GetHeader2(0, startColumnHeader2, startRow2, a) // выбираем список проверок // выборка row с нужной точки и до пустой ячейки
	//getHeader2 := GetRow3(0, startColumnHeader2, startRow2, b) // выбираем список проверок // выборка row с нужной точки и до пустой ячейки
	fmt.Println("getHeader2  --   ", getHeader2)

	//Generate test data
	generateTestData2 := GetColumn2(GenerateTestData2, startRow2, a) //
	a.GenerateData2(generateTestData2)                               //генерация тестовых данных

	a.DataForRequest2(numTestLine2, getHeader2, RequestParameters2) // создание запроса

	// Если в PATH есть }/reg? то генерируем токен
	if strings.Contains(RequestParameters2[3][0], "/reg?") {
		Token()
	}

}
