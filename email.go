package main

import (
	"github.com/jordan-wright/email"
	"net/smtp"
)

func mail()  {
	e := email.NewEmail()
	e.AttachFile("rwandaReport.xlsx")
	e.From = "Aleks < alex.mywu.uae@gmail.com>"
	e.To = []string{"natalka57m@gmail.com"}
	//e.Bcc = []string{"test_bcc@example.com"}
	//e.Cc = []string{"test_cc@example.com"}
	e.Subject = "Report"
	e.Text = []byte("Text Body is, of course, supported!")
	e.HTML = []byte("<h1>Report API testing!</h1>")
	e.Send("smtp.gmail.com:587", smtp.PlainAuth("", " alex.mywu.uae@gmail.com", "zxczxc12", "smtp.gmail.com"))

}