package error

type BaseResponse struct {
	Data    interface{} `json:"data"`
	Success int         `json:"success"`
	Message string      `json:"message"`
}

//StandardMappedResponse standardize response
type StandardMappedResponse struct {
	Data       interface{} `json:"data"`
	Success    bool        `json:"success"`
	Message    string      `json:"message"`
	Code       int         `json:"code"`
	StatusCode int         `json:"status_code"`
}

//ValidationErrorsResponse pattern errors
type ValidationErrorsResponse struct {
	Field string `json:"field"`
	Code  int    `json:"code"`
	Msg   string `json:"msg"`
}
