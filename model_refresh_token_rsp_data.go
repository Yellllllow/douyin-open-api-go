/*
 * 抖音开放API
 *
 * douyin open api
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package douyin

// RefreshTokenRspData struct for RefreshTokenRspData
type RefreshTokenRspData struct {
	// 接口调用凭证
	AccessToken string `json:"access_token,omitempty"`
	// 错误码描述
	Description string `json:"description,omitempty"`
	// 错误码
	ErrorCode int64 `json:"error_code,omitempty"`
	// 过期时间，单位（秒
	ExpiresIn int64 `json:"expires_in,omitempty"`
	// 用户刷新
	RefreshToken string `json:"refresh_token,omitempty"`
	// 当前应用下，授权用户唯一标识
	OpenId string `json:"open_id,omitempty"`
	// 用户授权的作用域
	Scope string `json:"scope,omitempty"`
}
