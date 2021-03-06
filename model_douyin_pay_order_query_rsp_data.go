/*
 * 抖音开放API
 *
 * douyin open api
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package douyin

// DouyinPayOrderQueryRspData struct for DouyinPayOrderQueryRspData
type DouyinPayOrderQueryRspData struct {
	// 金额
	Amount int64 `json:"amount,omitempty"`
	// 外部订单号
	BizOrderNo string `json:"biz_order_no,omitempty"`
	// 订单描述
	OrderDesc string `json:"order_desc,omitempty"`
	// 订单名称
	OrderName string `json:"order_name,omitempty"`
	// 订单状态
	OrderStatus int64 `json:"order_status,omitempty"`
	// 开始时间
	PayCreateTime int64 `json:"pay_create_time,omitempty"`
	// 结束时间
	PayFinishTime int64 `json:"pay_finish_time,omitempty"`
	// 内部订单号
	PayOrderNo string `json:"pay_order_no,omitempty"`
	// 标记，长度小于512
	Remark string `json:"remark,omitempty"`
	// 返回码
	RetCode string `json:"ret_code,omitempty"`
	// 返回信息
	RetMsg string `json:"ret_msg,omitempty"`
	// 返回码
	SubCode string `json:"sub_code,omitempty"`
	// 返回信息
	SubMsg string `json:"sub_msg,omitempty"`
	// 错误描述
	Description string `json:"description,omitempty"`
	// 返回错误码
	ErrorCode int64 `json:"error_code,omitempty"`
}
