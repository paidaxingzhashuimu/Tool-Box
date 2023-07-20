package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"tool.com/Box/Other"
)

func main() {
	Stute := Other.Init_Config()
	if Stute == 400 {
		return
	}
	Other.My_App = app.New()
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
	Other.My_App.Run()
}

//			Control_Log.Text += ">>添加工具“" + Tmp.Tool_Name + "（id:" + strconv.Itoa(Tmp.Tool_Id) + "）”成功\n\n"
