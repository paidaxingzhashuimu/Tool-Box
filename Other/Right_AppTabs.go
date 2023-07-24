package Other

import (
	"bufio"
	"encoding/json"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"syscall"
)

func Think() (Think_Notes string) {
	//Open_Think_Notes, _ := os.Open("Config\\Think.txt")
	Read_Think_Notes, _ := os.ReadFile("Config\\Think.txt")
	Think_Notes = string(Read_Think_Notes)
	return
}

func About_Author() (Version_Data string) {
	Version_Data =
		">>By 今晚要吃三碗饭\n\n" +
			">>当前版本 V4.1\n\n" +
			">>优化检测json表的判断\n\n" +
			">>历史版本 V4.0\n\n" +
			">>系统设置处可直接添加工具\n\n" +
			">>优化了部份模块代码，弹窗提示\n\n" +
			">>历史版本 V3.1\n\n" +
			">>优化了点击工具自动生成READNOTES.md文件\n\n" +
			"仅在保存修改的时候才会生成此文件。\n\n" +
			">>优化了部份模块代码，感谢栏不再为静态代码。\n\n" +
			">>历史版本 V3.0\n\n" +
			">>修复了模块排列乱序问题。\n\n" +
			">>添加了Thinkphp多个利用工具。\n\n" +
			">>json表全参数可修改，左侧模块会读取json配置表的\n\n" +
			"第一梯度参数，进行升序排列显示。\n\n" +
			">>历史版本 V2.1\n\n" +
			">>修复了部份bug,Gui勾选导致运行按钮不显示问题。\n\n"
	return
}

func READNOTES_Handel(Tool_Path string) {
	Read_READNOTES, _ := os.ReadFile(Tool_Path + "\\READNOTES.md")
	Use_READNOTES_Label.SetText(string(Read_READNOTES))
}

func Add_Tool() (System_Setup_Contain *fyne.Container) {
	var Gui_Flag int
	var Tmp_Flag_1 string
	var Tmp_Flag_2 string

	Select_EXE := widget.NewMultiLineEntry()    //选择文件路径
	Select_Folder := widget.NewMultiLineEntry() //选择文件目录
	Select_EXE.SetMinRowsVisible(1)             //
	Select_Folder.SetMinRowsVisible(1)

	Select_Moudle := widget.NewSelect([]string{}, func(Moudle string) {
		//Tmp_Moudle = Moudle
	}) //选择容器标签
	Select_Tap := widget.NewSelect(ParentKeys, func(Tap string) {
		//Tmp_Tap = Tap
		Select_Moudle.Options = SonKeys[Tap]
		if Tap != Tmp_Flag_1 {
			Select_Moudle.ClearSelected()
		}
		Tmp_Flag_1 = Tap
	}) //选择模块

	Select_Moudle.PlaceHolder = "选择工具模块" //默认背景
	Select_Tap.PlaceHolder = "选择左侧容器标签"

	Tool_Name := widget.NewEntry() //设置显示工具名字
	Tool_Name.SetPlaceHolder("工具显示名字")

	Tool_Args := widget.NewMultiLineEntry()
	Tool_Args.SetPlaceHolder("填写快捷语句,”换行“隔开")

	//打开EXE文件
	Open_File := dialog.NewFileOpen(func(File fyne.URIReadCloser, err error) {
		if File != nil {
			Select_EXE.SetText(File.URI().Path())
		}
	}, MyWindow)
	Open_File.SetDismissText("取消")
	Open_File.SetConfirmText("打开")
	//打开目录
	Open_Folder := dialog.NewFolderOpen(func(Folder fyne.ListableURI, err error) {
		if Folder != nil {
			Select_Folder.SetText(Folder.Path())
		}
	}, MyWindow)
	Open_Folder.SetDismissText("取消")
	Open_Folder.SetConfirmText("打开")

	Select_EXE_Button := widget.NewButton("选择文件", func() {
		Open_File.SetFilter(storage.NewExtensionFileFilter([]string{".exe", ".jar"}))
		Open_File.Show()
	}) //选择文件按钮
	Select_Folder_Button := widget.NewButton("选择目录", func() {
		Open_Folder.Show()
	}) //选择目录按钮

	Add_Tool_Button := widget.NewButton("添加工具", func() {
		var Tmp Tool
		sort.Ints(Tool_Id)
		if Select_EXE.Text != "" || Select_Folder.Text != "" && Tool_Name.Text != "" && Select_Moudle.Selected != "" && Select_Tap.Selected != "" {
			if Gui_Flag == 1 {
				Tmp = Tool{
					Tool_Id[(len(Tool_Id)-1)] + 1,
					Tool_Name.Text,
					Select_Tap.Selected,
					Select_Moudle.Selected,
					true,
					filepath.Base(Select_EXE.Text),
					[]string{},
					filepath.Dir(Select_EXE.Text),
				}
			}
			if Gui_Flag == 0 {
				Tmp = Tool{
					Tool_Id[(len(Tool_Id)-1)] + 1,
					Tool_Name.Text,
					Select_Tap.Selected,
					Select_Moudle.Selected,
					false,
					"",
					strings.Split(Tool_Args.Text, "\n"),
					Select_Folder.Text,
				}
			}
			Tools_Config_Json[Select_Tap.Selected][Select_Moudle.Selected] = append(Tools_Config_Json[Select_Tap.Selected][Select_Moudle.Selected], Tmp)
			dialog.NewInformation("提示", "添加工具 “"+Tmp.Tool_Name+" ”成功", MyWindow).Show()
			Control_Log.Text += ">>添加工具“" + Tmp.Tool_Name + "（id:" + strconv.Itoa(Tmp.Tool_Id) + "）”成功\n\n"
			Control_Log.Refresh()
			Tool_Id = append(Tool_Id, Tmp.Tool_Id)
			sort.Ints(Tool_Id)
		} else {
			dialog.NewInformation("提示", "你还有选项栏没填写呢！", MyWindow).Show()
		}

	}) //添加工具按钮                                            //添加工具按钮
	Advanced_Setup := widget.NewButton("高级设置", func() {
		var Flag string
		var Tmp_Tool Tool
		var Search []Tool
		Newwindow := My_App.NewWindow("高级设置")
		Newwindow.Resize(fyne.NewSize(450, 500))

		Widget1 := widget.NewEntry()
		Widget1.SetPlaceHolder("请输入工具名进行搜索")
		Widget2 := widget.NewMultiLineEntry()
		Widget3 := widget.NewEntry()
		Widget5 := widget.NewSelect([]string{}, func(value string) {})
		Widget5.PlaceHolder = " "
		Widget4 := widget.NewSelect(ParentKeys, func(Tap string) {
			//Tmp_Tap = Tap
			Widget5.Options = SonKeys[Tap]
			if Tap != Tmp_Flag_2 {
				Widget5.ClearSelected()
			}
			Tmp_Flag_2 = Tap
		})
		Widget4.PlaceHolder = " "
		Widget6 := widget.NewEntry()
		Widget7 := widget.NewSelect([]string{"true", "false"}, func(value string) {})
		Widget7.PlaceHolder = " "
		Widget8 := widget.NewEntry()
		Widget9 := widget.NewSelect([]string{}, func(value string) {
			for _, Value := range Search {
				if Value.Tool_Name+"（id:"+strconv.Itoa(Value.Tool_Id)+"）" == value {
					Tmp_Tool = Value
					Widget2.SetText(strings.Join(Value.Tool_Args, "\n"))
					Widget3.SetText(Value.Tool_Name)
					Widget4.SetSelected(Value.Tab_Name)
					Widget5.SetSelected(Value.Moudle_Name)
					Widget6.SetText(Value.Tool_Exe)
					//Widget7.Options = []string{"true", "false"}
					Widget7.SetSelected(strconv.FormatBool(Value.Tool_Flag))
					Widget8.SetText(Value.Tool_Path)
				}
			}
		})
		Widget9.PlaceHolder = " "

		Add_SelectEntry := widget.NewSelectEntry(ParentKeys)
		Add_SelectEntry.PlaceHolder = "选择一个容器标签"
		Add_Entry := widget.NewEntry()

		Delete_Select2 := widget.NewSelect([]string{}, func(value string) {
		})
		Delete_Select1 := widget.NewSelect(ParentKeys, func(value string) {
			Delete_Select2.Options = SonKeys[value]
			if value != Flag {
				Delete_Select2.ClearSelected()
			}
			Flag = value
		})
		Delete_Select1.PlaceHolder = "选择一个容器标签"
		Delete_Select2.PlaceHolder = "选择一个模块标签"

		Add_Button := widget.NewButton("新增", func() {
			if Add_SelectEntry.Text != "" && Add_Entry.Text != "" {
				dialog.NewConfirm("提示", "添加容器标签"+" ”"+Add_SelectEntry.Text+"” 下的模块 “"+Add_Entry.Text+"” ?", func(Bool bool) {
					if Bool == true {
						_, ok := Tools_Config_Json[Add_SelectEntry.Text][Add_Entry.Text]
						if ok == true {
							dialog.NewInformation("提示", "该容器标签下的模块已存在，请勿再次添加。", Newwindow).Show()
						} else {
							Transform := make(map[string][]Tool)
							Transform[Add_Entry.Text] = []Tool{}
							Tools_Config_Json[Add_SelectEntry.Text] = Transform                                   //总表新增数据。
							SonKeys[Add_SelectEntry.Text] = append(SonKeys[Add_SelectEntry.Text], Add_Entry.Text) //Son表新增数据。

							dialog.NewInformation("提示", "添加容器标签"+" ”"+Add_SelectEntry.Text+"” 下的模块 “"+Add_Entry.Text+"” 成功", Newwindow).Show()
							Control_Log.Text += ">>添加容器标签" + " ”" + Add_SelectEntry.Text + "” 下的模块 “" + Add_Entry.Text + "” 成功\"\n\n"

							Add_SelectEntry.SetOptions(ParentKeys)
							Select_Tap.Options = ParentKeys
							Delete_Select1.Options = ParentKeys
							//实时同步工具添加栏的参数
							Select_Tap.ClearSelected()
							Select_Moudle.ClearSelected()
							//实时同步搜索栏的参数
							Widget1.SetText("")
							Widget2.SetText("")
							Widget3.SetText("")
							Widget4.ClearSelected()
							Widget5.ClearSelected()
							Widget6.SetText("")
							Widget7.ClearSelected()
							Widget8.SetText("")
							Widget9.Options = []string{}
							Widget9.ClearSelected()

							//实时同步删除栏
							Delete_Select1.ClearSelected()
							Delete_Select2.ClearSelected()
						}
					}
				}, Newwindow).Show()
			} else {
				dialog.NewInformation("提示", "容器标签或者模块没填入呢", Newwindow).Show()
			}
		})
		Delete_Button := widget.NewButton("删除", func() {
			if Delete_Select1.Selected != "" { //容器标签不为空。
				var Delete_Id []int
				if Delete_Select2.Selected != "" { //模块标签不为空。
					dialog.NewConfirm("提示", "确认删除 ”"+Delete_Select1.Selected+"“ 容器标签下的 ”"+Delete_Select2.Selected+"” 模块？\n\n这样会导致该模块下的所有工具都会删除。", func(Bool bool) {
						if Bool == true {
							if Tools_Config_Json[Delete_Select1.Selected][Delete_Select2.Selected] != nil {
								for _, value := range Tools_Config_Json[Delete_Select1.Selected][Delete_Select2.Selected] {
									Delete_Id = append(Delete_Id, value.Tool_Id)
								}
								for _, V := range Delete_Id {
									for k, v := range Tool_Id {
										if V == v {
											Tool_Id = append(Tool_Id[:k], Tool_Id[k+1:]...)
										}
									}
								}
								sort.Ints(Tool_Id)
							}

							//删除Son表中的模块数据
							for Key, Value := range SonKeys[Delete_Select1.Selected] {
								if Delete_Select2.Selected == Value {
									SonKeys[Delete_Select1.Selected] = append(SonKeys[Delete_Select1.Selected][:Key], SonKeys[Delete_Select1.Selected][Key+1:]...)
									sort.Strings(SonKeys[Delete_Select1.Selected])
								}
							}
							//删除总表中的模块数据
							delete(Tools_Config_Json[Delete_Select1.Selected], Delete_Select2.Selected)

							//Select_Tap.Options = ParentKeys        //同步添加工具的容器标签数据。
							//Delete_Select1.Options = ParentKeys    //同步删除栏的容器标签数据。
							//Add_SelectEntry.SetOptions(ParentKeys) //同步添加容器标签数据。
							dialog.NewInformation("提示", "删除 ”"+Delete_Select1.Selected+"“ 容器标签下的 ”"+Delete_Select2.Selected+"” 模块成功", Newwindow)
							Control_Log.Text += ">>删除 ”" + Delete_Select1.Selected + "“ 容器标签下的 ”" + Delete_Select2.Selected + "” 模块成功\n\n"
							//实时同步工具添加栏的参数
							Select_Tap.ClearSelected()
							Select_Moudle.ClearSelected()
							//实时同步搜索栏的参数

							Widget2.SetText("")
							Widget3.SetText("")
							Widget4.ClearSelected()
							Widget5.ClearSelected()
							Widget6.SetText("")
							Widget7.ClearSelected()
							Widget8.SetText("")
							Widget9.Options = []string{}
							Widget9.ClearSelected()
							//实时同步添加栏的参数
							Add_SelectEntry.SetText("")
							Add_SelectEntry.Refresh()
							Add_Entry.SetText("")
							//实时同步删除栏
							Delete_Select1.ClearSelected()
							Delete_Select2.ClearSelected()
						}
					}, Newwindow).Show()

				} else {
					dialog.NewConfirm("提示", "确认删除 ”"+Delete_Select1.Selected+"“ 容器标签？\n\n这样会导致容器标签下所有工具都会删除。", func(Bool bool) {
						if Bool == true {
							for _, Value := range Tools_Config_Json[Delete_Select1.Selected] {
								if Value != nil {
									for _, v := range Value {
										Delete_Id = append(Delete_Id, v.Tool_Id)
									}
								}
							} //遍历json表，查看标签容器下是否有工具存在，然后将需要删除的工具放入delete数组。
							if Delete_Id != nil {
								sort.Ints(Delete_Id)
								for _, V := range Delete_Id {
									for k, v := range Tool_Id {
										if V == v {
											Tool_Id = append(Tool_Id[:k], Tool_Id[k+1:]...)
										}
									}
								}
							} //删除工具id数组参数

							sort.Ints(Tool_Id)
							delete(Tools_Config_Json, Delete_Select1.Selected) //删除总表容器标签数据。
							delete(SonKeys, Delete_Select1.Selected)           //删除Son表容器标签数据。
							for k, v := range ParentKeys {
								if v == Delete_Select1.Selected {
									ParentKeys = append(ParentKeys[:k], ParentKeys[k+1:]...)
								}
							} //删除Parent容器标签数据。

							Select_Tap.Options = ParentKeys        //同步添加工具的容器标签数据。
							Delete_Select1.Options = ParentKeys    //同步删除栏的容器标签数据。
							Add_SelectEntry.SetOptions(ParentKeys) //同步添加容器标签数据。
							dialog.NewInformation("提示", "删除 ”"+Delete_Select1.Selected+"“ 容器标签成功", Newwindow).Show()
							Control_Log.Text += ">>删除 ”" + Delete_Select1.Selected + "“ 容器标签成功\n\n"
							//实时同步工具添加栏的参数
							Select_Tap.ClearSelected()
							Select_Moudle.ClearSelected()
							//实时同步搜索栏的参数

							Widget2.SetText("")
							Widget3.SetText("")
							Widget4.ClearSelected()
							Widget5.ClearSelected()
							Widget6.SetText("")
							Widget7.ClearSelected()
							Widget8.SetText("")
							Widget9.Options = []string{}
							Widget9.ClearSelected()
							//实时同步添加栏的参数
							Add_SelectEntry.SetText("")
							Add_SelectEntry.Refresh()
							Add_Entry.SetText("")
							//实时同步删除栏
							Delete_Select1.ClearSelected()
							Delete_Select2.ClearSelected()
						}
					}, Newwindow).Show()
				}

			} else {
				dialog.NewInformation("提示", "你还没选择容器标签呢", Newwindow).Show()
			}
		})

		Newwindow.SetContent(container.NewVBox(
			container.NewCenter(container.NewHBox(container.NewGridWrap(fyne.NewSize(160, 30), Widget1),
				widget.NewButton("搜索", func() {
					if Widget1.Text == "" {
						dialog.NewInformation("提示", "搜索框没输入参数", Newwindow).Show()
					} else {
						Search = nil
						var Name []string
						Widget2.SetText("")
						Widget3.SetText("")
						Widget4.ClearSelected()
						Widget5.ClearSelected()
						Widget6.SetText("")
						Widget7.ClearSelected()
						Widget8.SetText("")
						for _, Value := range Tools_Config_Json {
							for _, value := range Value {
								for _, v := range value {
									if strings.Contains(v.Tool_Name, Widget1.Text) {
										Search = append(Search, v)
										Name = append(Name, v.Tool_Name+"（id:"+strconv.Itoa(v.Tool_Id)+"）")
									}
								}
							}
						}
						if Name != nil {
							Widget9.Options = Name
							Widget9.PlaceHolder = "搜索“" + Widget1.Text + "”成功,下拉查看结果"
							Widget9.ClearSelected()
						} else {
							Widget9.PlaceHolder = "未查询到“" + Widget1.Text + "”信息"
							Widget9.Options = []string{}
							Widget9.ClearSelected()
						}
					}
				}))),
			widget.NewForm(
				widget.NewFormItem("搜索结果", Widget9),
				widget.NewFormItem("工具名", Widget3),
				widget.NewFormItem("容器标签", Widget4),
				widget.NewFormItem("模块名称", Widget5),
				widget.NewFormItem("执行exe", Widget6),
				widget.NewFormItem("交互方式标志位", Widget7),
				widget.NewFormItem("工具路径", Widget8),
				widget.NewFormItem("快捷语句", Widget2)),
			container.NewGridWithColumns(2,
				widget.NewButton("修改参数", func() {
					if Widget9.Selected != "" {
						dialog.NewConfirm("确认修改提示", ">>工具名：“"+Tmp_Tool.Tool_Name+"”->“"+Widget3.Text+"”\n\n"+
							">>容器标签：“"+Tmp_Tool.Tab_Name+"”->“"+Widget4.Selected+"”\n\n"+
							">>模块名称：“"+Tmp_Tool.Moudle_Name+"”->“"+Widget5.Selected+"”\n\n"+
							">>执行exe：“"+Tmp_Tool.Tool_Exe+"”->“"+Widget6.Text+"”\n\n"+
							">>交互方式标志位：“"+strconv.FormatBool(Tmp_Tool.Tool_Flag)+"”->“"+Widget7.Selected+"”\n\n"+
							">>工具路径：“"+Tmp_Tool.Tool_Path+"”->“"+Widget8.Text+"”\n\n"+
							">>快捷语句：“"+strings.Join(Tmp_Tool.Tool_Args, "\n")+"”->“"+Widget2.Text+"”\n\n",
							func(Bool bool) {
								if Bool == true {
									var New_Tool Tool
									New_Tool.Tool_Id = Tmp_Tool.Tool_Id
									New_Tool.Tool_Name = Widget3.Text
									New_Tool.Tab_Name = Widget4.Selected
									New_Tool.Moudle_Name = Widget5.Selected
									New_Tool.Tool_Flag, _ = strconv.ParseBool(Widget7.Selected)
									New_Tool.Tool_Exe = Widget6.Text
									New_Tool.Tool_Args = strings.Split(Widget2.Text, "\n")
									New_Tool.Tool_Path = Widget8.Text
									if Tmp_Tool.Tab_Name == Widget4.Selected && Tmp_Tool.Moudle_Name == Widget5.Selected {
										for Key, Value := range Tools_Config_Json[Tmp_Tool.Tab_Name][Tmp_Tool.Moudle_Name] {
											if Tmp_Tool.Tool_Id == Value.Tool_Id {
												Tools_Config_Json[Tmp_Tool.Tab_Name][Tmp_Tool.Moudle_Name][Key] = New_Tool
											}
										}
									} else if Tmp_Tool.Tab_Name == Widget4.Selected && Tmp_Tool.Moudle_Name != Widget5.Selected {
										for Key, Value := range Tools_Config_Json[Tmp_Tool.Tab_Name][Tmp_Tool.Moudle_Name] {
											if Tmp_Tool.Tool_Id == Value.Tool_Id {
												Tools_Config_Json[Tmp_Tool.Tab_Name][Widget5.Selected] = append(Tools_Config_Json[Tmp_Tool.Tab_Name][Widget5.Selected], New_Tool)
												Tools_Config_Json[Tmp_Tool.Tab_Name][Tmp_Tool.Moudle_Name] = append(Tools_Config_Json[Tmp_Tool.Tab_Name][Tmp_Tool.Moudle_Name][:Key], Tools_Config_Json[Tmp_Tool.Tab_Name][Tmp_Tool.Moudle_Name][Key+1:]...)
											}
										}
									} else if Tmp_Tool.Tab_Name != Widget4.Selected {
										for Key, Value := range Tools_Config_Json[Tmp_Tool.Tab_Name][Tmp_Tool.Moudle_Name] {
											if Tmp_Tool.Tool_Id == Value.Tool_Id {
												Tools_Config_Json[Widget4.Selected][Widget5.Selected] = append(Tools_Config_Json[Widget4.Selected][Widget5.Selected], New_Tool)
												Tools_Config_Json[Tmp_Tool.Tab_Name][Tmp_Tool.Moudle_Name] = append(Tools_Config_Json[Tmp_Tool.Tab_Name][Tmp_Tool.Moudle_Name][:Key], Tools_Config_Json[Tmp_Tool.Tab_Name][Tmp_Tool.Moudle_Name][Key+1:]...)

											}
										}
									}
									dialog.NewInformation("修改提示", "修改“"+Tmp_Tool.Tab_Name+"（id:"+strconv.Itoa(Tmp_Tool.Tool_Id)+"）“"+"成功", Newwindow).Show()
									Control_Log.Text += ">>修改“" + Tmp_Tool.Tab_Name + "（id:" + strconv.Itoa(Tmp_Tool.Tool_Id) + "）“" + "成功\n\n"
									Control_Log.Refresh()
								}
							}, Newwindow).Show()
					} else {
						dialog.NewInformation("提示", "还没选择需要修改的工具呢", Newwindow).Show()
					}

				}),
				widget.NewButton("删除工具", func() {
					if Widget9.Selected != "" {
						dialog.NewConfirm("删除提示", "确认删除 “"+Widget9.Selected+"” ?", func(Bool bool) {
							if Bool == true {
								for Key, Value := range Tools_Config_Json[Tmp_Tool.Tab_Name][Tmp_Tool.Moudle_Name] {
									if Tmp_Tool.Tool_Id == Value.Tool_Id {
										//删除总表中的工具
										Tools_Config_Json[Tmp_Tool.Tab_Name][Tmp_Tool.Moudle_Name] = append(Tools_Config_Json[Tmp_Tool.Tab_Name][Tmp_Tool.Moudle_Name][:Key], Tools_Config_Json[Tmp_Tool.Tab_Name][Tmp_Tool.Moudle_Name][Key+1:]...)
										for k, v := range Widget9.Options {
											if Tmp_Tool.Tool_Name+"（id:"+strconv.Itoa(Tmp_Tool.Tool_Id)+"）" == v {
												Widget9.Options = append(Widget9.Options[:k], Widget9.Options[k+1:]...)
											}
										} //删除查询结果
										for key, value := range Tool_Id {
											if Tmp_Tool.Tool_Id == value {
												Tool_Id = append(Tool_Id[:key], Tool_Id[key+1:]...)
												sort.Ints(Tool_Id)
											}
										} //删除id号

										dialog.NewInformation("删除提示", "删除 “"+Widget9.Selected+"” 成功", Newwindow).Show()
										Control_Log.Text += ">>删除 “" + Widget9.Selected + "” 成功\n\n"
										Control_Log.Refresh()
										Widget2.SetText("")
										Widget3.SetText("")
										Widget4.ClearSelected()
										Widget5.ClearSelected()
										Widget6.SetText("")
										Widget7.ClearSelected()
										Widget8.SetText("")
										Widget9.PlaceHolder = " "
										Widget9.ClearSelected()
									}
								}
							}
						}, Newwindow).Show()
					} else {
						dialog.NewInformation("修改提示", "请选择工具删除", Newwindow).Show()
					}
				})),
			container.NewGridWithColumns(3, Add_SelectEntry, Add_Entry, Add_Button),
			container.NewGridWithColumns(3, Delete_Select1, Delete_Select2, Delete_Button),
			widget.NewButton("重置", func() {
				Add_SelectEntry.SetText("")
				Add_Entry.SetText("")
				Delete_Select1.ClearSelected()
				Delete_Select2.ClearSelected()
				//dialog.NewInformation("提示","",Newwindow).Show()
			}),
		))
		Newwindow.Show()

	}) //高级设置按钮

	//判断是否有Gui界面
	CheckGui := widget.NewRadioGroup([]string{"Gui界面", "命令行界面"}, func(Value string) {
		if Value == "Gui界面" {
			Gui_Flag = 1
			Select_Folder.Disable()
			Select_Folder_Button.Disable()
			Tool_Args.Disable()

			Select_EXE.Enable()
			Select_EXE_Button.Enable()
		}
		if Value == "命令行界面" {
			Gui_Flag = 0
			Select_EXE.Disable()
			Select_EXE_Button.Disable()

			Select_Folder.Enable()
			Select_Folder_Button.Enable()
			Tool_Args.Enable()
		}
	})
	CheckGui.SetSelected("命令行界面")

	Use_Setup := widget.NewButton("应用保存", func() {
		dialog.NewConfirm("提示", "是否应用保存？", func(Check bool) {
			if Check == true {
				Open_Tools_Config, _ := os.OpenFile("Config\\Tools_Config.json", os.O_RDWR|os.O_TRUNC, 0666)
				defer Open_Tools_Config.Close()
				Write_Tools_Config := bufio.NewWriter(Open_Tools_Config)
				Transform, _ := json.MarshalIndent(Tools_Config_Json, "", "\t")
				_, _ = Write_Tools_Config.WriteString(string(Transform))
				_ = Write_Tools_Config.Flush()
				dialog.NewInformation("提示", "应用成功,重启生效", MyWindow).Show()
			}
		}, MyWindow).Show()
	})

	Box_Root_Folder := widget.NewEntry()
	Box_Root_Folder.SetPlaceHolder("可选择工具箱根目录路径")
	Box_Config_Json := widget.NewEntry()
	Box_Config_Json.SetPlaceHolder("可选择Json配置表路径")
	Box_Root_Folder_Open := dialog.NewFileOpen(func(File fyne.URIReadCloser, err error) {
		if File != nil {
			Box_Config_Json.SetText(File.URI().Path())
		}
	}, MyWindow)
	Box_Root_Folder_Open.SetFilter(storage.NewExtensionFileFilter([]string{".json"}))

	System_Setup_Contain = container.NewBorder(nil, Use_Setup, nil, nil, Use_Setup, container.NewVBox(
		widget.NewForm(widget.NewFormItem("工具交互类型：", CheckGui)),
		container.NewHBox(container.NewGridWrap(fyne.NewSize(340, 30), Select_EXE), container.NewGridWrap(fyne.NewSize(50, 30), Select_EXE_Button)),
		container.NewHBox(container.NewGridWrap(fyne.NewSize(340, 30), Select_Folder), container.NewGridWrap(fyne.NewSize(50, 30), Select_Folder_Button)),
		widget.NewForm(
			widget.NewFormItem("工具名：", Tool_Name),
			widget.NewFormItem("容器标签：", Select_Tap),
			widget.NewFormItem("模块名：", Select_Moudle),
			widget.NewFormItem("快捷语句：", Tool_Args)),
		container.NewGridWithColumns(2, Add_Tool_Button, Advanced_Setup),
		container.NewHBox(container.NewGridWrap(fyne.NewSize(340, 30), Box_Root_Folder), container.NewGridWrap(fyne.NewSize(50, 30), widget.NewButton("选择目录",
			func() {
				dialog.NewFolderOpen(func(Folder fyne.ListableURI, err error) {
					if Folder != nil {
						Box_Root_Folder.SetText(Folder.Path())
						Open_File.SetLocation(Folder)
						Open_Folder.SetLocation(Folder)
					}
				}, MyWindow).Show()
			}))),
		container.NewHBox(container.NewGridWrap(fyne.NewSize(340, 30), Box_Config_Json), container.NewGridWrap(fyne.NewSize(50, 30), widget.NewButton("选择文件",
			func() {
				Box_Root_Folder_Open.Show()
			}))),
	))
	return
}

func Right_AppTabs() (Right_contain *fyne.Container) {
	Use_READNOTES_Label = widget.NewEntry() //创建右窗口，备注模块。
	Control_Log = widget.NewMultiLineEntry()
	Use_stute_label = widget.NewLabel("当前使用工具>>")        //工具状态栏描述。
	Use_option = widget.NewSelectEntry([]string{})       //工具快捷语句选择框，先初始化置空。
	Use_option.PlaceHolder = "快捷语句~"                     //工具快捷语句选择框,默认背景语句。
	Use_option.MultiLine = true                          //工具快捷语句选择框，可多行输入。
	Use_READNOTES_Label.MultiLine = true                 //备注模块，可多行输入。
	Use_READNOTES_Label.SetPlaceHolder("这个人很懒，什么都没有留下~") //备注模块，默认背景语句。
	Use_run_button = widget.NewButton("", func() {       //初始化有窗口按钮。
	})
	Use_Gui_Check = widget.NewCheck("Gui界面", func(flag bool) {}) //设置有窗口右上角确认按钮，功能点已弃用。

	//右窗口下方，确认保存修改按钮，弹窗确定保存修改，避免误点。
	Use_button_1 = widget.NewButtonWithIcon("保存修改", theme.CheckButtonCheckedIcon(), func() { //增加类型的按钮容器，可添加图片。
		//确认窗口，表题为消息框。
		Save_Confirm := dialog.NewConfirm("消息框", "是否确定保存修改？", func(Flag bool) {
			if Flag == true {
				Open_READNOTES, _ := os.OpenFile(Current_tool.Tool_Path+"\\READNOTES.md", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0666)
				defer Open_READNOTES.Close()
				Write_READNOTES := bufio.NewWriter(Open_READNOTES) //打开当前工具的tip文件。
				data := Use_READNOTES_Label.Text                   //获取备注栏中的内容。
				_, _ = Write_READNOTES.WriteString(data)           //将备注栏的内容写入tip文件。
				_ = Write_READNOTES.Flush()
			}
		}, MyWindow)
		Save_Confirm.SetConfirmText("确定")
		Save_Confirm.SetDismissText("取消")
		Save_Confirm.Show()
	})

	//右窗口下方，弹窗工具目录。
	Use_button_2 = widget.NewButtonWithIcon("工具目录", theme.HomeIcon(), func() {
		cmd := exec.Command("cmd", "/c", "explorer", Current_tool.Tool_Path)
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
		_ = cmd.Start()
	})

	//有窗口模块容器组装。
	Use_contain := container.NewBorder(
		container.NewHBox(container.NewGridWrap(fyne.NewSize(300, 20), Use_stute_label), container.NewGridWrap(fyne.NewSize(90, 20), Use_Gui_Check)),
		container.NewVBox(container.NewHBox(container.NewGridWrap(fyne.NewSize(300, 60), Use_option),
			container.NewGridWrap(fyne.NewSize(90, 60), Use_run_button)),
			container.NewGridWithColumns(2, Use_button_1, Use_button_2),
		),
		nil, nil,
		Use_READNOTES_Label,
	)

	//工具多页面容器。
	tabs := container.NewAppTabs(
		container.NewTabItem("工具使用&备注", Use_contain),
		container.NewTabItem("工具添加", Add_Tool()),
		container.NewTabItem("操作日志", Control_Log),
		container.NewTabItem("关于工具", widget.NewLabel(About_Author())),
		container.NewTabItem("感谢", widget.NewLabel(Think())),
	)
	//多页面容器，排布方式。
	tabs.SetTabLocation(container.TabLocationTop)

	//有窗口多容器组装。
	Right_contain = container.NewBorder(canvas.NewLine(color.Black),
		canvas.NewLine(color.Black),
		canvas.NewLine(color.Black),
		canvas.NewLine(color.Black),
		tabs,
	)

	return
}

//	func Get_Tool_Informatain(Tool Tools, args ...interface{}) {
//		params := make([]reflect.Value, len(args))
//		for i, _ := range args {
//			params[i] = reflect.ValueOf(args[i])
//		}
//		myinstance := Tools{}
//		ref := reflect.ValueOf(myinstance)
//		method := ref.MethodByName(Tool.Tool_Name)
//		method.Call(params)
//	}
