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

//var unSdk *Config

//	func init() {
//		unSdk = NewConfig(&Config{
//			ApiURL:    "https://gwapi.10646.cn/api",
//			AppID:     "PNqMsNLn11",
//			AppSecret: "c25epLsMZIFJVasfxhfVxBthc5Ivl3",
//			Version:   "/V1/1Main/vV1.1",
//		})
//	}http://ww.asyw.tianiot.com/index/index/ydToken?appid=bbO7egzjW5&chl_id=14
//func init() {
//	unSdk = NewConfig(&Config{
//		ApiURL:    "https://gwapi.10646.cn/api",
//		AppID:     "bbO7egzjW5",
//		AppSecret: "wwpsyL5eACfU3RKgeqOmCNS8MwIotd",
//		Version:   "/V1/1Main/vV1.1",
//	})
//}

//	func init() {
//		unSdk = NewConfig()
//	}
func (c *Config) ApiRequest(url string, data any) (res string, err error) {
	reqData, err := json.Marshal(Request{
		AppID:     c.AppID,
		Timestamp: c.Timestamp,
		TransID:   c.TransID,
		Token:     c.Sign,
		Data:      data,
	})
	//fmt.Println(string(reqData))
	fmt.Println("c", c)
	if err != nil {
		log.Printf("Error Marshal: %s", err)
		return
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // 不安全，仅用于测试
	}
	apiUrl := c.ApiURL + url + c.Version
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
