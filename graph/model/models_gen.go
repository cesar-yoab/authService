// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Authenticate struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RefreshToken struct {
	OldToken string `json:"oldToken"`
}

type RegisterInput struct {
	Fname           string `json:"fname"`
	Lname           string `json:"lname"`
	Email           string `json:"email"`
	Username        string `json:"username"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}

type Token struct {
	Jwt string `json:"jwt"`
}

type User struct {
	ID       string `json:"_id"`
	Username string `json:"username"`
}
