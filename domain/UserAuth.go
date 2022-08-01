package domain

//type for LoginAuth struct
type LoginAuth struct {
	Email    string `json:"email" from:"email"`
	Password string `json:"password" from:"password"`
}
