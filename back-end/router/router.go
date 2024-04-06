package router

import (
	"fmt"
	"os"
	v1 "webserver/router/api/v1"

	"github.com/gin-gonic/gin"
)

func Init() {
	port := os.Getenv("PORT")
	router := gin.Default()
	//router.Use(static.Serve("/", static.LocalFile("../front-end", true)))
	// router.GET("/", func(c *gin.Context) {
	// 	c.String(http.StatusOK, "Hello")
	// })

	apiV1 := router.Group("/v1")
	{
		apiV1.POST("/login", v1.Login)
		apiV1.POST("/register", v1.Register)
	}
	webPort := fmt.Sprintf("0.0.0.0:%s", port)
	router.Run(webPort)
}
