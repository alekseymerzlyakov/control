package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	log "github.com/sirupsen/logrus"
	_"regexp"
	"strconv"
	"strings"
	"time"
)


//switch {
//case FuncName == "ParseOtp":
//	re := regexp.MustCompile("[0-9]+")
//	s := re.FindAllString(valueTemp, -1)
//	if len(s) == 0 {
//		s = re.FindAllString("9999", -1)
//	}
//	SetPropValue(key,  s[0])
//default:
//	SetPropValue(key,  valueTemp)
//
//}



func Authorization_ozer (get string) {

	switch {

	case strings.Contains(get, "Ukraine"):
		msg := GetPropValue("data.country.ukraine.clientid") + ":" + GetPropValue("data.country.ukraine.secret_key")
		SetPropValue("data.country.ukraine.authorization", base64.StdEncoding.EncodeToString([]byte(msg)))

	case strings.Contains(get, "Kenya"):
		msg := GetPropValue("data.country.kenya.clientid") + ":" + GetPropValue("data.country.kenya.secret_key")
		SetPropValue("data.country.kenya.authorization", base64.StdEncoding.EncodeToString([]byte(msg)))


	default:
		fmt.Println("Authorization_ozer -> default")
	}
}

func Authorization () {
		msg := GetPropValue("clientid") + ":" + GetPropValue("secret_key")
		SetPropValue("Authorization", base64.StdEncoding.EncodeToString([]byte(msg)))
}


//
func ApiTest (get string) {
	fmt.Println("<<ApiTest>>   ----------->>>    ", get)
	apiTestName := strings.Split(get, ">>")
	getNumTestLine2 := GetNumTestLine2("BefoAftTest", apiTestName[1]) // поиск номера строки начала теста
	fmt.Println("getNumTestLine2   ----------->>>    ", getNumTestLine2)
	RendomData.ApiSheetName2 = "BefoAftTest"
	RendomData.ApiTest2("BefoAftTest", getNumTestLine2) //
}



func JWT() {
	//Token
	//Access_Token = ReqGetAccess_Token(URLq, Encoded(Aus))

	//Authorize
	authorize_code := GetPropValue("authorization_code")
	fmt.Println("authorize_code ---- >>>   ",authorize_code)
	var header string = "{\"typ\":\"JWT\",\"alg\":\"HS256\"}"
	var payload string = "{\"authorization_code\":\"" + authorize_code + "\"}"
	JWT_header_payload := Encoded(header) + "." + Encoded(payload)
	//var secretKey string = GetString("secretKey")

	HashSt := Hash(JWT_header_payload, GetPropValue("secretkey"))
	SetPropValue("jwt", HashSt)
}

// Hash generates a Hmac256 hash of a string using a secret


func Hash(src string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(src))
	var jw = src + "." + base64.StdEncoding.EncodeToString(h.Sum(nil))
	log.Info("JWT      ", jw)
	return jw
}

func Encoded(base string) string {
	encoded := base64.StdEncoding.EncodeToString([]byte(base))
	log.Println("encoded      " + encoded)
	return encoded
}

func UnixTime()  {
	currentTime := strconv.FormatInt(time.Now().Unix(), 10)
	log.Println("unixtime - - - >      " + currentTime)

	SetPropValue("unixtime", string(currentTime))

}

func Crop(exp string)  {
	cropExpression := strings.Split(exp, ">>")
	fmt.Println("cropExpression   ------>       ", cropExpression[1])

	//Last_4##data.url##systrans_id;
	cropArr := strings.Split(cropExpression[1], "##")
	varFromProrerty := GetPropValue(cropArr[1])
	varFromProrertyLen := len(varFromProrerty)

	switch {
	case strings.Contains(cropArr[0], "Last_"):
		lenNewVar := strings.Split(cropArr[0], "_")
		lenNewVar_ , _ :=  strconv.Atoi(lenNewVar[1])
		newVar := varFromProrerty[varFromProrertyLen-lenNewVar_:varFromProrertyLen]
		fmt.Println("varFromProrerty   ------>       ", varFromProrerty)
		fmt.Println("varFromProrertyLen   ------>       ", varFromProrertyLen)
		fmt.Println("newVar   ------>       ", newVar)
		SetPropValue(cropArr[2], newVar)

	case strings.Contains(cropArr[0], "Split_"):
		lenNewVar := strings.Split(cropArr[0], "_") // получили масив значений Split_id=_1
		splitValue := lenNewVar[1] // взяли по чем разбивать
		numElement , _ :=  strconv.Atoi(lenNewVar[2]) // какой елемент сохранять Split_id=_ -> 1
		splitValueArray := strings.Split(varFromProrerty, splitValue) // Массив елементов после разбивки
		fmt.Println("splitValueArray[1]   ------>       ", splitValueArray[1])
		SetPropValue(cropArr[2], splitValueArray[numElement])


	default:
		fmt.Println("switch выбора функций   ----------->>>  отработал default - значит не нашли функцию -->>   ", exp)

	}
}



