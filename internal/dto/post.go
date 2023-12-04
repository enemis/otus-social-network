package dto

import (
	"otus-social-network/internal/model"
)

type PostInput struct {
	Title  string           `json:"name" binding:"required" required:"$field is required"`
	Post   string           `json:"post" binding:"required" required:"$field is required"`
	Status model.PostStatus `json:"status" binding:"required,post_status"`
}
