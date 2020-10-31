package main

import (
	"flag"
	"fmt"
	"gogit/pkg/base"
	"gogit/pkg/data"
	"os"

	//"flag"
	//	"io"
	"bufio"
	"io/ioutil"
)

func main() {
	//TODO flagSets, right now the flags are on the gogit, not the subcommand
	var expectedFlag string
	flag.StringVar(&expectedFlag, "expected", "blob", "Expected Type of Object")
	flag.Parse()
	expected := []byte(expectedFlag)

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

func check(e error) {
	if e != nil {
		panic(e)
	}
}
