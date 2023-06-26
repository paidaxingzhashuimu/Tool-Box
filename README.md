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
            },