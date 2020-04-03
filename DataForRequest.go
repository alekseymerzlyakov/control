package main

import (
	"fmt"
	"github.com/alekseymerzlyakov/control/Requests"
	"github.com/tidwall/gjson"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Error struct {
	errorCount int
}


// Функция по созданию запросов из XLSX
func (a Data) DataForRequest(rownum int,header map[int][]string, RequestParameters map[int][]string)  {
	startRow := rownum + startRowTest
	startColumn := startColumn
	startBodyColumn := startColumn + startBodyCol
	startRownew :=startRow
	defaultRow := startRow-1
	templateRow := startRow-2
	templabeBodyRow := startRow-3
	a.Error.errorCount = 0


	for zIdx := 0; zIdx < 3000; zIdx++ {

		startBodyColumnnew := startBodyColumn
		requestBody := a.GetCell(startBodyColumn, templabeBodyRow)
		if a.GetCell(0, startRownew) == "End" { break } //if cell == nil { break }

		for xIdx := 0; xIdx < 30; xIdx++ {
			template := a.GetCell(startBodyColumnnew, templateRow)
			bodyROWdata := a.GetCell(startBodyColumnnew, startRownew)

			if bodyROWdata == "" {
				def := a.GetCell(startBodyColumnnew, defaultRow)
				defData := Replace(def)
				requestBody = strings.Replace(requestBody, template, defData, -1) // замены шаблонов на переменные
			} else {
				bodyROWdatarep := Replace(bodyROWdata)
				requestBody = strings.Replace(requestBody, template, bodyROWdatarep, -1) // замены шаблонов на переменные
			}

			if a.GetCell(startBodyColumnnew, templateRow) == "End" || a.GetCell(startBodyColumnnew, templateRow) == "" { break } //if cell == nil { break }
			startBodyColumnnew++
		}

		//==============Вызываем функции ожидания OTP=================
		// если в ответе есть определенный текст - вызываем функцию по его обработке
		if strings.Contains(requestBody,"{\"ops\":[{\"limit\":10,\"offset\":0,\"type\":\"list\",\"obj\":\"node\"") {
			time.Sleep(30 * time.Second)
		}
		time.Sleep(1 * time.Second)
		resp := Requests.Request(requestBody, header, RequestParameters)

		fmt.Println("---------------------------------------\n\n\n\n\n")
		fmt.Println("a.ApiTestName        -->>>>>>>>>   ", a.ApiTestName)
		fmt.Println("URL        -->>>>>>>>>   ", resp.URL)
		fmt.Println("requestBody        -->>>>>>>>>   ", requestBody)
		fmt.Println("resp.ResponseCode  -->>>>>>>>>   ", resp.ResponseCode)
		fmt.Println("resp.ResponseBody  -->>>>>>>>>   \n", resp.ResponseBody)

		a.SetSell(2, startRownew, requestBody)
		a.SetSell(4, startRownew, strconv.Itoa(resp.ResponseCode))
		a.SetSell(3, startRownew, resp.ResponseBody)

		BodyRespons := resp.ResponseBody

		////==============Вызываем функции обработки ответов=================
		//// если в ответе есть определенный текст - вызываем функцию по его обработке
		//if strings.Contains(BodyRespons,"\"request_proc\":\"ok\",\"ops\"") {
		//	//time.Sleep(30 * time.Second)
		//	otp(BodyRespons)
		//}

		if strings.Contains(resp.URL,"api/version") {
			SetPropValue("apiversion", BodyRespons)
		}


		//------------------------------------------------
		//-----------------ASSERTION----------------------
		//------------------------------------------------

		//Response Code Assertion
		assertMessage := ""
		fmt.Println("assertMerrage, assertError   :::::::::::>>>>>>>>>>>>>>>>>             ", assertMessage)

		ExpRespCode := a.GetCell(5, startRownew)
		fmt.Println("ExpRespCode   :::::::::::>>>>>>>>>>>>>>>>>             ", ExpRespCode)

		if strings.Contains(ExpRespCode, strconv.Itoa(resp.ResponseCode)) {
			assertMessage = assertMessage + "ResponseCode - PASS\n"
		} else {
			assertMessage = assertMessage + "ResponseCode - FAILL\n"
			a.Error.errorCount++
		}




// Response Message Assertion
		parseAssertion := a.GetCell(6, startRownew)
		if parseAssertion != "" {
			parseAssert := strings.Replace(parseAssertion, "\n", "", -1)
			fmt.Println("parseAssert   :::::::::::>>>>>>>>>>>>>>>>>             ", parseAssert)
			parseAssertArr := strings.Split(parseAssert, "##")

			for mIdx := 0; mIdx < len(parseAssertArr); mIdx++ {
				tempArr := parseAssertArr[mIdx]
				if parseAssertArr[mIdx] == ""  { break }
				fmt.Println("parseAssert   :::::::::::>>>>>>>>>>>>>>>>>             ", parseAssertArr[mIdx])
				if strings.Contains(BodyRespons, tempArr) {
					//assertMessage = assertMessage + tempArr + " - PASS\n"
				} else {
					assertMessage = assertMessage + tempArr + " - FAIL\n"
					a.Error.errorCount++
				}
				}
			}
		//}

		// Save Result (PASS/Fail)
		a.SetSell(7, startRownew, assertMessage)


		//------------------------------------------------
		//-----------------ParseResponse----------------------
		//------------------------------------------------
		//	data::data.#.language_name;data.#.language_code;
		//	data:: - это дополнительный путь который будет в писан в ./data.json file
		//  парсим обрезаем и записываем
		//  var newelement string


		parseData := a.GetCell(8, startRownew)
		if parseData != "" {
			parseData = strings.Replace(parseData, "\n", "", -1)
			all_parseRes := strings.Split(parseData, ";")
			for vIdx := 0; vIdx < len(all_parseRes); vIdx++ {
				FuncName := ""
				if all_parseRes[vIdx] == ""  { break }

				switch {
				case  strings.Contains(all_parseRes[vIdx], "::"):
					Dop_parseRes = strings.Split(all_parseRes[vIdx], "::")

//otp1.text<<ParseOtp##Your one-time WU password is>>
					if strings.Contains(Dop_parseRes[1], "<<") && strings.Contains(Dop_parseRes[1], ">>") {
						tempPers := strings.Split(Dop_parseRes[1], "<<")
						Dop_parseRes[1] = tempPers[0]
						fmt.Println("------------  >>>>    Dop_parseRes[1]  ", Dop_parseRes[1])
						FuncName = strings.Trim(tempPers[1], ">>")
						fmt.Println("------------  >>>>    funcName  ", FuncName)
					}

					fmt.Println("------------  >>>>    Dop_parseRes[1]  ", Dop_parseRes[1])
					valueTemp = gjson.Get(resp.ResponseBody,Dop_parseRes[1]).String()
					fmt.Println("------------  >>>>    valueTemp  ", valueTemp)
					if valueTemp == "[]" {
						fmt.Println("В данном Response нет такого ключа  ", Dop_parseRes[1])
						valueTemp = "В данном Response нет " + resp.ResponseBody + " такого ключа  " + Dop_parseRes[1]
					}

					if strings.Contains(valueTemp, "[\"") && strings.Contains(valueTemp, "\"]") {
						valueTemp = strings.Trim(valueTemp, "[\"")
						valueTemp = strings.Trim(valueTemp, "\"]")
					}
					parseDataArr := strings.Split(Dop_parseRes[1], ".")
					key := parseDataArr[len(parseDataArr)-1]

					if Dop_parseRes[0] != "" {
						key = Dop_parseRes[0] + "." + key
					}

					switch {
					case FuncName == "ParseOtp":
						re := regexp.MustCompile("[0-9]+")
						s := re.FindAllString(valueTemp, -1)
						if len(s) == 0 {
							s = re.FindAllString("9999", -1)
							}
						SetPropValue(key,  s[0])
					default:
						SetPropValue(key,  valueTemp)

					}

					Dop_parseRes[0] = ""
					fmt.Println("len(parseDataArr)        -->>>>>>>>>   ", key, valueTemp)

				default:
					Dop_parseRes = strings.Split(all_parseRes[vIdx], "::")

					//otp1.text<<ParseOtp##Your one-time WU password is>>
					//Dop_parseRes[0] = all_parseRes[vIdx]
					if strings.Contains(Dop_parseRes[0], "<<") && strings.Contains(Dop_parseRes[0], ">>") {
						tempPers := strings.Split(Dop_parseRes[0], "<<")
						Dop_parseRes[0] = tempPers[0]
						fmt.Println("------------  >>>>    Dop_parseRes[0]  ", Dop_parseRes[0])
						FuncName = strings.Trim(tempPers[1], ">>")
						fmt.Println("------------  >>>>    funcName  ", FuncName)
					}


					valueTemp = gjson.Get(resp.ResponseBody,Dop_parseRes[0]).String()
					if valueTemp == "[]" {
						fmt.Println("В данном Response нет такого ключа  ", Dop_parseRes[0])
						valueTemp = "В данном Response нет " + resp.ResponseBody + " такого ключа  " + Dop_parseRes[0]
					}

					if strings.Contains(valueTemp, "[\"") && strings.Contains(valueTemp, "\"]") {
						valueTemp = strings.Trim(valueTemp, "[\"")
						valueTemp = strings.Trim(valueTemp, "\"]")
					}
					parseDataArr := strings.Split(all_parseRes[vIdx], ".")
					key := parseDataArr[len(parseDataArr)-1]
					if valueTemp != "" {

						switch {
						case FuncName == "ParseOtp":
							re := regexp.MustCompile("[0-9]+")
							s := re.FindAllString(valueTemp, -1)
							SetPropValue(key,  s[0])
						default:
							SetPropValue(key,  valueTemp)
						}
						//SetPropValue(key,  valueTemp)
					}

					Dop_parseRes[0] = ""
					fmt.Println("len(parseDataArr)        -->>>>>>>>>   ", key, valueTemp)
				}
			}
		}

		fmt.Println("---------------------------------------\n\n\n\n\n")
		startRownew++
	}
}
