package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	log "github.com/sirupsen/logrus"
	"gopkg.in/ezzarghili/recaptcha-go.v4"
	"os"
	"regexp"
	_ "regexp"
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

var waitTime string

func Authorization_ozer(get string) {

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

func Authorization() {
	msg := GetPropValue("clientid") + ":" + GetPropValue("secret_key")
	SetPropValue("Authorization", base64.StdEncoding.EncodeToString([]byte(msg)))
}

//
func ApiTest(get string) {
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
	fmt.Println("authorize_code ---- >>>   ", authorize_code)
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

func UnixTime() {
	currentTime := strconv.FormatInt(time.Now().Unix(), 10)
	log.Println("unixtime - - - >      " + currentTime)
	SetPropValue("unixtime", string(currentTime))
}

func makeTimestamp() string {
	return strconv.FormatInt(time.Now().UnixNano()/1000000, 10)

}

func Crop(exp string) {
	cropExpression := strings.Split(exp, ">>")
	fmt.Println("cropExpression   ------>       ", cropExpression[1])

	//Last_4##data.url##systrans_id;
	cropArr := strings.Split(cropExpression[1], "##")
	varFromProrerty := GetPropValue(cropArr[1])
	varFromProrertyLen := len(varFromProrerty)

	switch {
	case strings.Contains(cropArr[0], "Last_"):
		lenNewVar := strings.Split(cropArr[0], "_")
		lenNewVar_, _ := strconv.Atoi(lenNewVar[1])
		newVar := varFromProrerty[varFromProrertyLen-lenNewVar_ : varFromProrertyLen]
		fmt.Println("varFromProrerty   ------>       ", varFromProrerty)
		fmt.Println("varFromProrertyLen   ------>       ", varFromProrertyLen)
		fmt.Println("newVar   ------>       ", newVar)
		SetPropValue(cropArr[2], newVar)

	case strings.Contains(cropArr[0], "Split_"):
		lenNewVar := strings.Split(cropArr[0], "_")                   // получили масив значений Split_id=_1
		splitValue := lenNewVar[1]                                    // взяли по чем разбивать
		numElement, _ := strconv.Atoi(lenNewVar[2])                   // какой елемент сохранять Split_id=_ -> 1
		splitValueArray := strings.Split(varFromProrerty, splitValue) // Массив елементов после разбивки
		fmt.Println("splitValueArray[1]   ------>       ", splitValueArray[1])
		SetPropValue(cropArr[2], splitValueArray[numElement])

	default:
		//fmt.Println("switch выбора функций   ----------->>>  отработал default - значит не нашли функцию -->>   ", exp)

	}
}

func Otp() {
	//apilogin := GetPropValue("otp.apilogin")
	//obj_id := GetPropValue("otp.obj_id")
	//conv_id := GetPropValue("otp.conv_id")
	apisecret := GetPropValue("otp.apisecret")
	body := GetPropValue("otp.body")
	epoch := makeTimestamp()
	SetPropValue("otp.epoch", epoch)
	request := epoch + apisecret + body + apisecret
	request_byte := []byte(request)
	fmt.Println("request_byte   ------>       ", string(request_byte[:]))

	h := sha1.New()
	h.Write([]byte(request))
	bs := h.Sum(nil)
	md5HashInString := hex.EncodeToString(bs[:])
	fmt.Println("SHA1   ------>       ", md5HashInString)
	SetPropValue("otp.hex", md5HashInString)
}

func ParseIntFromText(get string) {
	fmt.Println("<<ParseIntFromText>>   ----------->>>    ", get)

	get_path_array := strings.Split(get, ">>")
	pathToText := get_path_array[1]
	msg := "Func ParseIntFromText не содержит путь к тексту: корректный пример - <<ParseIntFromText>>otp.otp1;"
	if pathToText == "" {
		telegram(msg)
		os.Exit(0)
	}
	fmt.Println("pathToText   ----------->>>    ", pathToText)

	text := GetPropValue(pathToText)
	fmt.Println("text   ----------->>>    ", text)
	msg_ := "нет переменной -> otp.otp1 -> OTP небыло в ответе"
	if len(text) <= 1 {
		telegram(msg_)
		os.Exit(0)
	}

	fmt.Println("text   ------>       ", text)
	re := regexp.MustCompile("[0-9]+")
	s := re.FindAllString(text, -1)
	if len(s) == 0 {
		s = re.FindAllString("9999", -1)
	}
	SetPropValue(pathToText, s[0])
}

func Wait(get string) {
	get_path_array := strings.Split(get, ">>")
	waitTime := get_path_array[1]
	if waitTime == "" {
		waitTime = "5s"
	}

	comp, _ := time.ParseDuration(waitTime)
	fmt.Println("<<Wait>>   ----------->>>    ", comp)
	time.Sleep(comp)
}

func ReCaptcha(get string) {

	recaptchaSecret := GetPropValue("testdata.pkey")

	captcha, _ := recaptcha.NewReCAPTCHA(recaptchaSecret, recaptcha.V2, 10*time.Second) // for v2 API get your secret from https://www.google.com/recaptcha/admin

	err := captcha.Verify("03AGdBq24F-gUmho-vP5M7KN6sVy7hEeSBps6_z907DbrnsLeGfbyc9gWUAqtfTJrwr7dT9m-w_fHBCT5yxGf8qkWkvYxo-Ot8qCThpMjQ1UuBQtSJP4eWj3eQQwqSA0soEZjfWi2_fSRIEisQi-6s64PJJt2O6VBcTnivP4ORy5drouOmPpHUhA0mNK8rgga4Q42g7xqlewMTTLeYTy7hmbt8Uvgc1o6iTYOdEOR5z2P4GpOQ2fonHN0imnYqz4GFod5xr7KTI82WDkRvk-QRjxc3K0HQLDWIS-rQyERTb27fpePKrdGl4-a51LBvg_4FYOaTD1POyAXjnznSMNF_rKujwARBuydYeSvOyBjUsMEoC68a1ejqREHMepuDhDNbiIIE1p1fxiXcaJiolwgpL0-cGIy3NSy0fQ")
	if err != nil {
		// do something with err (log?)
		// Example check error codes array if they exist: (err.(*recaptcha.Error)).ErrorCodes
	}

}
