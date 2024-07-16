package controllers

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"stickyNote/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()
var apiKey string

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	apiKey = os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatal("API_KEY environment variable is not set")
	}
}

func setHeaders(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*") // Consider specifying allowed origins
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, x-api-key")
}

func validateAPIKey(c *gin.Context) bool {
	requestApiKey := c.GetHeader("x-api-key")
	if requestApiKey != apiKey {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid API key"})
		return false
	}
	return true
}

func forwardRequest(url string, user models.User) (*http.Response, error) {
	payload, err := json.Marshal(user)
	if err != nil {
		log.Println("Error marshalling user:", err)
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		log.Println("Error creating request:", err)
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", apiKey)

	client := &http.Client{}
	return client.Do(req)
}

func handleResponse(c *gin.Context, resp *http.Response) {
	defer resp.Body.Close()
	var response map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		log.Println("Error decoding response:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response from frontend"})
		return
	}
	c.JSON(resp.StatusCode, response)
}

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		setHeaders(c)
		if !validateAPIKey(c) {
			return
		}

		var user models.User
		if err := c.BindJSON(&user); err != nil {
			log.Println("Error binding JSON:", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
			return
		}

		if err := validate.Struct(user); err != nil {
			log.Println("Validation error:", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Validation failed"})
			return
		}

		signupURL := "http://10.10.10.83:5000/api/auth/signup"
		resp, err := forwardRequest(signupURL, user)
		if err != nil {
			log.Println("Error forwarding request:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to forward request to frontend"})
			return
		}

		handleResponse(c, resp)
	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		setHeaders(c)
		if !validateAPIKey(c) {
			return
		}

		var user models.User
		if err := c.BindJSON(&user); err != nil {
			log.Println("Error binding JSON:", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
			return
		}

		loginURL := "http://10.10.10.83:5000/api/auth/signin"
		resp, err := forwardRequest(loginURL, user)
		if err != nil {
			log.Println("Error forwarding request:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to forward request to frontend"})
			return
		}

		handleResponse(c, resp)
	}
}
