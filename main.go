package main

import (
	"fmt"
	_ "github.com/antchfx/htmlquery"
	"github.com/plandem/xlsx"
	log "github.com/sirupsen/logrus"
	"os"
	_ "strings"
	"time"
)
//
type Data struct {
	sheet xlsx.Sheet
	sheetBefoAftTest xlsx.Sheet
	empty bool
	Countryname string
	Filename string
	ApiSheetName string
	ApiSheetName2 string
	ApiTestName string
	Requestparam
	Cell
	Cell2
	Error
	Response
	Response2
}

type Requestparam struct {
	requestparam map[int][]string
}

var (
	developerKey string
	imei string
	apiversion string
	urlxlsxfile string
    ErrorCount int
	XL xlsx.Spreadsheet
    //Countryname string
)

var RendomData = new(Data)
var Filename = os.Getenv("FILENAME")

//var Countryname = os.Getenv("COUNTRY")



func main() {

	//Countryname = "jordan"
	//RendomData.Countryname = Countryname
	//DataStart.Country = Countryname
	//SetPropValue("countryname",Countryname)

	var filename string = "logfile.log"
	f, _ := os.Create(filename)
	// Create the log file if doesn't exist. And append to it if it already exists.
	f, err := os.OpenFile(filename, os.O_WRONLY | os.O_APPEND | os.O_CREATE, 0644)
	Formatter := new(log.TextFormatter)
	// You can change the Timestamp format. But you have to use the same date and time.
	// "2006-02-02 15:04:06" Works. If you change any digit, it won't work
	// ie "Mon Jan 2 15:04:05 MST 2006" is the reference time. You can't change it
	Formatter.TimestampFormat = "02-01-2006 15:04:05"
	Formatter.FullTimestamp = true
	log.SetFormatter(Formatter)
	if err != nil {
		// Cannot open log file. Logging to stderr
		fmt.Println(err)
	}else{
		log.SetOutput(f)
	}

	//doc, err := htmlquery.LoadURL("https://www.bing.com/search?q=golang")
	//if err != nil {
	//	panic(err)
	//}
	//// Find all news item.
	//list, err := htmlquery.QueryAll(doc, "//ol/li")
	//if err != nil {
	//	panic(err)
	//}
	//for i, n := range list {
	//	a := htmlquery.FindOne(n, "//a")
	//	fmt.Printf("%d %s(%s)\n", i, htmlquery.InnerText(a), htmlquery.SelectAttr(a, "href"))
	//}



//	s := `<html>
//<body onload="document.frmLaunch.submit();">
//<form name="frmLaunch" method="POST" action="https://test.benefit-gateway.com:6443/Gateway/hppaction?formAction=com.aciworldwide.commerce.gateway.payment.action.HostedPaymentPageAction">
//   Processing payment request, Do not close the browser, press back or refresh the page.
//   <input type="hidden" name="PaymentID" value="2200690021701160" />
//</form>
//</body>
//</html>`
//	doc, err := htmlquery.Parse(strings.NewReader(s))
//	list := htmlquery.FindOne(doc, ".//*[contains(@name,'PaymentID')]")
//	//   .//form/input[@name='PaymentID']/@value
//	//   .//*[contains(@name,'PaymentID')]/
//	//fmt.Println(list.Attr[0])
//	//htmlquery.SelectAttr(list, "href")
//	fmt.Println(htmlquery.SelectAttr(list, "value"))
//


////выводит окно формы
//	var inputform string = `
//<html>
//<body>
//<form action="/action_page.php">
//<input type="text" name="username" id="username">
//<input type="text" name="address" id="address">
//<input type="submit" onclick="golangfunc()">
//</form>
//</body>
//</html>
//`
//
//		ui, _ := lorca.New("data:text/html,"+url.PathEscape(inputform), "", 480, 320)
//		ui.Bind("golangfunc", func() {
//			username := ui.Eval(`document.getElementById('username').value`)
//			SetPropValue("paymantprocessor.username",username.String())
//			address := ui.Eval(`document.getElementById('address').value`)
//			SetPropValue("paymantprocessor.address",address.String())
//			fmt.Println(username, address)
//		})
//		defer ui.Close()
//		<-ui.Done()



	//log.Info("Some info. Earth is not flat")
	//log.Warning("This is a warning")
	//log.Error("Not fatal. An error. Won't stop execution")
	//log.Fatal("MAYDAY MAYDAY MAYDAY")
	//log.Panic("Do not panic")



	//fmt.Println(os.Getenv(Filename))
	//SetPropValue("countryname",Countryname)
	//gotenv.OverLoad()


	fmt.Println("Filename  --->   ", Filename)
	startData := Start(Filename)
	RendomData.Filename = Filename
	fmt.Println("RendomData.Filename  --->   ", RendomData.Filename)
	SetPropValue("countryname",RendomData.Countryname)

	//Countryname = startData.Country
	//RendomData.Countryname = os.Getenv("COUNTRY")

	//RendomData.Countryname = startData.Country

	SetPropValue("wiatsms","31")
	SetPropValue("phone","380711111111")
	SetPropValue("phonenumber","711111111")
	SetPropValue("email","alex.mywu.qa@gmail.com")
	SetPropValue("password","Password_01")
	SetPropValue("developerKey","30d9c37d7fa3a9ea534c90251771a3c272571e69924b318e9f13a87105bb31b6")
	SetPropValue("developer_key","df81a5132a0c8e14d3a531475933b0f8a76b2b4271616083283deca2b0111669")


	//Payment
	//SetPropValue("payment.currency_get","RWF")
	//SetPropValue("payment.country_send_code","FR")
	//SetPropValue("payment.amount_get","12")

	Downloading(startData, RendomData.Filename) //download xlsx file from google
	UDID()
	//drivers()
	//checkEmail() // check EMail

	// ------------        выбираем список проверок // выборка row с нужной точки и до пустой ячейки
	GetAPIList := RendomData.GetRow2(1, 1, "APIList", 1)
	XL, err := xlsx.Open(CountryXlsx)
	if err != nil {
		fmt.Println("xlsxmmmmmm not opened")
	}






	// Выполнение теста - перебираем список Requests из вкладки APIlist
	for rIdx := 0; rIdx < len(GetAPIList); rIdx++ {
		//fmt.Println(GetAPIList[rIdx][0],GetAPIList[rIdx][1])
		RendomData.ApiSheetName, RendomData.ApiTestName = GetAPIList[rIdx][0],GetAPIList[rIdx][1]
		getNumTestLine := GetNumTestLine(RendomData.ApiSheetName, RendomData.ApiTestName) // поиск номера строки начала теста
		SetPropValue("sheetname",RendomData.ApiSheetName)   // Сохраняем в дата файл
		SetPropValue("apitestname",RendomData.ApiTestName)
		RendomData.sheet = XL.SheetByName(GetPropValue("sheetname"))
		RendomData.sheetBefoAftTest = XL.SheetByName("BefoAftTest")
		RendomData.ApiTest(RendomData.ApiSheetName, getNumTestLine) //
	}




	//Save report to local disk
	XL.SaveAs("./Report/" + Filename + "Report.xlsx")


	//t := time.Now()
	t := time.Now().String()
	fmt.Println("t := time.Now() ->  ", t)


	//Save report
	upload(startData, t)

	mail(RendomData.Filename)  //Send Email with report
	//telegram("mes string")

	//Send report link to telegram
	Reports(startData, t)














	//
	//fmt.Fprintf (os.Stderr,"*** 開始 ***\n")
	//driver := agouti.ChromeDriver()
	//driver.Start()
	//page, _ := driver.NewPage()
	//url := GetPropValue("paymentprocessor.url")
	//page.Navigate(url)
	//log.Printf(page.Title())
	////
	////tag := page.FindByID("reserve_month")
	////tag.Clear()
	////tag.SendKeys("11")
	////
	////tag = page.FindByID("reserve_day")
	////tag.Clear()
	////tag.SendKeys("10")
	////
	//tag := page.FindByXPath("//input[@id='Ecom_Payment_Card_Number_id']")
	////tag.Clear()
	//tag.Fill("5554555455545555")
	////
	////tag = page.FindByID("headcount")
	////tag.Clear()
	////tag.SendKeys("2")
	////
	//tag = page.FindByXPath("//input[contains(@id,'Ecom_Payment_Card_Name_id')]")
	////tag.Clear()
	//tag.Fill("5554555455545511")
	////
	//tag = page.FindByXPath("//input[contains(@name,'Ecom_Payment_Pin')]")
	////tag.Clear()
	//tag.Click()
	//
	//tag = page.FindByXPath("//input[@name='1']")
	////tag.Clear()
	//tag.Click()
	//tag = page.FindByXPath("//input[@name='2']")
	////tag.Clear()
	//tag.Click()
	//tag = page.FindByXPath("//input[@name='3']")
	////tag.Clear()
	//tag.Click()
	//tag = page.FindByXPath("//input[@name='4']")
	////tag.Clear()
	//tag.Click()
	//
	//
	//tag = page.FindByXPath("//select[contains(@id,'id')][@name='Ecom_Payment_Card_ExpDate_Month']")
	//tag.SendKeys("12")
	//
	//tag = page.FindByXPath("//select[contains(@id,'id')][@name='Ecom_Payment_Card_ExpDate_Year']")
	//tag.SendKeys("2027")
	//
	//tag = page.FindByXPath("//input[contains(@name,'Enter')]")
	////tag.Clear()
	//tag.Click()
	//
	//tag = page.FindByXPath("//input[contains(@id,'EntrySubmitAction_id')]")
	////tag.Clear()
	//tag.Click()
	//
	//tag = page.FindByXPath("//select[contains(@id,'id')][@name='Ecom_Payment_Card_ExpDate_Year']")
	////tag.Clear()
	//tag.Click()
	//
	//tag = page.FindByXPath("//input[contains(@id,'ConfirmAction_id')]")
	////tag.Clear()
	//tag.Click()
	//
	//fmt.Println(page.Session())
	//time.Sleep(60 * time.Second)
	//page.Forward()
	//page.Navigate("https://bahrain.uat.wuamerigo.com/send_money/review/wait?PaymentID=1793551071601160")
	//time.Sleep(60 * time.Second)



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