/*
 * 抖音开放API
 *
 * douyin open api
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package douyin

// RenewRefreshTokenRsp
type RenewRefreshTokenRsp struct {
	Data    RenewRefreshTokenRspData `json:"data,omitempty"`
	Message string                   `json:"message,omitempty"`
	Extra   ConnectRspExtra          `json:"extra,omitempty"`
}
