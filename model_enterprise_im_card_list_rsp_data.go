/*
 * 抖音开放API
 *
 * douyin open api
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package douyin

// EnterpriseImCardListRspData struct for EnterpriseImCardListRspData
type EnterpriseImCardListRspData struct {
	// 错误码
	ErrorCode int64 `json:"error_code,omitempty"`
	// 错误码描述
	Description string                             `json:"description,omitempty"`
	Cursor      int64                              `json:"cursor,omitempty"`
	HasMore     bool                               `json:"has_more,omitempty"`
	Cards       []EnterpriseImCardListRspDataCards `json:"cards,omitempty"`
}