/*
 * 抖音开放API
 *
 * douyin open api
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package douyin

// VideoSearchCommentReplyListRsp
type VideoSearchCommentReplyListRsp struct {
	Data  ItemCommentListRspData    `json:"data,omitempty"`
	Extra FollowingListRspDataExtra `json:"extra,omitempty"`
}