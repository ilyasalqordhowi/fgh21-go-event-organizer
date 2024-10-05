package dtos

type User struct {
	Id       int    `json:"id"`
	Email    string `json:"email" form:"email" binding:"required,email" db:"email"`
	Password string `json:"-" form:"password" db:"password"`
	Username *string `json:"username" form:"username" db:"username"`
}
type ChangePassword struct{
	OldPassword string `form:"oldPassword" json:"oldPassword"`
	NewPassword string `form:"newPassword" json:"newPassword"`
}
