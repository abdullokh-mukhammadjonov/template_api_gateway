package var_api_gateway

import "mime/multipart"

type EntityFilesSwag struct {
	Comment string `json:"comment"`
}

type File struct {
	File *multipart.FileHeader `form:"file" binding:"required"`
}

type FilesResponse struct {
	FilePath string `json:"file_path"`
	URL      string `json:"file_url"`
}
type ImageResponse struct {
	FileID   string `json:"file_id"`
	FilePath string `json:"file_path"`
}

type AllFilesResponse struct {
	Count int64           `json:"count"`
	Files []FilesResponse `json:"files"`
}