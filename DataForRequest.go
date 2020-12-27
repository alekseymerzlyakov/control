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

type Error struct {
	errorCount int
}

var BodyRespons string

// Функция по созданию запросов из XLSX
func (a *Data) DataForRequest(rownum int, header map[int][]string, RequestParameters map[int][]string) {
	startRow := rownum + startRowTest
	startColumn := startColumn
	startBodyColumn := startColumn + startBodyCol
	startRownew := startRow
	defaultRow := startRow - 1
	templateRow := startRow - 2
	templabeBodyRow := startRow - 3
	//ErrorCount = 0

	for zIdx := 0; zIdx < 3000; zIdx++ {

		startBodyColumnnew := startBodyColumn
		requestBody := a.GetCell(startBodyColumn, templabeBodyRow)

		// Before function
		before := a.GetCell(BeforeColumn, startRownew)
		before_after(before)
		before = ""

		// After function
		after := a.GetCell(BeforeColumn+1, startRownew)
		if a.GetCell(0, startRownew) == "End" {
			break
		} //if cell == nil { break }

		for xIdx := 0; xIdx < 30; xIdx++ {
			template := a.GetCell(startBodyColumnnew, templateRow)
			bodyROWdata := a.GetCell(startBodyColumnnew, startRownew)

			if bodyROWdata == "" {
				def := a.GetCell(startBodyColumnnew, defaultRow)
				defData := a.Replace(def)
				requestBody = strings.Replace(requestBody, template, defData, -1) // замены шаблонов на переменные
			} else {
				bodyROWdatarep := a.Replace(bodyROWdata)
				requestBody = strings.Replace(requestBody, template, bodyROWdatarep, -1) // замены шаблонов на переменные
			}

			if a.GetCell(startBodyColumnnew, templateRow) == "End" || a.GetCell(startBodyColumnnew, templateRow) == "" {
				break
			} //if cell == nil { break }
			startBodyColumnnew++
		}

		//==============Вызываем функции ожидания OTP=================
		// если в ответе есть определенный текст - вызываем функцию по его обработке
		//if strings.Contains(requestBody,"{\"ops\":[{\"limit\":10,\"offset\":0,\"type\":\"list\",\"obj\":\"node\"") {
		//	time.Sleep(30 * time.Second)
		//}

		fmt.Println("\n\n---------------------------START NEW API-----------------------------")
		fmt.Println("ApiSheetName + ApiTestName  	-->      ", a.ApiSheetName, a.ApiTestName)
		log.Info("ApiSheetName + ApiTestName  	-->      ", a.ApiSheetName, a.ApiTestName)

		resp := a.Request(requestBody, header, RequestParameters) //Request

		SetPropValue("last_response", resp.ResponseBody)

		if resp.ResponseCode == 502 || resp.ResponseCode == 500 {
			fmt.Println("GW is DOWN resp.ResponseCode  ->  ", resp.ResponseCode)
			fmt.Println("Выполнение теста остановленно resp.ResponseBody  ->    ", resp.ResponseBody)
			msg := "GW is DOWN resp.ResponseCode  ->    " + string(resp.ResponseCode)
			msg = msg + "\nВыполнение теста остановленно resp.ResponseBody  ->    \n" + string(resp.ResponseBody)
			msg = msg + a.Response.URL
			telegram(msg)
			RendomData.XL.SaveAs("./Report/" + Filename + "Report.xlsx")
			a.mail(RendomData.Filename) //Send Email
			os.Exit(0)
		}

		log.Info("------------------------------------------------------------------")
		fmt.Println("---------------------------------------\n\n")
		a.SetSell(2, startRownew, requestBody)
		a.SetSell(4, startRownew, strconv.Itoa(resp.ResponseCode))
		a.SetSell(3, startRownew, resp.ResponseBody)

		BodyRespons = resp.ResponseBody

		////==============Вызываем функции обработки ответов=================
		//// если в ответе есть определенный текст - вызываем функцию по его обработке
		//if strings.Contains(BodyRespons,"\"request_proc\":\"ok\",\"ops\"") {
		//	//time.Sleep(30 * time.Second)
		//	otp(BodyRespons)
		//}

		// Функция сохранения всего ответа в переменную
		if strings.Contains(resp.URL, "api/version") {
			SetPropValue("apiversion", BodyRespons)
		}

		//------------------------------------------------
		//-----------------ASSERTION----------------------
		//------------------------------------------------

		//Response Code Assertion
		assertMessage := ""
		a.errorCount = 0

		//fmt.Println("assertMessage, assertError   :::::::::::>>>>>>>>>>>>>>>>>             ", assertMessage)

		ExpRespCode := a.GetCell(5, startRownew)
		//fmt.Println("ExpRespCode   :::::::::::>>>>>>>>>>>>>>>>>             ", ExpRespCode)

		if ExpRespCode != "" {
			if strings.Contains(ExpRespCode, strconv.Itoa(resp.ResponseCode)) {
				assertMessage = assertMessage + "ResponseCode - PASS\n"
			} else {
				assertMessage = assertMessage + "ResponseCode - FAILL\n"
				ErrorCount++
				a.errorCount++
				fmt.Println("ErrorCount       ----->             ", string(ErrorCount))
			}
		}

		// Response Message Assertion
		parseAssertion := a.GetCell(6, startRownew)

		//if parseAssertion != "" && ExpRespCode != "" {assertMessage = " - Please set assertion\n"}
		if parseAssertion != "" {
			parseAssert := strings.Replace(parseAssertion, "\n", "", -1)
			//fmt.Println("parseAssert       ----->             ", parseAssert)
			parseAssertArr := strings.Split(parseAssert, "##")

			for mIdx := 0; mIdx < len(parseAssertArr); mIdx++ {

				tempArr := parseAssertArr[mIdx]
				if parseAssertArr[mIdx] == "" {
					break
				}
				fmt.Println("parseAssert       ----->             ", parseAssertArr[mIdx])

				assertRequired := strings.Split(tempArr, "<<")

				if strings.Contains(BodyRespons, assertRequired[0]) {
					assertMessage = assertMessage + tempArr + " - PASS\n"
				} else {
					assertMessage = assertMessage + tempArr + " - FAIL\n"
					ErrorCount++
					a.errorCount++

					if strings.Contains(tempArr, "Required") {
						fmt.Println("Нет обязательного assertion в ответе ->  ", assertRequired[0])
						fmt.Println("Выполнение теста остановленно")

						msg := "Нет обязательного assertion в ответе ->  " + assertRequired[0]
						msg = msg + "\nВыполнение теста остановленно\n\n"
						msg = msg + " - ApiTestName --->  " + a.ApiTestName
						msg = msg + a.Response.URL + "\n"
						msg = msg + " ResponseBody    --->>>    " + resp.ResponseBody + "\n"
						msg = msg + "Страна --->   " + a.Countryname + "\n"
						telegram(msg)
						//RendomData.XL.SaveAs("./Report/IssueReport.xlsx")
						//RendomData.mail(RendomData.Filename)  //Send Email
						os.Exit(0)
					}
				}
			}
		}

		// Save Result (PASS/Fail)
		a.SetSell(7, startRownew, assertMessage)

		//------------------------------------------------
		//-----------------ParseResponse----------------------
		//------------------------------------------------
		//	data::data.#.language_name;data.#.language_code;
		//  парсим обрезаем и записываем
		//  var newelement string

		parseData := a.GetCell(8, startRownew)
		if parseData != "" {
			parseData = strings.Replace(parseData, "\n", "", -1)
			parseData = strings.Replace(parseData, "\r", "", -1)
			parseData = strings.Replace(parseData, " ", "", -1)
			parseData = strings.Replace(parseData, "##}", "}", -1)
			all_parseRes := strings.Split(parseData, ";")
			for vIdx := 0; vIdx < len(all_parseRes); vIdx++ {
				//FuncName := ""
				if all_parseRes[vIdx] == "" {
					break
				}

				// Пармис HTML
				switch {
				case strings.Contains(all_parseRes[vIdx], "::") && strings.Contains(all_parseRes[vIdx], "htmlquery_value"):
					Dop_parseRes = strings.Split(all_parseRes[vIdx], "::")
					xpath := strings.Split(Dop_parseRes[1], "##")
					doc, _ := htmlquery.Parse(strings.NewReader(resp.ResponseBody))
					list := htmlquery.FindOne(doc, xpath[0])
					//   .//form/input[@name='PaymentID']/@value   --  //   .//*[contains(@name,'PaymentID')]/   ---   //htmlquery.SelectAttr(list, "href")
					valueTemp = htmlquery.SelectAttr(list, "value")
					SetPropValue(Dop_parseRes[0], valueTemp)

					//Парсим JSON
				case strings.Contains(all_parseRes[vIdx], "<<FOR>>") && strings.Contains(all_parseRes[vIdx], "if("):

					//получаем ifTemplate_ ->   Ukraine
					fmt.Println("all_parseRes[vIdx] ->  ", all_parseRes[vIdx])
					re := regexp.MustCompile(`(<<)(.*?)(>>)`)
					forTemplate_ := re.FindAllString(all_parseRes[vIdx], -1)
					forTemplate := forTemplate_[0]
					fmt.Println("forTemplate ->  ", forTemplate)
					ifVaule := regexp.MustCompile(`(==")(.*?)("\))`)
					ifTemplate := ifVaule.FindAllString(all_parseRes[vIdx], -1)
					ifTemplates := strings.Trim(ifTemplate[0], "==\"")
					ifTemplates = strings.Trim(ifTemplates, "\")")
					ifTemplates = a.Replace(ifTemplates)
					fmt.Println("ifTemplate_ ->  ", ifTemplates)

					numm := 0
					//if(data.<<FOR>>.transaction_date=="Ukraine"){aaaaaaa.payment::data.<<FOR>>.transaction_date##bbbbbbb.payment::data.<<FOR>>.transaction_date##}
					for {
						numString := strconv.Itoa(numm)
						fmt.Println("numm  ->  ", numString)
						newall_parseRes := strings.Replace(all_parseRes[vIdx], forTemplate, numString, -1)

						//получаем data.0.transaction_date
						ifPath_ := regexp.MustCompile(`(if\()(.*?)(==")`)
						ifPath := ifPath_.FindAllString(newall_parseRes, -1)
						ifPaths := strings.Trim(ifPath[0], "if(")
						ifPaths = strings.Trim(ifPaths, "==\"")

						//Получаем список для парсинга
						listParse_ := regexp.MustCompile(`(\){)(.*?)(})`)
						listParse := listParse_.FindAllString(newall_parseRes, -1)
						listParses := strings.Trim(listParse[0], "){")
						listParses = strings.Trim(listParses, "}")
						listParsesArray := strings.Split(listParses, "##")

						fmt.Println("newall_parseRes  ->  ", newall_parseRes)
						fmt.Println("forTemplate      ->  ", forTemplate)
						fmt.Println("ifTemplates      ->  ", ifTemplates)
						fmt.Println("ifPaths          ->  ", ifPaths)
						fmt.Println("listParses_       ->  ", listParses)
						fmt.Println("listParsesArray   ->  ", listParsesArray, len(listParsesArray))

						result := gjson.Get(resp.ResponseBody, ifPaths).String()
						fmt.Println("result          ->  ", result)

						if result == ifTemplates {

							for j := 0; j <= len(listParsesArray)-1; j++ {
								//if listParsesArray[j] == "" { break }
								fmt.Println("listParsesArray   ->  ", listParsesArray[j], len(listParsesArray))
								//fmt.Println(listParsesArray[j])

								Dop_parsePath := strings.Split(listParsesArray[j], "::")
								dopPath := len(Dop_parsePath) - 1
								fmt.Println("Dop_parseRes[dopPath] -> ", Dop_parsePath[dopPath])

								valueTemp = gjson.Get(resp.ResponseBody, Dop_parsePath[dopPath]).String()
								if valueTemp == "[]" {
									fmt.Println("В данном Response нет такого ключа  ", Dop_parsePath[dopPath])
									valueTemp = "В данном Response нет " + resp.ResponseBody + " такого ключа  " + Dop_parsePath[dopPath]
								}

								if strings.Contains(valueTemp, "[\"") && strings.Contains(valueTemp, "\"]") {
									valueTemp = strings.Trim(valueTemp, "[\"")
									valueTemp = strings.Trim(valueTemp, "\"]")
								}
								parseDataArr := strings.Split(Dop_parsePath[dopPath], ".")

								key := parseDataArr[len(parseDataArr)-1]
								if len(key) == 1 {
									key = parseDataArr[len(parseDataArr)-2]
								}

								if dopPath == 1 {
									key = Dop_parsePath[0] + "." + key
								}
								SetPropValue(key, valueTemp)
								Dop_parsePath[0] = ""
								dopPath = 0

							}
							break
						}

						numm++
						if numm >= 100 {
							break
						}

					}

				//Парсим JSON
				case !strings.Contains(all_parseRes[vIdx], "htmlquery_value") && !strings.Contains(all_parseRes[vIdx], "<<FOR>>") && !strings.Contains(all_parseRes[vIdx], "if("):

					fmt.Println("all_parseRes[vIdx] ->  ", all_parseRes[vIdx])

					Dop_parseRes = strings.Split(all_parseRes[vIdx], "::")
					dopPath := len(Dop_parseRes) - 1

					valueTemp = gjson.Get(resp.ResponseBody, Dop_parseRes[dopPath]).String()
					if valueTemp == "[]" {
						fmt.Println("В данном Response нет такого ключа  ", Dop_parseRes[dopPath])
						valueTemp = "В данном Response нет " + resp.ResponseBody + " такого ключа  " + Dop_parseRes[dopPath]
					}

					if strings.Contains(valueTemp, "[\"") && strings.Contains(valueTemp, "\"]") {
						valueTemp = strings.Trim(valueTemp, "[\"")
						valueTemp = strings.Trim(valueTemp, "\"]")
					}
					parseDataArr := strings.Split(Dop_parseRes[dopPath], ".")

					key := parseDataArr[len(parseDataArr)-1]
					if len(key) == 1 {
						key = parseDataArr[len(parseDataArr)-2]
					}

					if dopPath == 1 {
						key = Dop_parseRes[0] + "." + key
					}
					SetPropValue(key, valueTemp)
					Dop_parseRes[0] = ""
					dopPath = 0

				default:
					fmt.Println("len(parseDataArr)        -->>>>>>>>>   ", valueTemp)
				}
			}
		}

		fmt.Println("---------------------------------------\n\n\n\n\n")
		startRownew++

		//

		before_after(after)
		//fmt.Println("after -> ", after)
		after = ""
	}

}

func before_after(bef_aft string) {

	if len(bef_aft) >= 1 {
		bef_aft = strings.Replace(bef_aft, "\n", "", -1)
		beforeArr := strings.Split(bef_aft, ";")

		for fIdx := 0; fIdx < len(beforeArr); fIdx++ {
			if beforeArr[fIdx] != "" {
				switch {
				// берем otp.otp1.tex и находим число и записываем назад вместо текста
				case strings.Contains(beforeArr[fIdx], "<<ParseIntOtp>>"):
					otp1 := GetPropValue("otp.otp1.text")
					otp2 := GetPropValue("otp.otp2.text")
					re := regexp.MustCompile("[0-9]+")
					fmt.Println("otp1 otp2   ----------->>>    ", otp1, otp2)
					otp1_int := re.FindAllString(otp1, -1)
					if len(otp1_int) == 0 {
						otp1_int = re.FindAllString("OTP не найден в otp.otp1.text", -1)
					}
					SetPropValue("otp.otp1.text", otp1_int[0])

					otp2_int := re.FindAllString(otp2, -1)
					if len(otp1_int) == 0 {
						otp2_int = re.FindAllString("OTP не найден в otp.otp2.text", -1)
					}
					SetPropValue("otp.otp2.text", otp2_int[0])

				case strings.Contains(beforeArr[fIdx], "<<ApiTest>>"):
					ApiTest(beforeArr[fIdx])

				case strings.Contains(beforeArr[fIdx], "<<Authorization>>"):
					Authorization()

				case strings.Contains(beforeArr[fIdx], "<<JWT>>"):
					JWT()

				case strings.Contains(beforeArr[fIdx], "<<UnixTime>>"):
					UnixTime()

				case strings.Contains(beforeArr[fIdx], "<<Crop>>"):
					Crop(beforeArr[fIdx])

				case strings.Contains(beforeArr[fIdx], "<<Authorization_ozer>>"):
					Authorization_ozer(beforeArr[fIdx])

				case strings.Contains(beforeArr[fIdx], "<<Token>>"):
					Token()

				case strings.Contains(beforeArr[fIdx], "<<udid>>"):
					UDID()

				//case strings.Contains(beforeArr[fIdx], "<<OTP>>"):
				//	Otp()

				case strings.Contains(beforeArr[fIdx], "<<ParseIntFromText>>"):
					ParseIntFromText(beforeArr[fIdx])

				case strings.Contains(beforeArr[fIdx], "<<Wait>>"):
					Wait(beforeArr[fIdx])

				case strings.Contains(beforeArr[fIdx], "<<ReCaptcha>>"):
					ReCaptcha(beforeArr[fIdx])

				default:
					fmt.Println("switch выбора функций   ----------->>>  отработал default - значит не нашли функцию -->>   ", beforeArr[fIdx])

				}
			}
		}
		beforeArr = nil
	}
}
