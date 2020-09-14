package main

import (
	"fmt"
	"github.com/antchfx/htmlquery"
	log "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
	"os"
	"regexp"
	"strconv"
	"strings"
)

//type Error struct {
//	errorCount int
//}

var BodyRespons2 string

// Функция по созданию запросов из XLSX
func (b *Data) DataForRequest2 (rownum2 int,header2 map[int][]string, RequestParameters2 map[int][]string)  {
	startRow2 := rownum2 + startRowTest2
	startColumn := startColumn2
	startBodyColumn2 := startColumn + startBodyCol2
	startRownew2 :=startRow2
	defaultRow2 := startRow2-1
	templateRow2 := startRow2-2
	templabeBodyRow2 := startRow2-3
	b.Error.errorCount = 0

	for zIdx := 0; zIdx < 3000; zIdx++ {

		startBodyColumnnew2 := startBodyColumn2
		requestBody2 := b.GetCell2(startBodyColumn2, templabeBodyRow2)


		// Before function
		//
		//
		//

		if b.GetCell2(0, startRownew2) == "End" { break } //if cell == nil { break }



		for sIdx := 0; sIdx < 30; sIdx++ {
			template2 := b.GetCell2(startBodyColumnnew2, templateRow2)
			fmt.Println("template2  	-->      ", template2)

			bodyROWdata2 := b.GetCell2(startBodyColumnnew2, startRownew2)
			fmt.Println("bodyROWdata2  	-->      ", bodyROWdata2)

			if bodyROWdata2 == "" {
				def2 := b.GetCell2(startBodyColumnnew2, defaultRow2)
				defData2 := b.Replace2(def2)
				requestBody2 = strings.Replace(requestBody2, template2, defData2, -1) // замены шаблонов на переменные
			} else {
				bodyROWdatarep2 := b.Replace2(bodyROWdata2)
				fmt.Println("1 requestBody2  	-->      ", requestBody2)
				fmt.Println("2 template2  	-->      ", template2)
				fmt.Println("3 bodyROWdatarep2  	-->      ", bodyROWdatarep2)
				requestBody2 = strings.Replace(requestBody2, template2, bodyROWdatarep2, -1) // замены шаблонов на переменные
			}

			if b.GetCell2(startBodyColumnnew2, templateRow2) == "End" || b.GetCell2(startBodyColumnnew2, templateRow2) == "" { break } //if cell == nil { break }
			startBodyColumnnew2++
		}

		//==============Вызываем API =================

		resp2 := Request2(requestBody2, header2, RequestParameters2)  //Request
		fmt.Println("ApiTest Name  	-->      ", b.ApiTestName)
		log.Info("ApiTest Name        -->      ", b.ApiTestName)
		fmt.Println("Request URL         -->      ", resp2.URL2)
		log.Info("Request URL         -->      ", resp2.URL2)
		fmt.Println("Request Body        -->      ", requestBody2)
		log.Info("Request Body        -->      ", requestBody2)
		fmt.Println("Response Code 	  -->       ", resp2.ResponseCode2)
		log.Info("Response Code       -->       ", resp2.ResponseCode2)
		fmt.Println("Response Body       -->     ", resp2.ResponseBody2)
		log.Info("Response Body       -->     ", resp2.ResponseBody2)
		log.Info("------------------------------------------------------------------")
		fmt.Println("---------------------------------------\n\n")
		b.SetSell2(2, startRownew2, requestBody2)
		b.SetSell2(4, startRownew2, strconv.Itoa(resp2.ResponseCode2))
		b.SetSell2(3, startRownew2, resp2.ResponseBody2)


		if resp2.ResponseCode2 == 502 || resp2.ResponseCode2 == 500  {
			fmt.Println("GW is DOWN resp.ResponseCode  ->  ", resp2.ResponseCode2)
			fmt.Println("Выполнение теста остановленно resp.ResponseBody  ->    ", resp2.ResponseBody2)
			msg := "GW is DOWN resp.ResponseCode  ->    " + string(resp2.ResponseCode2)
			msg = msg + "\nВыполнение теста остановленно resp.ResponseBody  ->    \n" + string(resp2.ResponseBody2)
			msg = msg + b.Response.URL
			telegram(msg)
			XL.SaveAs("./Report/" + Filename + "Report.xlsx")
			mail(RendomData.Filename)  //Send Email
			os.Exit(0)
		}

		BodyRespons = resp2.ResponseBody2

		////==============Вызываем функции обработки ответов=================
		//// если в ответе есть определенный текст - вызываем функцию по его обработке
		//if strings.Contains(BodyRespons,"\"request_proc\":\"ok\",\"ops\"") {
		//	//time.Sleep(30 * time.Second)
		//	otp(BodyRespons)
		//}


		// Функция сохранения всего ответа в переменную
		if strings.Contains(resp2.URL2,"api/version") {
			SetPropValue2("apiversion", BodyRespons)
		}


		//------------------------------------------------
		//-----------------ASSERTION----------------------
		//------------------------------------------------

		//Response Code Assertion
		assertMessage2 := ""
		fmt.Println("assertMessage, assertError   :::::::::::>>>>>>>>>>>>>>>>>             ", assertMessage2)

		ExpRespCode2 := b.GetCell2(5, startRownew2)
		fmt.Println("ExpRespCode2   :::::::::::>>>>>>>>>>>>>>>>>             ", ExpRespCode2)

		if strings.Contains(ExpRespCode2, strconv.Itoa(resp2.ResponseCode2)) {
			assertMessage2 = assertMessage2 + "ResponseCode - PASS\n"
		} else {
			assertMessage2 = assertMessage2 + "ResponseCode - FAILL\n"
			ErrorCount++
			b.errorCount++
		}


// Response Message Assertion
		parseAssertion := b.GetCell2(6, startRownew2)
		if parseAssertion != "" {
			parseAssert := strings.Replace(parseAssertion, "\n", "", -1)
			fmt.Println("parseAssert   ------->>>>             ", parseAssert)
			parseAssertArr := strings.Split(parseAssert, "##")

			for mIdx := 0; mIdx < len(parseAssertArr); mIdx++ {
				tempArr2 := parseAssertArr[mIdx]
				if parseAssertArr[mIdx] == ""  { break }
				fmt.Println("parseAssert   ------->>>>             ", parseAssertArr[mIdx])

				assertRequired2 := strings.Split(tempArr2, "<<")

				if strings.Contains(BodyRespons2, tempArr2) {
					//assertMessage = assertMessage + tempArr + " - PASS\n"
				} else {
					assertMessage2 = assertMessage2 + tempArr2 + " - FAIL\n"
					ErrorCount++
					b.errorCount++

					if strings.Contains(tempArr2, "Required") {
						fmt.Println("Нет обязательного assertion в ответе ->  ", assertRequired2[0])
						fmt.Println("Выполнение теста остановленно")

						msg := "Нет обязательного assertion в ответе ->  " +assertRequired2[0]
						msg = msg + "\nВыполнение теста остановленно\n\n"
						msg = msg + " - ApiTestName --->  " + b.ApiTestName
						msg = msg + b.Response.URL + "\n"

						//msg = msg + " ResponseBody    --->>>    " + resp2.ResponseBody2 + "\n"
						msg = msg + "Страна --->   " + b.Countryname + "\n"
						telegram(msg)
						//XL.SaveAs("./Report/IssueReport.xlsx")
						//mail2(RendomData.Filename)  //Send Email
						os.Exit(0)
					}
				}
				}
			}

		// Save Result (PASS/Fail)
		b.SetSell2(7, startRownew2, assertMessage2)








		//------------------------------------------------
		//-----------------ParseResponse----------------------
		//------------------------------------------------
		//	data::data.#.language_name;data.#.language_code;
		//  парсим обрезаем и записываем
		//  var newelement string


		parseData := b.GetCell2(8, startRownew2)
		if parseData != "" {
			parseData = strings.Replace(parseData, "\n", "", -1)
			all_parseRes := strings.Split(parseData, ";")
			for vIdx := 0; vIdx < len(all_parseRes); vIdx++ {
				funcName2 := ""
				if all_parseRes[vIdx] == ""  { break }

				switch {
				case  strings.Contains(all_parseRes[vIdx], "::") && strings.Contains(all_parseRes[vIdx], "htmlquery_value"):
					Dop_parseRes = strings.Split(all_parseRes[vIdx], "::")

					xpath := strings.Split(Dop_parseRes[1], "##")

					doc, _ := htmlquery.Parse(strings.NewReader(resp2.ResponseBody2))
					list := htmlquery.FindOne(doc, xpath[0])
					fmt.Println("xpath[0]        -->     ", xpath[0])
					//   .//form/input[@name='PaymentID']/@value
					//   .//*[contains(@name,'PaymentID')]/
					//fmt.Println(list.Attr[0])
					//htmlquery.SelectAttr(list, "href")
					valueTemp = htmlquery.SelectAttr(list, "value")
					fmt.Println("valueTemp  --->>>   ", valueTemp)
					SetPropValue2(Dop_parseRes[0],  valueTemp)




				case  strings.Contains(all_parseRes[vIdx], "::") && !strings.Contains(all_parseRes[vIdx], "htmlquery_value"):
					Dop_parseRes = strings.Split(all_parseRes[vIdx], "::")

//otp1.text<<ParseOtp##Your one-time WU password is>>
					if strings.Contains(Dop_parseRes[1], "<<") && strings.Contains(Dop_parseRes[1], ">>") {
						tempPers := strings.Split(Dop_parseRes[1], "<<")
						Dop_parseRes[1] = tempPers[0]
						fmt.Println("------------  >>>>    Dop_parseRes[1]  ", Dop_parseRes[1])
						funcName2 = strings.Trim(tempPers[1], ">>")
						fmt.Println("------------  >>>>    funcName2  ", funcName2)
					}

					fmt.Println("------------  >>>>    Dop_parseRes[1]  ", Dop_parseRes[1])
					valueTemp = gjson.Get(resp2.ResponseBody2,Dop_parseRes[1]).String()
					fmt.Println("------------  >>>>    valueTemp  ", valueTemp)
					if valueTemp == "[]" {
						fmt.Println("В данном Response нет такого ключа  ", Dop_parseRes[1])
						valueTemp = "В данном Response нет " + resp2.ResponseBody2 + " такого ключа  " + Dop_parseRes[1]
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
					case funcName2 == "ParseOtp":
						re := regexp.MustCompile("[0-9]+")
						s := re.FindAllString(valueTemp, -1)
						if len(s) == 0 {
							s = re.FindAllString("9999", -1)
							}
						SetPropValue2(key,  s[0])
					default:
						SetPropValue2(key,  valueTemp)

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
						funcName2 = strings.Trim(tempPers[1], ">>")
						fmt.Println("------------  >>>>    funcName  ", funcName2)
					}


					valueTemp = gjson.Get(resp2.ResponseBody2,Dop_parseRes[0]).String()
					if valueTemp == "[]" {
						fmt.Println("В данном Response нет такого ключа  ", Dop_parseRes[0])
						valueTemp = "В данном Response нет " + resp2.ResponseBody2 + " такого ключа  " + Dop_parseRes[0]
					}

					if strings.Contains(valueTemp, "[\"") && strings.Contains(valueTemp, "\"]") {
						valueTemp = strings.Trim(valueTemp, "[\"")
						valueTemp = strings.Trim(valueTemp, "\"]")
					}
					parseDataArr := strings.Split(all_parseRes[vIdx], ".")
					key := parseDataArr[len(parseDataArr)-1]
					if valueTemp != "" {

						switch {
						case funcName2 == "ParseOtp":
							re := regexp.MustCompile("[0-9]+")
							s := re.FindAllString(valueTemp, -1)
							SetPropValue2(key,  s[0])
						default:
							SetPropValue2(key,  valueTemp)
						}
						//SetPropValue(key,  valueTemp)
					}

					Dop_parseRes[0] = ""
					fmt.Println("len(parseDataArr)        -->>>>>>>>>   ", key, valueTemp)
				}
			}
		}

		fmt.Println("---------------------------------------\n\n\n\n\n")
		startRownew2++
	}



	// After function



}
