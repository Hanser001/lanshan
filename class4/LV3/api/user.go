package api

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"lanshan/class4/LV3/api/middleware"
	"lanshan/class4/LV3/dao"
	"lanshan/class4/LV3/model"
	"lanshan/class4/LV3/utils"
	"time"
)

// 注册
func register(c *gin.Context) {
	if err := c.ShouldBind(&model.User{}); err != nil {
		utils.RespSuccess(c, "verification failed")
		return
	}
	// 传入用户名和密码
	username := c.PostForm("username")
	password := c.PostForm("password")

	//检验用户名是否重复
	flag := dao.SelectUserByName(username, 0)
	if flag {
		utils.RespFail(c, "user already exists")
		return
	}

	dao.Adduser(username, password)
	utils.RespSuccess(c, "add successful")
}

// 登录
func login(c *gin.Context) {
	if err := c.ShouldBind(&model.User{}); err != nil {
		utils.RespFail(c, "verification failed")
		return
	}

	// 传入用户名和密码
	username := c.PostForm("username")
	password := c.PostForm("password")

	//检验用户名是否存在
	flag := dao.SelectUserByName(username, 0)
	//不存在则退出
	if !flag {
		// 以 JSON 格式返回信息
		utils.RespFail(c, "user doesn't exists")
		return
	}

	//查找正确的密码
	selectPassword := dao.SelectPasswordFromUsername(username)
	if password != selectPassword {
		utils.RespFail(c, "wrong password")
		return
	}
	// 正确则登录成功
	// 创建一个我们自己的声明
	claim := model.MyClaims{
		Username: username, // 自定义字段
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(), // 过期时间
			Issuer:    "Yxh",                                // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	tokenString, _ := token.SignedString(middleware.Secret)
	utils.RespSuccess(c, tokenString)
}

func getUsernameFromToken(c *gin.Context) {
	username, _ := c.Get("username")
	utils.RespSuccess(c, username.(string))
}

// 添加好友
func Add(c *gin.Context) {
	//传入用户名，密码，好友名
	username := c.PostForm("username")
	password := c.PostForm("password")
	friendname := c.PostForm("friendname")

	//检索用户
	flag := dao.SelectUserByName(username, 0)
	//不存在则报错
	if !flag {
		utils.RespFail(c, "user doesn't exist")
		return
	}

	//检索密码是否正确
	//不正确则报错
	selectPassword := dao.SelectPasswordFromUsername(username)
	if password != selectPassword {
		utils.RespFail(c, "wrong password")
		return
	}

	//检索想要添加的用户是否存在
	flag2 := dao.SelectUserByName(friendname, 0)
	//不存在则报错
	if !flag2 {
		utils.RespFail(c, "friend doesn't exist")
		return
	}

	//得到双方ID
	userId := dao.GetIdFromUsername(username)
	if userId == 0 {
		utils.RespFail(c, "some bad thing happened:)")
		return
	}
	friendId := dao.GetIdFromUsername(friendname)

	if friendId == 0 {
		utils.RespFail(c, "some bad thing happened:)")
		return
	}

	//检索是否已经添加过好友
	flag3 := dao.ExamineFriend(userId, friendId)
	if flag3 {
		utils.RespFail(c, "already add this friend")
		return
	}

	dao.AddFriend(userId, friendId)
	utils.RespSuccess(c, "add friend successful")
}

// 删除好友
func Delete(c *gin.Context) {
	//传入用户名，密码，好友名
	username := c.PostForm("username")
	password := c.PostForm("password")
	friendname := c.PostForm("friendname")

	//检索用户
	flag := dao.SelectUserByName(username, 0)
	//不存在则报错
	if !flag {
		utils.RespFail(c, "user doesn't exist")
		return
	}

	//检索密码是否正确
	//不正确则报错
	selectPassword := dao.SelectPasswordFromUsername(username)
	if password != selectPassword {
		utils.RespFail(c, "wrong password")
		return
	}

	//检索用户是否存在
	flag2 := dao.SelectUserByName(friendname, 0)
	//不存在则报错
	if !flag2 {
		utils.RespFail(c, "friend doesn't exist")
		return
	}

	//得到用双方ID
	userId := dao.GetIdFromUsername(username)
	if userId == 0 {
		utils.RespFail(c, "some bad thing happened:)")
		return
	}
	friendId := dao.GetIdFromUsername(friendname)

	if friendId == 0 {
		utils.RespFail(c, "some bad thing happened:)")
		return
	}

	//检索是否有这个好友
	flag3 := dao.ExamineFriend(userId, friendId)
	if !flag3 {
		utils.RespFail(c, "do not have this friend")
		return
	}

	dao.DeleteFriend(userId, friendId)
	utils.RespSuccess(c, "delete successful")
}

// 修改好友分组
func Change(c *gin.Context) {
	//输入用户名和密码
	username := c.PostForm("username")
	password := c.PostForm("password")
	friendname := c.PostForm("friendname")
	newgroup := c.PostForm("newgroup")

	//检验用户名是否存在
	flag := dao.SelectUserByName(username, 0)
	//不存在则退出
	if !flag {
		// 以 JSON 格式返回信息
		utils.RespFail(c, "user doesn't exists")
		return
	}

	//查找正确的密码
	selectPassword := dao.SelectPasswordFromUsername(username)
	if password != selectPassword {
		utils.RespFail(c, "wrong password")
		return
	}

	//检索用户是否存在
	flag2 := dao.SelectUserByName(friendname, 0)
	//不存在则报错
	if !flag2 {
		utils.RespFail(c, "friend doesn't exist")
		return
	}

	//得到双方的ID
	userId := dao.GetIdFromUsername(username)
	if userId == 0 {
		utils.RespFail(c, "some bad thing happened:)")
		return
	}
	friendId := dao.GetIdFromUsername(friendname)

	if friendId == 0 {
		utils.RespFail(c, "some bad thing happened:)")
		return
	}

	//检索是否有这个好友
	flag3 := dao.ExamineFriend(userId, friendId)
	if !flag3 {
		utils.RespFail(c, "do not have this friend")
		return
	}

	//进行分组
	dao.ChangeGroup(newgroup, userId, friendId)
	utils.RespSuccess(c, "change group successfully")
}

// 查看所有好友
func SeeFriend(c *gin.Context) {
	//输入用户名和密码
	username := c.PostForm("username")
	password := c.PostForm("password")

	//检验用户名是否存在
	flag := dao.SelectUserByName(username, 0)
	//不存在则退出
	if !flag {
		// 以 JSON 格式返回信息
		utils.RespFail(c, "user doesn't exists")
		return
	}

	//查找正确的密码
	selectPassword := dao.SelectPasswordFromUsername(username)
	if password != selectPassword {
		utils.RespFail(c, "wrong password")
		return
	}
	// 正确则登录成功,查找好友
	userId := dao.GetIdFromUsername(username)
	if userId == 0 {
		utils.RespFail(c, "get message failed")
		return
	}

	friends := dao.SelectAllFriends(userId)
	if len(friends) == 0 {
		utils.RespFail(c, "no friends")
		return
	}

	utils.RespFriends(c, friends)
}

// 搜索好友
func SelectFriend(c *gin.Context) {
	//输入用户名,密码,关键词
	username := c.PostForm("username")
	password := c.PostForm("password")
	keyword := c.PostForm("keyword")

	//检验用户是否存在
	flag := dao.SelectUserByName(username, 0)
	//不存在则退出
	if !flag {
		// 以 JSON 格式返回信息
		utils.RespFail(c, "user doesn't exists")
		return
	}

	//查找正确的密码
	selectPassword := dao.SelectPasswordFromUsername(username)
	if password != selectPassword {
		utils.RespFail(c, "wrong password")
		return
	}

	//若有错就终止进程
	userId := dao.GetIdFromUsername(username)
	if userId == 0 {
		utils.RespFail(c, "get message failed")
		return
	}

	// 无误就进行查找
	friends := dao.SelectAllFriends(userId)
	if len(friends) == 0 {
		utils.RespFail(c, "no friends")
		return
	}

	SelectedUsers := dao.SelectSomeUsers(keyword)
	if len(SelectedUsers) == 0 {
		utils.RespFail(c, "no such users")
		return
	}

	//将既在好友列表中又符合搜索条件的用户ID取出来
	SelectedFriends := make([]int, 100)
	for i := 0; i < len(friends); i++ {
		for j := 0; j < len(SelectedUsers); j++ {
			if SelectedUsers[j] == friends[i] {
				SelectedFriends = append(SelectedFriends, friends[i])
				break
			}
		}
	}

	if len(SelectedFriends) == 0 {
		utils.RespFail(c, "no such friends")
		return
	}

	//以json格式返回信息
	utils.RespFriends(c, SelectedFriends)
}
