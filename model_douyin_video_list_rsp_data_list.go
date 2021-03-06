/*
 * 抖音开放API
 *
 * douyin open api
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package douyin

// DouyinVideoListRspDataList struct for DouyinVideoListRspDataList
type DouyinVideoListRspDataList struct {
	Cover       string                           `json:"cover,omitempty"`
	CreateTime  int64                            `json:"create_time,omitempty"`
	IsReviewed  bool                             `json:"is_reviewed,omitempty"`
	IsTop       bool                             `json:"is_top,omitempty"`
	ItemId      string                           `json:"item_id,omitempty"`
	ShareUrl    string                           `json:"share_url,omitempty"`
	Title       string                           `json:"title,omitempty"`
	VideoStatus int64                            `json:"video_status,omitempty"`
	Statistics  DouyinVideoListRspDataStatistics `json:"statistics,omitempty"`
}
