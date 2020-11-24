package data

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

const GoDir string = ".gogit"

func Init() {
	os.Mkdir(GoDir, 0755)
	os.Mkdir(GoDir+"/objects", 0755)
}

func HashObject(dat []byte, type_ []byte) string {
	splitter := byte(0)

	//TODO get back to this once i understand different vars better
	new := append(type_[:], splitter)
	hi := append(new[:], dat...)

	h := sha1.New()
	h.Write(hi)
	oid := h.Sum(nil)

	hash := hex.EncodeToString(oid)
	err := ioutil.WriteFile("./"+GoDir+"/objects/"+hash, hi, 0644)
	check(err)

	return hash
}

func GetObject(oid string, expected string) ([]byte, error) {
	fileName := "./" + GoDir + "/objects/" + oid
	file, err := ioutil.ReadFile(fileName)
	check(err)

	//Split file on the zero byte using for loop
	//TODO better splitting so i can just trim from file
	expectedLength := len(expected)

	type_ := file[:expectedLength]
	remainder := file[expectedLength+1:]

	if expected != "none" && string(type_) != expected {
		return nil, fmt.Errorf("Expected %s, got %s", expected, type_)
	}

	return remainder, nil
}

func UpdateRef(oid string, ref string) {

	filelocation := "./" + GoDir + "/" + ref
	newpath := filepath.Dir(filelocation)
	fmt.Println(newpath)
	dirErr := os.MkdirAll(newpath, os.ModePerm)
	if dirErr != nil {
		panic(dirErr)
	}

	writeErr := ioutil.WriteFile(filelocation, []byte(oid), 0644)

	if writeErr != nil {
		panic(writeErr)
	}
}

func GetRef(ref string) string {
	oid, err := ioutil.ReadFile("./" + GoDir + "/" + ref)

	if err != nil {
		return ""
	}

	return string(oid)
}

func IterRefs() map[string]string {

	RefMap := make(map[string]string)
	RefMap["HEAD"] = GetOid("HEAD")

	RefDir := fmt.Sprintf("./%s/refs", GoDir)

	err := filepath.Walk(RefDir,
		func(path string, info os.FileInfo, err error) error {

			if err != nil {
				return err
			}

			if info.IsDir() {
				return nil
			}
			RefMap[info.Name()] = GetOid(info.Name())
			return nil
		})
	if err != nil {
		panic(err)
	}

	return RefMap
}

func GetOid(name string) string {

	if name == "" {
		return "HEAD"
	}

	refLocations := []string{
		fmt.Sprintf("%s", name),
		fmt.Sprintf("refs/%s", name),
		fmt.Sprintf("refs/tags/%s", name),
		fmt.Sprintf("refs/heads/%s", name),
	}

	for _, location := range refLocations {
		ref := GetRef(location)

		if ref != "" {
			return ref
		}
	}

	return name

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
