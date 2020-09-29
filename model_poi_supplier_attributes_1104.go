/*
 * 抖音开放API
 *
 * douyin open api
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package douyin

// PoiSupplierAttributes1104 酒店政策
type PoiSupplierAttributes1104 struct {
	// 离店时间(HH:mm)
	CheckOutTime string `json:"check_out_time,omitempty"`
	// 儿童政策(<=500字)
	Child string `json:"child,omitempty"`
	// 宠物政策(<=500字)
	Pet string `json:"pet,omitempty"`
	// 入住时间(HH:mm)
	CheckInTime string                             `json:"check_in_time,omitempty"`
	Breakfast   PoiSupplierAttributes1104Breakfast `json:"breakfast,omitempty"`
}
