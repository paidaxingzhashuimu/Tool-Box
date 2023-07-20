package Other

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"os/exec"
	"sort"
	"syscall"
)

// 遍历字典中第二梯度中的一个工具数组，生成按钮集合容器。
func Button_handle(Title_name string, Tool_parms []Tool, flag int) (Area_contain *fyne.Container) {
	Title := widget.NewLabel(Title_name)
	var Box []fyne.CanvasObject
	for _, Value := range Tool_parms {
		This_Value := Value
		Tmp_Button := widget.NewButton(This_Value.Tool_Name, func() {
			Current_tool.Tool_Path = This_Value.Tool_Path
			READNOTES_Handel(This_Value.Tool_Path)
			Use_stute_label.SetText("当前使用工具>> " + This_Value.Tool_Name)
			if This_Value.Tool_Flag == true {
				Use_Gui_Check.SetChecked(true)
				Use_option.SetText("")
				Use_option.Disable()
				Use_run_button.SetText("运行工具")
				Use_run_button.OnTapped = func() {
					cmd := exec.Command("cmd", "/c", This_Value.Tool_Exe)
					cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
					cmd.Dir = This_Value.Tool_Path
					_ = cmd.Start()
				}
			}
			if This_Value.Tool_Flag == false {
				Use_Gui_Check.SetChecked(false)
				Use_option.SetOptions(This_Value.Tool_Args)
				Use_option.SetText("")
				Use_option.Enable()
				Use_run_button.SetText("执行语句")
				Use_option.Hidden = false
				Use_run_button.OnTapped = func() {
					cmd := exec.Command("cmd", "/c", "start cmd /k "+Use_option.Text)
					cmd.Dir = This_Value.Tool_Path
					_ = cmd.Start()
				}
			}
		})
		//根据json中flag标志，生成按钮执行模式。
		Box = append(Box, Tmp_Button)
		//单个工具模块的容器集合。
	}

	//解决最后一个工具模块不添加下横线的为题。
	if flag == 1 {
		Area_contain = container.New(layout.NewVBoxLayout(), //容器排列模式，竖着叠加。
			container.NewCenter(Title),
			container.NewGridWrap(fyne.NewSize(108, 50), Box...), //网格排列容器，能设置每个容器的大小。
		)
	}
	if flag == 0 {
		Area_contain = container.New(layout.NewVBoxLayout(),
			container.NewCenter(Title),
			container.NewGridWrap(fyne.NewSize(108, 50), Box...),
			canvas.NewLine(color.Black),
		)
	}

	return
}

// 对多个数组的按钮集合容器进行加工，返回左侧一个页面。
func RETURN_AppTabs(Config_Key string) *fyne.Container {
	var Tmp []fyne.CanvasObject
	var Keys []string
	//获取json配置表中第二梯度的参数，加入数组中，进行升序排列。
	for Key, _ := range Tools_Config_Json[Config_Key] {
		Keys = append(Keys, Key)
	}
	sort.Strings(Keys)
	num := len(Keys)

	//根据keys数组进行工具模块排列，确保工具模块不会乱序排列，解决map底层，随机遍历问题。
	for Num, Value := range Keys {
		var tmp *fyne.Container
		if Num == num-1 {
			//判断是否是最后一个工具模块，设置flag为1。
			tmp = Button_handle(Value, Tools_Config_Json[Config_Key][Value], 1)
		} else {
			tmp = Button_handle(Value, Tools_Config_Json[Config_Key][Value], 0)
		}
		//将每一个返回的单个数组按钮容器，放进一个数组中。
		Tmp = append(Tmp, tmp)
	}

	//居中排列多个数组按钮容器。
	Model_Contain := container.NewBorder(canvas.NewLine(color.Black),
		canvas.NewLine(color.Black),
		canvas.NewLine(color.Black),
		canvas.NewLine(color.Black),
		container.NewVBox(
			Tmp..., //竖着排列容器
		),
	)

	return Model_Contain
}

func Left_AppTabs() (Left_contain *fyne.Container) {
	var TabItem []*container.TabItem

	//遍历json配置表中的key，放入keys数组中，进行排列。

	//根据keys数组进行排列左侧模块页面，保证不会乱序。
	for _, Value := range ParentKeys {
		Tmp := container.NewTabItem(Value, RETURN_AppTabs(Value))
		TabItem = append(TabItem, Tmp)
	}

	//添加进左侧页面容器生成。
	tabs := container.NewAppTabs(
		TabItem...,
	)
	tabs.SetTabLocation(container.TabLocationLeading)
	Left_contain = container.New(layout.NewGridLayout(1), tabs)
	return
}
