/*
 * 抖音开放API
 *
 * douyin open api
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package douyin

// PoiExtHotelOrderStatusRspData struct for PoiExtHotelOrderStatusRspData
type PoiExtHotelOrderStatusRspData struct {
	// 金额
	Balance int64 `json:"balance,omitempty"`
	// 附加信息
	Ext string `json:"ext,omitempty"`
	// 错误描述
	Description string `json:"description,omitempty"`
	// 返回错误码
	ErrorCode int64 `json:"error_code,omitempty"`
	// 订单确认状态。0 - 订单确认, 1 - 价格变动, 2 - 库存不足, 3 - 确认中
	Status int64 `json:"status,omitempty"`
}
