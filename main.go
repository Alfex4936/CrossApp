package main

import (
	"embed"
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
	TextMenu.AddText("재시작", keys.CmdOrCtrl("i"), func(_ *menu.CallbackData) {
		runtime.WindowReloadApp(app.ctx)
	})
	FileMenu := AppMenu.AddSubmenu("옵션")
	FileMenu.AddText("종료", keys.CmdOrCtrl("q"), func(_ *menu.CallbackData) {
		runtime.Quit(app.ctx)
	})

	// Create application with options
	err := wails.Run(&options.App{
		Title:             "아주대학교",
		Width:             1280,
		Height:            786,
		Menu:              AppMenu,
		DisableResize:     false,
		Fullscreen:        false,
		Frameless:         true,
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
