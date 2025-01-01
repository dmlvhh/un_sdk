package un_sdk

import (
	"fmt"
	"testing"
)

func TestApiRequest(t *testing.T) {
	request, err := ApiRequest("/wsGetTerminalDetails", WsGetTerminalDetails{
		MessageId: "1",
		//OpenId:    "35858ouPtS8WqMB",
		OpenId: "32913oucTnkZIeP",
		//OpenId:  "37018oumTfzh79h",
		Version: "1.0",
		//Iccids:  []string{"89860624650023884322"},
		Iccids: []string{"89860622510012775034"},
		//Iccids: []string{"89860623580065903063"},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(request)
}
