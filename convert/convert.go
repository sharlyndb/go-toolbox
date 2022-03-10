/**
 * @Time: 2022/3/10 14:50
 * @Author: yt.yin
 */

package convert

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

// BytesToHex 字节数组转hex
// []byte{0x55, 0xAA} 被转成 55AA
func BytesToHex(data []byte) string {
	return strings.ToUpper(hex.EncodeToString(data))
}

// HexToBytes 将hex 字符串转成 byte数组
// AABBCC 转成字节数组 []byte{0xAA, 0xBB, 0xCC}
func HexToBytes(hexStr string) ([]byte) {
	decodeString, _ := hex.DecodeString(hexStr)
	return decodeString
}

// HexBCC 计算BCC校验码
func HexBCC(hexStr string) string {
	hexToBytes := HexToBytes(hexStr)
	length := len(hexToBytes)
	if length < 1 {
		return ""
	}
	bcc := hexToBytes[0]
	if length > 1 {
		for i := 1; i < length; i++ {
			bcc = bcc ^ hexToBytes[i]
		}
	}
	return BytesToHex([]byte{bcc & 0xFF})
}

// BytesBCC 计算 BCC
func BytesBCC(bytes []byte) byte {
	bcc := bytes[0]
	if len(bytes) > 1 {
		for i := 1; i < len(bytes); i++ {
			bcc = bcc ^ bytes[i]
		}
	}
	return bcc & 0xFF
}

// DecToHex 十进进制转16进制
func DecToHex(n uint64) string {
	s := strconv.FormatUint(n, 16)
	s = strings.ToUpper(s)
	length := len(s)
	if length % 2 ==1 {
		s = "0"+s
	}
	return s
}

// HexToDec 十六进制转10进制
func HexToDec(h string) uint64 {
	n, err := strconv.ParseUint(h, 16, 64)
	if err != nil{
		return 0
	}
	return n
}

// DecToBin 十进制转二进制
func DecToBin(n uint64) string{
	s := strconv.FormatUint(n, 2)
	length:= len(s)
	mod := length % 8
	if mod != 0 {
		prefixNum := 8-mod
		var sb strings.Builder
		for i := 0;i < prefixNum; i++ {
			sb.WriteString("0")
		}
		s = sb.String()+s
	}
	return s
}

// HexToBin 16进制转二进制
func HexToBin(h string) string {
	n, err := strconv.ParseUint(h, 16, 64)
	if err != nil{
		return ""
	}
	return DecToBin(n)
}

// ByteToBinStr 将byte 以8个bit位的形式展示
func ByteToBinStr(b byte) string {
	return fmt.Sprintf("%08b", b)
}

// BytesToBinStr 将byte数组转成8个bit位一组的字符串
func BytesToBinStr(bs []byte) string {
	if len(bs) <= 0 {
		return ""
	}
	buf := bytes.NewBuffer([]byte{})
	for _, v := range bs {
		buf.WriteString(fmt.Sprintf("%08b", v))
	}
	return buf.String()
}

// BytesToBinStrWithSplit 将byte数组转8个bit一组的字符串并且带分割符
func BytesToBinStrWithSplit(bs []byte,split string) string {
	length := len(bs)
	if length <= 0 {
		return ""
	}
	buf := bytes.NewBuffer([]byte{})
	for i := 0; i < length-1; i++ {
		v := bs[i]
		buf.WriteString(fmt.Sprintf("%08b", v))
		buf.WriteString(split)
	}
	buf.WriteString(fmt.Sprintf("%08b",bs[length-1]))
	return buf.String()
}

// HexSuffixZero hex 后补位
func HexSuffixZero(hex string, byteSize int) string {
	data1 := HexToBytes(hex)
	data2 := make([]byte, byteSize)
	copy(data2, data1)
	return BytesToHex(data2)
}

func HexPrefixZero(hex string, byteSize int) string {
	data1 := HexToBytes(hex)
	data2 := make([]byte, byteSize-len(data1))
	for _, v := range data1 {
		data2 = append(data2, v)
	}
	return BytesToHex(data2)
}

// GBKSuffixZero GBK 编码按字节右补0
func GBKSuffixZero(gbkStr string, byteSize int) string {
	data1, _ := ioutil.ReadAll(transform.NewReader(bytes.NewReader([]byte(gbkStr)), simplifiedchinese.GBK.NewEncoder()))
	data2 := make([]byte, byteSize)
	copy(data2, data1)
	return BytesToHex(data2)
}

// GBKSuffixSpace 编码按字节右补空格
func GBKSuffixSpace(chinese string, byteSize int) (hex string) {
	data1, _ := ioutil.ReadAll(transform.NewReader(bytes.NewReader([]byte(chinese)), simplifiedchinese.GBK.NewEncoder()))
	data2 := make([]byte, byteSize)
	copy(data2, data1)
	for i:= len(data1);i < len(data2);i++ {
		data2[i]= 0x20
	}
	return string(data2)
}

// ReverseString 反转字符串
func ReverseString(s string) string {
	runes := []rune(s)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}

// StrSuffixZero 字符串后补0
func StrSuffixZero(str string, length int) string {
	if len(str) == length {
		return str
	}
	if len(str) > length{
		return str[:length]
	}
	var sb strings.Builder
	sb.WriteString(str)
	for i := 0; i < length - len(str); i++{
		sb.WriteString("0")
	}
	return sb.String()
}

// StrSuffixF 字符串后补F
func StrSuffixF(str string, length int) string {
	if len(str) == length {
		return str
	}
	if len(str) > length{
		return str[:length]
	}
	var sb strings.Builder
	sb.WriteString(str)
	for i := 0; i < length - len(str); i++{
		sb.WriteString("F")
	}
	return sb.String()
}

// StrPrefixZero 字符串前补0
func StrPrefixZero(str string, length int) string {
	if len(str) == length {
		return str
	}
	if len(str) > length{
		return str[:length]
	}
	var sb strings.Builder
	for i := 0; i < length - len(str); i++{
		sb.WriteString("0")
	}
	sb.WriteString(str)
	return sb.String()
}


// StrPrefixSpace 字符串前补空格
func StrPrefixSpace(str string, length int) string {
	if len(str) == length {
		return str
	}
	if len(str) > length{
		return str[:length]
	}
	var sb strings.Builder
	for i := 0; i < length - len(str); i++{
		sb.WriteString(" ")
	}
	sb.WriteString(str)
	return sb.String()
}

// StrSuffixSpace 字符串后补空格
func StrSuffixSpace(str string, length int) string {
	if len(str) == length {
		return str
	}
	if len(str) > length{
		return str[:length]
	}
	var sb strings.Builder
	sb.WriteString(str)
	for i := 0; i < length - len(str); i++{
		sb.WriteString(" ")
	}
	return sb.String()
}

// HexReverse 字节颠倒
func HexReverse(hex string) string {
	toBytes := HexToBytes(hex)
	length := len(toBytes)
	if length <= 1 {
		return hex
	}
	for i := range toBytes {
		a := toBytes[i]
		toBytes[i] = toBytes[length-1-i]
		toBytes[length-1-i] = a
		if i == length/2{
			break
		}
	}
	return BytesToHex(toBytes)
}

// AsString 其他类型转String
func AsString(src interface{}) string {
	switch v := src.(type) {
	case string:
		return v
	case []byte:
		return string(v)
	case int:
		return strconv.Itoa(v)
	case int8:
		return strconv.Itoa(int(v))
	case int16:
		return strconv.Itoa(int(v))
	case int32:
		return strconv.FormatInt(int64(v), 10)
	case int64:
		return strconv.FormatInt(v, 10)
	case uint:
		return strconv.FormatUint(uint64(v), 10)
	case uint8:
		return strconv.FormatUint(uint64(v), 10)
	case uint16:
		return strconv.FormatUint(uint64(v), 10)
	case uint32:
		return strconv.FormatUint(uint64(v), 10)
	case uint64:
		return strconv.FormatUint(v, 10)
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



