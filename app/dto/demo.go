package dto

type ReqAddDemo struct {
	Name  string `json:"name" form:"name" binding:"required"`
	Email string `json:"email" form:"email" binding:"required"`
}

type ReqUpdateDemo struct {
	Id    int    `json:"id" form:"id" binding:"required"`
	Name  string `json:"name" form:"name" binding:"required"`
	Email string `json:"email" form:"email" binding:"required"`
}

type ReqDeleteDemo struct {
	Id int `json:"id" form:"id" binding:"required"`
}
