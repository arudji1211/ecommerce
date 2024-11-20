package response

type ResponseSuccess struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}