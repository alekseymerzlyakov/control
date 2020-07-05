package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

var  (
	Value string
	OldValue string
)

func (a Data) Replace(value string) (string)  {

	Value := value
	re := regexp.MustCompile(`([${{])(.*?){*}([}}])`)
	submatchall := re.FindAllString(value, -1)

	for _, element := range submatchall {
		if strings.Contains(element, "${{") && strings.Contains(element, "}}") {
			element = strings.Trim(element, "${{")
			element = strings.Trim(element, "}}")
			OldValue := "${{" + element + "}}"

			switch {
			case strings.Contains(element, "udid"):
				//fmt.Println("Такой переменной нет в файле data.json ->  ", element)
				if GetPropValue(element) == "" {
					fmt.Println("Такой переменной нет в файле data.json ->  ", element)
					fmt.Println("Выполнение остановленно - Смотреть Replase.go", element)

					msg := "Такой переменной нет в файле data.json ->  " +element
					msg = msg + "\nВыполнение остановленно - Смотреть Replase.go\n"
					msg = msg + a.Response.URL
						telegram(msg)
					os.Exit(0)
				}
				Value = strings.Replace(Value, OldValue, GetPropValue("udid"), -1)

			case strings.Contains(element, "RanString"):
				Value = strings.Replace(Value, OldValue, Random(element), -1) // замены шаблонов на переменные
			case strings.Contains(element, "RanInt"):
				Value = strings.Replace(Value, OldValue, Random(element), -1)
			case strings.Contains(element, "RanStrInt"):
				Value = strings.Replace(Value, OldValue, Random(element), -1)

			default:
				Value = strings.Replace(Value, OldValue, GetPropValue(element), -1)
				if GetPropValue(element) == "" {
					fmt.Println("Такой переменной нет в файле data.json ->  ", element)
					fmt.Println("Выполнение остановленно - Смотреть Replase.go", element)
					os.Exit(0)
				}
			}
		} else {
			Value = element
		}
	}
	return Value
}
