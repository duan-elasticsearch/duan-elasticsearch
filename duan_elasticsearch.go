package duan_elasticsearch

type DuanElasticsearch struct {
	Query *Query `json:"query,omitempty"`

	From string `json:"from,omitempty"`
	Size string `json:"size,omitempty"`

	// 未实现的预留接口
	Aggs interface{} `json:"aggs,omitempty"`
	Source interface{} `json:"source,omitempty"`
	Sort interface{} `json:"sort,omitempty"`
}
