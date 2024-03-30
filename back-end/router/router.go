package router

import (
	"fmt"
	"net/http"
	"os"
	v1 "webserver/router/api/v1"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func Init() {
	port := os.Getenv("PORT")
	router := gin.Default()
	router.Use(static.Serve("/", static.LocalFile("../front-end", true)))
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	apiV1 := router.Group("/api/v1")
	{
		apiV1.POST("/login", v1.Login)
		apiV1.POST("/register", v1.Register)
	}
	webPort := fmt.Sprintf("127.0.0.1:%s", port)
	router.Run(webPort)
}
