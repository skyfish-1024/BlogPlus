package main

import (
	"BLOGplus/model"
	"BLOGplus/routers"
)

func main() {
	//引用数据库
	model.InitDb()
	routers.InitRouter()
}
