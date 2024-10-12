package TencentCos

import (
	"context"
	"gsync/logger"
	"strings"
)

func (f *FileManager) CreateFolder(key string) {
	key = FolderFormater(key)
	if _, err := f.client.Object.Put(context.Background(), key, strings.NewReader(""), nil); err != nil {
		logger.Error("文件夹创建失败[ %s ] %s", key, err)
	} else {
		logger.Info("文件夹创建成功[ %s ]", key)
	}
}

func (f *FileManager) DeleteFolder(key string) {
	key = FolderFormater(key)
	if _, err := f.client.Object.Delete(context.Background(), key); err != nil {
		logger.Error("文件夹删除失败[ %s ] %s", key, err)
	} else {
		logger.Info("文件夹删除成功[ %s ]", key)
	}
}

func FolderFormater(key string) string {
	if !strings.HasSuffix(key, "/") {
		return key + "/"
	}
	return key
}
