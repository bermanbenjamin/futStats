package handlers

type SignInRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Age      int    `json:"age" binding:"required"`
	Slug     string `json:"slug" binding:"required"`
}
