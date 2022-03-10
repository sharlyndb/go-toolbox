/**
 * @Time: 2022/3/6 00:46
 * @Author: yt.yin
 */

package request

// ID 接收单个请求体json格式的id
type ID struct {

	/** 主键id */
	ID            string     `json:"id"            form:"id"`

	/** 表名 */
	TableName     string     `json:"tableName"     form:"tableName"`
}

// NumIds 接收数字型的id数组
type NumIds struct {

	/** 主键id集合 */
	Ids           []int      `json:"ids"           form:"ids"`

	/** 表名 */
	TableName     string     `json:"tableName"     form:"tableName"`
}

// StrIds 字符串型的id数组
type StrIds struct {

	Ids           []string   `json:"ids"           form:"ids"`

	/** 表名 */
	TableName     string     `json:"tableName"     form:"tableName"`
}

// TableReq 接收单个请求体json格式的表名
type TableReq struct {
	TableName     string     `json:"tableName"     form:"tableName"`
}