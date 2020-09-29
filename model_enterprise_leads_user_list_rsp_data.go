/*
 * 抖音开放API
 *
 * douyin open api
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package douyin

// EnterpriseLeadsUserListRspData struct for EnterpriseLeadsUserListRspData
type EnterpriseLeadsUserListRspData struct {
	// 错误码
	ErrorCode int64 `json:"error_code,omitempty"`
	// 错误码描述
	Description string                                `json:"description,omitempty"`
	Cursor      int64                                 `json:"cursor,omitempty"`
	HasMore     bool                                  `json:"has_more,omitempty"`
	List        []string                              `json:"list,omitempty"`
	Users       []EnterpriseLeadsUserListRspDataUsers `json:"users,omitempty"`
}
