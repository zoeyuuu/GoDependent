package main

import (
	"fmt"
	"os"
	"path/filepath"
)

var root string

func getFileName() {
	/*
		// 获取命令行参数作为目录路径
		if len(os.Args) < 2 {
			fmt.Println("Usage: go run main.go <dir>")
			return
		}
		root := os.Args[1]
	*/

	root = "D:\\projects\\go_projects\\src\\GoDependent\\file"

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
				files = append(files, absPath)
			}

		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	var relPath string
	// 打印所有文件的相对路径
	for _, file := range files {
		relPath, _ = filepath.Rel(root, file)
		fmt.Println(relPath)
	}

}
