package file

import (
	"fmt"
	"os"
	"path/filepath"
)

var Name string
var root string

// 所有文件的绝对路径名
var Files []string

func GetFileName(root1 string) {
	/*
		// 获取命令行参数作为目录路径
		if len(os.Args) < 2 {
			fmt.Println("Usage: go run main.go <dir>")
			return
		}
		root := os.Args[1]
	*/

	root = root1
	Name = filepath.Base(root)

	// 遍历目录下的所有文件和子目录
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			// 如果是文件，则添加到结果列表
			absPath, err := filepath.Abs(path)
			if err != nil {
				return err
			}
			if filepath.Ext(absPath) == ".go" {
				absPath = filepath.ToSlash(absPath)
				Files = append(Files, absPath)
			}

		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
		return
	}
}

// 打印所有文件的相对路径
func printAllFileRelPath() {
	var relPath string
	fmt.Println("所有文件的相对路径")
	for _, file := range Files {
		relPath, _ = filepath.Rel(root, file)
		relPath = filepath.ToSlash(relPath)
		fmt.Println(relPath)
	}
	fmt.Println("--------")
}
