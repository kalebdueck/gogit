package base

import (
	"fmt"
	"gogit/pkg/data"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func WriteTree(directory string) string {
	var tree string
	files, err := ioutil.ReadDir(directory)

	if err != nil {
		panic(err)
	}

	//Todo not happy with all the reading and writing side effects in this function
	// consider splitting out read and write into two functions for easier testing
	for _, file := range files {
		fullPath := directory + "/" + file.Name()

		if IsIgnored(fullPath) {
			continue
		}

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

	//trim the last \n off the back
	tree = strings.TrimSuffix(tree, "\n")

	return data.HashObject([]byte(tree), []byte("tree"))
}

func IsIgnored(path string) bool {
	//TODO look at and parse out files from .gitignore

	if strings.Contains(path, ".gogit") {
		return true
	}

	if strings.Contains(path, ".git") {
		return true
	}
	return false
}

// "Reads" a tree into the working directory
//TODO file and folder permissions?
func ReadTree(tree_oid string, base_path string) {
	emptyCurrentDirectory(".")
	saveTreeToDir(tree_oid, base_path)
}

func saveTreeToDir(tree_oid string, base_path string) {
	fmt.Println(tree_oid)
	for _, info := range iterTreeEntries(tree_oid) {
		//TODO temporarily ignoring the directories
		fmt.Println(info)
		path := base_path + info[2]
		switch info[0] {
		case "blob":
			file, _ := data.GetObject(info[1], "blob")
			ioutil.WriteFile(path, file, 0644)
			fmt.Println(path)

		case "tree":
			os.MkdirAll(path, 0755)
			saveTreeToDir(info[1], path+"/")
		default:
			//TODO err out
			fmt.Println("Invalid oid object type")
		}
	}
}

func emptyCurrentDirectory(directory string) {
	//Step one remove all files not ignored
	walkErr := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		fullPath := directory + "/" + path
		if IsIgnored(fullPath) {
			return nil
		}

		fmt.Println("removing")
		fmt.Println(fullPath)
		os.Remove(fullPath)
		return nil
	})

	if walkErr != nil {
		panic(walkErr)
	}

	//Step 2 remove all empty directories (We know these don't have files in them)
	filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		fullPath := directory + "/" + path
		if IsIgnored(fullPath) || info.Name() == "." {
			return nil
		}

		if !info.IsDir() {
			return nil
		}

		fmt.Println("Looking at directory")
		fmt.Println(fullPath)

		isEmpty, _ := isDirEmpty(fullPath)
		fmt.Println(isEmpty)
		if !info.IsDir() || !isEmpty {
			return nil
		}

		fmt.Println("removing directory")
		fmt.Println(fullPath)
		os.Remove(fullPath)
		return nil
	})

}

func isDirEmpty(name string) (bool, error) {
	f, err := os.Open(name)
	if err != nil {
		return false, err
	}
	defer f.Close()

	// read in ONLY one file
	_, err = f.Readdir(1)

	// and if the file is EOF... well, the dir is empty.
	if err == io.EOF {
		return true, nil
	}
	return false, err
}

func iterTreeEntries(oid string) [][]string {
	tree, _ := data.GetObject(oid, "tree")
	var result [][]string
	for _, line := range strings.Split(string(tree), "\n") {
		fmt.Println(line)
		split := strings.Split(line, " ")
		result = append(result, split)
	}

	return result
}

func Commit(message string) string {
	if message == "" {
		//TODO actually STDERR?
		return "message flag is required, and must be non empty"
	}

	treeName := WriteTree(".")
	var commit string
	commit = fmt.Sprintf("tree %s\n", treeName)
	commit += fmt.Sprintf("parent %s\n", data.GetRef("HEAD", true).Value)
	commit += "\n"
	commit += fmt.Sprintf("%s\n", message)

	oid := data.HashObject([]byte(commit), []byte("commit"))

	data.UpdateRef("HEAD", data.RefValue{Value: oid}, true)

	return oid
}

type CommitData struct {
	Tree    string
	Parent  string
	Message string
}

func GetCommit(oid string) CommitData {
	commit, err := data.GetObject(oid, "commit")

	commitLines := strings.Split(string(commit), "\n")

	treeLine := strings.Split(commitLines[0], " ")
	parentLine := strings.Split(commitLines[1], " ")
	message := commitLines[3]

	if err != nil {
		panic(err)
	}
	return CommitData{
		Tree:    treeLine[1],
		Parent:  parentLine[1],
		Message: message,
	}
}

func CreateTag(tagName string, oid string) string {
	data.UpdateRef("refs/tags/"+tagName, data.RefValue{Value: oid}, true)
	return tagName
}

//GetOid Converts a string to an oid if exists
//Otherwise it asssumes the name was an oid and returns it
func GetOid(name string) data.RefValue {
	return data.GetOid(name)
}

//Iterate through each oid in oids finding every parent,
//Filtering out the duplicates
func IterCommitsAndParents(oids []string) []string {
	var visited []string

	for {
		if len(oids) == 0 {
			break
		}

		//Pop the last var off
		last := oids[len(oids)-1]
		oids = oids[:len(oids)-1]

		if last == "" {
			break
		}

		//if last in visited, loop again
		var isVisited bool
		for _, visitedOid := range visited {
			if last == visitedOid {
				isVisited = true
				break
			}
		}
		if isVisited {
			break
		}

		visited = append(visited, last)

		fmt.Println(last)
		fmt.Println("que")
		commit := GetCommit(last)
		fmt.Println(commit)

		oids = append(oids, commit.Parent)
	}

	return visited
}

func CreateBranch(name string, oid string) {
	data.UpdateRef("refs/heads/"+name, data.RefValue{Value: oid}, true)
}

func IsBranch(branch string) bool {
	return data.GetRef(fmt.Sprintf("refs/heads/%s", branch), true).Value != ""
}

func Init() {
	data.Init()
	data.UpdateRef(
		"HEAD",
		data.RefValue{
			Symbolic: true,
			Value:    "refs/head/master",
		},
		true,
	)
}
