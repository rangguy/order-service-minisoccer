package dto

type KafkaEvent struct {
	Name string `json:"name"`
}

type KafkaMetaData struct {
	Sender    string `json:"sender"`
	SendingAt string `json:"sendingAt"`
}

type DataType string

type KafkaBody[T any] struct {
	Type string `json:"type"`
	Data T      `json:"data"`
}

type KafkaMessage[T any] struct {
	Event    KafkaEvent    `json:"event"`
	Metadata KafkaMetaData `json:"metadata"`
	Body     KafkaBody[T]  `json:"body"`
}
