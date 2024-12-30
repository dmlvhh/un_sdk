package un_sdk

import (
	"fmt"
	"testing"
)

func TestGenerateSignature(t *testing.T) {
	timestamp := GetCurrentTimestamp()
	transID := GenerateSerialNumber()
	appID := "PNqMsNLn11"
	appSecret := "c25epLsMZIFJVasfxhfVxBthc5Ivl3"
	signature := GenerateSignature(appID, appSecret, timestamp, transID)
	fmt.Println(timestamp)
	fmt.Println(transID)
	fmt.Println(signature)
}

func TestGetCurrentTimestamp(t *testing.T) {
	fmt.Println(GetCurrentTimestamp())
}

func TestGenerateSerialNumber(t *testing.T) {
	fmt.Println(GenerateSerialNumber())
}

func TestGetMictime(t *testing.T) {
	//fmt.Println(GetMictime(1)) // 输出类型 1 的时间戳
	//fmt.Println(GetMictime(2)) // 输出类型 2 的时间戳
	//fmt.Println(GetMictime(0)) // 输出默认的时间戳
}
