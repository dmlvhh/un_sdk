package un_sdk

import (
	"encoding/hex"
	"fmt"
	"github.com/tjfoc/gmsm/sm3"
	"math/rand"
	"time"
)

func GenerateSignature(appID, appSecret string) string {
	timestamp := GetCurrentTimestamp()
	transID := GenerateSerialNumber()
	signatureString := fmt.Sprintf("app_id%stimestamp%strans_id%s%s", appID, timestamp, transID, appSecret)
	hash := sm3.Sm3Sum([]byte(signatureString))
	signature := hex.EncodeToString(hash[:])
	return signature
}

func GetCurrentTimestamp() string {
	// 获取当前时间
	currentTime := time.Now()

	// 获取毫秒部分
	milliseconds := currentTime.Nanosecond() / 1_000_000 // 转换为毫秒

	// 格式化为所需的字符串格式
	return fmt.Sprintf("%s %03d", currentTime.Format("2006-01-02 15:04:05"), milliseconds)
}
func GenerateSerialNumber() string {
	// 获取当前时间
	now := time.Now()

	// 格式化日期时间为 YYYYMMDDHHMMSS
	dateTimeStr := now.Format("20060102150405")

	// 获取毫秒部分并格式化为三位字符串
	milliseconds := now.Nanosecond() / 1_000_000
	millisecondsStr := fmt.Sprintf("%03d", milliseconds)

	// 生成六位随机数
	rand.Seed(time.Now().UnixNano())                     // 初始化随机数种子
	randomNumber := rand.Intn(1000000)                   // 生成一个 0 到 999999 的随机数
	randomNumberStr := fmt.Sprintf("%06d", randomNumber) // 格式化为六位字符串

	// 拼接序列号
	serialNumber := fmt.Sprintf("%s%s%s", dateTimeStr, millisecondsStr, randomNumberStr)

	return serialNumber
}
