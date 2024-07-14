package main

import (
	"context"
	"log"
	"os"

	"github.com/getlantern/systray"
)

// App is the main application structure
type App struct{}

// NewApp creates a new instance of App
func NewApp() *App {
	return &App{}
}

// startup initializes the system tray icon and menu.
func (a *App) startup(ctx context.Context) {
	// Initialize the systray
	systray.Run(onReady, func() {})
}

func onReady() {
	// Load the .ico file
	icon, err := LoadIcon("icon.ico") // Ensure this path is correct
	if err != nil {
		log.Fatalf("Failed to load icon: %v", err)
		return
	}

	// Set the tray icon
	systray.SetIcon(icon)
	// Set the tray icon's tooltip
	systray.SetTooltip("Sticky Note Application")

	// Add menu items
	mQuit := systray.AddMenuItem("Quit", "Quit the application")

	// Handle menu item clicks
	go func() {
		for range mQuit.ClickedCh{
			systray.Quit()
		}
	}()
}

// LoadIcon loads the .ico file from the specified path
func LoadIcon(path string) ([]byte, error) {
	// Read the .ico file content
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return content, nil
}
