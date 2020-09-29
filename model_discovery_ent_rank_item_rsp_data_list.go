/*
 * 抖音开放API
 *
 * douyin open api
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package douyin

// DiscoveryEntRankItemRspDataList struct for DiscoveryEntRankItemRspDataList
type DiscoveryEntRankItemRspDataList struct {
	Actors        []string `json:"actors,omitempty"`
	Areas         []string `json:"areas,omitempty"`
	Directors     []string `json:"directors,omitempty"`
	DiscussionHot int64    `json:"discussion_hot,omitempty"`
	Hot           int64    `json:"hot,omitempty"`
	Id            string   `json:"id,omitempty"`
	InfluenceHot  int64    `json:"influence_hot,omitempty"`
	MaoyanId      string   `json:"maoyan_id,omitempty"`
	Name          string   `json:"name,omitempty"`
	NameEn        string   `json:"name_en,omitempty"`
	Poster        string   `json:"poster,omitempty"`
	ReleaseDate   string   `json:"release_date,omitempty"`
	SearchHot     int64    `json:"search_hot,omitempty"`
	Tags          []string `json:"tags,omitempty"`
	TopicHot      int64    `json:"topic_hot,omitempty"`
	Type          int32    `json:"type,omitempty"`
}