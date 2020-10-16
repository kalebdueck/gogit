package data

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
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

func GetObject(oid string, expected []byte) ([]byte, error) {
	fileName := "./" + GoDir + "/objects/" + oid
	file, err := ioutil.ReadFile(fileName)
	check(err)

	//Split file on the zero byte using for loop
	//TODO better splitting so i can just trim from file
	type_ := file[:4]
	remainder := file[5:]

	if string(expected) != "none" && string(type_) != string(expected) {
		return nil, fmt.Errorf("Expected %s, got %s", expected, type_)
	}

	return remainder, nil
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
