package jetstream

import (
	"time"
)

type GlobalJetStream struct {
	Streams        int              `json:"streams"`
	Consumers      int              `json:"consumers"`
	Messages       int              `json:"messages"`
	Bytes          int              `json:"bytes"`
	AccountDetails []AccountDetails `json:"account_details"`
}

type AccountDetails struct {
	Name         string         `json:"name"`
	StreamDetail []StreamDetail `json:"stream_detail"`
}

type StreamDetail struct {
	Name           string           `json:"name"`
	State          StreamState      `json:"state"`
	ConsumerDetail []ConsumerDetail `json:"consumer_detail"`
}

type StreamState struct {
	Messages      int       `json:"messages"`
	Bytes         int       `json:"bytes"`
	FirstSeq      int       `json:"first_seq"`
	LastSeq       int       `json:"last_seq"`
	FirstTS       time.Time `json:"first_ts"`
	LastTS        time.Time `json:"last_ts"`
	NumSubjects   int       `json:"num_subjects"`
	ConsumerCount int       `json:"consumer_count"`
}

type ConsumerDetail struct {
	StreamName string    `json:"stream_name"`
	Name       string    `json:"name"`
	Created    time.Time `json:"created"`
}
