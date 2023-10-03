package dto

type UserId struct {
	Id string `uri:"id" json:"id" binding:"required,uuid"`
}
