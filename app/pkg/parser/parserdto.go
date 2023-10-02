package parser

// import (
// 	"time"
// )

type ParserPayload struct {
	Data     interface{} `json:"data"`
	ConfigId string      `json:"ConfigId"`
}
