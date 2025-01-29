package dto

type Register struct{
	Username string `json:"username" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type ResetPassword struct{
	Password string `json:"password" binding:"required,min=6"`
	Email string `json:"email" binding:"required,email"`
}

type Login struct{
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}