package un_sdk

import (
	"encoding/hex"
	"fmt"
	"github.com/tjfoc/gmsm/sm3"
	"math/rand"
	"strings"
	"time"
)

func GenerateSignature(appID, appSecret, timestamp, transID string) string {
	signatureString := fmt.Sprintf("app_id%stimestamp%strans_id%s%s", appID, timestamp, transID, appSecret)
	fmt.Println(signatureString)
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

func GenerateSerialNumber1(timeStr string) string {
	// 原始时间字符串
	//timeStr := "2017-06-20 00:00:00 041"

	// 分离日期和序号
	parts := strings.Split(timeStr, " ")
	dateTimePart := parts[0] + " " + parts[1] // "2017-06-20 00:00:00"
	seqPart := parts[2]                       // "041"

	// 解析时间
	t, _ := time.Parse("2006-01-02 15:04:05", dateTimePart)

	// 格式化时间，去掉分隔符
	formattedTime := t.Format("20060102150405")
	// 拼接序号
	result := formattedTime + seqPart
	rand.Seed(time.Now().UnixNano())                     // 初始化随机数种子
	randomNumber := rand.Intn(1000000)                   // 生成一个 0 到 999999 的随机数
	randomNumberStr := fmt.Sprintf("%06d", randomNumber) // 格式化为六位字符串
	return result + randomNumberStr
}

//func GetMictime(t int) string {
//	// 获取当前时间的微秒级时间戳
//	t1 := time.Now().UnixNano() / int64(time.Millisecond) // 毫秒级时间戳
//	t2 := time.Now().UnixNano() % int64(time.Millisecond) // 微秒级时间戳
//
//	switch t {
//	case 1:
//		// 格式化为 Y-m-d H:i:s 并附加毫秒
//		return fmt.Sprintf("%s%03d", time.Unix(t1/1000, t2).Format("2006-01-02 15:04:05"), t1%1000)
//	case 2:
//		// 格式化为 YmdHis 并附加毫秒和随机数
//		return fmt.Sprintf("%s%03d%06d", time.Unix(t1/1000, t2).Format("20060102150405"), t1%1000, rand.Intn(999999-600000)+600000)
//	default:
//		// 返回微秒级时间戳
//		return fmt.Sprintf("%.0f", float64(t1)+float64(t2)/1000000)
//	}
//}
