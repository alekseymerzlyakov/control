package main

import (
	"fmt"
	"github.com/plandem/xlsx"
	"os"
)

//
type Data struct {
	sheet xlsx.Sheet
	empty bool
	ApiSheetName string
	ApiTestName string
	Requestparam
	Cell
	Error
	Response
}

type Requestparam struct {
	requestparam map[int][]string
}

var (
	countryname string
	developerKey string
	imei string
	apiversion string
	rwandareport string
	urlxlsxfile string
)

	var RendomData = new(Data)

func main() {
	SetPropValue("countryname",os.Getenv(countryname))
	SetPropValue("urlxlsxfile","https://docs.google.com/spreadsheets/d/1vEomALrKNbteAE1uEr2iSyenRpDRQt06EVKi53vDiVU/export?format=xlsx")



	SetPropValue("wiatsms","30")
	SetPropValue("phone","380711111111")
	SetPropValue("phonenumber","711111111")
	SetPropValue("email","alex.mywu.qa@gmail.com")
	SetPropValue("password","Password_01")
	SetPropValue("developerKey","30d9c37d7fa3a9ea534c90251771a3c272571e69924b318e9f13a87105bb31b6")
	SetPropValue("developer_key","df81a5132a0c8e14d3a531475933b0f8a76b2b4271616083283deca2b0111669")
	//SetPropValue("rwandareport","rwandaReport")

	//Payment
	SetPropValue("payment.currency_get","RWF")
	SetPropValue("payment.country_send_code","FR")
	SetPropValue("payment.amount_get","11")
	Downloading() //download xlsx file from google
	UDID()
	//drivers()
	//telegram("dfgfrg")
	//checkEmail() // check EMail
	//telegram()

	// ------------        выбираем список проверок // выборка row с нужной точки и до пустой ячейки
	GetAPIList := RendomData.GetRow2(1, 1, "APIList", 1)

	xl, err := xlsx.Open(CountryName)
	if err != nil {
		fmt.Println("xlsxmmmmmm not opened")
	}

	// Выполнение теста - перебираем список Requests из вкладки APIlist
	for rIdx := 0; rIdx < len(GetAPIList); rIdx++ {
		//fmt.Println(GetAPIList[rIdx][0],GetAPIList[rIdx][1])
		RendomData.ApiSheetName, RendomData.ApiTestName = GetAPIList[rIdx][0],GetAPIList[rIdx][1]
		getNumTestLine := GetNumTestLine(RendomData.ApiSheetName, RendomData.ApiTestName)

		SetPropValue("sheetname",RendomData.ApiSheetName)   // Сохраняем в дата файл
		SetPropValue("apitestname",RendomData.ApiTestName)

		RendomData.sheet = xl.SheetByName(GetPropValue("sheetname"))
		RendomData.ApiTest(RendomData.ApiSheetName, getNumTestLine) //
	}


	xl.SaveAs("./Report/" + GetPropValue("countryname") + "Report.xlsx")



	mail()  //Send Email with report

	//get, err := viper.Get(key).(string)
	//if get == "" {
	//	msg := GetPropValue("countryname") + "\nТакой переменной нет в файле - data.json ->  " + key
	//	telegram(msg)
	//	fmt.Println("Такой переменной нет в файле - data.json ->  ", key, err)
	//	//os.Exit(0)
	//}












	//fmt.Fprintf (os.Stderr,"*** 開始 ***\n")
	//driver := agouti.ChromeDriver()
	//driver.Start()
	//
	//page, _ := driver.NewPage()
	//
	//url := "http://example.selenium.jp/reserveApp/"
	//page.Navigate(url)
	//log.Printf(page.Title())
	//
	//tag := page.FindByID("reserve_month")
	//tag.Clear()
	//tag.SendKeys("11")
	//
	//tag = page.FindByID("reserve_day")
	//tag.Clear()
	//tag.SendKeys("10")
	//
	//tag = page.FindByID("reserve_term")
	//tag.Clear()
	//tag.SendKeys("3")
	//
	//tag = page.FindByID("headcount")
	//tag.Clear()
	//tag.SendKeys("2")
	//
	//str_select := "input[type='radio'][name='bf'][value='off']"
	//item := page.Find(str_select)
	//item.Click()
	//
	//str_select = "input[type='checkbox'][id='plan_b']"
	//item = page.Find(str_select)
	//item.Click()
	//
	//tag = page.FindByID("guestname")
	//tag.Clear()
	//tag.SendKeys("明智小五郎")
	//
	//tag = page.FindByID("goto_next")
	//tag.Click()
	//
	//fmt.Fprintf (os.Stderr,"*** 終了 ***\n")
	//defer driver.Stop()
}







//////////////
//package main
//
//import (
//	"github.com/fedesog/webdriver"
//)
//
//func main() {
//	chromeDriver := webdriver.NewChromeDriver("./vendor/chromedriver")
//	err := chromeDriver.Start()
//	if err != nil {
//		panic(err)
//	}
//	desired := webdriver.Capabilities{"Platform": "Mac"}
//	required := webdriver.Capabilities{}
//	session, err := chromeDriver.NewSession(desired, required)
//	if err != nil {
//		panic(err)
//	}
//
//	session.Url("http://stackoverflow.com")
//	el, err := session.FindElement("tag name", "body")
//	el2, err := session.FindElement("tag name", "body")
//	if err != nil {
//		panic(err)
//	}
//	err = el.SendKeys("ctrl t")
//	if err != nil {
//		panic(err)
//	}
//}