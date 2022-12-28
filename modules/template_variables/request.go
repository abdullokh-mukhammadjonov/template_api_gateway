package template_variables

type GetRequest struct {
	ID string `json:"id"`
}
type GetAllRequest struct {
	Limit uint32
	Page  uint32
}

type DeleteRequest struct {
	ID        string `json:"id"`
	SessionID string `json:"session_id"`
}
