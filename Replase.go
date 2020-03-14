package main

import (
	"regexp"
	"strings"
)

var  (
	Value string
	OldValue string
)

func Replace(value string) (string)  {
	Value := value
	re := regexp.MustCompile(`([${{])(.*?){*}([}}])`)
	submatchall := re.FindAllString(value, -1)

	for _, element := range submatchall {
		element = strings.Trim(element, "${{")
		element = strings.Trim(element, "}}")
		OldValue := "${{" + element + "}}"

		switch {
		case strings.Contains(element, "RanString"):
			Value = strings.Replace(Value, OldValue, Random(element), -1) // замены шаблонов на переменные
		case strings.Contains(element, "RanInt"):
			Value = strings.Replace(Value, OldValue, Random(element), -1)
		case strings.Contains(element, "RanStrInt"):
			Value = strings.Replace(Value, OldValue, Random(element), -1)
		default:
			Value = strings.Replace(Value, OldValue, GetPropValue(element), -1)
		}
	}
	return Value
}
