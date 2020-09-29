/*
 * 抖音开放API
 *
 * douyin open api
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package douyin

// DevtoolMicappIsLegalRspData struct for DevtoolMicappIsLegalRspData
type DevtoolMicappIsLegalRspData struct {
	// 错误描述
	Description string `json:"description,omitempty"`
	// 返回错误码
	ErrorCode int64 `json:"error_code,omitempty"`
	// 是否合法的信息
	IsLegal bool `json:"is_legal,omitempty"`
}
