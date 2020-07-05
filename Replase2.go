package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

var  (
	Value2 string
	OldValue2 string
)

func (b Data) Replace2(value string) (string)  {

	Value2 := value
	re := regexp.MustCompile(`([${{])(.*?){*}([}}])`)
	submatchall := re.FindAllString(value, -1)

	for _, element2 := range submatchall {
		if strings.Contains(element2, "${{") && strings.Contains(element2, "}}") {
			element2 = strings.Trim(element2, "${{")
			element2 = strings.Trim(element2, "}}")
			OldValue2 := "${{" + element2 + "}}"

			switch {
			case strings.Contains(element2, "udid"):
				//fmt.Println("Такой переменной нет в файле data.json ->  ", element)
				if GetPropValue2(element2) == "" {
					fmt.Println("Такой переменной нет в файле data.json ->  ", element2)
					fmt.Println("Выполнение остановленно - Смотреть Replase.go", element2)

					msg := "Такой переменной нет в файле data.json ->  " +element2
					msg = msg + "\nВыполнение остановленно - Смотреть Replase.go\n"
					msg = msg + b.Response2.URL2
						telegram(msg)
					os.Exit(0)
				}
				Value2 = strings.Replace(Value2, OldValue2, GetPropValue2("udid"), -1)

			case strings.Contains(element2, "RanString"):
				Value2 = strings.Replace(Value2, OldValue2, Random2(element2), -1) // замены шаблонов на переменные
			case strings.Contains(element2, "RanInt"):
				Value2 = strings.Replace(Value2, OldValue2, Random2(element2), -1)
			case strings.Contains(element2, "RanStrInt"):
				Value2 = strings.Replace(Value2, OldValue2, Random2(element2), -1)

			default:
				Value2 = strings.Replace(Value2, OldValue2, GetPropValue2(element2), -1)
				if GetPropValue2(element2) == "" {
					fmt.Println("Такой переменной нет в файле data.json ->  ", element2)
					fmt.Println("Выполнение остановленно - Смотреть Replase.go", element2)
					os.Exit(0)
				}
			}
		} else {
			Value2 = element2
		}
	}
	return Value2
}
