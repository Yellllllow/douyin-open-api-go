/*
 * 抖音开放API
 *
 * douyin open api
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package douyin

// DouyinVideoListRspDataStatistics struct for DouyinVideoListRspDataStatistics
type DouyinVideoListRspDataStatistics struct {
	CommentCount  int64 `json:"comment_count,omitempty"`
	DiggCount     int64 `json:"digg_count,omitempty"`
	DownloadCount int64 `json:"download_count,omitempty"`
	ForwardCount  int64 `json:"forward_count,omitempty"`
	PlayCount     int64 `json:"play_count,omitempty"`
	ShareCount    int64 `json:"share_count,omitempty"`
}
