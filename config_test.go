package un_sdk_test

import (
	"fmt"
	"github.com/dmlvhh/un_sdk"
	"testing"
)

func TestNewConfig(t *testing.T) {
	unSdk := un_sdk.NewConfig(&un_sdk.Config{
		ApiURL:    "dasdas",
		AppID:     "dasdsa",
		AppSecret: "asdadada",
		Version:   "/V1/1Main/vV1.1",
	})
	fmt.Println(unSdk)
}
