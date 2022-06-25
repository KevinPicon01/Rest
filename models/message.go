package models

type WebSocketMessage struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}
