package common

import "bytes"

type StdService struct {
	Code    int
	Message string
	Data    interface{}
	Buf     *bytes.Buffer
}

type StdResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Count   *int        `json:"count"`
}