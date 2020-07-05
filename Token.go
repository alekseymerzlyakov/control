package main

import (
	"crypto"
	"crypto/hmac"
	"encoding/hex"
	"fmt"
)
func ComputeHmac256(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(crypto.SHA256.New, key)
	h.Write([]byte(message))
	sha256 := hex.EncodeToString(h.Sum(nil))
	return sha256//base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func Token()  {

	developer_key := GetPropValue("developer_key")
	fmt.Println("developer_key  -->    ", developer_key)
	devicekey := GetPropValue("devicekey")
	fmt.Println("devicekey  -->    ", devicekey)

	challenge := GetPropValue("challenge")
	fmt.Println("challenge  -->    ", challenge)
	udid := GetPropValue("udid")
	fmt.Println("udid  -->    ", udid)

	master_key := ComputeHmac256(udid + devicekey, developer_key)
	fmt.Println("master_key  -->    ", master_key)
	token := ComputeHmac256(challenge, master_key)
	fmt.Println("token  -->    ", token)
	SetPropValue("token",token)


}

func UDID()  {
	developerKey := GetPropValue("developerkey")
	imei = Random("RanStrInt##10")
	key := []byte(developerKey)
	h := hmac.New(crypto.SHA256.New, key)
	h.Write([]byte(imei))
	sha256 := hex.EncodeToString(h.Sum(nil))

	fmt.Println("UDID  -->    ", sha256)
	SetPropValue("UDID",sha256)
}
