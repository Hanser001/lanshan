package api

import (
	"github.com/gin-gonic/gin"
	"lanshan/class4/LV3/api/middleware"
)

func InitRouter() {
	r := gin.Default()
	r.Use(middleware.CORS())

	r.POST("/register", register)
	r.POST("/login", login)

	FriendRouter := r.Group("/friends")
	{
		FriendRouter.POST("/addFriend", Add)       //添加好友
		FriendRouter.POST("/deleteFriend", Delete) //删除好友
		FriendRouter.POST("/group", Change)
		FriendRouter.GET("/seeFriends", SeeFriend)      //查看所有好友
		FriendRouter.GET("/selectFriend", SelectFriend) //搜索好友
	}

	UserRouter := r.Group("/user")
	{
		UserRouter.Use(middleware.JWTAuthMiddleware())
		UserRouter.GET("/get", getUsernameFromToken)
	}

	r.Run(":8088")
}
