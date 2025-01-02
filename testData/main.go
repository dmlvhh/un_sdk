package main

import (
	"fmt"
	"github.com/dmlvhh/un_sdk"
	"log"
)

func main() {
	// 初始化 SDK 配置
	unSdk := un_sdk.NewConfig("https://gwapi.10646.cn", "bbO7egzjW5", "wwpsyL5eACfU3RKgeqOmCNS8MwIotd", "/V1/1Main/vV1.1", "32913oucTnkZIeP")
	//unSdk := un_sdk.NewConfig("https://gwapi.10646.cn", "PNqMsNLn11", "c25epLsMZIFJVasfxhfVxBthc5Ivl3", "/V1/1Main/vV1.1", "35858ouPtS8WqMB")
	fmt.Println(unSdk)
	res, err := unSdk.WsGetTerminalDetails(&un_sdk.WsGetTerminalDetailsReq{
		MessageId: "123456",
		OpenId:    unSdk.OpenId,
		Version:   unSdk.Version,
		Iccids:    []string{"89860623580065927849", "89860623580065927708"},
		//Iccids: []string{"89860624650023884322"},
	})
	if err != nil {
		return
	}
	if err != nil {
		log.Fatal("Request failed:", err)
	}
	fmt.Println("API response:", res)
}
