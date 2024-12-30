package un_sdk

type Request struct {
	AppID     string      `json:"app_id"`
	Timestamp string      `json:"timestamp"`
	TransID   string      `json:"trans_id"`
	Token     string      `json:"token"`
	Data      interface{} `json:"data"`
}

type WsGetTerminalUsageReq struct {
}
