/*
 * 抖音开放API
 *
 * douyin open api
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package douyin

// AuthorizeV2Rsp
type AuthorizeV2Rsp struct {
	Data    AuthorizeV2RspData `json:"data,omitempty"`
	Message string             `json:"message,omitempty"`
	Extra   ConnectRspExtra    `json:"extra,omitempty"`
}
