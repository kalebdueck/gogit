package main

import (
	"flag"
	"fmt"
	"gogit/pkg/base"
	"gogit/pkg/data"
	"os"
	"strings"

	//"flag"
	//	"io"
	"bufio"
	"io/ioutil"
)

func main() {
	//TODO flagSets, right now the flags are on the gogit, not the subcommand
	var expectedFlag string
	flag.StringVar(&expectedFlag, "expected", "blob", "Expected Type of Object")
	var messageFlag string
	flag.StringVar(&messageFlag, "m", "", "Your Commit Message")
	var oidFlag string
	flag.StringVar(&oidFlag, "oid", "", "Your Commit OID")
	flag.Parse()
	expected := []byte(expectedFlag)

	fmt.Println(messageFlag)

	switch flag.Arg(0) {
	case "init":
		Init()
	case "hash-object":
		HashObject(flag.Arg(1), expected)
	case "cat-file":
		CatFile(flag.Arg(1), expected)
	case "write-tree":
		WriteTree(flag.Arg(1))
	case "read-tree":
		ReadTree(flag.Arg(1))
	case "commit":
		Commit(messageFlag)
	case "log":
		Log(oidFlag)
	}
}

func Init() {
	data.Init()
}

func HashObject(file string, type_ []byte) {
	dat, err := ioutil.ReadFile(file)
	check(err)
	fmt.Println(data.HashObject(dat, type_))
}

func CatFile(object string, expected []byte) {
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	resp, err := data.GetObject(object, expected)

	if err != nil {
		fmt.Println(err)
		return
	}

	f.Write(resp)
}

func WriteTree(directory string) {
	result := base.WriteTree(directory)
	fmt.Println(result)
}

func ReadTree(oid string) {
	base.ReadTree(oid, "./")
}

func Commit(message string) {
	result := base.Commit(message)

	fmt.Println(result)
}

func Log(oid string) {
	if oid == "" {
		oid = data.GetHead()
	}

	for oid != "" {
		commit := base.GetCommit(oid)

		fmt.Printf("commit: %s\n", oid)

		var newOid string = ""
		for _, line := range strings.Split(commit, "\n") {
			split := strings.Split(line, " ")
			if split[0] == "parent" {
				newOid = split[1]
			}
		}

		oid = newOid
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
