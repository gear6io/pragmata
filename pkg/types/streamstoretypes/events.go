package streamstoretypes

import "time"

// Event is a single ingested data record with a nanosecond-precision timestamp.
type Event struct {
	ObservedAt int64          `json:"observedAt"` // unix nanoseconds
	DataSource string         `json:"dataSource"`
	Data       map[string]any `json:"data"`
}

// NewEvent creates an Event stamped at the current time.
func NewEvent(dataSource string, data map[string]any) *Event {
	return &Event{
		ObservedAt: time.Now().UnixNano(),
		DataSource: dataSource,
		Data:       data,
	}
}
