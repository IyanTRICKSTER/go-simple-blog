package requests

import (
	"mime/multipart"
)

type CreatePostRequest struct {
	Title   string                `validate:"required,min=3,max=255" json:"title" form:"title"`
	Content string                `validate:"required,min=3" json:"content" form:"content"`
	Image   *multipart.FileHeader `validate:"required" json:"image" form:"image"`
}
