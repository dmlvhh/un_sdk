package un_sdk

var unSdk *Config

func init() {
	unSdk = NewConfig(&Config{
		ApiURL:    "https://gwapi.10646.cn/api",
		AppID:     "PNqMsNLn11",
		AppSecret: "c25epLsMZIFJVasfxhfVxBthc5Ivl3",
	})
}
