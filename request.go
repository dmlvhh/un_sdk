package un_sdk

type Request struct {
	AppID     string      `json:"app_id"`
	Timestamp string      `json:"timestamp"`
	TransID   string      `json:"trans_id"`
	Token     string      `json:"token"`
	Data      interface{} `json:"data"`
}

type WsGetTerminalDetailsReq struct {
	MessageId string   `json:"messageId"`
	OpenId    string   `json:"openId"`
	Version   string   `json:"version"`
	Iccids    []string `json:"iccids"`
}
type WsGetTerminalUsageReq struct {
	MessageId    string `json:"messageId"`
	OpenId       string `json:"openId"`
	Version      string `json:"version"`
	Iccid        string `json:"iccid"`
	BillingCycle string `json:"billingCycle"`
}

type WsGetTerminalDetails struct {
	MessageId string   `json:"messageId"`
	OpenId    string   `json:"openId"`
	Version   string   `json:"version"`
	Iccids    []string `json:"iccids"`
}

type WsEditTerminalReq struct {
	MessageId    string `json:"messageId"`
	OpenId       string `json:"openId"`
	Version      string `json:"version"`
	Asynchronous string `json:"asynchronous"`
	Iccid        string `json:"iccid"`
	ChangeType   string `json:"changeType"`
	TargetValue  string `json:"targetValue"`
}

type WsGetTerminalUsageDataDetailsReq struct {
	MessageId      string `json:"messageId"`
	OpenId         string `json:"openId"`
	Version        string `json:"version"`
	Iccid          string `json:"iccid"`
	CycleStartDate string `json:"cycleStartDate"`
	PageNumber     string `json:"pageNumber"`
}

type WsQueryAcctTableReq struct {
	MessageId  string `json:"messageId"`
	OpenId     string `json:"openId"`
	Version    string `json:"version"`
	PageNumber string `json:"pageNumber"`
}

type WsGetTerminalEventsReq struct {
	MessageId string `json:"messageId"`
	OpenId    string `json:"openId"`
	Version   string `json:"version"`
	Iccid     string `json:"iccid"`
	//Since      string `json:"since"`
	//PageNumber string `json:"pageNumber"`
}

type LocationBasedServiceReq struct {
	Type string `json:"type"`
	Msid string `json:"msid"`
}
