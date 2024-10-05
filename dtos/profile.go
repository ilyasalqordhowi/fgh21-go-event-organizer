package dtos

type Profile struct {
	Id            int     `json:"id"`
	Picture       *string `json:"picture"  db:"picture"`
	FullName      string  `json:"full_name"  binding:"required" form:"full_name" db:"full_name"`
	Birthdate     *string `json:"birtdate" form:"birthdate" db:"birth_date"`
	Gender        *int    `json:"gender" form:"gender" db:"gender"`
	PhoneNumber   *string `json:"phoneNumber" form:"phone_number" db:"phone_number"`
	Profession    *string `json:"profession" form:"profession" db:"profession"`
	NationalityId *int    `json:"nationalityId" form:"nationalityId" db:"nationality_id"`
	UserId        int     `json:"userId" form:"userId" db:"user_id"`
}
type JoinRegist struct {
	Id       int     `json:"id"`
	Email    *string `json:"email"  binding:"required,email" form:"email" db:"email"`
	Password string  `json:"-" form:"password"  binding:"required" db:"password"`
	Results  Profile
}