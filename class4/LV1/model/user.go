package model

import "github.com/dgrijalva/jwt-go"

type User struct {
	Id       int
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Question string
	Answer   string
}

type MyClaims struct {
	Username string
	jwt.StandardClaims
}
