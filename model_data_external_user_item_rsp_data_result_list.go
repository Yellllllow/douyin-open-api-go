/*
 * 抖音开放API
 *
 * douyin open api
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package douyin

// DataExternalUserItemRspDataResultList struct for DataExternalUserItemRspDataResultList
type DataExternalUserItemRspDataResultList struct {
	// 日期
	Date     string `json:"date,omitempty"`
	NewIssue int64  `json:"new_issue,omitempty"`
	// 每天新增视频播放
	NewPlay int64 `json:"new_play,omitempty"`
	// 每日内容总数
	TotalIssue int64 `json:"total_issue,omitempty"`
}
