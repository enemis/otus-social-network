package dto

type FriendId struct {
	Id string `form:"friend_id" json:"friend_id" binding:"required,uuid"`
}
