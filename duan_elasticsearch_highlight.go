package duan_elasticsearch

type Highlight struct {
	PreTags  string                            `json:"pre_tags,omitempty"`
	PostTags string                            `json:"post_tags,omitempty"`
	Fields   map[string]map[string]interface{} `json:"fields,omitempty"`
}
