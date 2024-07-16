package main

import (
	"embed"
	"fmt"
	"stickyNote/controllers"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS



func main() {
	// Initialize the Gin router
	router := gin.Default()

	// Configure CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://10.10.83:5000"},  
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept", "x-api-key"},
	}))

	// Define routes
	router.POST("/api/login", controllers.Login())
	router.POST("/api/signup", controllers.SignUp())

	// Run the Gin server in a separate goroutine
	go func() {
		if err := router.Run(":8081"); err != nil {
			fmt.Printf("Failed to run Gin server: %v\n", err)
		}
	}()

	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "stickyNote",
		Width:  700,
		Height: 700,

		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,  // Fixed: Removed the need to pass a function with context
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
