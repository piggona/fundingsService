package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/piggona/fundingsView/api/handlers"
)

func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("origin")
		if len(origin) == 0 {
			origin = c.Request.Header.Get("Origin")
		}
		c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST")
		c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func createServer() *http.Server {
	router := gin.New()

	// 注册插件
	router.Use(cors())
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

	router.GET("api/category/detail/:uuid", handlers.GetCategoryDetail)
	router.GET("api/category/resfish/:uuid", handlers.GetCategoryRelatedFund)
	router.GET("api/category/restech/:uuid", handlers.GetCategoryRelatedTech)
	router.GET("api/category/resorg/:uuid", handlers.GetCategoryRelatedOrg)
	router.GET("api/category/resindu/:uuid", handlers.GetCategoryRelatedIndu)

	router.GET("api/indu/detail/:uuid", handlers.GetIndustryDetail)
	router.GET("api/indu/resfish/:uuid", handlers.GetIndustryRelatedFund)
	router.GET("api/indu/restech/:uuid", handlers.GetIndustryRelatedTech)
	router.GET("api/indu/resorg/:uuid", handlers.GetIndustryRelatedOrg)
	router.GET("api/indu/rescate/:uuid", handlers.GetIndustryRelatedDiv)

	router.GET("api/org/detail/:uuid", handlers.GetOrganizationDetail)
	router.GET("api/org/resfish/:uuid", handlers.GetOrganizationRelatedFund)
	router.GET("api/org/rescate/:uuid", handlers.GetOrganizationRelatedDiv)
	router.GET("api/org/restech/:uuid", handlers.GetOrganizationRelatedTech)
	router.GET("api/org/resindustry/:uuid", handlers.GetOrganizationRelatedIndu)

	router.GET("api/tech/detail/:uuid", handlers.GetTechnologyDetail)
	router.GET("api/tech/resfish/:uuid", handlers.GetTechnologyRelatedFund)
	router.GET("api/tech/rescate/:uuid", handlers.GetTechnologyRelatedDiv)
	router.GET("api/tech/resorg/:uuid", handlers.GetTechnologyRelatedOrg)
	router.GET("api/tech/resindustry/:uuid", handlers.GetTechnologyRelatedIndu)

	// 配置服务器
	srv := &http.Server{
		Addr:    ":8070",
		Handler: router,
	}
	return srv
}
