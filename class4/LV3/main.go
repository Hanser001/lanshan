package main

import (
	"lanshan/class3/api"
	"lanshan/class4/LV3/dao"
)

func main() {
	dao.InitDB()
	api.InitRouter()
}
