package main

import "gsync/TencentCos"

func main() {
	fileManager := TencentCos.NewFileManager("backup")
	fileManager.DefaultUpload("upload", "C:\\Users\\Tomato\\Downloads\\图吧工具箱2409安装程序.exe")
	fileManager.DefaultDownload("upload/图吧工具箱2409安装程序.exe", "./图吧工具箱2409安装程序.exe")
}
