package dao

import (
	"context"
	"fmt"
	"time"
)

type UserCache struct {
	Id       int
	Username string
	Context  context.Context
}

func NewUserCache(Id int, Username string, Context context.Context) *UserCache {
	return &UserCache{
		Id:       Id,
		Username: Username,
		Context:  Context,
	}
}

func (u *UserCache) StringSet() {
	err := Rdb.Set(u.Context, u.Username, u.Id, 24*time.Hour)
	if err != nil {
		fmt.Println(err)
	}
}

func (u *UserCache) GetStringValue() {
	val := Rdb.Get(u.Context, u.Username)
	if val.Err() != nil {
		fmt.Println(val.Err())
	}
	fmt.Println("key", val)
}

// 用set实现点赞,也可以取消点赞
func (u *UserCache) SetLikes(beLikedUserId int) {
	err := Rdb.SAdd(u.Context, string(beLikedUserId), u.Id)
	if err != nil {
		fmt.Println(err)
	}

	num, err2 := Rdb.SCard(u.Context, string(beLikedUserId)).Result()
	if err != nil {
		fmt.Println(err2)
	}
	fmt.Println(num)
}

func (u *UserCache) DeleteLikes(beLikedUserId int) {
	err := Rdb.SRem(u.Context, string(beLikedUserId), u.Id)
	if err != nil {
		fmt.Println(err)
	}

	num, err2 := Rdb.SCard(u.Context, string(beLikedUserId)).Result()
	if err != nil {
		fmt.Println(err2)
	}
	fmt.Println(num)
}
