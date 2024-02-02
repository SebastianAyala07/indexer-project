package models

type Search struct {
	SearchType string `json:"search_type"`
	Query      struct {
		Term  string `json:"term"`
		Field string `json:"field"`
	} `json:"query"`
	SortFields []string `json:"sort_fields"`
	From       int      `json:"from"`
	MaxResults int      `json:"max_results"`
	Source     []string `json:"_source"`
}
