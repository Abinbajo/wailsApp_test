package main

import (
	"context"
	"log"
	"os"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"github.com/getlantern/systray"
)


type App struct {
	ctx context.Context
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	go systray.Run(a.onReady, a.onExit)
}

func (a *App) onReady() {
	icon, err := LoadIcon("assets/image/icon.ico") 
	if err != nil {
		log.Fatalf("Failed to load icon: %v", err)
		return
	}

	systray.SetIcon(icon)

	systray.SetTooltip("Sticky Note Application")

	mQuit := systray.AddMenuItem("Quit", "Quit the application")

	go func() {
		for range mQuit.ClickedCh {
			systray.Quit()
			a.CloseWindow()
		}
	}()
}

func (a *App) onExit() {
	
}

func (a *App) CloseWindow() {
	runtime.Quit(a.ctx)
}

func LoadIcon(path string) ([]byte, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return content, nil
}