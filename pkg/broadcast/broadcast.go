package broadcast

import (
	"encoding/json"
	"time"
)

// Broadcast : Struct that indicates a broadcast message with variaty of information.
type Broadcast struct {
	Message     string        `json:"message"`
	Information []interface{} `json:"info,omitempty"`
}

func (o Broadcast) MarshalJSON() ([]byte, error) {
	type AliasBroadcast Broadcast
	return json.Marshal(&struct {
		*AliasBroadcast
		TimeStamp string `json:"timestamp"`
	}{
		AliasBroadcast: (*AliasBroadcast)(&o),
		TimeStamp:      time.Now().Local().Format("2006-01-02 15:04:05"),
	})
}
