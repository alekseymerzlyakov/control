package main

import (
	"fmt"
	"github.com/sclevine/agouti"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var PassFail int = 11

// Open Page
func (a Data) OpenPage(url string) {

	//driver :=  agouti.ChromeDriver()

	driver := agouti.ChromeDriver(
		agouti.ChromeOptions("args", []string{"--kiosk"}),
		//agouti.ChromeOptions("args", []string{"--headless", "--disable-gpu", "--no-sandbox"}),
	)

	driver.Start()
	page, err := driver.NewPage()
	if err != nil {
		log.Fatal("Failed to open page:", err)
	}

	RendomData.page = page
	RendomData.page.Navigate(url)
	currentUrl, _ := RendomData.page.URL()
	fmt.Println("currentUrl -->  ", currentUrl)
	SetPropValue("web.url", currentUrl)
	//return a.page;
	RendomData.MessageWeb = RendomData.MessageWeb + "OpenPage -> PASS\n"
	a.SetSell(PassFail, NewcurrentRow, RendomData.MessageWeb)
}

// Wait URL opened
func (a Data) WaitURL(url string) {

	for zIdx := 0; zIdx < 120; zIdx++ {
		currentUrl, _ := RendomData.page.URL()
		SetPropValue("web.url", currentUrl)
		fmt.Println("WaitURL -->  ", currentUrl)
		if strings.Contains(currentUrl, url) {
			RendomData.MessageWeb = RendomData.MessageWeb + "WaitURL -> PASS\n"
			fmt.Println("MessageWeb -->  ", RendomData.MessageWeb)
			a.SetSell(PassFail, NewcurrentRow, RendomData.MessageWeb)
			break
		}

		time.Sleep(1 * time.Second)
	}
	currentUrl, _ := RendomData.page.URL()
	fmt.Println("WaitURL -->  ", currentUrl)
	if !strings.Contains(currentUrl, url) {
		RendomData.MessageWeb = "Страница ->  " + url + " не открылась" +
			"\nВыполнение теста остановленно" +
			"\n - Web Test Name --->  " + RendomData.ApiTestName +
			"\n - Api Sheet Name --->  " + RendomData.ApiSheetName +
			"\nCountry --->   " + RendomData.Countryname
		a.Errors(RendomData.MessageWeb)
	}

}

// Wait Path opened
func (a Data) WaitPath(path string) {
	// Wait path
	for zIdx := 0; zIdx < 60; zIdx++ {
		xpathText, err := RendomData.page.FindByXPath(path).Text()
		fmt.Println("WaitPath -->  ", xpathText)
		if err != nil {
			log.Println("WaitPath - Failed to Wait Path:", err)
			time.Sleep(1 * time.Second)
		}

		RendomData.page.Screenshot("qew")
		if err == nil {
			RendomData.MessageWeb = RendomData.MessageWeb + "WaitPath -> PASS\n"
			a.SetSell(PassFail, NewcurrentRow, RendomData.MessageWeb)
			break
		}
	}

	xpathText, err := RendomData.page.FindByXPath(path).Text()
	if err != nil {
		log.Println("WaitPath - Failed to Wait Path:", xpathText, err)
		RendomData.MessageWeb = RendomData.MessageWeb + "WaitPath ->  " + path + " не найден\n" +
			"\nВыполнение теста остановленно" +
			"\n - Web Test Name --->  " + RendomData.ApiTestName +
			"\n - Api Sheet Name --->  " + RendomData.ApiSheetName +
			"\nCountry --->   " + RendomData.Countryname
		a.Errors(RendomData.MessageWeb)
	}
}

// Click Button
func (a Data) ClickButton(path string) {
	// Wait path
	for zIdx := 0; zIdx < 60; zIdx++ {
		xpathText, err := RendomData.page.FindByXPath(path).Text()
		fmt.Println("ClickButton -->  ", xpathText)

		if err != nil {
			log.Println("ClickButton - Failed to find Path:", err)
			time.Sleep(1 * time.Second)
		}

		if err == nil {
			RendomData.page.FindByXPath(path).Click()
			RendomData.MessageWeb = RendomData.MessageWeb + "ClickButton -> PASS\n"
			a.SetSell(PassFail, NewcurrentRow, RendomData.MessageWeb)
			break
		}
	}

	xpathText, err := RendomData.page.FindByXPath(path).Text()
	if err != nil {
		log.Println("ClickButton - Failed to find Path:", xpathText, err)
		RendomData.MessageWeb = RendomData.MessageWeb + "ClickButton ->  " + path + " кнопка не найдена\n" +
			"\nВыполнение теста остановленно" +
			"\n - Web Test Name --->  " + RendomData.ApiTestName +
			"\n - Api Sheet Name --->  " + RendomData.ApiSheetName +
			"\nCountry --->   " + RendomData.Countryname
		a.Errors(RendomData.MessageWeb)
	}
}

// Fill Field
func (a Data) FillField(path string, text string) {

	for zIdx := 0; zIdx < 60; zIdx++ {
		xpathText := RendomData.page.FindByXPath(path)
		count, _ := xpathText.Count()
		if count == 0 {
			log.Println("FillField - Failed to find Path:", count)
			time.Sleep(1 * time.Second)

		}

		if count >= 1 {
			xpathText.Click()
			xpathText.Fill(text)
			RendomData.MessageWeb = RendomData.MessageWeb + "FillField -> PASS\n"
			a.SetSell(PassFail, NewcurrentRow, RendomData.MessageWeb)
			break
		}
	}

	xpathText2 := RendomData.page.FindByXPath(path)
	count, _ := xpathText2.Count()
	if count == 0 {
		RendomData.MessageWeb = RendomData.MessageWeb + "FillField ->  " + path + " поле не заполнено\n" +
			"\nВыполнение теста остановленно" +
			"\n - Web Test Name --->  " + RendomData.ApiTestName +
			"\n - Api Sheet Name --->  " + RendomData.ApiSheetName +
			"\nCountry --->   " + RendomData.Countryname
		a.Errors(RendomData.MessageWeb)
	}
}

// Properties
func (a Data) Properties(path string, text string) {
	// Find path
	for zIdx := 0; zIdx < 60; zIdx++ {
		xpathText := RendomData.page.FindByXPath(path)
		count, _ := xpathText.Count()
		if count == 0 {
			log.Println("Properties - Failed to find Path:", count)
			time.Sleep(1 * time.Second)
		}

		if count >= 1 {

			if text == "Enabled" {
				props, _ := xpathText.Enabled()
				if props {
					RendomData.MessageWeb = RendomData.MessageWeb + "Properties Enabled -> PASS\n"
					a.SetSell(PassFail, NewcurrentRow, RendomData.MessageWeb)
					break
				} else {
					RendomData.MessageWeb = RendomData.MessageWeb + "Properties ->  " + strconv.FormatBool(props) + " не соответствует свойству Enabled\n" +
						"\nВыполнение теста остановленно" +
						"\n - Web Test Name --->  " + RendomData.ApiTestName +
						"\n - Api Sheet Name --->  " + RendomData.ApiSheetName +
						"\nCountry --->   " + RendomData.Countryname
					a.Errors(RendomData.MessageWeb)
				}
			} // Enabled

			if text == "Disabled" {
				props, _ := xpathText.Enabled()
				if !props {
					RendomData.MessageWeb = RendomData.MessageWeb + "Properties Disabled -> PASS\n"
					a.SetSell(PassFail, NewcurrentRow, RendomData.MessageWeb)
					break
				} else {
					RendomData.MessageWeb = RendomData.MessageWeb + "Properties ->  " + strconv.FormatBool(props) + " не соответствует свойству Disabled\n" +
						"\nВыполнение теста остановленно" +
						"\n - Web Test Name --->  " + RendomData.ApiTestName +
						"\n - Api Sheet Name --->  " + RendomData.ApiSheetName +
						"\nCountry --->   " + RendomData.Countryname
					a.Errors(RendomData.MessageWeb)
				}
			}

			break
		}

	}

	xpathText := RendomData.page.FindByXPath(path)
	count, _ := xpathText.Count()
	if count == 0 {
		RendomData.MessageWeb = RendomData.MessageWeb + "Properties ->  " + string(count) +
			" путь не найден\n" +
			"\nВыполнение теста остановленно" +
			"\n - Web Test Name --->  " + RendomData.ApiTestName +
			"\n - Api Sheet Name --->  " + RendomData.ApiSheetName +
			"\nCountry --->   " + RendomData.Countryname
		a.Errors(RendomData.MessageWeb)
	}
}

// Select
func (a Data) Select(path string, text string) {

	for zIdx := 0; zIdx < 60; zIdx++ {
		xpathText := RendomData.page.FindByXPath(path)
		count, _ := xpathText.Count()
		if count == 0 {
			log.Println("Select - Failed to find Path:", count)
			time.Sleep(1 * time.Second)

		}

		if count >= 1 {

			xpathText.Click()
			xpathText2 := RendomData.page.FindByXPath("//div[@class='select__box active']")
			xpathText2.FindByName(string(text)).Click()
			//time.Sleep(1 * time.Second)

			//xpathText.Fill(text)
			RendomData.MessageWeb = RendomData.MessageWeb + "Select -> PASS\n"
			a.SetSell(PassFail, NewcurrentRow, RendomData.MessageWeb)
			break
		}
	}

	xpathText2 := RendomData.page.FindByXPath(path)
	count, _ := xpathText2.Count()
	if count == 0 {
		RendomData.MessageWeb = RendomData.MessageWeb + "Select ->  " + path + " значение не выбрано из списка\n" +
			"\nВыполнение теста остановленно" +
			"\n - Web Test Name --->  " + RendomData.ApiTestName +
			"\n - Api Sheet Name --->  " + RendomData.ApiSheetName +
			"\nCountry --->   " + RendomData.Countryname
		a.Errors(RendomData.MessageWeb)
	}
}

//

func (a Data) MessageW(message string) {
	a.SetSell(PassFail, NewcurrentRow, message)
	telegram(message)
	RendomData.XL.SaveAs("./Report/" + Filename + "Report.xlsx")
	a.mail(RendomData.Filename) //Send Email

}

func (a Data) Errors(message string) {
	Screen()
	ErrorCount++
	RendomData.errorCount++
	a.errorCount++
	a.MessageW(message)
	CloseWindow()
	os.Exit(0)
}
