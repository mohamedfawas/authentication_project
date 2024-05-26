package handlers

import (
	"fmt"
	data "ginserver/data"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func DeleteCookie(c *gin.Context) {
	c.SetCookie("session", "", -1, "/", "", false, true)
	cookie, err := c.Cookie("session")
	if err != nil {
		fmt.Println("Error in DeleteCookie function", err)
	}
	delete(data.Sessions, cookie)
	fmt.Println("cookie deleted")
}

func CreateCookie(c *gin.Context) {
	fmt.Printf("cookie created")
	id := uuid.NewString() // Generates random UUID
	email := c.PostForm("email")
	fmt.Println("cookie id :", id, "email is :", email)

	// Set cookie and session datas
	session := data.SessionData{SessionId: id, Email: email}
	data.Sessions[id] = session
	c.SetCookie("session", id, 3600, "/", "", false, true)
	fmt.Println(data.Sessions)
}
