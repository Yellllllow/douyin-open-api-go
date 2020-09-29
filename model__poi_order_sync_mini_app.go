/*
 * 抖音开放API
 *
 * douyin open api
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package douyin

// PoiOrderSyncMiniApp 小程序形式对接抖音时，该字段必传
type PoiOrderSyncMiniApp struct {
	// 外部商户ID
	UserOpenId string `json:"user_open_id,omitempty"`
	// 抖音商户ID
	AppId string `json:"app_id,omitempty"`
}
