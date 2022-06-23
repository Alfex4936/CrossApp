package main

import (
	"embed"
	"fmt"
	"time"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

//go:embed frontend/dist
var assets embed.FS

func main() {
	currentTime := time.Now()
	date := currentTime.Format("2006-01-02")

	// Create an instance of the app structure
	app := NewApp()

	AppMenu := menu.NewMenu()
	TextMenu := AppMenu.AddSubmenu(date)
	TextMenu.AddText("Rerun", keys.CmdOrCtrl("i"), func(_ *menu.CallbackData) {
		fmt.Print("Rerun\n")
		runtime.WindowReloadApp(app.ctx)
	})
	FileMenu := AppMenu.AddSubmenu("File")
	FileMenu.AddText("Quit", keys.CmdOrCtrl("q"), func(_ *menu.CallbackData) {
		runtime.Quit(app.ctx)
	})

	// Create application with options
	err := wails.Run(&options.App{
		Title:             "아주대학교 앱",
		Width:             1280,
		Height:            786,
		Menu:              AppMenu,
		DisableResize:     false,
		Fullscreen:        false,
		Frameless:         false,
		StartHidden:       false,
		HideWindowOnClose: false,
		Assets:            assets,
		OnStartup:         app.startup,
		OnShutdown:        app.shutdown,
		Bind: []interface{}{
			app,
		},
		Windows: &windows.Options{
			WebviewIsTransparent: false,
			WindowIsTranslucent:  false,
			DisableWindowIcon:    false,
		},
	})

	if err != nil {
		println("Error:", err)
	}
}
