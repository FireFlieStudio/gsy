package TencentCos

import (
	"context"
	"gsync/logger"
	"strings"
)

func (f *FileManager) CreateFolder(key string) {
	if !FolderFormatCheck(key) {
		logger.Error("文件夹格式错误,无法创建[ %s ]", key)
	}
	if _, err := f.client.Object.Put(context.Background(), key, strings.NewReader(""), nil); err != nil {
		logger.Error("文件夹创建失败[ %s ] %s", key, err)
	} else {
		logger.Info("文件夹创建成功[ %s ]", key)
	}
}

func (f *FileManager) DeleteFolder(key string) {
	if !FolderFormatCheck(key) {
		logger.Error("文件夹格式错误,无法删除[ %s ]", key)
	}
	if _, err := f.client.Object.Delete(context.Background(), key); err != nil {
		logger.Error("文件夹删除失败[ %s ] %s", key, err)
	} else {
		logger.Info("文件夹删除成功[ %s ]", key)
	}
}

func FolderFormatCheck(key string) bool {
	return strings.HasSuffix(key, "/")
}
