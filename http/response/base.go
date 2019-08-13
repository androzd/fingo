package response

type ErrorResponse struct {
	Errors map[string]string `json:"errors"`
}

type SuccessResponse struct {
	Status string `json:"status"`
	Data struct{} `json:"data"`
}