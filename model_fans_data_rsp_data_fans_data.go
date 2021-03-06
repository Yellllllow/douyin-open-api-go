/*
 * 抖音开放API
 *
 * douyin open api
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package douyin

// FansDataRspDataFansData struct for FansDataRspDataFansData
type FansDataRspDataFansData struct {
	ActiveDaysDistributions   []FansDataRspDataFansDataActiveDaysDistributions `json:"active_days_distributions,omitempty"`
	AgeDistributions          []FansDataRspDataFansDataActiveDaysDistributions `json:"age_distributions,omitempty"`
	AllFansNum                int64                                            `json:"all_fans_num,omitempty"`
	DeviceDistributions       []FansDataRspDataFansDataActiveDaysDistributions `json:"device_distributions,omitempty"`
	FlowContributions         []FansDataRspDataFansDataFlowContributions       `json:"flow_contributions,omitempty"`
	GenderDistributions       []FansDataRspDataFansDataActiveDaysDistributions `json:"gender_distributions,omitempty"`
	GeographicalDistributions []FansDataRspDataFansDataActiveDaysDistributions `json:"geographical_distributions,omitempty"`
	InterestDistributions     []FansDataRspDataFansDataActiveDaysDistributions `json:"interest_distributions,omitempty"`
}
