package main

import (
	"fmt"
	handlers "login-gin/handler"
	"login-gin/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

var Users = map[string]models.User{}

func main() {

	//Create a server
	r := gin.Default()
	store := cookie.NewStore([]byte("12345"))
	r.Use(sessions.Sessions("mysession", store))

	r.LoadHTMLGlob("templates/*html") //Load templates
	r.Static("/static", "./static")   //Load static files like CSS,HTML

	//Handle Requests
	r.GET("/", handlers.Home)
	r.GET("/login", handlers.Login)
	r.POST("/login", handlers.LoginPost)
	r.GET("/signup", handlers.Signup)
	r.POST("/signup", handlers.SignupPost)
	r.POST("/logout", handlers.Logout)

	//Start sever
	fmt.Println("Server Started at PORT:8080")
	r.Run(":8080")

}
