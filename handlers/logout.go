package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Logout(c *gin.Context) {
	DeleteCookie(c)
	c.Redirect(http.StatusFound, "/login")
}
