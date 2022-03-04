/**
 * @Time: 2022/3/4 19:02
 * @Author: yt.yin
 */

package page

import (
	"bytes"
	"log"
	"net/url"
	"strconv"
	"strings"
	"unicode"

	"github.com/gin-gonic/gin"
)

const (

	/** ------- and 条件 ------  */
	/** 小于 < */
	lt = "lt:"

	/** 大于 > */
	gt = "gt:"

	/** 小于 <= */
	lte = "lte:"

	/** 大于 >= */
	gte = "gte:"

	/** 默认是等于 */
	eq = "eq:"

	/** 模糊查询 */
	lk = "lk:"

	/** ------- or 条件 ------  */

	/** 小于 */
	orlt = "orlt:"

	/** 大于  */
	orgt = "orgt:"

	/** 小于 */
	orlte = "orlte:"

	/** 大于  */
	orgte = "orgte:"

	/** 默认是等于 */
	oreq = "oreq:"

	/** 模糊查询 */
	orlk = "orlk:"

	/** ------- 排序 ------  */

	/** 降序 */
	pd = ":pd:"

	/** 升序 */
	pa = ":pa:"
)

// PageBean 全局分页对象
type PageBean struct {

	/** 当前页  */
	Page int `json:"page"`

	/** 当前页的行数 */
	PageSize int `json:"pageSize"`

	/** 总记录数 */
	Total int64 `json:"total"`

	/** 每行的数据 */
	Rows interface{} `json:"rows"`
}

type PageInfo struct {

	/** 当前页 */
	Current int

	/** 每页显示的最大行数 */
	RowCount int

	/** 表名 仅限于指定表名去查询 */
	TableName string

	/** 查询 and 条件参数 */
	AndParams map[string]interface{}

	/** 查询 or 条件参数 */
	OrParams map[string]interface{}

	/** 排序 */
	OrderStr string
}

// PageParam 获取url查询参数
func PageParam(c *gin.Context) *PageInfo {
	s := c.Request.URL.RawQuery
	paramStr, err := url.QueryUnescape(s)
	if err != nil {
		log.Println("url参数decode异常：" + err.Error())
		return nil
	}
	pageInfo := PageInfo{}
	andParams := make(map[string]interface{})
	orParams := make(map[string]interface{})
	paramArr := strings.Split(paramStr, "&")
	for _, v := range paramArr {
		ky := strings.Split(v, "=")
		if len(ky) != 2 {
			continue
		}
		if ky[0] == "current" {
			current, err := strconv.Atoi(ky[1])
			if err != nil {
				current = 1
			}
			if current < 1 {
				current = 1
			}
			pageInfo.Current = current
			continue
		} else if ky[0] == "rowCount" {
			rowCount, err := strconv.Atoi(ky[1])
			if err != nil {
				rowCount = 10
			}
			if rowCount < 1 {
				rowCount = 10
			} else if rowCount > 100 {
				rowCount = 100
			}
			pageInfo.RowCount = rowCount
			continue
		} else if ky[0] == "orderStr" {
			pageInfo.OrderStr = ky[1]
			continue
		} else if ky[0] == "tableName" {
			pageInfo.TableName = ky[1]
			continue
		}
		key := ky[0]
		value := ky[1]
		if key == "_t" || key == "_time" || key == "_timestamp" {
			continue
		} else if strings.Index(value, lt) == 0 {
			value = strings.Replace(value, lt, "", 1)
			if value == "" {
				continue
			}
			key = CamelToCase(key)
			andParams[key+" < ?"] = value
			continue
		} else if strings.Index(value, lte) == 0 {
			value = strings.Replace(value, lte, "", 1)
			if value == "" {
				continue
			}
			key = CamelToCase(key)
			andParams[key+" <= ?"] = value
			continue
		} else if strings.Index(value, gt) == 0 {
			value = strings.Replace(value, gt, "", 1)
			if value == "" {
				continue
			}
			key = CamelToCase(key)
			andParams[key+" > ?"] = value
			continue
		} else if strings.Index(value, gte) == 0 {
			value = strings.Replace(value, gte, "", 1)
			if value == "" {
				continue
			}
			key = CamelToCase(key)
			andParams[key+" >= ?"] = value
			continue
		} else if strings.Index(value, lk) == 0 {
			value = strings.Replace(value, lk, "", 1)
			if value == "" {
				continue
			}
			key = CamelToCase(key)
			andParams[key+" LIKE ?"] = value + "%"
			continue
		} else if strings.Index(value, eq) == 0 {
			value = strings.Replace(value, eq, "", 1)
			if value == "" {
				continue
			}
			key = CamelToCase(key)
			andParams[key+" = ?"] = value
			continue
		} else if strings.Index(value, orlt) == 0 {
			value = strings.Replace(value, orlt, "", 1)
			if value == "" {
				continue
			}
			key = CamelToCase(key)
			orParams[key+" < ?"] = value
		} else if strings.Index(value, orlte) == 0 {
			value = strings.Replace(value, orlte, "", 1)
			if value == "" {
				continue
			}
			key = CamelToCase(key)
			orParams[key+" <= ?"] = value
		} else if strings.Index(value, orgt) == 0 {
			value = strings.Replace(value, orgt, "", 1)
			if value == "" {
				continue
			}
			key = CamelToCase(key)
			orParams[key+" > ?"] = value
		} else if strings.Index(value, orlk) == 0 {
			value = strings.Replace(value, oreq, "", 1)
			if value == "" {
				continue
			}
			key = CamelToCase(key)
			orParams[key+" LIKE ?"] = value + "%"
		} else if strings.Index(value, oreq) == 0 {
			value = strings.Replace(value, oreq, "", 1)
			if value == "" {
				continue
			}
			key = CamelToCase(key)
			orParams[key+" = ?"] = value
		} else {
			if value == "" {
				continue
			}
			key = CamelToCase(key)
			andParams[key+" = ?"] = value
		}
	}
	if pageInfo.OrderStr != "" {
		v := CamelToCase(pageInfo.OrderStr)
		v = strings.ReplaceAll(v, pd, " desc,")
		v = strings.ReplaceAll(v, pa, " asc,")
		v = strings.TrimSuffix(v, ",")
		pageInfo.OrderStr = v
	}
	pageInfo.AndParams = andParams
	pageInfo.OrParams = orParams
	return &pageInfo
}

func CamelToCase(name string) string {
	buffer := NewBuffer()
	for i, r := range name {
		if unicode.IsUpper(r) {
			if i != 0 {
				buffer.Append('_')
			}
			buffer.Append(unicode.ToLower(r))
		} else {
			buffer.Append(r)
		}
	}
	return buffer.String()
}

// Buffer 内嵌bytes.Buffer，支持连写
type Buffer struct {
	*bytes.Buffer
}

func NewBuffer() *Buffer {
	return &Buffer{Buffer: new(bytes.Buffer)}
}

func (b *Buffer) Append(i interface{}) *Buffer {
	switch val := i.(type) {
	case string:
		b.append(val)
	case int:
		b.append(strconv.Itoa(val))
	case int64:
		b.append(strconv.FormatInt(val, 10))
	case uint:
		b.append(strconv.FormatUint(uint64(val), 10))
	case uint64:
		b.append(strconv.FormatUint(val, 10))
	case []byte:
		_, _ = b.Write(val)
	case rune:
		_, _ = b.WriteRune(val)
	}
	return b
}

func (b *Buffer) append(s string) *Buffer {
	defer func() {
		if err := recover(); err != nil {
			log.Println("*****内存不够了！******")
		}
	}()
	_, _ = b.WriteString(s)
	return b
}

// CheckPageRows 获取页数和行数
func CheckPageRows(currentStr, rowCountStr string) (current, rowCount int) {
	current, err := strconv.Atoi(currentStr)
	if err != nil {
		current = 1
	}
	if current < 1 {
		current = 1
	}
	rowCount, err = strconv.Atoi(rowCountStr)
	if err != nil {
		rowCount = 10
	}
	if rowCount < 1 {
		rowCount = 10
	} else if rowCount > 500 {
		rowCount = 500
	}
	return current, rowCount
}
