package node

//Node is the structure which holds the file/dir information for a given path.
type Node struct {
	Name     string
	Path     string
	IsDir    bool
	Children []*Node
}


func (node *Node) ToMap() interface{}{
	if len(node.Children) < 1 {
		return "true"
	}
	yamlMap := make(map[string]interface{})
	for _, child := range node.Children {
		yamlMap[child.Name] = child.ToMap()
	}
	return yamlMap
}