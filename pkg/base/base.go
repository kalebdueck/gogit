package base

import (
	"fmt"
	"gogit/pkg/data"
	"io/ioutil"
	"strings"
	"encoding/hex"
)

func WriteTree(directory string) string {

	var tree string
	files, err := ioutil.ReadDir(directory)

	if err != nil {
		panic(err)
	}

	for _, file := range files {
		if IsIgnored(file.Name()) {
			continue
		}

		fullPath := directory + "/" + file.Name()

		var otype string
		var oid string
		if file.IsDir() {
			otype = "tree"
			oid = WriteTree(fullPath)
		} else {
			otype = "blob"
			if file.Mode().IsRegular() {
				fileData, err := ioutil.ReadFile(fullPath)
				if err != nil {
					panic(err)
				}
				oid = data.HashObject(fileData, []byte(otype))
			}
		}

		tree += otype + " " + oid + " " + file.Name() + "\n"
	}
	//TODO sort tree entries

	return data.HashObject([]byte(tree), []byte("tree"))
}

func IsIgnored(path string) bool {
	if strings.Contains(path, ".gogit") {
		return true
	}
	return false
}

// "Reads" a tree into the working directory
func ReadTree(oid string) {
	iterTreeEntries(oid)
}

func getTree(oid string, basePath string)  {

}

func iterTreeEntries(oid string) {
	tree, _ := data.GetObject(oid, []byte("tree"))
	var also []byte
	hex.Decode(tree, also)

	fmt.Println(also)
}