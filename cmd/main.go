package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

func main() {
	IteratePath("/Users/prashant/go/src/github.com/prashant182/filetree/test")
}

//File is the type of structure which holds the key information regarding the Files on a disk
type File struct {
	Name     string
	Path     string
	IsDir    bool
	Children []*File
}

//IteratePath iterates over the files in a given directory and creates a JSON/Yaml structure out of it.
func IteratePath(path string) {
	rootFile, err := os.Stat(path)
	if err != nil {
		logrus.Error("Path : ", path, " doesn't exist. ", err)
	}
	rootfile := ToFile(rootFile, path)
	stack := []*File{rootfile}
	for len(stack) > 0 {
		file := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		children, _ := ioutil.ReadDir(file.Path)
		for _, chld := range children {
			child := ToFile(chld, filepath.Join(file.Path, chld.Name())) //turn it into a File object
			file.Children = append(file.Children, child)                 //append it to the children of the current file popped
			stack = append(stack, child)
		}
	}

	output, err := yaml.Marshal(rootfile)
	if err != nil {
		logrus.Error("YAML marshal error ", err)
	}
	fmt.Println(string(output))
}

//ToFile converts the fileInfo object into a File object for filetree package
func ToFile(file os.FileInfo, path string) *File {
	outFile := File{
		Name:     file.Name(),
		IsDir:    file.IsDir(),
		Children: []*File{},
		Path:     path,
	}

	return &outFile
}