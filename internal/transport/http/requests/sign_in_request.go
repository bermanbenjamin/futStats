package requests

type SignInRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Age      int    `json:"age" binding:"required"`
	Slug     string `json:"slug" binding:"required"`
}
