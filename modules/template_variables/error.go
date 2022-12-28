package template_variables

type Error struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
	Reason  string `json:"reason"`
}
