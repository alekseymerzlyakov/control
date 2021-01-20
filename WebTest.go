package main

import (
	"fmt"
	"strings"
)

const (
	webStartColumn int = 1
)

var NewcurrentRow int

func (a Data) WebTest(ApiSheetName string, startRow int) {

	//Выбираем все строки теста до пробела
	currentRow := startRow

	for rIdx := currentRow; rIdx < 3000; rIdx++ {
		webFunc_ := a.GetCell(webStartColumn, rIdx)
		webFunc := a.Replace(webFunc_)
		NewcurrentRow = rIdx
		if string(webFunc) == "" {
			break
		}

		webFunc = strings.Replace(webFunc, "\n", "", -1)
		webFunc = strings.Replace(webFunc, "\r", "", -1)
		funcLine := strings.Split(webFunc, ";")

		RendomData.MessageWeb = ""

		for zIdx := 0; zIdx < len(funcLine); zIdx++ {

			Func := strings.Replace(funcLine[zIdx], ";", "", -1)

			// Выполняем дейтстия перед и/или после выполнения основного теста
			before_header := ""
			before_header = a.GetCell(BeforeColumn, rIdx)
			if len(before_header) >= 4 || before_header == "" {
				before_after(before_header)
				before_header = ""
			}


			//Generate test data
			generateTestData := GetColumn(GenerateTestData, rIdx, a) //
			a.GenerateData(generateTestData)

			//Выбор функции
			switch {
			case strings.Contains(Func, "<<OpenPage>>"):
				fmt.Println("<<OpenPage>> --->   ", Func)
				url := strings.Split(Func, ">>")
				//Вызов функции
				a.OpenPage(string(url[1]))

			case strings.Contains(Func, "<<WaitURL>>"):
				fmt.Println("<<WaitURL>> --->   ", Func)
				data := strings.Split(Func, ">>")
				a.WaitURL(data[1])

			case strings.Contains(Func, "<<WaitPath>>"):
				fmt.Println("<<WaitPath>> --->   ", Func)
				data := strings.Split(Func, ">>")
				a.WaitPath(data[1])

			case strings.Contains(Func, "<<ClickButton>>"):
				fmt.Println("<<ClickButton>> --->   ", Func)
				data := strings.Split(Func, ">>")
				a.ClickButton(data[1])

			case strings.Contains(Func, "<<PageToTop>>"):
				fmt.Println("<<PageToTop>> --->   ", Func)
				//data := strings.Split(Func, ">>")
				a.PageToTop()

			case strings.Contains(Func, "<<FillField>>"):
				fmt.Println("<<FillField>> --->   ", Func)
				data_ := strings.Split(Func, ">>")
				data := strings.Split(data_[1], "::")
				a.FillField(data[0], data[1])

			case strings.Contains(Func, "<<Select>>"):
				fmt.Println("<<Select>> --->   ", Func)
				data_ := strings.Split(Func, ">>")
				data := strings.Split(data_[1], "::")
				a.Select(data[0], data[1])

			case strings.Contains(Func, "<<Check>>"):
				fmt.Println("<<Check>> --->   ", Func)
				data_ := strings.Split(Func, ">>")
				data := strings.Split(data_[1], "::")
				a.Check(data[0], data[1])

			//
			default:
			}
			after_header := ""
			fmt.Println("after_header --->   ", after_header)
			after_header = a.GetCell(BeforeColumn+1, rIdx)
			fmt.Println("after_header2 --->   ", after_header)
			if len(after_header) >= 2  || after_header == "" {
				before_after(after_header)
				after_header = ""
				fmt.Println("after_header3 --->   ", after_header)
			}


		}

	}

}
