package pathwalker

import (
	"fmt"
	"os"
	"path/filepath"
)

func GenerateRoutesFromFileSystem() error {
	dirs, err := os.ReadDir("./")
	if err != nil {
		return err
	}

	for _, d := range dirs {
		fmt.Println(d)
	}
	return nil
}

func GenerateRoutesFromFileSystem2(root string) error {
	files, err := FilePathWalkDir(root)
	if err != nil {
		return err
	}

	for _, d := range files {
		fmt.Println(d)
	}
	return nil
}

func FilePathWalkDir(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		fmt.Println(path, info)
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}
