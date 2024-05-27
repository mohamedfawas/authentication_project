package handlers

import (
	"fmt"
	data "ginserver/data"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	ClearCache(c)

	// checking cookie exist also fetchig user details
	var name string
	cookie, err := c.Cookie("session")
	if err != nil {
		fmt.Println("error in index rooter :", err)
	}

	session, exist := data.Sessions[cookie]
	if exist {
		data, exist := data.UserData[session.Email]
		if exist {
			name = data.Name
			c.HTML(http.StatusOK, "index.html", name)
		}
	} else {
		c.Redirect(http.StatusFound, "/login")
	}

}
