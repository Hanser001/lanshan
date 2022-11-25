package api

import (
	"bufio"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"io"
	"lanshan/class3/api/middleware"
	"lanshan/class3/dao"
	"lanshan/class3/model"
	"lanshan/class3/utils"
	"os"
	"time"
)

// 注册
func register(c *gin.Context) {
	form := model.User{}
	if err := c.ShouldBind(&form); err != nil {
		fmt.Println(err)
		utils.RespSuccess(c, "verification failed")
		return
	}
	// 传入用户名和密码
	username := c.PostForm("username")
	password := c.PostForm("password")

	//打开dao目录的names.txt文件，不存在则创建
	file1, err := os.OpenFile("class3/dao/names.txt", os.O_APPEND|os.O_CREATE, os.ModePerm)
	if err != nil {
		utils.RespFail(c, "file wrong")
		return
	}
	defer file1.Close()
	reader := bufio.NewReader(file1)
	for {
		line, _, err := reader.ReadLine()
		// 验证用户名是否重复
		if username == string(line) {
			utils.RespFail(c, "user already exist")
			return
		}
		if err == io.EOF {
			file1.WriteString("\n")
			file1.WriteString(username) //写入用户名
			break
		}
		if err != nil {
			utils.RespFail(c, "file wrong")
			return
		}
	}

	//打开dao目录的passwords.txt文件，不存在则创建
	file2, err := os.OpenFile("class3/dao/passwords.txt", os.O_APPEND|os.O_CREATE, os.ModePerm)
	if err != nil {
		utils.RespFail(c, "file wrong")
		return
	}
	defer file2.Close()
	file2.WriteString("\n")
	file2.WriteString(password) //写入密码

	dao.AddUser(username, password)
	// 以 JSON 格式返回信息
	utils.RespSuccess(c, "add user successful")
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

	// 验证用户名是否存在
	// 不存在则退出
	file1, err := os.Open("class3/dao/names.txt")
	if err != nil {
		utils.RespFail(c, "open file wrong")
		return
	}
	defer file1.Close()
	reader1 := bufio.NewReader(file1)
	for {
		line, _, err := reader1.ReadLine()
		if err != nil {
			utils.RespFail(c, "user doesn't exist")
			return
		}
		if username == string(line) {
			break
		}
	}

	// 查找正确的密码
	// 若不正确则传出错误
	file2, err := os.Open("class3/dao/passwords.txt")
	if err != nil {
		utils.RespFail(c, "file wrong")
		return
	}
	defer file2.Close()
	reader2 := bufio.NewReader(file2)
	for {
		line, _, err := reader2.ReadLine()
		if err != nil {
			utils.RespFail(c, "wrong password")
			return
		}
		if password == string(line) {
			break
		}
	}

	dao.AddUser(username, password)
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

// 留言(保存ID和留言内容)
func leave(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	message := c.PostForm("message")
	// 验证用户名是否存在
	flag := dao.SelectUser(username)
	// 不存在则退出
	if !flag {
		// 以 JSON 格式返回信息
		utils.RespFail(c, "user doesn't exists")
		return
	}
	// 查找正确的密码
	selectPassword := dao.SelectPasswordFromUsername(username)
	// 若不正确则传出错误
	if selectPassword != password {
		// 以 JSON 格式返回信息
		utils.RespFail(c, "wrong password")
		return
	}
	dao.LeaveMessage(message)
	//打开dao目录的messages.txt文件，不存在则创建
	file, err := os.OpenFile("class3/dao/messages.txt", os.O_APPEND|os.O_CREATE, os.ModePerm)
	if err != nil {
		utils.RespFail(c, "file wrong")
		return
	}
	defer file.Close()
	//写入用户名和留言内容
	file.WriteString("\n")
	file.WriteString(username)
	file.WriteString(":")
	file.WriteString(message)
	utils.RespSuccess(c, "已留言")
}

// 添加密保问题
func addQuestion(c *gin.Context) {
	//传入用户名和密码，问题和答案
	username := c.PostForm("username")
	password := c.PostForm("password")
	question := c.PostForm("question")
	answer := c.PostForm("answer")

	// 验证用户名是否存在
	flag := dao.SelectUser(username)
	if !flag {
		// 以 JSON 格式返回信息
		utils.RespFail(c, "user doesn't exist")
		return
	}
	// 查找正确的密码
	selectPassword := dao.SelectPasswordFromUsername(username)
	// 若不正确则传出错误,正确则进入下一步
	if selectPassword != password {
		utils.RespSuccess(c, "set question successful")
		return
	}
	dao.AddPwdQuestion(username, question)
	dao.AddPwdAnswer(username, answer)
}

// 找回密码
func findpassword(c *gin.Context) {
	//传入用户名和答案与新密码
	username := c.PostForm("username")
	answer := c.PostForm("answer")
	newPassword := c.PostForm("newPassword")

	// 验证用户名是否存在
	flag1 := dao.SelectUser(username)
	fmt.Println(flag1)
	if !flag1 {
		// 以 JSON 格式返回信息
		utils.RespFail(c, "user doesn't exist")
		return
	}
	//查找该用户名是否设置密保问题
	flag2 := dao.SelectQuestion(username)
	if !flag2 {
		utils.RespFail(c, "the user doesn't set question")
		return
	}
	//查找密保问题答案
	trueAnswer := dao.SelectAnswerFromQuestion(username)
	//答案不正确就报错
	if answer != trueAnswer {
		utils.RespFail(c, "answer wrong")
		return
	}
	//正确就修改密码
	dao.FindPassword(username, newPassword)
	utils.RespSuccess(c, "change password successful")
}

func AddLike(c *gin.Context) {
	//传入用户名，密码，想要点赞的用户名
	username := c.PostForm("username")
	password := c.PostForm("password")
	friendname := c.PostForm("friendname")

	//登录
	flag := dao.SelectUser(username)
	if !flag {
		// 以 JSON 格式返回信息
		utils.RespFail(c, "user doesn't exist")
		return
	}
	// 查找正确的密码
	selectPassword := dao.SelectPasswordFromUsername(username)
	// 若不正确则传出错误,正确则进入下一步
	if selectPassword != password {
		utils.RespSuccess(c, "set question successful")
		return
	}

	//检索好友是否存在
	flag2 := dao.SelectUser(friendname)
	if !flag2 {
		utils.RespFail(c, "friend doesn't exist")
		return
	}
	//检查用户是否已经为该用户点过赞
	ret := dao.SelectLike(friendname)
	for i := 0; i < len(ret); i++ {
		if username == ret[i] {
			utils.RespFail(c, "已经点过了")
			return
		}
	}
	//没点过赞就为此用户点赞
	dao.Like(username, friendname)
	utils.RespSuccess(c, "点赞成功")
}
