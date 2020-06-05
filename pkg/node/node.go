package node

import (
	"github.com/prashant182/filetree/pkg/util"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

//Node is the structure which holds the file/dir information for a given path.
type Node struct {
	Name     string
	Path     string
	IsDir    bool
	Size     int64
	ModTime  time.Time
	Children []*Node
}

//ToFile converts the fileInfo object into a File object for filetree package
func ToFile(file os.FileInfo, path string) *Node {
	outFile := Node{
		Name:     file.Name(),
		Children: []*Node{},
		Path:     path,
		IsDir: file.IsDir(),
		ModTime:  file.ModTime(),
	}
	return &outFile
}

//RemoveExtension removes the extension from the filename. i.e file_test.go becomes file_test
func (nd *Node) RemoveExtension() *Node {
	nd.Name = util.RemoveExtension(nd.Name)
	return nd
}

func (nd *Node) ConvertCamelCase() *Node {
	nd.Name = util.Camelize(nd.Name)
	return nd
}

func (nd *Node) ToMap(contains string, camelCase bool, removeExtn bool) interface{} {
	if removeExtn && camelCase == true {
		nd = nd.RemoveExtension().ConvertCamelCase()
	}else if removeExtn == true && camelCase ==false{
		nd = nd.RemoveExtension()
	}else if camelCase==true && removeExtn == false{
		nd = nd.ConvertCamelCase()
	}

	if len(nd.Children) < 1 {
		return "true"
	}
	mp := make(map[string]interface{})
	for _, child := range nd.Children {
		if !child.IsDir {
			if strings.Contains(child.Name, contains) {
				mp[child.Name] = child.ToMap(contains, camelCase, removeExtn)
			}
		} else {
			mp[child.Name] = child.ToMap(contains, camelCase, removeExtn)
		}

	}
	return mp
}

//DFS walks the given file and gennerate the tree of a dir structure.
func DFS(path string) (node *Node) {
	rootFile, err := os.Stat(path)
	if err != nil {
		log.Println("Path : ", path, " doesn't exist. ")
		os.Exit(0)
	}
	rootfile := ToFile(rootFile, path)
	stack := []*Node{rootfile}
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
	return rootfile
}
