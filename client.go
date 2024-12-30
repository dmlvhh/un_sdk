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

func ApiRequest(url string, data any) (res string, err error) {
	reqData, err := json.Marshal(Request{
		AppID:     unSdk.AppID,
		Timestamp: GetCurrentTimestamp(),
		TransID:   GenerateSerialNumber(),
		Token:     unSdk.Sign,
		Data:      data,
	})
	fmt.Println(reqData)
	if err != nil {
		log.Printf("Error Marshal: %s", err)
		return
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // 不安全，仅用于测试
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Post(unSdk.ApiURL+url+unSdk.Version, "application/json", bytes.NewBuffer(reqData))
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
