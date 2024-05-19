package handlers

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// InitSession initializes the session middleware

func InitSession() gin.HandlerFunc {
	store := cookie.NewStore([]byte("12345"))
	return sessions.Sessions("mysession", store)
}

// CreateSession creates a new session
func CreateSession(c *gin.Context, key string, value interface{}) error {
	session := sessions.Default(c)
	session.Set(key, value)

	// Save the session and check for errors
	if err := session.Save(); err != nil {
		return err
	}

	return nil
}

// GetSessionValue gets the value of a given key from the session
func GetSessionValue(c *gin.Context, key string) interface{} {
	session := sessions.Default(c)
	return session.Get(key)
}

// DestroySession destroys the session
func DestroySession(c *gin.Context) error {
	session := sessions.Default(c)
	session.Clear()
	return session.Save()
}

// SessionValidate checks if the session is valid (optional, depending on your use case)
func SessionValidate() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userID := session.Get("userID") // Replace with your actual key

		// Example: If the user is not logged in, redirect to the login page
		if userID == nil {
			c.Redirect(http.StatusFound, "/login") // Adjust the redirect URL as needed
			c.Abort()
			return
		}

		c.Next()
	}
}
