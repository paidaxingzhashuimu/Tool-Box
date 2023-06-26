config.json文件下（可自己自行修改json表，添加工具）
{
    "简简单单一顿扫": {//第一个键是容器名(例如下面的OA_反序列化，就是gui左边的键)，不能修改(需要添加容器，找作者修改)
        "目录扫描": [
            {
                "Tool_Name": "7kbscan",                                           //按钮显示的名字
                "Tool_Path": "Tools\\简简单单一顿扫\\目录扫描\\7kbscan",                //工具路径
                "Tool_Flag": true,                                                    //标志位，若是工具有gui界面，填true，否则false
                "Tool_Exe": "7kbscan-WebPathBrute.exe",             //若flag=true，则填写启动exe名称
                "Tool_Args": []
            },
            {
                "Tool_Name": "dirsearch_403",
                "Tool_Path": "Tools\\简简单单一顿扫\\目录扫描\\dirsearch_403",
                "Tool_Flag": false,
                "Tool_Exe": "",
                "Tool_Args": [
                                        "python dirsearch.py -b yes -u "        //若flag=false，则填写启动快捷语句
                                     ]
            }
}
拥有gui界面的工具可点击运行工具，右边可修改工具备注，保存。
![图片](https://github.com/paidaxingzhashuimu/Tool-Box/assets/103090032/6ba82e04-d416-4d3d-8975-577d57a4c9c1)

命令行工具，可自行在json表，添加快捷语句，执行。
![图片](https://github.com/paidaxingzhashuimu/Tool-Box/assets/103090032/b7e8ffff-1eb9-4002-a19f-4087f2a1b3a1)

            
