package un_sdk

import (
	"encoding/json"
	"log"
)

/*
*
  - 查询设备详情
  - simStatus:SIM卡状态:
    "0": 可测试,
    "1": 可激活,
    "2": 已激活,
    "3": 已停用,
    "4": 已失效,
    "5"": 已清除,
    "6": 已更换,
    "7": 库存,
    "8": 开始
    *deviceId:设备ID
  - ratePlan:资费计划
  - realNameStatus:实名制状态(0-不需实名未实名,1-需要实名未实名,2-需要实名已实名,3-不需要实名已实名).
  - lockStatus:锁定状态 0:否，1:是
  - dateActivated:激活日期
    *
    *
*/
//WsGetTerminalDetails 查询设备详情
func (c *Config) WsGetTerminalDetails(req *WsGetTerminalDetailsReq) (res *WsGetTerminalDetailsRes, err error) {
	request, err := ApiRequest("/wsGetTerminalDetails", req)
	if err != nil {
		log.Fatalf("wsGetTerminalDetails: %s", err)
		return
	}
	err = json.Unmarshal([]byte(request), &res)
	return
}

// WsEditTerminal 0: 可测试 1:
// 可激活 2: 已激活 3: 已停用 4: 已失效 5: 已清
// 除 6: 已更换 7: 库存 8: 开始 9:预清除
// WsEditTerminal 为给定设备更改属性的值
func (c *Config) WsEditTerminal(req *WsEditTerminalReq) (res *WsEditTerminalRes, err error) {
	request, err := ApiRequest("/wsEditTerminal", req)
	if err != nil {
		log.Fatalf("wsEditTerminal: %s", err)
		return
	}
	err = json.Unmarshal([]byte(request), &res)
	return
}

func (c *Config) WsGetTerminalUsageDataDetails(req *WsGetTerminalUsageDataDetailsReq) (res *WsGetTerminalUsageDataDetailsRes, err error) {
	request, err := ApiRequest("/wsGetTerminalUsageDataDetails", req)
	if err != nil {
		log.Fatalf("wsGetTerminalUsageDataDetails: %s", err)
		return
	}
	err = json.Unmarshal([]byte(request), &res)
	return
}

// WsGetTerminalUsage 查询设备用量
func (c *Config) WsGetTerminalUsage(req *WsGetTerminalUsageReq) (res *WsGetTerminalUsageRes, err error) {
	request, err := ApiRequest("/wsGetTerminalUsage", req)
	if err != nil {
		log.Fatalf("wsGetTerminalUsage: %s", err)
		return
	}
	err = json.Unmarshal([]byte(request), &res)
	return
}

func (c *Config) WsQueryAcctTable(req *Request) (res *WsQueryAcctTableRes, err error) {
	request, err := ApiRequest("/wsQueryAcctTable", req)
	if err != nil {
		log.Fatalf("WsQueryAcctTable: %s", err)
		return
	}
	err = json.Unmarshal([]byte(request), &res)
	return
}

// WsGetTerminalEvents 查询指定设备关联的事件资费计划
func (c *Config) WsGetTerminalEvents(req *WsGetTerminalEventsReq) (res *WsGetTerminalEventsRes, err error) {
	request, err := ApiRequest("/wsGetTerminalEvents", req)
	if err != nil {
		log.Fatalf("wsGetTerminalEvents: %s", err)
		return
	}
	err = json.Unmarshal([]byte(request), &res)
	return
}

// 立即定位服务https://gwapi.10646.cn/api/LocationBasedService/V1.0
func (c *Config) LocationBasedService(req *Request) (res *LocationBasedServiceRes, err error) {
	request, err := ApiRequest("/LocationBasedService/V1.0", req)
	if err != nil {
		log.Fatalf("LocationBasedService: %s", err)
		return
	}
	err = json.Unmarshal([]byte(request), &res)
	return
}
