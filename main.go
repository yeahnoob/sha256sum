package main

import (
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	filename := os.Args[1]
	// fmt.Println("Hashing the file: ", filename)
	data, err := ioutil.ReadFile(filename)
	check(err)
	fmt.Printf("%x\n", sha256.Sum256(data))
}
