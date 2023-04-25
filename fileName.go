package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func getFileName() {
	root := "." // 目录名

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			// 如果是目录，则递归遍历
			if path != root {
				subfiles, err := listFiles(path)
				if err != nil {
					return err
				}
				files = append(files, subfiles...)
				return filepath.SkipDir
			}
		} else {
			// 如果是文件，则添加到结果列表
			files = append(files, path)
		}
		return nil
	})

	if err != nil {
		fmt.Println(err)
		return
	}
	/*
		for _, file := range files {
			fmt.Println(file)
		}
	*/
}

// 列出指定目录下的所有文件和子目录
func listFiles(dir string) ([]string, error) {
	files := []string{}
	list, err := ioutil.ReadDir(dir)
	if err != nil {
		return files, err
	}
	for _, item := range list {
		path := filepath.Join(dir, item.Name())
		if item.IsDir() {
			// 如果是子目录，则递归遍历
			subfiles, err := listFiles(path)
			if err != nil {
				return files, err
			}
			files = append(files, subfiles...)
		} else {
			// 如果是文件，则添加到结果列表
			files = append(files, path)
		}
	}
	return files, nil
}
