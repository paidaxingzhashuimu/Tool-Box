package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"tool.com/Gui/Other_go"
)

func main() {
	Stute := Other_go.Init_Config()
	if Stute == 400 {
		return
	}
	Other_go.My_App = app.New()
	Logo, _ := fyne.LoadResourceFromPath("四次元口袋.ico")
	Other_go.My_App.SetIcon(Logo)
	myWindow := Other_go.My_App.NewWindow("四次元口袋  by 今晚要吃三碗饭")
	myWindow.Resize(fyne.NewSize(1000, 600))
	myWindow.SetContent(container.NewHBox(
		container.NewGridWrap(fyne.NewSize(570, 580), Other_go.Left_AppTabs()),
		container.NewGridWrap(fyne.NewSize(410, 580), Other_go.Right_AppTabs()),
	),
	)
	myWindow.Show()
	Other_go.My_App.Run()
}
