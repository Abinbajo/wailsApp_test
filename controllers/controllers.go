package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"stickyNote/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

const apiKey = "your-api-key"

func VerifyPassword(userPassword string, givenPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(givenPassword), []byte(userPassword))
	valid := err == nil
	msg := ""
	if !valid {
		msg = "Login or Password is Incorrect"
	}
	return valid, msg
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check API key
		requestApiKey := c.GetHeader("API-Key")
		if requestApiKey != apiKey {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid API key"})
			return
		}

		var user models.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	
		
		loginURL := "http://your-frontend-url/api/login"  
		payload, _ := json.Marshal(user)
		resp, err := http.Post(loginURL, "application/json", bytes.NewBuffer(payload))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to forward request to frontend"})
			return
		}
		defer resp.Body.Close()

		
		var response map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response from frontend"})
			return
		}

		
		c.JSON(resp.StatusCode, response)
	}
}
