package main

const (
	startColumnReq    int = 1
	startColumnHeader int = 5
	GenerateTestData  int = 7
)

func (a Data) ApiTest(ApiSheetName string, numTestLine int) {
	startRow := numTestLine

	before_header := a.GetCell(BeforeColumn, startRow)
	before_after(before_header)
	before_header = ""

	// получаем данные  Metod	Protocol	Domain	Path
	RequestParameters := GetColumn(startColumnReq, startRow, a) //MetodProtocolDomainPath := GetColumn(4, 1, ap)
	//fmt.Println(RequestParameters)

	// получаем Header
	//getHeader := GetRow(0, startColumnHeader, startRow, a) // выбираем список проверок // выборка row с нужной точки и до пустой ячейки
	getHeader := GetHeader(0, startColumnHeader, startRow, a) // выбираем список проверок // выборка row с нужной точки и до пустой ячейки
	//fmt.Println("test1  --   ", getHeader)

	//Generate test data
	generateTestData := GetColumn(GenerateTestData, startRow, a) //
	a.GenerateData(generateTestData)                             //генерация тестовых данных

	a.DataForRequest(numTestLine, getHeader, RequestParameters) // создание запроса

	// After function
	after_header := a.GetCell(BeforeColumn+1, startRow)
	before_after(after_header)
	after_header = ""

	//// Если в PATH есть }/reg? то генерируем токен
	//if strings.Contains(RequestParameters[3][0], "/reg?") {
	//	Token()
	//}

}
