package main

import (
	"fmt"
	"strconv"
)

func Reports(data *Starts, t string)  {
	reportUrl := "https://drive.google.com/drive/u/1/folders/" + data.ReportFolderName
	fmt.Println("a.errorCount ->  ", ErrorCount)

	if ErrorCount >=1 {
	msg := "Report for ->  " + Filename + " Country" +
	"\n\nTest contain error\nTotal error:    " + strconv.Itoa(ErrorCount) + "\n" +
	"\nFolder URL:    " + reportUrl + "\n" +
	"\nReport file name:    " + Filename + "_Report_"+ t + ".xlsx" +
	"\nData file name:    " + DataStart.Urlxlsxfile

	telegram(msg)
	}
}


