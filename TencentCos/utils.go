package TencentCos

import (
	"path"
	"path/filepath"
	"strings"
)

func FolderFormater(key string) string {
	if !strings.HasSuffix(key, "/") {
		return key + "/"
	}
	return key
}

// AddressMergingAndConv Merge The File And Path
// Convert Windows-style path to Linux-style
func AddressMergingAndConv(fileWithPath, pathOnly string) (string, string) {
	fileWithPath = PathConv(fileWithPath)
	fileName := filepath.Base(fileWithPath)
	mergedAddress := path.Join(pathOnly, fileName)
	return PathConv(mergedAddress), fileWithPath
}

// PathConv Use For Path Covert
// \\ -> /
func PathConv(filePath string) string {
	return strings.Replace(filePath, "\\", "/", -1)
}

// upLoadPathConv Use For Upload Path Convert
// D:/a -> D/a
func upLoadPathConv(filePath string) string {
	return strings.Replace(PathConv(filePath), ":", "", -1)
}

// downloadPathConv Use For Download Path Convert
// D/a -> D:/a
func downloadPathConv(filePath string) string {
	return strings.Replace(PathConv(filePath), "/", ":/", 1)
}
