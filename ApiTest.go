package main

import (
	"fmt"
	"strings"
)

const (
	startColumnReq int = 1
	startColumnHeader int = 5
	GenerateTestData int = 7
)


func (a Data) ApiTest(ApiSheetName string, rownum int) {
	startRow := rownum

	// получаем данные  Metod	Protocol	Domain	Path
	RequestParameters := GetColumn(startColumnReq, startRow, a) //MetodProtocolDomainPath := GetColumn(4, 1, ap)
	fmt.Println(RequestParameters)

	// получаем Header
	getHeader := GetRow(0, startColumnHeader, startRow, a) // выбираем список проверок // выборка row с нужной точки и до пустой ячейки
	fmt.Println("test1  --   ",getHeader)

	//Generate test data
	generateTestData := GetColumn(GenerateTestData, startRow, a) //
	a.GenerateData(generateTestData) //генерация тестовых данных


	a.DataForRequest(rownum, getHeader, RequestParameters)  // создание запроса

	// Если в PATH есть }/reg? то генерируем токен
	if strings.Contains(RequestParameters[3][0],"/reg?") {Token()}

}