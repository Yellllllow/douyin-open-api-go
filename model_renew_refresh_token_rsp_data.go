/*
 * 抖音开放API
 *
 * douyin open api
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package douyin

// RenewRefreshTokenRspData struct for RenewRefreshTokenRspData
type RenewRefreshTokenRspData struct {
	// 错误码描述
	Description string `json:"description,omitempty"`
	// 错误码
	ErrorCode int64 `json:"error_code,omitempty"`
	// 过期时间，单位（秒
	ExpiresIn int64 `json:"expires_in,omitempty"`
	// 用户刷新
	RefreshToken string `json:"refresh_token,omitempty"`
}
