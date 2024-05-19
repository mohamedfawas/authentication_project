package handlers

import (
	"fmt"
	"login-gin/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Logout(c *gin.Context) {

	//Deleting existing Session
	err := DestroySession(c)
	if err != nil {
		fmt.Println("Error occured:", err)
		return
	}
	//Clear cache
	utils.Clearache(c)
	//Load login
	c.HTML(http.StatusOK, "login.html", "Logout Success")
}
