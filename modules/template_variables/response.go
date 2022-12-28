package template_variables

type Response struct {
	SessionID  string `json:"session_id"`
	ID         string `json:"id"`
	StatusCode int    `json:"response_code"`
	Error      Error  `json:"error"`
}
type FailureResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
	Message string `json:"message"`
}

type AuctionResponse struct {
	ResultCode int    `json:"result_code"`
	ResultText string `json:"result_text"`
}
type CreateResponse struct {
	ID string `json:"id"`
}
type EmptyResponse struct{}
type ExistsResponse struct {
	Exists bool `json:"exists"`
}

type SuccessResponse struct {
	Success bool `json:"success" example:"true"`
}
type ResponseError struct {
	Error interface{}
}
type InternalServerError struct {
	Code    string
	Message string
}
