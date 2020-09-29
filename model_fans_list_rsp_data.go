/*
 * 抖音开放API
 *
 * douyin open api
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package douyin

// FansListRspData struct for FansListRspData
type FansListRspData struct {
	// 错误码
	ErrorCode int64 `json:"error_code,omitempty"`
	HasMore   bool  `json:"has_more,omitempty"`
	// 粉丝总数
	Total  int64 `json:"total,omitempty"`
	Cursor int64 `json:"cursor,omitempty"`
	// 错误码描述
	Description string `json:"description,omitempty"`
	List        []User `json:"list,omitempty"`
}