package un_sdk

import "encoding/json"

type Config struct {
	ApiURL    string `json:"api_url"`
	AppID     string `json:"app_id"`
	AppSecret string `json:"app_secret"`
	Sign      string `json:"sign"`
	Version   string `json:"version"`
	Timestamp string `json:"timestamp"`
	TransID   string `json:"trans_id"`
	OpenId    string `json:"openId"`
}

// NewConfig 创建一个新的配置实例
func NewConfig2(conf *Config) *Config {
	timestamp := GetCurrentTimestamp()
	transID := GenerateSerialNumber()
	sign := GenerateSignature(conf.AppID, conf.AppSecret, timestamp, transID)
	return &Config{
		ApiURL:    conf.ApiURL,
		AppID:     conf.AppID,
		AppSecret: conf.AppSecret,
		Sign:      sign,
		Version:   conf.Version,
		Timestamp: timestamp,
		TransID:   transID,
		OpenId:    conf.OpenId,
	}
}
func NewConfig(apiURL, appID, appSecret, version, OpenId string) *Config {
	timestamp := GetCurrentTimestamp() // 获取当前时间戳
	transID := GenerateSerialNumber()  // 生成一个交易
	return &Config{
		ApiURL:    apiURL,
		AppID:     appID,
		OpenId:    OpenId,
		AppSecret: appSecret,
		Version:   version,
		Timestamp: timestamp,                                               // 假设你有一个函数来生成当前时间戳
		TransID:   transID,                                                 // 假设你有一个函数生成唯一的 TransID
		Sign:      GenerateSignature(appID, appSecret, timestamp, transID), // 签名生成
	}
}
func (c *Config) String() string {
	data, _ := json.MarshalIndent(c, "", "  ")
	//data, _ := json.Marshal(c)
	return string(data)
}
