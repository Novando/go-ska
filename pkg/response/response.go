package response

type StdResponse struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

type CounterResponse struct {
	Count   int         `json:"count"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}
