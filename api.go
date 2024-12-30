package un_sdk

import (
	"encoding/json"
	"fmt"
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
func (c *Config) WsGetTerminalDetails(req *WsGetTerminalUsageReq) (res *WsGetTerminalDetailsRes, err error) {
	request, err := ApiRequest("/wsGetTerminalDetails", req)
	if err != nil {
		log.Fatalf("wsGetTerminalDetails: %s", err)
		return
	}
	err = json.Unmarshal([]byte(request), &res)
	fmt.Println(err)
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
	fmt.Println(err)
	return
}

func (c *Config) WsGetTerminalUsageDataDetails(req *WsGetTerminalUsageDataDetailsReq) (res *WsGetTerminalUsageDataDetailsRes, err error) {
	request, err := ApiRequest("/wsGetTerminalUsageDataDetails", req)
	if err != nil {
		log.Fatalf("wsGetTerminalUsageDataDetails: %s", err)
		return
	}
	fmt.Println(request)
	err = json.Unmarshal([]byte(request), &res)
	fmt.Println(err)
	return
}
