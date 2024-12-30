package un_sdk

import "encoding/json"

type Config struct {
	ApiURL    string `json:"api_url"`
	AppID     string `json:"app_id"`
	AppSecret string `json:"app_secret"`
	Sign      string `json:"sign"`
	Version   string `json:"version"`
}

// NewConfig 创建一个新的配置实例
func NewConfig(conf *Config) *Config {
	sign := GenerateSignature(conf.AppID, conf.AppSecret)
	return &Config{
		ApiURL:    conf.ApiURL,
		AppID:     conf.AppID,
		AppSecret: conf.AppSecret,
		Sign:      sign,
		Version:   conf.Version,
	}
}
func (c *Config) String() string {
	data, _ := json.MarshalIndent(c, "", "  ")
	//data, _ := json.Marshal(c)
	return string(data)
}
