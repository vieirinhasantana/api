package model

import "time"

type Response struct {
	MetaData MetaData `json:"metadata"`
	Message  Message  `json:data`
}

type MetaData struct {
	IDRequest string    `json:"id_request"`
	Date      time.Time `json:"date"`
	Status    bool      `json:"status"`
}

type Message struct {
	Response string `json:"response"`
}
