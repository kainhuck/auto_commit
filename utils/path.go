package utils

import (
	"os"
)

// 检查路径是否存在
func CheckPathExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return false
}

// 检查文件是否存在
func CheckFileExist(file string) bool {
	f, err := os.Stat(file)
	if err == nil && !f.IsDir() {
		return true
	}
	return false
}

// 检查目录是否存在
func CheckDirectoryExist(dir string) bool {
	f, err := os.Stat(dir)
	if err == nil && f.IsDir() {
		return true
	}
	return false
}

// 用户家目录
func HomeDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return "~"
	}
	return home
}
