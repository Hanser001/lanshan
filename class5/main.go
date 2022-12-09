package main

import (
	"context"
	"lanshan/class5/dao"
)

func main() {
	dao.InitDB()
	dao.InitRedis()
	SetCache()
	Likes()
}

func SetCache() {
	c := dao.NewUserCache(1, "ovo", context.Background())
	c.StringSet()
	c.GetStringValue()
}

func Likes() {
	c := dao.NewUserCache(2, "^v^", context.Background())
	c.SetLikes(1)
	c.DeleteLikes(1)
}
