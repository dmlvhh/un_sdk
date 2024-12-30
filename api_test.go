package un_sdk

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestConfig_WsGetTerminalDetails(t *testing.T) {
	res, err := unSdk.WsGetTerminalDetails(&WsGetTerminalUsageReq{
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
