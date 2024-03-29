package v1

import (
	"net/http"
	"webserver/models"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var user models.User
	c.Bind(&user)
	isLogin := user.CheckUser()
	if isLogin {
		c.String(http.StatusOK, "Hello %s", user.User)
	} else {
		c.String(http.StatusUnauthorized, "")
	}
}
