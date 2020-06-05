package util

import (
	"strings"
)

//RemoveExtension removes the extension from a given file. e.g file_test.go becomes file_test
func RemoveExtension(s string) string {
	splits := strings.Split(s, ".")
	return splits[0]
}
