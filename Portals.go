package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"tool.com/Box/Other"
)

func main() {
	Other.My_App = app.New()
	Checkwindow := Other.My_App.NewWindow("错误")
	Checkwindow.Resize(fyne.NewSize(500, 400))
	State := Other.Init_Config()
	if State == 400 {
		Checkwindow.Show()
		dialog.NewConfirm("错误", "“Config”文件下未发现”Tools_Config1.json“文件！", func(Bool bool) {
			Checkwindow.Close()
		}, Checkwindow).Show()

	} else {
		Logo, _ := fyne.LoadResourceFromPath("四次元口袋.ico")
		Other.My_App.SetIcon(Logo)
		Other.MyWindow = Other.My_App.NewWindow("四次元口袋  by 今晚要吃三碗饭")
		Other.MyWindow.Resize(fyne.NewSize(1000, 600))
		Other.MyWindow.SetContent(container.NewHBox(
			container.NewGridWrap(fyne.NewSize(580, 580), Other.Left_AppTabs()),
			container.NewGridWrap(fyne.NewSize(410, 580), Other.Right_AppTabs()),
		),
		)
		Other.MyWindow.Show()
	}
	Other.My_App.Run()
}
