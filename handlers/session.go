package handlers

import (
	"fmt"
	data "ginserver/data"
	"net/http"

	"github.com/gin-gonic/gin"
)

func VerifyCookie(c *gin.Context) bool {
	cookie, err := c.Cookie("session")
	if err != nil {
		fmt.Println("error in verifycookie checking at session file : ", err)
	}

	_, exist := data.Sessions[cookie]
	if exist {
		c.Redirect(http.StatusFound, "/index")
		return true
	}
	return false
}
