package un_sdk

import (
	"fmt"
	"testing"
)

func TestGenerateSignature(t *testing.T) {
	appID := "PNqMsNLn11"
	appSecret := "c25epLsMZIFJVasfxhfVxBthc5Ivl3"
	signature := GenerateSignature(appID, appSecret)
	fmt.Println(signature)
}

func TestGetCurrentTimestamp(t *testing.T) {
	fmt.Println(GetCurrentTimestamp())
}

func TestGenerateSerialNumber(t *testing.T) {
	fmt.Println(GenerateSerialNumber())
}
