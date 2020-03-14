package main

import (
	"fmt"
	_ "github.com/plandem/xlsx/types"
)

//
//var ApiSheetName string

func main() {

	Downloading() //download xlsx file from google
	// ------------        выбираем список проверок // выборка row с нужной точки и до пустой ячейки
	GetAPIList := GetRow(1, 1, "APIList", "")

	// Выполнение теста - перебираем список Requests из вкладки APIlist
	for rIdx := 0; rIdx < len(GetAPIList); rIdx++ {
		fmt.Println(GetAPIList[rIdx][0],GetAPIList[rIdx][1])
		ApiSheetName, ApiTestName := GetAPIList[rIdx][0],GetAPIList[rIdx][1]
		ApiTest(ApiSheetName, ApiTestName) //
	}




	//cell := sheet.Cell(1, 2)
	//cell2 := sheet.Cell(1, 3)
	//cell3 := sheet.Cell(1, 4)
	//cell4 := sheet.Cell(1, 5)

	//fmt.Println("cell 2-3", cell, cell2, cell3 , cell4)









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