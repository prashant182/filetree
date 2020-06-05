package main

import (
	"fmt"
	"github.com/prashant182/filetree/pkg/tree"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
	"github.com/sirupsen/logrus"
)

func main() {
	IteratePath("/Users/prashant/go/src/github.com/prashant182/filetree/sc")
}

//IteratePath iterates over the files in a given directory and creates a JSON/Yaml structure out of it.
func IteratePath(path string) {
	rootFile, err := os.Stat(path)
	if err != nil {
		logrus.Error("Path : ", path, " doesn't exist. ", err)
	}
	rootfile := ToFile(rootFile, path)
	stack := []*tree.Node{rootfile}
	for len(stack) > 0 {
		file := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		children, _ := ioutil.ReadDir(file.Path)
		for _, chld := range children {
			child := ToFile(chld, filepath.Join(file.Path, chld.Name()))
			file.Children = append(file.Children, child)
			stack = append(stack, child)
		}
	}

	m := rootfile.ToMap()
	output, err := yaml.Marshal(m)
	if err != nil {
		logrus.Error("YAML marshal error ", err)
	}
	fmt.Println(string(output))
}

//ToFile converts the fileInfo object into a File object for filetree package
func ToFile(file os.FileInfo, path string) *tree.Node {
	outFile := tree.Node{
		Name:     file.Name(),
		Children: []*tree.Node{},
		Path:     path,
	}
	return &outFile
}