package main

import (
	"hr-api/auth"
	"hr-api/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func setRoutes() *gin.Engine {
	r := gin.New()
	r.SetTrustedProxies(nil)
	r.Use(gin.Recovery())
	r.Use(gin.Logger())
	cc := cors.DefaultConfig()
	cc.AllowAllOrigins = true
	r.Use(cors.New(cc))
	r.PUT("/api/token/check", func(c *gin.Context) { handlers.CheckTokenHandler(c) })
	r.GET("/api/psy/:name", func(c *gin.Context) { handlers.Psy_fetchOne(c) })
	r.GET("/api/:name", func(c *gin.Context) { handlers.Self_fetchOne(c) })
	r.GET("/api/psy", handlers.Psy_fetchAll())
	r.GET("/api", handlers.Self_fetchAll())
	authd := r.Group("/api").Use(auth.Auth())
	{
		authd.GET("/token/generate", func(c *gin.Context) { handlers.GenerateTokenHandler(c) })
		authd.DELETE("/:name", func(c *gin.Context) { handlers.Self_deleteOne(c) })
		authd.PATCH("/:name", func(c *gin.Context) { handlers.Self_updateOne(c) })
		authd.POST("/", func(c *gin.Context) { handlers.Self_insertOne(c) })
	}
	return r
}

func main() {
	gin.SetMode(gin.DebugMode)
	r := setRoutes()
	r.Run("localhost:8080")
}
