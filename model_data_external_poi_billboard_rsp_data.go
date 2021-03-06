/*
 * 抖音开放API
 *
 * douyin open api
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package douyin

// DataExternalPoiBillboardRspData struct for DataExternalPoiBillboardRspData
type DataExternalPoiBillboardRspData struct {
	// 错误描述
	Description string `json:"description,omitempty"`
	// 返回错误码
	ErrorCode  int64                                       `json:"error_code,omitempty"`
	ResultList []DataExternalPoiBillboardRspDataResultList `json:"result_list,omitempty"`
}
