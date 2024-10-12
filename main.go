package main

import "gsync/TencentCos"

func main() {
	fileManager := TencentCos.NewFileManager("backup")
	fileManager.DefaultUpload("upload", "C:\\Users\\Tomato\\Downloads\\图吧工具箱2409安装程序.exe")
}
