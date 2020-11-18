package main

import (
	"gogit/cmd"
	//"flag"
	//	"io"
)

func main() {
	cmd.Execute()
}

/*
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
*/
