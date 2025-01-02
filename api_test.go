package un_sdk

import (
	"encoding/json"
	"fmt"
	"testing"
)

var unSdk Config

func init() {
	unSdk = *NewConfig(
		"https://gwapi.10646.cn/api",
		"bbO7egzjW5",
		"wwpsyL5eACfU3RKgeqOmCNS8MwIotd",
		"/V1/1Main/vV1.1",
		"dasdas",
	)
}

func TestConfig_WsGetTerminalDetails(t *testing.T) {
	res, err := unSdk.WsGetTerminalDetails(&WsGetTerminalDetailsReq{
		MessageId: "123456",
		OpenId:    "35858ouPtS8WqMB",
		Version:   "1.0",
		Iccids:    []string{"89860624650023884322"},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	marshal, _ := json.Marshal(res)
	fmt.Println(string(marshal))
}
func TestConfig_WsEditTerminal(t *testing.T) {
	res, err := unSdk.WsEditTerminal(&WsEditTerminalReq{
		MessageId:    "123456",
		OpenId:       "35858ouPtS8WqMB",
		Version:      "1.0",
		Asynchronous: "0",
		Iccid:        "89860624650023884322",
		ChangeType:   "3",
		TargetValue:  "2",
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	marshal, _ := json.Marshal(res)
	fmt.Println(string(marshal))
}

func TestConfig_WsGetTerminalUsageDataDetails(t *testing.T) {
	res, err := unSdk.WsGetTerminalUsageDataDetails(&WsGetTerminalUsageDataDetailsReq{
		MessageId:      "123456",
		OpenId:         "35858ouPtS8WqMB",
		Version:        "1.0",
		Iccid:          "89860624650023884322",
		CycleStartDate: "2024-10-01 17:35:55",
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	marshal, _ := json.Marshal(res)
	fmt.Println(string(marshal))
}

func TestConfig_WsGetTerminalUsage(t *testing.T) {
	res, err := unSdk.WsGetTerminalUsage(&WsGetTerminalUsageReq{
		MessageId:    "123456",
		OpenId:       "35858ouPtS8WqMB",
		Version:      "1.0",
		Iccid:        "89860624650023884322",
		BillingCycle: "202412",
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	marshal, _ := json.Marshal(res)
	fmt.Println(string(marshal))
}

func TestConfig_WsQueryAcctTable(t *testing.T) {
	timestamp := GetCurrentTimestamp()
	transID := GenerateSerialNumber()
	appID := "PNqMsNLn11"
	appSecret := "c25epLsMZIFJVasfxhfVxBthc5Ivl3"
	signature := GenerateSignature(appID, appSecret, timestamp, transID)
	res, err := unSdk.WsQueryAcctTable(&Request{
		AppID:     "PNqMsNLn11",
		Timestamp: timestamp,
		TransID:   transID,
		Token:     signature,
		Data: WsQueryAcctTableReq{
			MessageId: "123456",
			OpenId:    "35858ouPtS8WqMB",
			Version:   "1.0",
			//PageNumber: "1",
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	marshal, _ := json.Marshal(res)
	fmt.Println(string(marshal))
}

func TestConfig_WsGetTerminalEvents(t *testing.T) {
	res, err := unSdk.WsGetTerminalEvents(&WsGetTerminalEventsReq{
		MessageId: "123456",
		OpenId:    "35858ouPtS8WqMB",
		Version:   "1.0",
		Iccid:     "89860624650023884322",
		//Since:      "",
		//PageNumber: "",
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	marshal, _ := json.Marshal(res)
	fmt.Println(string(marshal))
}

func TestConfig_LocationBasedService(t *testing.T) {
	timestamp := GetCurrentTimestamp()
	transID := GenerateSerialNumber()
	appID := "PNqMsNLn11"
	appSecret := "c25epLsMZIFJVasfxhfVxBthc5Ivl3"
	signature := GenerateSignature(appID, appSecret, timestamp, transID)
	res, err := unSdk.LocationBasedService(&Request{
		AppID:     "PNqMsNLn11",
		Timestamp: timestamp,
		TransID:   transID,
		Token:     signature,
		Data: LocationBasedServiceReq{
			Type: "msisdn",
			Msid: "8986062465002388432",
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	marshal, _ := json.Marshal(res)
	fmt.Println(string(marshal))
}
