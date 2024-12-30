package un_sdk

import (
	"fmt"
	"testing"
)

func TestApiRequest(t *testing.T) {
	request, err := ApiRequest("/wsGetTerminalUsage", WsGetTerminalUsageReq{})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(request)
}
