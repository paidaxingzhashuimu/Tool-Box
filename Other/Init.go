package Other

import (
	"encoding/json"
	"fmt"
	"github.com/flopp/go-findfont"
	"github.com/goki/freetype/truetype"
	"io"
	"os"
	"sort"
)

func init() {
	fontPath, err := findfont.Find("Fonts\\STXINWEI.TTF")
	if err != nil {
		panic(err)
	}
	fontData, err := os.ReadFile(fontPath)
	if err != nil {
		panic(err)
	}
	_, err = truetype.Parse(fontData)
	if err != nil {
		panic(err)
	}
	_ = os.Setenv("FYNE_FONT", fontPath)
} //初始化字体，使用中文字体。

func Init_Config() int {
	SonKeys = make(map[string][]string)
	Open_Tools_Config, err := os.Open("Config\\Tools_Config.json")
	defer Open_Tools_Config.Close()
	if err != nil {
		fmt.Println("[-]Tools_config文件路径出错，请检测！")
		return 400
	}

	Read_Tools_Config, _ := io.ReadAll(Open_Tools_Config)
	_ = json.Unmarshal(Read_Tools_Config, &Tools_Config_Json)

	for Key, Value := range Tools_Config_Json {
		ParentKeys = append(ParentKeys, Key)
		for key, value := range Value {
			SonKeys[Key] = append(SonKeys[Key], key)
			for _, v := range value {
				Tool_Id = append(Tool_Id, v.Tool_Id)
			}
		}
		sort.Strings(SonKeys[Key])
	}
	sort.Ints(Tool_Id)
	sort.Strings(ParentKeys)
	return 200
} //读取json配置表。
