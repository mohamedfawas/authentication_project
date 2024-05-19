package handlers

import (
	"fmt"
	"login-gin/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	//Clear Cache of Browser
	utils.Clearache(c)

	//check if user already loged in
	userSession := GetSessionValue(c, "login-session")
	if userSession != nil {
		//Redirect to home
		fmt.Println("User session is", userSession)
		fmt.Println("Redirected to Home")
		c.HTML(http.StatusSeeOther, "index.html", userSession)
		return
	}
	//Load Login
	fmt.Println("Login Loaded")
	c.HTML(http.StatusOK, "login.html", nil)
}

func LoginPost(c *gin.Context) {
	//Clear Cache of Browser
	utils.Clearache(c)

	//Check If the user already login
	userSession := GetSessionValue(c, "login-session")
	if userSession != nil {
		//Redirect to home
		fmt.Println("Redirected to Home")
		c.HTML(http.StatusBadRequest, "index.html", userSession)
		return
	}
	// Retrieve data from Login Form
	c.Request.ParseForm()
	email := c.Request.FormValue("email")
	pwd := c.Request.Form.Get("password")

	// Check if user exists
	user, exist := GetUser(email)
	if exist != nil {
		// if not, Redirect to login
		c.HTML(http.StatusBadRequest, "login.html", exist)
		return
	}
	//Check if entere email and password is valid
	if user.Email != email {
		c.HTML(http.StatusBadRequest, "login.html", "Invalid Email")
		return
	} else if user.Password != pwd {
		c.HTML(http.StatusBadRequest, "login.html", "Worong Password")
		return
	}
	//Create a session user
	err := CreateSession(c, "login-session", "username")
	if err != nil {
		c.Error(err)
		return
	}
	//Load Home
	// c.HTML(http.StatusOK, "index.html", user)
	c.Redirect(http.StatusSeeOther, "/")

}
