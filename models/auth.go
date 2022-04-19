package models

type Authentication struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Token struct {
	Email       *string `json:"email"`
	TokenString string  `json:"token"`
}
