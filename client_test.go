package un_sdk

import (
	"fmt"
	"testing"
)

func TestApiRequest(t *testing.T) {
	request, err := ApiRequest("/wsGetTerminalDetails", WsGetTerminalDetails{
		MessageId: "1",
		OpenId:    "35858ouPtS8WqMB",
		Version:   "1.0",
		Iccids:    []string{"89860624650023884322"},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(request)
}
