package api

import (
	"shortener-url/api/docs"
	"shortener-url/api/handlers"
	"shortener-url/config"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// SetUpRouter godoc
// @description This is a api gateway
// @termsOfService amiin_ticker
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func SetUpRouter(h handlers.Handler, cfg config.Config) (r *gin.Engine) {
	r = gin.New()

	r.Use(gin.Logger(), gin.Recovery())

	docs.SwaggerInfo.Title = cfg.ServiceName
	docs.SwaggerInfo.Version = cfg.Version
	// docs.SwaggerInfo.Host = cfg.ServiceHost + cfg.HTTPPort
	docs.SwaggerInfo.Schemes = []string{cfg.HTTPScheme}

	// r.Use(customCORSMiddleware())
	// r.Use(h.HasAccessCheck)

	// Default
	r.GET("/ping", h.Ping)
	r.GET("/config", h.GetConfig)

	v1 := r.Group("/v1")
	{

		// User
		v1.POST("/user", h.CreateUsers)
		v1.PUT("/user-token", h.UpdateUser)
		v1.GET("/user/:id", h.GetUserByID)
		v1.DELETE("/user/:id", h.DeleteUsers)
		v1.GET("/user", h.GetUserList)

		// Login
		v1.POST("/login", h.Login)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return
}

func customCORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Max-Age", "3600")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
