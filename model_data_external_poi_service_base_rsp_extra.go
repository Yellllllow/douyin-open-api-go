/*
 * 抖音开放API
 *
 * douyin open api
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package douyin

// DataExternalPoiServiceBaseRspExtra struct for DataExternalPoiServiceBaseRspExtra
type DataExternalPoiServiceBaseRspExtra struct {
	// 子错误码描述
	SubDescription string `json:"sub_description,omitempty"`
	// 子错误码
	SubErrorCode int64 `json:"sub_error_code,omitempty"`
	// 错误码描述
	Description string `json:"description,omitempty"`
	// 错误码
	ErrorCode int64 `json:"error_code,omitempty"`
	// 标识请求的唯一id
	Logid string `json:"logid,omitempty"`
	// 毫秒级时间戳
	Now int64 `json:"now,omitempty"`
}
