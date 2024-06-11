package v1

import (
	"net/http"
	"webserver/middlewares"
	"webserver/models"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var user models.User
	c.Bind(&user)
	user.Password = middlewares.Sha256(user.Password)
	if user.CheckUser() {
		c.String(http.StatusOK, "Hello %s", user.User)
	} else {
		c.String(http.StatusUnauthorized, "")
	}
}

func Register(c *gin.Context) {
	var user models.User
	c.Bind(&user)
	user.Password = middlewares.Sha256(user.Password)
	if user.RegisterUser() {
		c.String(http.StatusOK, "Hello %s", user.User)
	} else {
		c.String(http.StatusUnauthorized, "")
	}
}
