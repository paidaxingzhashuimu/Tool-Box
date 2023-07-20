package Other

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

var My_App fyne.App
var SonKeys map[string][]string
var ParentKeys []string
var Tool_Id []int
var Control_Log *widget.Entry
var Current_tool Tool
var MyWindow fyne.Window
var Tools_Config_Json map[string]map[string][]Tool
var Use_READNOTES_Label *widget.Entry
var Use_button_1 *widget.Button
var Use_button_2 *widget.Button
var Use_run_button *widget.Button
var Use_Gui_Check *widget.Check
var Use_stute_label *widget.Label
var Use_option *widget.SelectEntry

type Tool struct {
	Tool_Id     int
	Tool_Name   string
	Tab_Name    string
	Moudle_Name string
	Tool_Flag   bool
	Tool_Exe    string
	Tool_Args   []string
	Tool_Path   string
}
