/*
 * 抖音开放API
 *
 * douyin open api
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package douyin

// UserinfoRspData struct for UserinfoRspData
type UserinfoRspData struct {
	Province string `json:"province,omitempty"`
	// 用户在当前开发者账号下的唯一标识（未绑定开发者账号没有该字段）
	UnionId string `json:"union_id,omitempty"`
	Avatar  string `json:"avatar,omitempty"`
	// 类型 `EAccountM` - 普通企业号 `EAccountS` - 认证企业号 `EAccountK` - 品牌企业号
	EAccountRole string `json:"e_account_role,omitempty"`
	// 错误码
	ErrorCode int64 `json:"error_code,omitempty"`
	// 用户在当前应用的唯一标识
	OpenId   string `json:"open_id,omitempty"`
	Nickname string `json:"nickname,omitempty"`
	City     string `json:"city,omitempty"`
	Country  string `json:"country,omitempty"`
	// 错误码描述
	Description string `json:"description,omitempty"`
	// 性别`0` - 未知 `1` - 男性 `2` - 女性
	Gender int64 `json:"gender,omitempty"`
}
