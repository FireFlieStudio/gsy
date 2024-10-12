package main

import (
	"fmt"
	"gsync/TencentCos"
)

func main() {
	fileManger := TencentCos.NewFileManager("backup")
	r := fileManger.AutoList("upload")
	for _, v := range r {
		for _, i := range v.Contents {
			fmt.Println(i.Key, i.ETag, i.Size)
		}
	}
}
