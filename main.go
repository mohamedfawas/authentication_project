package main

import (
	"fmt"
	"ginserver/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	// prerequirements
	r := gin.Default()
	r.LoadHTMLGlob("view/*.html")
	r.Static("/static", "./static")

	// routes
	r.GET("/login", handlers.Login)
	r.POST("/login", handlers.PostLogin)
	r.GET("/index", handlers.Index)
	r.GET("/logout", handlers.Logout)

	// server creation
	err := r.Run(":7000")
	if err != nil {
		fmt.Println("error in server start ,", err)
	}
}
