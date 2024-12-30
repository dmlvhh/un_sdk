package un_sdk

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
type WsGetTerminalDetailsRes struct {
	Message string `json:"message"`
	Data    struct {
		MessageId  string `json:"messageId"`
		Version    string `json:"version"`
		ResultCode string `json:"resultCode"`
		Terminals  []struct {
			MonthToDateSMSUsage     string `json:"monthToDateSMSUsage"`
			FieldValue9             string `json:"fieldValue9"`
			FieldValue10            string `json:"fieldValue10"`
			MonthToDateDataUsage    string `json:"monthToDateDataUsage"`
			MonthToDateVoiceUsage   string `json:"monthToDateVoiceUsage"`
			CopyRulePwd             string `json:"copyRulePwd"`
			FixedIpAddress          string `json:"fixedIpAddress"`
			RatePlan                int    `json:"ratePlan"`
			OperatorCustom1         string `json:"operatorCustom1"`
			Iccid                   string `json:"iccid"`
			RealNameStatus          string `json:"realNameStatus"`
			OperatorAccountId       string `json:"operatorAccountId"`
			DateActivated           string `json:"dateActivated"`
			OverageLimitReached     string `json:"overageLimitReached"`
			TestableDataUsage       string `json:"testableDataUsage"`
			FieldValue3             string `json:"fieldValue3"`
			OperatorCustom2         string `json:"operatorCustom2"`
			MonthToDateSMSUsageMO   string `json:"monthToDateSMSUsageMO"`
			OperatorCustom4         string `json:"operatorCustom4"`
			Imsi                    string `json:"imsi"`
			NetworkType             string `json:"networkType"`
			CtdSessionCount         int    `json:"ctdSessionCount"`
			FieldValue5             string `json:"fieldValue5"`
			Imei                    string `json:"imei"`
			ProvinceCusutom5        string `json:"provinceCusutom5"`
			LockStatus              string `json:"lockStatus"`
			OverageLimitOverride    string `json:"overageLimitOverride"`
			FieldValue6             string `json:"fieldValue6"`
			OperatorCustom3         string `json:"operatorCustom3"`
			OperatorCustom5         string `json:"operatorCustom5"`
			ProvinceCusutom3        string `json:"provinceCusutom3"`
			Msisdn                  string `json:"msisdn"`
			Customer                string `json:"Customer"`
			CustomerCustom5         string `json:"customerCustom5"`
			CustomerCustom2         string `json:"customerCustom2"`
			MonthToDateVoiceUsageMO string `json:"monthToDateVoiceUsageMO"`
			FieldValue8             string `json:"fieldValue8"`
			CustomerCustom3         string `json:"customerCustom3"`
			CustomerCustom4         string `json:"customerCustom4"`
			CustomerCustom1         string `json:"customerCustom1"`
			FieldValue7             string `json:"fieldValue7"`
			DeviceId                string `json:"deviceId"`
			DateReceivied           string `json:"dateReceivied"`
			FieldValue4             string `json:"fieldValue4"`
			SecureSimId             string `json:"secureSimId"`
			TotalTimeValue          string `json:"totalTimeValue"`
			FieldValue2             string `json:"fieldValue2"`
			FieldValue1             string `json:"fieldValue1"`
			DateAdded               string `json:"dateAdded"`
			ProvinceCusutom4        string `json:"provinceCusutom4"`
			ProvinceCusutom2        string `json:"provinceCusutom2"`
			ProvinceCusutom1        string `json:"provinceCusutom1"`
			ModemId                 string `json:"modemId"`
			AccountId               string `json:"accountId"`
			DateShipped             string `json:"dateShipped"`
			SimStatus               string `json:"simStatus"`
			MonthToDateUsage        string `json:"monthToDateUsage"`
			SimNotes                string `json:"simNotes"`
			DateModified            string `json:"dateModified"`
		} `json:"terminals"`
		Timestamp string `json:"timestamp"`
	} `json:"data"`
	Status string `json:"status"`
}
type WsGetTerminalUsageRes struct {
}

type WsEditTerminalRes struct {
	Message string `json:"message"`
	Data    struct {
		MessageId     string `json:"messageId"`
		Version       string `json:"version"`
		ResultCode    string `json:"resultCode"`
		Timestamp     string `json:"timestamp"`
		Iccid         string `json:"iccid"`
		EffectiveDate string `json:"effectiveDate"`
	} `json:"data"`
	Status string `json:"status"`
}

type WsGetTerminalUsageDataDetailsRes struct {
}
