package dao

import (
	"fmt"
	"lanshan/class4/LV1/model"
	"log"
)

func Adduser(username string, password string) {
	sqlStr := "insert into user(name,password) values (?,?)"
	_, err := MyDb.Exec(sqlStr, username, password)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	log.Println("insert success")
}

func SelectUser(username string, id int) bool {
	sqlStr := "select name from user where id>?"
	rows, err := MyDb.Query(sqlStr, id)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return false
	}
	// 非常重要：关闭rows释放持有的数据库链接
	defer rows.Close()

	// 循环读取结果集中的数据
	for rows.Next() {
		var u model.User
		err := rows.Scan(&u.Username)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return false
		}
		if u.Username == username {
			return true
		}
	}
	return false
}

func SelectPasswordFromUsername(username string) string {
	sqlStr := "select password from user where name=?"
	var u model.User
	MyDb.QueryRow(sqlStr, username).Scan(&u.Password)
	return u.Password
}

// 用更新数据的方法来添加密保问题和问题答案
func AddPwdQuestion(username string, question string) {
	sqlStr := "update user set question=? where name=?"
	MyDb.Exec(sqlStr, question, username)
}

func AddPwdAnswer(username string, answer string) {
	sqlStr := "update user set answer=? where name=?"
	MyDb.Exec(sqlStr, answer, username)
}

// 检索问题是否存在与问题答案
func SelectQuestion(username string) bool {
	var u model.User
	sqlStr := "select question from user where name=?"
	MyDb.QueryRow(sqlStr, username).Scan(&u.Question)
	if u.Question == "" {
		return false
	}
	return true
}

func SelectAnswerFromQuestion(username string) string {
	var u model.User
	sqlStr := "select answer from user where name=?"
	MyDb.QueryRow(sqlStr, username).Scan(&u.Answer)
	return u.Answer
}

// 更新方法找回密码
func FindPassword(username string, newpassword string) {
	sqlStr := "update user set password=? where name=? "
	MyDb.Exec(sqlStr, newpassword, username)
}
