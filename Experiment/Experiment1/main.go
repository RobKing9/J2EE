package main

import (
	"Experiment1/dao"
	"Experiment1/router"
	"github.com/gin-gonic/gin"
)

func main() {
	//创建数据库
	//连接数据库
	db := dao.GetDB()
	defer func() { _ = db.Close() }()
	r := gin.Default()

	//加载HTML文件
	r.Static("/static", "static")
	//r.StaticFS("/static", http.Dir("./static"))
	r.LoadHTMLGlob("template/*")

	r = router.CollectRouters(r)

	r.Run(":9999")
}
