package main

import (
	"embed"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

//go:embed 1.txt 2.txt 3/3.txt
var hoge embed.FS

func main() {
	newDir, err := os.MkdirTemp("", "play")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(newDir)

	fmt.Println("new dir", newDir)
	e := fs.FS(hoge)
	err = fs.WalkDir(e, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		fmt.Println("walking", path, d.Name())
		var newPath string = filepath.Join(newDir, path)
		if d.IsDir() {
			if d.Name() == "." {
				return nil
			}
			return os.Mkdir(newPath, 0755)
		}
		srcFile, err := e.Open(path)
		if err != nil {
			return err
		}
		defer srcFile.Close()
		dstFile, err := os.Create(newPath)
		if err != nil {
			return err
		}
		_, err = io.Copy(dstFile, srcFile)
		if err != nil {
			return err
		}
		fmt.Println("copied", path, "to", newPath)
		return nil
	})
	if err != nil {
		panic(err)
	}
}
