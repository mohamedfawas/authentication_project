package handlers

import (
	"fmt"
	"login-gin/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	//Clear Cache of the browser
	utils.Clearache(c)

	//check if User logged in

	userSession := GetSessionValue(c, "login-session")
	User, _ := GetUser(GlobalEmail)
	fmt.Println("", GlobalEmail)
	if userSession != nil {
		//Load Home
		c.HTML(http.StatusOK, "index.html", User)
		return
	}

	//Redirect Login
	fmt.Println("Redirect to Login")
	c.HTML(http.StatusOK, "login.html", nil)

}
