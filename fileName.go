package main

import (
	"fmt"
	"os"
	"path/filepath"
)

var root string

// 所有文件的绝对路径名
var files []string

func getFileName() {
	/*
		// 获取命令行参数作为目录路径
		if len(os.Args) < 2 {
			fmt.Println("Usage: go run main.go <dir>")
			return
		}
		root := os.Args[1]
	*/

	root = "D:/projects/go_projects/src/github.com/hugo"

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
				files = append(files, absPath)
			}

		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	//printAllFileRelPath()
}

// 打印所有文件的相对路径
func printAllFileRelPath() {
	var relPath string
	fmt.Println("所有文件的相对路径")
	for _, file := range files {
		relPath, _ = filepath.Rel(root, file)
		relPath = filepath.ToSlash(relPath)
		fmt.Println(relPath)
	}
	fmt.Println("--------")
}
