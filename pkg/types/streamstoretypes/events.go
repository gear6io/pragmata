package streamstoretypes

import "time"

type Event struct {
	ObservedAt int64          `json:"observedAt"` // unix timestamp in nanoseconds
	DataSource string         `json:"dataSource"`
	Data       map[string]any `json:"data"`
}

func NewEvent(dataSource string, data map[string]any) *Event {
	return &Event{
		ObservedAt: time.Now().UnixNano(),
		DataSource: dataSource,
		Data:       data,
	}
}
