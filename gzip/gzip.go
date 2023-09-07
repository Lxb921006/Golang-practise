package main

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	sourceDir := "C:\\Users\\Administrator\\Desktop\\ccc" // 源目录路径
	gzName := filepath.Base(sourceDir) + ".tar.gz"
	targetFile := filepath.Join("C:\\Users\\Administrator\\Desktop", gzName) // 压缩后的目标文件名

	// 创建目标文件
	target, err := os.Create(targetFile)
	if err != nil {
		fmt.Println("无法创建目标文件:", err)
		return
	}
	defer target.Close()

	// 创建 gzip.Writer
	gzWriter := gzip.NewWriter(target)
	defer gzWriter.Close()

	// 创建 tar.Writer
	tarWriter := tar.NewWriter(gzWriter)
	defer tarWriter.Close()

	// 遍历源目录并压缩其中的文件和子目录
	err = filepath.Walk(sourceDir, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 创建 tar 文件头
		header, err := tar.FileInfoHeader(info, info.Name())
		if err != nil {
			return err
		}

		// 修改文件头中的名称，以相对路径存储
		relPath, _ := filepath.Rel(sourceDir, filePath)
		header.Name = relPath

		// 写入 tar 文件头
		if err = tarWriter.WriteHeader(header); err != nil {
			return err
		}

		// 如果是文件，复制文件内容到 tar.Writer
		if !info.IsDir() {
			file, err := os.Open(filePath)
			if err != nil {
				return err
			}
			defer file.Close()

			_, err = io.Copy(tarWriter, file)
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		fmt.Println("压缩目录时出错:", err)
		return
	}

	fmt.Println("目录已成功压缩到", targetFile)
}
