package router

import (
	"Experiment1/controler"
	"github.com/gin-gonic/gin"
)

func CollectRouters(r *gin.Engine) *gin.Engine {
	r.GET("/", controler.Login)
	r.POST("/logsuccess", controler.LoginSuccess)

	r.GET("/register", controler.Register)
	r.POST("/regsuccess", controler.RegSuccess)

	r.GET("/index", controler.Index)

	r.GET("/resume", controler.Resume)

	r.GET("/mydream", controler.Mydream)

	r.GET("/mylife", controler.Mylife)

	r.GET("/learn", controler.Learn)
	return r
}
