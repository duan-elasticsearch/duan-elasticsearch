package duan_elasticsearch

type DuanElasticsearch struct {
	Query *Query `json:"query,omitempty"`

	From string `json:"from,omitempty"`
	Size string `json:"size,omitempty"`

	// 准确count使用
	TrackTotalHits bool `json:"track_total_hits,omitempty"`

	// 未实现的预留接口
	Aggs   interface{} `json:"aggs,omitempty"`
	Source interface{} `json:"_source,omitempty"`
	Sort   interface{} `json:"sort,omitempty"`
}
