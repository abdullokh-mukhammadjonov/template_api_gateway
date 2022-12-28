package var_content_service

type Handbook struct {
	Active      bool   `json:"active"`
	Title       string `json:"title"`
	UrlLink     string `json:"url_link"`
	Description string `json:"description"`
	Lang        string `json:"lang"`
	Date        string `json:"date"`
	AccessType  string `json:"access_type"`
}

type HandbookSlim struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Slug        string `json:"slug"`
	UrlLink     string `json:"url_link"`
	Description string `json:"description"`
	Active      bool   `json:"active"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	Lang        string `json:"lang"`
	Date        string `json:"date"`
	AccessType  string `json:"access_type"`
}

type ListHandbookResponse struct {
	Count     int64           `json:"count,string"`
	Handbooks []*HandbookSlim `json:"handbooks"`
}
