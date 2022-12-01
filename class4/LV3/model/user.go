package model

import "github.com/dgrijalva/jwt-go"

type User struct {
	Id       int
	Name     string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Age      int
}

type Friend struct {
	UserId   int
	FriendId int
	Group    string
}

type MyClaims struct {
	Username string
	jwt.StandardClaims
}
