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
func NewConfig(conf *Config) *Config {
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
func (c *Config) String() string {
	data, _ := json.MarshalIndent(c, "", "  ")
	//data, _ := json.Marshal(c)
	return string(data)
}
