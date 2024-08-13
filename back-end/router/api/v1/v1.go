package v1

import (
	"net/http"
	"webserver/middlewares"
	"webserver/models"

	"github.com/gin-gonic/gin"
)

// @BasePath /login

// @Summary User login
// @Schemes
// @Description Test user login
// @Tags Login
// @Accept mpfd
// @Produce mpfd
// @Param user_name formData string true "user"
// @param password formData string true "password"
// @Success 200 {string} Login
// @Router /login [post]
func Login(c *gin.Context) {
	var user models.User
	c.Bind(&user)
	user.Password = middlewares.Sha256(user.Password)
	if user.CheckUser() {
		var jwtSecret = []byte("secret")
		claims := middlewares.CreateDemoToken(user.User, jwtSecret)
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello " + user.User,
			"token":   claims,
		})
	} else {
		c.String(http.StatusUnauthorized, "")
	}
}

// @BasePath /register

// @Summary User register
// @Schemes
// @Description Test user register
// @Tags Register
// @Accept mpfd
// @Produce mpfd
// @Param user_name formData string true "user"
// @param password formData string true "password"
// @Success 200 {string} Register
// @Router /register [post]
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
