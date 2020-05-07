package api

import (
	"riji/service"
	"riji/utils"

	"github.com/gin-contrib/pprof"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func Install(r *gin.Engine, s *service.RijiServer) *gin.Engine {

	// 性能分析工具
	pprof.Register(r)

	// 中间件
	// r.Use(utils.HandleErrors()) // 日志中间件
	r.Use(utils.CORSMiddleware())
	r.Use(sessions.Sessions("rijiSession", cookie.NewStore([]byte("session"))))
	r.Use(utils.AuthHandler())

	// 处理不存在的API
	r.NoRoute(utils.NoRoute)

	// 注册路由
	api := r.Group("/api/v1")

	// hello world
	api.POST("/hello", s.HelloWorld)
	api.DELETE("/picture/:id", s.PicDelete)
	api.GET("/picture", s.PicList)
	api.GET("/video", s.VideoList)
	api.GET("/message", s.MessageList)
	api.POST("/message", s.MessageCreate)
	api.DELETE("/message/:id", s.MessageDelete)
	api.GET("/upload/credentials", s.GetCredentials)
	api.POST("/upload", s.CosCallback)

	// 生成背景图片
	api.POST("/back_img_update", s.UpdateBackImg)
	// 获取背景图片
	api.GET("/back_img_url", s.GetBackImgUrl)

	user := r.Group("/api/user")
	user.POST("/login", s.UserLogin)
	user.POST("/logout", s.UserLogout)
	user.POST("/register", s.UserRegister)

	return r
}
