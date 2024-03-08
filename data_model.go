package arrowdb

type Row struct {
	Metric    string            `json:"metric"`
	Tags      map[string]string `json:"tags"`
	Timestamp int64             `json:"timestamp"`
	Value     float64           `json:"value"`
}

type DataPoint struct {
	Timestamp int64   `json:"ts"`
	Value     float64 `json:"v"`
}

// SeriesDataPoints return from a range query
type SeriesDataPoints struct {
	Metric     string            `json:"metric"`
	Tags       map[string]string `json:"tags"`
	DataPoints []DataPoint       `json:"data_points"`
}

// TopNData return from a topN query
type TopNData struct {
	Metric      string            `json:"metric"`
	Tags        map[string]string `json:"tags"`
	GroupByName string            `json:"name"`
	Value       float64           `json:"value"`
}

type TagMatcher struct {
	Name     string
	Value    string
	IsRegexp bool // if support regexp
}

type TagMatcherSet []TagMatcher
