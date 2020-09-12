package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/piggona/fundingsView/api/handlers"
)

func createServer() *http.Server {
	router := gin.New()

	// 注册插件
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// 配置路由,思考是否需要为某些路由配置中间件middleware
	router.POST("api/user", handlers.CreateUser)
	router.POST("api/user/:username", handlers.Login)
	router.GET("api/user/:username", handlers.GetUserInfo)
	router.DELETE("api/user/:username", handlers.DeleteUser)

	router.GET("api/search/basic", handlers.HomePage)
	router.GET("api/search/rank/:activeAmount", handlers.SwitchRank)
	router.GET("api/search/tech/:activeTech", handlers.SwitchTech)
	router.POST("api/search", handlers.GetSearch)

	router.GET("api/fund/detail/:uuid", handlers.GetDetail)
	router.GET("api/fund/cop/:uuid", handlers.GetCopTree)
	router.GET("api/fund/word/:uuid", handlers.GetWordTree)
	router.GET("api/fund/similar/:uuid/:page", handlers.GetSimilar)

	// 配置服务器
	srv := &http.Server{
		Addr:    ":8070",
		Handler: router,
	}
	return srv
}
