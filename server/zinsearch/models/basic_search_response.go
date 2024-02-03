package models

type SearchResponse struct {
	Took     int  `json:"took"`
	TimedOut bool `json:"timed_out"`
	Shards   struct {
		Total       int `json:"total"`
		Succcessful int `json:"succcessful"`
		Skipped     int `json:"skipped"`
		Failed      int `json:"failed"`
	}
	Hits struct {
		Total struct {
			Value int `json:"value"`
		} `json:"total"`
		MaxScore float64 `json:"max_score"`
		Hits     []struct {
			Index     string  `json:"_index"`
			Type      string  `json:"_type"`
			ID        string  `json:"_id"`
			Score     float64 `json:"_score"`
			Timestamp string  `json:"@timestamp"`
			Source    struct {
				Subject  string `json:"Subject"`
				BodyData string `json:"bodyData"`
				From     string `json:"From"`
				To       string `json:"To"`
				Date     string `json:"Date"`
			} `json:"_source"`
		} `json:"hits"`
	}
}
