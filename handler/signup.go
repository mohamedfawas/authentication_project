package handlers

import (
	"fmt"
	"log"
	"login-gin/models"
	"login-gin/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

var GlobalEmail string
var User models.User

func Signup(c *gin.Context) {
	// Clear Cache
	utils.Clearache(c)

	// Check if user is already logged in
	userSession := GetSessionValue(c, "login-session")
	if userSession != nil {
		// Redirect to Home
		fmt.Println("Redirected to Home")
		c.HTML(http.StatusSeeOther, "index.html", userSession)
		return
	}

	// Load Signup
	fmt.Println("Signup Loaded")
	c.HTML(http.StatusOK, "signup.html", nil)
}

// Global Email is a global variable for email accessible across other Go files

func SignupPost(c *gin.Context) {
	// Clear Cache
	utils.Clearache(c)

	// Check if user is already logged in
	userSession := GetSessionValue(c, "login-session")
	if userSession != nil {
		// Redirect to Home
		fmt.Println("Redirected to Home")
		c.HTML(http.StatusSeeOther, "index.html", userSession)
		return
	}

	// Retrieve data from HTML Form
	c.Request.ParseForm()
	username := c.Request.FormValue("name")
	GlobalEmail = c.Request.FormValue("email") // Set the global email variable
	pwd := c.Request.Form.Get("password")
	confirmPwd := c.Request.Form.Get("confirm-password")

	// Check Passwords
	if pwd != confirmPwd {
		c.HTML(http.StatusBadRequest, "signup.html", "Password must match")
		return
	}

	// Create User
	exist := CreateUser(username, GlobalEmail, pwd)
	if exist != nil {
		c.HTML(http.StatusBadRequest, "signup.html", exist)
		return
	}

	// Create Session and Login User
	User, _ := GetUser(GlobalEmail)
	err := CreateSession(c, "login-session", "username")
	if err != nil {
		c.Error(err)
	}
	log.Println(User)

	// Load home
	c.Redirect(http.StatusSeeOther, "/")
}
