package main

import (
	"fmt"
	"os"
)

type Starts struct {
	ReportFolderName string
	Urlxlsxfile string
	Country string
}

	var DataStart = new(Starts)

func Start(country string) *Starts {

	//	WU_API_TEST_Doc
	// https://drive.google.com/drive/u/0/folders/1lFEzB1e9c4sXOicRra5ggkGDVdEuGObK

	//	WU_API_TEST_Doc
	// https://drive.google.com/drive/u/0/folders/1lFEzB1e9c4sXOicRra5ggkGDVdEuGObK


	switch {

	case  country == "ukraine_hosted":
		DataStart.Urlxlsxfile = "https://docs.google.com/spreadsheets/d/1KwvG7gvu9711I6ASe3GlavdQxWOcLjHk8VipHSwY8BI/export?format=xlsx"
		SetPropValue("urlxlsxfile",DataStart.Urlxlsxfile)
		DataStart.ReportFolderName = "1FmVA85w-jsd74AcGo7zKKXiaLzRpAWaU"
		SetPropValue("reportfoldername",DataStart.ReportFolderName)
		RendomData.Countryname = "ukraine"
		//admin Agent API
		SetPropValue("clientid","UA713575")
		SetPropValue("secret_key","GPBqjL2Fa4evxT9TMH#n3G0l")

		//admin credential for Kenya Agent API
		SetPropValue("data.country.kenya.clientid","KE932150")
		SetPropValue("data.country.kenya.secret_key","JWorVd7A#Qk@aK936D3ZOmn!")

		//Merchant
		SetPropValue("public_key","vRpwp")
		SetPropValue("secretkey","SvgcnbWOd9dkYRuJ@g5HMkQ2qwq7ih")

		SetPropValue("payment.country_send_code","KE") // код страны получателя
		SetPropValue("payment.amount_get","1")  //сумма отправления



	case  country == "kenya_hosted":
		DataStart.Urlxlsxfile = "https://docs.google.com/spreadsheets/d/1sbgE5RRKkrQbneVgeu0x0efzmbstOBR9IueFbd4cI6U/export?format=xlsx"
		SetPropValue("urlxlsxfile",DataStart.Urlxlsxfile)
		DataStart.ReportFolderName = "1Yl4nz1o6GFMW23IwR3ga3YmsIKBUbnJJ"
		SetPropValue("reportfoldername",DataStart.ReportFolderName)
		RendomData.Countryname = "kenya"
		//admin Agent API
		SetPropValue("clientid","KE932150")
		SetPropValue("secret_key","JWorVd7A#Qk@aK936D3ZOmn!")
		//admin credential for Urkaine  Agent API
		SetPropValue("data.country.ukraine.clientid","UA713575")
		SetPropValue("data.country.ukraine.secret_key","GPBqjL2Fa4evxT9TMH#n3G0l")

		//Merchant
		SetPropValue("public_key","iW0HE")
		SetPropValue("secretkey","BCfPeJcSuk!WBuRZbF#WmQgyil6dz8")

		SetPropValue("payment.country_send_code","UA") // код страны получателя
		SetPropValue("payment.amount_get","118")  //сумма отправления




	case  country == "rwanda":
		DataStart.Urlxlsxfile = "https://docs.google.com/spreadsheets/d/1vEomALrKNbteAE1uEr2iSyenRpDRQt06EVKi53vDiVU/export?format=xlsx"
		SetPropValue("urlxlsxfile",DataStart.Urlxlsxfile)
		DataStart.ReportFolderName = "1weiJWHewtYR4AHp05XFTF7ggrj0xwGcw"
		SetPropValue("reportfoldername",DataStart.ReportFolderName)
		RendomData.Countryname = "rwanda"
		SetPropValue("payment.country_send_code","UA") // код страны получателя
		SetPropValue("payment.amount_get","11")  //сумма отправления

	case  country == "bahrain":
		DataStart.Urlxlsxfile = "https://docs.google.com/spreadsheets/d/1XFLdYIV--f6zBUJ4-R71Bh2CvpPpY96Rtk-9CZw4y-M/export?format=xlsx"
		SetPropValue("urlxlsxfile",DataStart.Urlxlsxfile)
		DataStart.ReportFolderName = "1fZCIHD3LSB4N02aQ-xJbDXOGRP4lb5fO"
		SetPropValue("reportfoldername",DataStart.ReportFolderName)
		RendomData.Countryname = "bahrain"
		SetPropValue("payment.country_send_code","FR") // код страны получателя
		SetPropValue("payment.amount_get","11")  //сумма отправления

	case  country == "albania":
		DataStart.Urlxlsxfile = "https://docs.google.com/spreadsheets/d/1qkSFhXLBPGmDcbWCb938-phmfaoeNO9-iA5BMI4PivM/export?format=xlsx"
		SetPropValue("urlxlsxfile",DataStart.Urlxlsxfile)
		DataStart.ReportFolderName = "1gpYiO3Qii998lI0e_4OXY3G0egXWf4MX"
		SetPropValue("reportfoldername",DataStart.ReportFolderName)
		RendomData.Countryname = "albania"
		SetPropValue("payment.country_send_code","FR") // код страны получателя
		SetPropValue("payment.amount_get","222")  //сумма отправления

	case  country == "aruba":
		DataStart.Urlxlsxfile = "https://docs.google.com/spreadsheets/d/1GjgU86Tybyda0YuWGwaMKI5tAgTkkhkCw4BvSDjWWQ0/export?format=xlsx"
		SetPropValue("urlxlsxfile",DataStart.Urlxlsxfile)
		DataStart.ReportFolderName = "11_-itZPO8uVRM2SBU--opdKt7ksDVf6B"
		SetPropValue("reportfoldername",DataStart.ReportFolderName)
		RendomData.Countryname = "aruba"
		SetPropValue("payment.country_send_code","UA") // код страны получателя
		SetPropValue("payment.amount_get","11")  //сумма отправления

	case  country == "bahamas":
		DataStart.Urlxlsxfile = "https://docs.google.com/spreadsheets/d/1BLsr6KDHKaobhLGRLV80nv9kp7J0WBIssqfHuMuNbBg/export?format=xlsx"
		SetPropValue("urlxlsxfile",DataStart.Urlxlsxfile)
		DataStart.ReportFolderName = "1gER5CgEN1_SwR_oBGVhbAOLTL7q94uBs"
		SetPropValue("reportfoldername",DataStart.ReportFolderName)
		RendomData.Countryname = "bahamas"
		SetPropValue("payment.country_send_code","UA") // код страны получателя
		SetPropValue("payment.amount_get","11")  //сумма отправления

	case  country == "benin":
		DataStart.Urlxlsxfile = "https://docs.google.com/spreadsheets/d/1V7f1JxK0wpoyvQ8VSjV2CuqM1HoZ_2_LJQuf4aLGvPY/export?format=xlsx"
		SetPropValue("urlxlsxfile",DataStart.Urlxlsxfile)
		DataStart.ReportFolderName = "1yHyVddh6FIX2uVOGySsA0Grx4sx2SOHc"
		SetPropValue("reportfoldername",DataStart.ReportFolderName)
		RendomData.Countryname = "benin"
		SetPropValue("payment.country_send_code","FR") // код страны получателя
		SetPropValue("payment.amount_get","11")  //сумма отправления

	case  country == "bvi":
		DataStart.Urlxlsxfile = "https://docs.google.com/spreadsheets/d/15I-rdxjiYbtvLjpa11V5D17Ws0mr4vpa7SaXxYnSETU/export?format=xlsx"
		SetPropValue("urlxlsxfile",DataStart.Urlxlsxfile)
		DataStart.ReportFolderName = "1t_3y1KncUUi6hlVOcVKABOC2tin9lLrw"
		SetPropValue("reportfoldername",DataStart.ReportFolderName)
		RendomData.Countryname = "bvi"
		SetPropValue("payment.country_send_code","FR") // код страны получателя
		SetPropValue("payment.amount_get","11")  //сумма отправления

	case  country == "belarus":
		DataStart.Urlxlsxfile = "https://docs.google.com/spreadsheets/d/1hn4aNeOLUFTzHSBvegIhh9BGWQ1u738PctggYlmSW2c/export?format=xlsx"
		SetPropValue("urlxlsxfile",DataStart.Urlxlsxfile)
		DataStart.ReportFolderName = "1r1U60qBtE6u5zg7PRac_lWxtULlttcPY"
		SetPropValue("reportfoldername",DataStart.ReportFolderName)
		RendomData.Countryname = "belarus"
		SetPropValue("payment.country_send_code","UA") // код страны получателя
		SetPropValue("payment.amount_get","11")  //сумма отправления

	case  country == "brunei":
		DataStart.Urlxlsxfile = "https://docs.google.com/spreadsheets/d/1owGSEb2gfhsjGHlEHOw1lkHRhsbFxC6uB4FOz0OzKso/export?format=xlsx"
		SetPropValue("urlxlsxfile",DataStart.Urlxlsxfile)
		DataStart.ReportFolderName = "1uXFXD3wH4jOW5flvcECfYjfUkLdnO3LL"
		SetPropValue("reportfoldername",DataStart.ReportFolderName)
		RendomData.Countryname = "brunei"
		SetPropValue("payment.country_send_code","FR") // код страны получателя
		SetPropValue("payment.amount_get","11")  //сумма отправления

	case  country == "burkinafaso":
		DataStart.Urlxlsxfile = "https://docs.google.com/spreadsheets/d/1pHmlSyc--hd2ofqxgisfIn0qL3k0NWL8VtySOCN1LHA/export?format=xlsx"
		SetPropValue("urlxlsxfile",DataStart.Urlxlsxfile)
		DataStart.ReportFolderName = "1JDOr8X5f5gSkCK2GvGDCFMJmqVmC9Z9f"
		SetPropValue("reportfoldername",DataStart.ReportFolderName)
		RendomData.Countryname = "burkinafaso"
		SetPropValue("payment.country_send_code","FR") // код страны получателя
		SetPropValue("payment.amount_get","11")  //сумма отправления

	case  country == "cameroun":
		DataStart.Urlxlsxfile = "https://docs.google.com/spreadsheets/d/1c0FLyZLwxQZJfjj_mBG1q_XbSgA6IGEABEYLVQ1bttA/export?format=xlsx"
		SetPropValue("urlxlsxfile",DataStart.Urlxlsxfile)
		DataStart.ReportFolderName = "1HPFLOa-dtEhlaMU1g42LCMWVZST9fwSm"
		SetPropValue("reportfoldername",DataStart.ReportFolderName)
		RendomData.Countryname = "cameroun"
		SetPropValue("payment.country_send_code","UA") // код страны получателя
		SetPropValue("payment.amount_get","11")  //сумма отправления

	case  country == "caymanislands":
		DataStart.Urlxlsxfile = "https://docs.google.com/spreadsheets/d/1RlFh-tKTyqGMLE6LIokw4tcA_r-3slCvXl-TfqctRH0/export?format=xlsx"
		SetPropValue("urlxlsxfile",DataStart.Urlxlsxfile)
		DataStart.ReportFolderName = "1PTBKdgk0RSkYfYCoiMqxkJm3zbYQXTn6"
		SetPropValue("reportfoldername",DataStart.ReportFolderName)
		RendomData.Countryname = "caymanislands"
		SetPropValue("payment.country_send_code","FR") // код страны получателя
		SetPropValue("payment.amount_get","11")  //сумма отправления

	case  country == "cookislands":
		DataStart.Urlxlsxfile = "https://docs.google.com/spreadsheets/d/1CrmcX0vYPoXlGnnsTZSApqWJaUk2ju45WUKWJUwRU5Q/export?format=xlsx"
		SetPropValue("urlxlsxfile",DataStart.Urlxlsxfile)
		DataStart.ReportFolderName = "1NdGB8dNcwr5NR1PHkLcSjSqO46Oh27C3"
		SetPropValue("reportfoldername",DataStart.ReportFolderName)
		RendomData.Countryname = "cookislands"
		SetPropValue("payment.country_send_code","FR") // код страны получателя
		SetPropValue("payment.amount_get","11")  //сумма отправления

	case  country == "cotedivoire":
		DataStart.Urlxlsxfile = "https://docs.google.com/spreadsheets/d/1gojRK7Qa_Fpqe-uYQ2Hcyq5zy7Qrzsjnv7IWTCrk3R4/export?format=xlsx"
		SetPropValue("urlxlsxfile",DataStart.Urlxlsxfile)
		DataStart.ReportFolderName = "1_TP3uSWeAZPgrKfvAj5kZVhd9Z1xQQtp"
		SetPropValue("reportfoldername",DataStart.ReportFolderName)
		RendomData.Countryname = "cotedivoire"
		SetPropValue("payment.country_send_code","FR") // код страны получателя
		SetPropValue("payment.amount_get","11")  //сумма отправления

	case  country == "dominicanrepublic":
		DataStart.Urlxlsxfile = "https://docs.google.com/spreadsheets/d/1Bl1yNn1M_avb2LHfXIGlqz3Z6GB_AmjGvcK7x-tlc0w/export?format=xlsx"
		SetPropValue("urlxlsxfile",DataStart.Urlxlsxfile)
		DataStart.ReportFolderName = "1tFXvxC1EHxWTpvkBbI3S_tU1b91s09NZ"
		SetPropValue("reportfoldername",DataStart.ReportFolderName)
		RendomData.Countryname = "dominicanrepublic"
		SetPropValue("payment.country_send_code","FR") // код страны получателя
		SetPropValue("payment.amount_get","11")  //сумма отправления

	case  country == "rdc":
		DataStart.Urlxlsxfile = "https://docs.google.com/spreadsheets/d/1A9g0_GcqU4ZE0RNcAnNTtmrUtsv_mjBOTplHhtrsCCU/export?format=xlsx"
		SetPropValue("urlxlsxfile",DataStart.Urlxlsxfile)
		DataStart.ReportFolderName = "10k5ePKFjQLQfWA8rUDadqYc9S0D2jz-V"
		SetPropValue("reportfoldername",DataStart.ReportFolderName)
		RendomData.Countryname = "rdc"
		SetPropValue("payment.country_send_code","UA") // код страны получателя
		SetPropValue("payment.amount_get","11")  //сумма отправления

	case  country == "easttimor":
		DataStart.Urlxlsxfile = "https://docs.google.com/spreadsheets/d/1CM15WLrkmrTPbRdMnE9PygAMxcJhQxOmxXHZ6zV3Hho/export?format=xlsx"
		SetPropValue("urlxlsxfile",DataStart.Urlxlsxfile)
		DataStart.ReportFolderName = "1UtpPcvBC-On0Aqo7oZc05BZw3tLz2dij"
		SetPropValue("reportfoldername",DataStart.ReportFolderName)
		RendomData.Countryname = "easttimor"
		SetPropValue("payment.country_send_code","FR") // код страны получателя
		SetPropValue("payment.amount_get","11")  //сумма отправления

	case  country == "ecuador":
		DataStart.Urlxlsxfile = "https://docs.google.com/spreadsheets/d/1hO9f8cygcspIAFiz8gQL_E6Ihe1-D8Zof7ZUYWJ8F6s/export?format=xlsx"
		SetPropValue("urlxlsxfile",DataStart.Urlxlsxfile)
		DataStart.ReportFolderName = "19O0uFGXalYkBsdSpnK9vjUNL-WQN5qFA"
		SetPropValue("reportfoldername",DataStart.ReportFolderName)
		RendomData.Countryname = "ecuador"
		SetPropValue("payment.country_send_code","FR") // код страны получателя
		SetPropValue("payment.amount_get","11")  //сумма отправления

	case  country == "fiji":
		DataStart.Urlxlsxfile = "https://docs.google.com/spreadsheets/d/1uT7oPnMHccWiMALVWAclxwgH3r-OIKVO7Lb6jxg-Cbk/export?format=xlsx"
		SetPropValue("urlxlsxfile",DataStart.Urlxlsxfile)
		DataStart.ReportFolderName = "1VnRO-mE2Ot4-X0owm6_NThb9IR6HMshH"
		SetPropValue("reportfoldername",DataStart.ReportFolderName)
		RendomData.Countryname = "fiji"
		SetPropValue("payment.country_send_code","UA") // код страны получателя
		SetPropValue("payment.amount_get","11")  //сумма отправления

	case  country == "gabon":
		DataStart.Urlxlsxfile = "https://docs.google.com/spreadsheets/d/1lgcHzY_BAMFeuHMkYaNjJ7HaL0GrGuzYAA2Sv1zRo18/export?format=xlsx"
		SetPropValue("urlxlsxfile",DataStart.Urlxlsxfile)
		DataStart.ReportFolderName = "1nf5iJi8_Tzo8E7ZAqDr-cZk9pw7FDWhd"
		SetPropValue("reportfoldername",DataStart.ReportFolderName)
		RendomData.Countryname = "gabon"
		SetPropValue("payment.country_send_code","FR") // код страны получателя
		SetPropValue("payment.amount_get","11")  //сумма отправления

	case  country == "georgia":
		DataStart.Urlxlsxfile = "https://docs.google.com/spreadsheets/d/1ef0skB1IB6E5gSx2ps_Bxs1x10pjSNsU1iYyxkNXy60/export?format=xlsx"
		SetPropValue("urlxlsxfile",DataStart.Urlxlsxfile)
		DataStart.ReportFolderName = "1KXzWJGDUGANZKJESJnT9riKnWlYA69e5"
		SetPropValue("reportfoldername",DataStart.ReportFolderName)
		RendomData.Countryname = "georgia"
		SetPropValue("payment.country_send_code","FR") // код страны получателя
		SetPropValue("payment.amount_get","11")  //сумма отправления

	case  country == "guinea":
		DataStart.Urlxlsxfile = "https://docs.google.com/spreadsheets/d/1U0J-gKUb_m62e1UUJtRpMWXZ2xsJhxQrVCt1Vlc7Kts/export?format=xlsx"
		SetPropValue("urlxlsxfile",DataStart.Urlxlsxfile)
		DataStart.ReportFolderName = "1zy_OY48GQwpUou8hz-XjYEMDeNfxW8b2"
		SetPropValue("reportfoldername",DataStart.ReportFolderName)
		RendomData.Countryname = "guinea"
		SetPropValue("payment.country_send_code","UA") // код страны получателя
		SetPropValue("payment.amount_get","11")  //сумма отправления

	case  country == "guineabissau":
		DataStart.Urlxlsxfile = "https://docs.google.com/spreadsheets/d/1TC6IzP0SZTt52mBXOe6TL-qy04MP2RjOeEpiRhw09Q8/export?format=xlsx"
		SetPropValue("urlxlsxfile",DataStart.Urlxlsxfile)
		DataStart.ReportFolderName = "16BqLWB6NQtEC6hTV3_DjXEXqvYqD9vBi"
		SetPropValue("reportfoldername",DataStart.ReportFolderName)
		RendomData.Countryname = "guineabissau"
		SetPropValue("payment.country_send_code","FR") // код страны получателя
		SetPropValue("payment.amount_get","11")  //сумма отправления

	case  country == "haiti":
		DataStart.Urlxlsxfile = "https://docs.google.com/spreadsheets/d/1hPqDGiVPJ1PZkeRFTDoN2oXO5l6UktUKvIrrWc0yKlo/export?format=xlsx"
		SetPropValue("urlxlsxfile",DataStart.Urlxlsxfile)
		DataStart.ReportFolderName = "1DhT6hsgzUaCB7S2jgjmZUrxeOQGDXz85"
		SetPropValue("reportfoldername",DataStart.ReportFolderName)
		RendomData.Countryname = "haiti"
		SetPropValue("payment.country_send_code","UA") // код страны получателя
		SetPropValue("payment.amount_get","11")  //сумма отправления


	case  country == "jordan":
		DataStart.Urlxlsxfile = "https://docs.google.com/spreadsheets/d/1Nv6-XsOcvhARdy2pnI3--RhwR14qOjzZt6uhDA62O_Q/export?format=xlsx"
		SetPropValue("urlxlsxfile",DataStart.Urlxlsxfile)
		DataStart.ReportFolderName = "1Qqy7XQlpYCY16UtwkMwqmAFwXugrCfHu"
		SetPropValue("reportfoldername",DataStart.ReportFolderName)
		RendomData.Countryname = "jordan"
		SetPropValue("payment.country_send_code","FR") // код страны получателя
		SetPropValue("payment.amount_get","11")  //сумма отправления



	case  country == "kazakhstan":
		DataStart.Urlxlsxfile = "https://docs.google.com/spreadsheets/d/1omOoS6avzbnyVcw5qT921evEPZmYcK_Y4oCIpZzUzbo/export?format=xlsx"
		SetPropValue("urlxlsxfile",DataStart.Urlxlsxfile)
		DataStart.ReportFolderName = "1y5fzXTvJT4pbRRU85FmbSq2GgxpAArDb"
		SetPropValue("reportfoldername",DataStart.ReportFolderName)
		RendomData.Countryname = "kazakhstan"
		SetPropValue("payment.country_send_code","UA") // код страны получателя
		SetPropValue("payment.amount_get","11")  //сумма отправления

	case  country == "kuwait":
		DataStart.Urlxlsxfile = "https://docs.google.com/spreadsheets/d/1RlnjiQKLZJvKYst1VDFjNskYMrvzakpjmlSEECnmhgY/export?format=xlsx"
		SetPropValue("urlxlsxfile",DataStart.Urlxlsxfile)
		DataStart.ReportFolderName = "1lxVUhfZu5FokRY_y5PWNTkzSztBjQ_m6"
		SetPropValue("reportfoldername",DataStart.ReportFolderName)
		RendomData.Countryname = "kuwait"
		SetPropValue("payment.country_send_code","FR") // код страны получателя
		SetPropValue("payment.amount_get","11")  //сумма отправления

	case  country == "lebanon":
		DataStart.Urlxlsxfile = "https://docs.google.com/spreadsheets/d/1DQAQxgdegaDPQ-X46qbS6m9nh3zldXaSB7vKviSIf1s/export?format=xlsx"
		SetPropValue("urlxlsxfile",DataStart.Urlxlsxfile)
		DataStart.ReportFolderName = "1gSOUGsaNxOSN_hLDpbKTezJArvIamhm8"
		SetPropValue("reportfoldername",DataStart.ReportFolderName)
		RendomData.Countryname = "lebanon"
		SetPropValue("payment.country_send_code","FR") // код страны получателя
		SetPropValue("payment.amount_get","11")  //сумма отправления

	case  country == "maldives":
		DataStart.Urlxlsxfile = "https://docs.google.com/spreadsheets/d/1LAdYuF8UpT3JIbzNy72dD2tWPfeC-plgqaprWdKjwqk/export?format=xlsx"
		SetPropValue("urlxlsxfile",DataStart.Urlxlsxfile)
		DataStart.ReportFolderName = "1SV-YGTTrHQbeqUsAMXNUZYuh-Qd8fbMC"
		SetPropValue("reportfoldername",DataStart.ReportFolderName)
		RendomData.Countryname = "maldives"
		SetPropValue("payment.country_send_code","FR") // код страны получателя
		SetPropValue("payment.amount_get","11")  //сумма отправления

	case  country == "mali":
		DataStart.Urlxlsxfile = "https://docs.google.com/spreadsheets/d/1L4zS2ak4nipOSuFT6t-zR6fmzw63GVzbN57V3bNaMs4/export?format=xlsx"
		SetPropValue("urlxlsxfile",DataStart.Urlxlsxfile)
		DataStart.ReportFolderName = "1-0s4mlWTg9RjwjcbSwnSZUrEEBuSPfIQ"
		SetPropValue("reportfoldername",DataStart.ReportFolderName)
		RendomData.Countryname = "mali"
		SetPropValue("payment.country_send_code","FR") // код страны получателя
		SetPropValue("payment.amount_get","11")  //сумма отправления

	case  country == "marshallislands":
		DataStart.Urlxlsxfile = "https://docs.google.com/spreadsheets/d/1uR6sKrprqN4d_4abZ4c39x9aWBummv8rHpan6k6UKDY/export?format=xlsx"
		SetPropValue("urlxlsxfile",DataStart.Urlxlsxfile)
		DataStart.ReportFolderName = "1I4ARThmr6JNoL8foSA_7oMorm7Ucl4bB"
		SetPropValue("reportfoldername",DataStart.ReportFolderName)
		RendomData.Countryname = "marshallislands"
		SetPropValue("payment.country_send_code","FR") // код страны получателя
		SetPropValue("payment.amount_get","11")  //сумма отправления

	case  country == "moldova":
		DataStart.Urlxlsxfile = "https://docs.google.com/spreadsheets/d/1okwvkbrOQHC6cVhQQKb9Fgw-mjBgRVx5WQsvJcdvsyE/export?format=xlsx"
		SetPropValue("urlxlsxfile",DataStart.Urlxlsxfile)
		DataStart.ReportFolderName = "1P_2AA2NNJDzDWn4y_jJUf-hWGwLN_UIv"
		SetPropValue("reportfoldername",DataStart.ReportFolderName)
		RendomData.Countryname = "moldova"
		SetPropValue("payment.country_send_code","UA") // код страны получателя
		SetPropValue("payment.amount_get","11")  //сумма отправления

	case  country == "montenegro":
		DataStart.Urlxlsxfile = "https://docs.google.com/spreadsheets/d/1q-_JCUdUeOBkHPU-QmZJgyNadQ9wS-MoquxEtXQ9hlY/export?format=xlsx"
		SetPropValue("urlxlsxfile",DataStart.Urlxlsxfile)
		DataStart.ReportFolderName = "1w7AwhmIomv-20IwCNmXLduKH0OBmSR1G"
		SetPropValue("reportfoldername",DataStart.ReportFolderName)
		RendomData.Countryname = "montenegro"
		SetPropValue("payment.country_send_code","FR") // код страны получателя
		SetPropValue("payment.amount_get","11")  //сумма отправления

	case  country == "niger":
		DataStart.Urlxlsxfile = "https://docs.google.com/spreadsheets/d/1piHHgT4kxKxQQYlVF4lUAe9R0ksip6e-gyjPHDdBUNo/export?format=xlsx"
		SetPropValue("urlxlsxfile",DataStart.Urlxlsxfile)
		DataStart.ReportFolderName = "17p_Bg96JA5wyIsGsikUCboSeOHVRKdRW"
		SetPropValue("reportfoldername",DataStart.ReportFolderName)
		RendomData.Countryname = "niger"
		SetPropValue("payment.country_send_code","FR") // код страны получателя
		SetPropValue("payment.amount_get","11")  //сумма отправления

	case  country == "nigeria":
		DataStart.Urlxlsxfile = "https://docs.google.com/spreadsheets/d/1y-J4Ft5Zewz59HyInxDyUIDRl_lPMH4Fq3vTljqFZYk/export?format=xlsx"
		SetPropValue("urlxlsxfile",DataStart.Urlxlsxfile)
		DataStart.ReportFolderName = "1735Q4LGpk5ePmFfRgTgULzwnux10PRI1"
		SetPropValue("reportfoldername",DataStart.ReportFolderName)
		RendomData.Countryname = "nigeria"
		SetPropValue("payment.country_send_code","FR") // код страны получателя
		SetPropValue("payment.amount_get","11")  //сумма отправления

	case  country == "oman":
		DataStart.Urlxlsxfile = "https://docs.google.com/spreadsheets/d/1Wzx2J_dMkoJw45pBcDPEhKoZzkM40NMW_InK1BworLA/export?format=xlsx"
		SetPropValue("urlxlsxfile",DataStart.Urlxlsxfile)
		DataStart.ReportFolderName = "1flG3ar91RUXNhQZPwCAZR7SVySKOFWqX"
		SetPropValue("reportfoldername",DataStart.ReportFolderName)
		RendomData.Countryname = "oman"
		SetPropValue("payment.country_send_code","FR") // код страны получателя
		SetPropValue("payment.amount_get","11")  //сумма отправления

	case  country == "paraguay":
		DataStart.Urlxlsxfile = "https://docs.google.com/spreadsheets/d/1Y1BV8t2ywMkT-AvyfdokCaJEqM92Vx-vdN-eXtpAimE/export?format=xlsx"
		SetPropValue("urlxlsxfile",DataStart.Urlxlsxfile)
		DataStart.ReportFolderName = "1II7gQeY0F6DujwJHFTetkqlAuCam7wc_"
		SetPropValue("reportfoldername",DataStart.ReportFolderName)
		RendomData.Countryname = "paraguay"
		SetPropValue("payment.country_send_code","UA") // код страны получателя
		SetPropValue("payment.amount_get","11")  //сумма отправления

	case  country == "philippines":
		DataStart.Urlxlsxfile = "https://docs.google.com/spreadsheets/d/1V-8SFL9IhzOUwDvYWB0FhjsoAg1lBHj37vCoGGrKMjo/export?format=xlsx"
		SetPropValue("urlxlsxfile",DataStart.Urlxlsxfile)
		DataStart.ReportFolderName = "1s-oi4wzkZlQnuNU7CSlzG-O7kByH1wn0"
		SetPropValue("reportfoldername",DataStart.ReportFolderName)
		RendomData.Countryname = "philippines"
		SetPropValue("payment.country_send_code","FR") // код страны получателя
		SetPropValue("payment.amount_get","11")  //сумма отправления

	case  country == "qatar":
		DataStart.Urlxlsxfile = "https://docs.google.com/spreadsheets/d/1CXSl3WJYoJr8lgbmaiKVSAjJnUuizex6UIXF810u_5E/export?format=xlsx"
		SetPropValue("urlxlsxfile",DataStart.Urlxlsxfile)
		DataStart.ReportFolderName = "1yMySdfkuBGXUOG9Pl0sUGqeQ5Rvploz1"
		SetPropValue("reportfoldername",DataStart.ReportFolderName)
		RendomData.Countryname = "qatar"
		SetPropValue("payment.country_send_code","FR") // код страны получателя
		SetPropValue("payment.amount_get","11")  //сумма отправления

	case  country == "samoa":
		DataStart.Urlxlsxfile = "https://docs.google.com/spreadsheets/d/1L-l0SbbNecRN28n96fMqhQdafFNsBjD5R-LbcDhvIvA/export?format=xlsx"
		SetPropValue("urlxlsxfile",DataStart.Urlxlsxfile)
		DataStart.ReportFolderName = "1eLLsbqi0XhZx3905XG8eVa_Dw4ff6Xjg"
		SetPropValue("reportfoldername",DataStart.ReportFolderName)
		RendomData.Countryname = "samoa"
		SetPropValue("payment.country_send_code","FR") // код страны получателя
		SetPropValue("payment.amount_get","11")  //сумма отправления

	case  country == "saudi":
		DataStart.Urlxlsxfile = "https://docs.google.com/spreadsheets/d/1Hu4z-5E-uew4xNQ92LoNGtd7pFDL2WfC3haGohomKAs/export?format=xlsx"
		SetPropValue("urlxlsxfile",DataStart.Urlxlsxfile)
		DataStart.ReportFolderName = "1BjBk_lepK9YFNuZYkocalaLOg3S2E76p"
		SetPropValue("reportfoldername",DataStart.ReportFolderName)
		RendomData.Countryname = "saudi"
		SetPropValue("payment.country_send_code","UA") // код страны получателя
		SetPropValue("payment.amount_get","11")  //сумма отправления

	case  country == "senegal":
		DataStart.Urlxlsxfile = "https://docs.google.com/spreadsheets/d/1XxRXJDKDWFlQkGmcoxmeF-biNMJsB5YSZeboB71icyg/export?format=xlsx"
		SetPropValue("urlxlsxfile",DataStart.Urlxlsxfile)
		DataStart.ReportFolderName = "1bGxhc0PwZy4V1ORjOpjx8QidyrL_PSBI"
		SetPropValue("reportfoldername",DataStart.ReportFolderName)
		RendomData.Countryname = "senegal"
		SetPropValue("payment.country_send_code","FR") // код страны получателя
		SetPropValue("payment.amount_get","11")  //сумма отправления

	case  country == "serbia":
		DataStart.Urlxlsxfile = "https://docs.google.com/spreadsheets/d/1JFXhWZRdirjeX8jB0_xV8fiyAHB1PL22BkTdu939yNs/export?format=xlsx"
		SetPropValue("urlxlsxfile",DataStart.Urlxlsxfile)
		DataStart.ReportFolderName = "1qy14zSJZzl8iGIr7T0dkevZndxO21yia"
		SetPropValue("reportfoldername",DataStart.ReportFolderName)
		RendomData.Countryname = "serbia"
		SetPropValue("payment.country_send_code","FR") // код страны получателя
		SetPropValue("payment.amount_get","11")  //сумма отправления

	case  country == "solomonislands":
		DataStart.Urlxlsxfile = "https://docs.google.com/spreadsheets/d/1oFI03FmhXic-x01tzxQ3ELGgwU_lrmffm_ByyCvcczk/export?format=xlsx"
		SetPropValue("urlxlsxfile",DataStart.Urlxlsxfile)
		DataStart.ReportFolderName = "1euAcruWUI_BItBhvybDNv7C3rosANyhm"
		SetPropValue("reportfoldername",DataStart.ReportFolderName)
		RendomData.Countryname = "solomonislands"
		SetPropValue("payment.country_send_code","FR") // код страны получателя
		SetPropValue("payment.amount_get","11")  //сумма отправления

	case  country == "srilanka":
		DataStart.Urlxlsxfile = "https://docs.google.com/spreadsheets/d/1o98cqLV8ClBsxQ86MPdeyi8mOKTpMjQI4XXbD6VTICE/export?format=xlsx"
		SetPropValue("urlxlsxfile",DataStart.Urlxlsxfile)
		DataStart.ReportFolderName = "1znDg-IIJ2OpYewt2fhLMyIx8Z7kYs9YG"
		SetPropValue("reportfoldername",DataStart.ReportFolderName)
		RendomData.Countryname = "srilanka"
		SetPropValue("payment.country_send_code","FR") // код страны получателя
		SetPropValue("payment.amount_get","11")  //сумма отправления

	case  country == "stkittsandnevis":
		DataStart.Urlxlsxfile = "https://docs.google.com/spreadsheets/d/1HjVdjmxbcPOoQy4UZgWb2W15Z8FXpXaH-GCIkQrdE4Y/export?format=xlsx"
		SetPropValue("urlxlsxfile",DataStart.Urlxlsxfile)
		DataStart.ReportFolderName = "1B4uxMYwOyaR_6yGm0bmsz3svb_s0Hj3k"
		SetPropValue("reportfoldername",DataStart.ReportFolderName)
		RendomData.Countryname = "stkittsandnevis"
		SetPropValue("payment.country_send_code","FR") // код страны получателя
		SetPropValue("payment.amount_get","11")  //сумма отправления

	case  country == "stlucia":
		DataStart.Urlxlsxfile = "https://docs.google.com/spreadsheets/d/1v38KXdl3oEZR1sznpTpZ5x4qkDVF6biYr4DbkK7hOLU/export?format=xlsx"
		SetPropValue("urlxlsxfile",DataStart.Urlxlsxfile)
		DataStart.ReportFolderName = "1e9lHRvm48L_5bfHMyAdmesNYxFenrecM"
		SetPropValue("reportfoldername",DataStart.ReportFolderName)
		RendomData.Countryname = "stlucia"
		SetPropValue("payment.country_send_code","FR") // код страны получателя
		SetPropValue("payment.amount_get","11")  //сумма отправления

	case  country == "stvincentandthegrenadines":
		DataStart.Urlxlsxfile = "https://docs.google.com/spreadsheets/d/1ifGY4CKgFciyWeueeIrYvYZjjc2sg_9JOPCE6bTv0vU/export?format=xlsx"
		SetPropValue("urlxlsxfile",DataStart.Urlxlsxfile)
		DataStart.ReportFolderName = "1CU5CAXeU1WeGO3_wy9mK3FXPBDdTp3zB"
		SetPropValue("reportfoldername",DataStart.ReportFolderName)
		RendomData.Countryname = "stvincentandthegrenadines"
		SetPropValue("payment.country_send_code","FR") // код страны получателя
		SetPropValue("payment.amount_get","11")  //сумма отправления

	case  country == "thailand":
		DataStart.Urlxlsxfile = "https://docs.google.com/spreadsheets/d/11utSGeVeqmHtC-gMolVMi0QFRJq4GTHGyl4-YRQx3H4/export?format=xlsx"
		SetPropValue("urlxlsxfile",DataStart.Urlxlsxfile)
		DataStart.ReportFolderName = "1BNPphnq4EvQIDUJ3JL0y4B6O1tsdI5Mi"
		SetPropValue("reportfoldername",DataStart.ReportFolderName)
		RendomData.Countryname = "thailand"
		SetPropValue("payment.country_send_code","FR") // код страны получателя
		SetPropValue("payment.amount_get","11")  //сумма отправления


	case  country == "togo":
		DataStart.Urlxlsxfile = "https://docs.google.com/spreadsheets/d/1k38N9ms-_fyjyz7TseLJDHRM_0YDi1vOdV7GkVscZKs/export?format=xlsx"
		SetPropValue("urlxlsxfile",DataStart.Urlxlsxfile)
		DataStart.ReportFolderName = "11l1Otoze-yBckYtIcdQiciQrGmmkit46"
		SetPropValue("reportfoldername",DataStart.ReportFolderName)
		RendomData.Countryname = "togo"
		SetPropValue("payment.country_send_code","FR") // код страны получателя
		SetPropValue("payment.amount_get","11")  //сумма отправления

	case  country == "tonga":
		DataStart.Urlxlsxfile = "https://docs.google.com/spreadsheets/d/18NQ7hHZ4i6Ntca1E2M17LmhwNUCSkXWAJT8jt_nP9Og/export?format=xlsx"
		SetPropValue("urlxlsxfile",DataStart.Urlxlsxfile)
		DataStart.ReportFolderName = "1N3vE23vwE4DY6umCx2nPQ3iGnNsP4tLt"
		SetPropValue("reportfoldername",DataStart.ReportFolderName)
		RendomData.Countryname = "tonga"
		SetPropValue("payment.country_send_code","FR") // код страны получателя
		SetPropValue("payment.amount_get","11")  //сумма отправления

	case  country == "tt":
		DataStart.Urlxlsxfile = "https://docs.google.com/spreadsheets/d/1rZZf7gmcxUXgskF9fKl0Hs38G9-xcmTUrIfMqbU26BI/export?format=xlsx"
		SetPropValue("urlxlsxfile",DataStart.Urlxlsxfile)
		DataStart.ReportFolderName = "1jNk7gBryk3rVevCVQuXVzDx7_fR5jPQo"
		SetPropValue("reportfoldername",DataStart.ReportFolderName)
		RendomData.Countryname = "tt"
		SetPropValue("payment.country_send_code","FR") // код страны получателя
		SetPropValue("payment.amount_get","11")  //сумма отправления

	case  country == "turksandcaicos":
		DataStart.Urlxlsxfile = "https://docs.google.com/spreadsheets/d/1CBb7ZR3y8qhme2bb4QW0DeVHvJTdn-vAcFec_0u1Jc8/export?format=xlsx"
		SetPropValue("urlxlsxfile",DataStart.Urlxlsxfile)
		DataStart.ReportFolderName = "1K_OXdc7vq7E2bkEcdLChc_GZRE_4Szes"
		SetPropValue("reportfoldername",DataStart.ReportFolderName)
		RendomData.Countryname = "turksandcaicos"
		SetPropValue("payment.country_send_code","FR") // код страны получателя
		SetPropValue("payment.amount_get","11")  //сумма отправления

	case  country == "uae":
		DataStart.Urlxlsxfile = "https://docs.google.com/spreadsheets/d/1sdusf7fDxY0NxkBOQZcWzKOTVlDM9YV8L9HKgQM1D9U/export?format=xlsx"
		SetPropValue("urlxlsxfile",DataStart.Urlxlsxfile)
		DataStart.ReportFolderName = "1QAxSGEWwhPYeIwbFPBzynxLUEhtovkad"
		SetPropValue("reportfoldername",DataStart.ReportFolderName)
		RendomData.Countryname = "uae"
		SetPropValue("payment.country_send_code","FR") // код страны получателя
		SetPropValue("payment.amount_get","11")  //сумма отправления

	case  country == "uganda":
		DataStart.Urlxlsxfile = "https://docs.google.com/spreadsheets/d/1eDH1jZQDu6SUdFL34a_eVeyO6HGR04DqbcHbbLW-r8A/export?format=xlsx"
		SetPropValue("urlxlsxfile",DataStart.Urlxlsxfile)
		DataStart.ReportFolderName = "1mYz1k-bqYHu208FVVK-q9o2VyS2o5coT"
		SetPropValue("reportfoldername",DataStart.ReportFolderName)
		RendomData.Countryname = "uganda"
		SetPropValue("payment.country_send_code","UA") // код страны получателя
		SetPropValue("payment.amount_get","11")  //сумма отправления

	case  country == "uruguay":
		DataStart.Urlxlsxfile = "https://docs.google.com/spreadsheets/d/1nPCqwBNRHGol95w8vjTcZZ4q8podXaw0lLI70WwhjAQ/export?format=xlsx"
		SetPropValue("urlxlsxfile",DataStart.Urlxlsxfile)
		DataStart.ReportFolderName = "1lh_dTPUYXN6F88kI8wRO9vGS8LiPxCPj"
		SetPropValue("reportfoldername",DataStart.ReportFolderName)
		RendomData.Countryname = "uruguay"
		SetPropValue("payment.country_send_code","FR") // код страны получателя
		SetPropValue("payment.amount_get","11")  //сумма отправления


	case  country == "vanuatu":
		DataStart.Urlxlsxfile = "https://docs.google.com/spreadsheets/d/1FR21InIwCUAwNgweqjpqsUf59Rr8uJwkWuJumN3Qh8E/export?format=xlsx"
		SetPropValue("urlxlsxfile",DataStart.Urlxlsxfile)
		DataStart.ReportFolderName = "1tZKRFvVHSRIP_9U6g3RtPHCnkAei0pSh"
		SetPropValue("reportfoldername",DataStart.ReportFolderName)
		RendomData.Countryname = "vanuatu"
		SetPropValue("payment.country_send_code","FR") // код страны получателя
		SetPropValue("payment.amount_get","11")  //сумма отправления


	default:
		DataStart.Urlxlsxfile = "Not Find Country\n Test stoped"
		fmt.Println("Not Find Country XLSX file\n Test stoped\n Please see Start.go")
		SetPropValue("urlxlsxfile",DataStart.Urlxlsxfile)
		os.Exit(0)
	}

	return DataStart
}

