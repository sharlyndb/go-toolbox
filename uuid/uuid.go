/**
 * @Time: 2022/3/4 15:20
 * @Author: yt.yin
 */

package uuid

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

const (
	NumSource = "0123456789"
	HexSource = "ABCDEF0123456789"
	StrSource = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
)

// RandomStr 随机一个字符串
func RandomStr(length int) string {
	var sb strings.Builder
	if length > 0 {
		for i := 0; i < length; i++ {
			sb.WriteString(string(StrSource[rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(StrSource))]))
		}
	}
	return sb.String()
}

// RandomNum 随机一个数字字符串
func RandomNum(length int) string {
	var sb strings.Builder
	if length > 0 {
		for i := 0; i < length; i++ {
			sb.WriteString(string(NumSource[rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(NumSource))]))
		}
	}
	return sb.String()
}

// RandomHex 随机一个hex字符串
func RandomHex(bytesLen int) string {
	var sb strings.Builder
	for i := 0; i < bytesLen<<1; i++ {
		sb.WriteString(string(HexSource[rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(HexSource))]))
	}
	return sb.String()
}

// SameSubStr 创建多少个相同 子字符串的字符串
func SameSubStr(subStr string, repeat int) string {
	var sb strings.Builder
	for i := 0; i < repeat; i++ {
		sb.WriteString(subStr)
	}
	return sb.String()
}

// UUID 生成uuid
func UUID() string {
	id := uuid.New()
	return strings.ReplaceAll(id.String(), "-", "")
}

// UniqueID 根据指定字段生成 uuid
func UniqueID(fields ...interface{}) string {
	if len(fields) == 0 {
		return UUID()
	}
	var buf strings.Builder
	for i := range fields {
		field := fields[i]
		buf.WriteString(asString(field))
	}
	s := strings.TrimSpace(buf.String())
	if s == "" {
		return UUID()
	}
	return Md5(s)
}

// Md5 md5加密
func Md5(src string) string {
	m := md5.New()
	m.Write([]byte(src))
	res := hex.EncodeToString(m.Sum(nil))
	return res
}

// 其他类型转String
func asString(src interface{}) string {
	switch v := src.(type) {
	case string:
		return v
	case []byte:
		return string(v)
	case int:
		return strconv.Itoa(v)
	case int32:
		return strconv.FormatInt(int64(v), 10)
	case int64:
		return strconv.FormatInt(v, 10)
	case float32:
		return strconv.FormatFloat(float64(v), 'f', -1, 64)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	case time.Time:
		return time.Time.Format(v, "2006-01-02 15:04:05")
	case bool:
		return strconv.FormatBool(v)
	default:
		{
			b, _ := json.Marshal(v)
			return string(b)
		}
	}
}
