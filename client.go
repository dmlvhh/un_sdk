package un_sdk

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var unSdk *Config

func init() {
	unSdk = NewConfig(&Config{
		ApiURL:    "https://gwapi.10646.cn/api",
		AppID:     "PNqMsNLn11",
		AppSecret: "c25epLsMZIFJVasfxhfVxBthc5Ivl3",
		Version:   "/V1/1Main/vV1.1",
	})
}
func ApiRequest(url string, data any) (res string, err error) {
	reqData, err := json.Marshal(Request{
		AppID:     unSdk.AppID,
		Timestamp: unSdk.Timestamp,
		TransID:   unSdk.TransID,
		Token:     unSdk.Sign,
		Data:      data,
	})
	fmt.Println(string(reqData))
	if err != nil {
		log.Printf("Error Marshal: %s", err)
		return
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // 不安全，仅用于测试
	}
	apiUrl := unSdk.ApiURL + url + unSdk.Version
	fmt.Println(apiUrl)
	client := &http.Client{Transport: tr}
	resp, err := client.Post(apiUrl, "application/json", bytes.NewBuffer(reqData))
	if err != nil {
		log.Printf("Error making POST request: %s", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %s", err)
		return
	}
	return string(body), nil
}
