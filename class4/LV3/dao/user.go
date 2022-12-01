package dao

import (
	"fmt"
	"lanshan/class4/LV3/model"
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

func SelectPasswordFromUsername(username string) string {
	sqlStr := "select password from user where name=?"
	var u model.User
	MyDb.QueryRow(sqlStr, username).Scan(&u.Password)
	return u.Password
}

func SelectUserByName(username string, startId int) bool {
	sqlStr := "select name from user where id>?"
	rows, _ := MyDb.Query(sqlStr, startId)

	// 非常重要：关闭rows释放持有的数据库链接
	defer rows.Close()

	// 循环读取结果集中的数据
	for rows.Next() {
		var u model.User
		rows.Scan(&u.Name)
		if u.Name == username {
			return true
		}
	}
	return false
}

// 通过用户名获取ID
func GetIdFromUsername(username string) int {
	sqlStr := "select id from user where name=?"
	var u model.User
	err := MyDb.QueryRow(sqlStr, username).Scan(&u.Id)
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		//有错误就返回一个不存在的ID
		return 0
	}
	return u.Id
}

// 检索好友是否已添加
func ExamineFriend(userId int, friendId int) bool {
	sqlStr := "select friendId from friends where userId=?"
	rows, _ := MyDb.Query(sqlStr, userId)

	defer rows.Close()

	//添加过返回true，没添加过返回false
	for rows.Next() {
		var f model.Friend
		rows.Scan(&f.FriendId)
		if friendId == f.FriendId {
			return true
		}
	}
	return false
}

// 添加好友时使用ID添加
// 好友关系是双向的，用一个事务进行添加好友
func AddFriend(userId int, friendId int) {
	tx, err := MyDb.Begin()
	if err != nil {
		if tx != nil {
			tx.Rollback() //回滚
		}
		log.Printf("begin trans failed,err:%v\n", err)
		return
	}

	sqlStr1 := "insert into friends (userId,friendId) values (?,?)"
	ret1, err := MyDb.Exec(sqlStr1, userId, friendId)
	if err != nil {
		tx.Rollback()
		log.Printf("exec sql1 failed,err:%v\n", err)
		return
	}
	affRow1, err := ret1.RowsAffected()
	if err != nil {
		tx.Rollback()
		log.Printf("exec ret1.RowsAffected() failed,err:%v\n", err)
		return
	}

	sqlStr2 := "insert into friends (userId,friendId) values (?,?)"
	ret2, err := MyDb.Exec(sqlStr2, friendId, userId)
	if err != nil {
		tx.Rollback()
		log.Printf("exec sql2 failed,err:%v\n", err)
		return
	}
	affRow2, err := ret2.RowsAffected()
	if err != nil {
		tx.Rollback()
		log.Printf("exec ret2.RowsAffected() failed,err:%v\n", err)
		return
	}

	if affRow1 == 1 && affRow2 == 1 {
		fmt.Println("事务提交")
		tx.Commit() // 提交事务
	} else {
		tx.Rollback()
		fmt.Println("事务回滚")
	}

	fmt.Println("exec trans success!")
}

// 删除好友同理，也用事务进行处理
func DeleteFriend(userId int, friendId int) {
	tx, err := MyDb.Begin()
	if err != nil {
		if tx != nil {
			tx.Rollback() //回滚
		}
		log.Printf("begin trans failed,err:%v\n", err)
		return
	}
	sqlStr1 := "delete from friends where userId=? and friendId=?"
	ret1, err := MyDb.Exec(sqlStr1, userId, friendId)
	if err != nil {
		tx.Rollback()
		log.Printf("exec sql1 failed,err:%v\n", err)
		return
	}
	affRow1, err := ret1.RowsAffected()
	if err != nil {
		tx.Rollback()
		log.Printf("exec ret1.RowsAffected() failed,err:%v\n", err)
		return
	}

	sqlStr2 := "delete from friends where userId=? and friendId=?"
	ret2, err := MyDb.Exec(sqlStr2, friendId, userId)
	if err != nil {
		tx.Rollback()
		log.Printf("exec sql2 failed,err:%v\n", err)
		return
	}
	affRow2, err := ret2.RowsAffected()
	if err != nil {
		tx.Rollback()
		log.Printf("exec ret2.RowsAffected() failed,err:%v\n", err)
		return
	}

	if affRow1 == 1 && affRow2 == 1 {
		fmt.Println("事务提交啦...")
		tx.Commit() // 提交事务
	} else {
		tx.Rollback()
		fmt.Println("事务回滚啦...")
	}

	fmt.Println("exec trans success!")
}

// 用更新的方式更改好友分组
func ChangeGroup(group string, userId int, friendId int) {
	sqlStr := "update friends set group=? where userId=? and friendId=?"
	_, err := MyDb.Exec(sqlStr, group, userId, friendId)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
	fmt.Printf("update success\n")
}

// 查看所有好友的id
func SelectAllFriends(userId int) []int {
	ids := make([]int, 100)
	sqlStr := "select friendId from friends where userId=?"
	rows, err := MyDb.Query(sqlStr, userId)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return []int{}
	}

	// 非常重要：关闭rows释放持有的数据库链接
	defer rows.Close()

	// 循环读取结果集中的数据
	for rows.Next() {
		var f model.Friend
		err := rows.Scan(&f.FriendId)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return []int{}
		}
		fmt.Printf("friendId:%v", f.FriendId)
		ids = append(ids, f.FriendId)
	}
	return ids
}

// 搜找用户名中包含查找关键词的用户id
func SelectSomeUsers(keyWord string) []int {
	ids := make([]int, 100)
	sqlStr := "select id from user where name like concat('%',?,'%')?"
	rows, err := MyDb.Query(sqlStr, keyWord)
	if err != nil {
		fmt.Printf("query failed,err:%v\n", err)
		return []int{}
	}

	defer rows.Close()

	for rows.Next() {
		var f model.Friend
		err := rows.Scan(&f.FriendId)
		if err != nil {
			fmt.Printf("scan failed,err:%v\n", err)
			return []int{}
		}
		fmt.Printf("friendId:%v", f.FriendId)
		ids = append(ids, f.FriendId)
	}
	return ids
}
