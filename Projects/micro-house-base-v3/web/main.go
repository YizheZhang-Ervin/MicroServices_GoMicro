package main

import (
	"micro-house-base-v3/web/controller"
	"micro-house-base-v3/web/model"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func LoginFilter() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 初始化 Session 对象
		s := sessions.Default(ctx)
		userName := s.Get("userName")

		if userName == nil {
			ctx.Abort() // 从这里返回, 不必继续执行了
		} else {
			ctx.Next() // 继续向下
		}
	}
}

// 添加gin框架开发3步骤
func main() {

	// 初始化 ModelFunc.go里的 Redis 链接池
	model.InitRedis()

	// 初始化 Model.go里的 MySQL 链接池
	model.InitDb()

	// 初始化路由
	router := gin.Default()

	// 初始化容器
	store, _ := redis.NewStore(10, "tcp", "192.168.6.108:6379", "", []byte("micro-house-base-v2"))

	// 使用容器，中间件会对后面的路由都生效
	router.Use(sessions.Sessions("mysession", store)) // 使用中间件! -- 指定容器.

	router.Static("/home", "view")

	// 添加路由分组
	r1 := router.Group("/api/v1.0")
	{
		r1.GET("/session", controller.GetSession)
		r1.GET("/imagecode/:uuid", controller.GetImageCd)
		r1.GET("/smscode/:phone", controller.GetSmscd)
		r1.POST("/users", controller.PostRet)
		r1.GET("/areas", controller.GetArea)
		r1.POST("/sessions", controller.PostLogin)

		r1.Use(LoginFilter()) //以后的路由,都不需要再校验 Session 了. 直接获取数据即可!

		r1.DELETE("/session", controller.DeleteSession)
		r1.GET("/user", controller.GetUserInfo)
		r1.PUT("/user/name", controller.PutUserInfo)
		r1.POST("/user/avatar", controller.PostAvatar)
	}

	// 启动运行
	router.Run(":8080")
}
