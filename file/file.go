package file

import (
	"path/filepath"
)

// GetFileName 获取文件路径下的文件名
func GetFileName(filePath string) string {
	_, realName := filepath.Split(filePath)
	return realName
}
