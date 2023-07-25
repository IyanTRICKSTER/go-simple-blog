package requests

import "mime/multipart"

type UpdatePostRequest struct {
	Title   string                `validate:"required,min=3,max=255" json:"title"`
	Content string                `validate:"required,min=3" json:"content"`
	Image   *multipart.FileHeader `validate:"required" json:"image"`
}
