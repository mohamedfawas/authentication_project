package handlers

import (
	data "ginserver/data"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	ClearCache(c)
	// check cookie exist
	cookieExist := VerifyCookie(c)
	if !cookieExist {
		c.HTML(http.StatusOK, "login.html", nil)
	}
}

func PostLogin(c *gin.Context) {
	ClearCache(c)

	//  data fetching from form
	email := c.PostForm("email")
	password := c.PostForm("password")

	// compair with previous data & cookie setting
	val, exist := data.UserData[email]
	if exist && val.Email == email && val.Password == password {
		CreateCookie(c)
		c.Redirect(http.StatusFound, "/index")
	} else {
		c.HTML(http.StatusOK, "login.html", "Enter valid data")
	}
}
