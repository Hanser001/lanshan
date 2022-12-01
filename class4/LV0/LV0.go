package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"lanshan/class4/LV0/model"
	"log"
)

var db *sql.DB

func initDB() {
	var err error
	dsn := "root:WJJ20040311@tcp(127.0.0.1:3306)/lesson6?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalln(err)
	}
	// 尝试与数据库建立连接（校验dsn是否正确）
	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("DB connect success")
	return
}

// 插入数据
func InsertStudent(st model.Student) {
	sqlStr := "insert into student(name,sex,age) values (?,?,?)"
	_, err := db.Exec(sqlStr, st.Name, st.Sex, st.Age)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	log.Println("insert success")
}

// 删除学生
func DeleteStduentByName() {
	sqlStr := "delete from student where name = ?"
	_, err := db.Exec(sqlStr, "xiaohu")
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return
	}
	log.Println("delete success")
}

// 批量查询
func MultiQueryById(id int) {
	sqlStr := "select id, name, age from student where id > ?"
	rows, err := db.Query(sqlStr, id)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	// 关闭rows释放持有的数据库链接
	defer rows.Close()

	// 循环读取结果集中的数据
	for rows.Next() {
		var st model.Student
		err := rows.Scan(&st.Id, &st.Name, &st.Age)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		log.Printf("MultiQueryById: id:%d name:%s age:%d\n", st.Id, st.Name, st.Age)
	}
}

func main() {
	initDB()

	//s1 := model.Student{
	//Name: "koi",
	//Sex:  "woman",
	//Age:  15,
	//}

	//InsertStudent(s1)

	//DeleteStduentByName()
	MultiQueryById(0)
}
