package dto

type UserId struct {
	Id string `uri:"id" json:"id" binding:"required,uuid"`
}

type FindUser struct {
	Name    string `form:"name" json:"name" binding:"required_without=Surname" _required_without:"$field should be filled if surname is empty"`
	Surname string `form:"surname" json:"surname" binding:"required_without=Name" _required_without:"$field should be filled if name is empty"`
}
