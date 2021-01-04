package main

import (
	"github.com/jordan-wright/email"
	"net/smtp"
)

var passFail string

func (a Data) mail(countryname string) {
	e := email.NewEmail()
	e.AttachFile("./Report/" + Filename + "Report.xlsx")
	e.AttachFile("./logfile.log")
	e.From = "Aleks < alex.mywu.uae@gmail.com>"
	e.To = []string{"natalka57m@gmail.com"}
	//e.Bcc = []string{"test_bcc@example.com"}
	//e.Cc = []string{"test_cc@example.com"}

	switch {
	case ErrorCount >= 1 || RendomData.errorCount >= 1:
		passFail = "FAILL"
		e.AttachFile("./error.png")

	default:
		passFail = "PASS"
	}

	e.Subject = "Report " + countryname + " - " + passFail
	e.Text = []byte("Text Body is, of course, supported!")
	e.HTML = []byte("<h1>Report Web/API testing!</h1>")
	e.Send("smtp.gmail.com:587", smtp.PlainAuth("", " alex.mywu.uae@gmail.com", "zxczxc12", "smtp.gmail.com"))

}
