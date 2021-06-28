package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

//	changelogFile, err := os.OpenFile("CHANGELOG1.md", os.O_APPEND|os.O_WRONLY, 0744)
//	CheckErr(err)
//
//	newchangelogFile, err := ioutil.ReadFile("CHANGELOG2.md")
//	CheckErr(err)
//
//	changelogFile.Write(newchangelogFile[11:])
	header, err := ioutil.ReadFile("template/changelogHeader.txt")
	CheckErr(err)
	fmt.Println(header[189:])
	file, err := os.Create("OLDCHANGELOG.md")
	CheckErr(err)
	file.Write(header)
	file.Close()
}

func CheckErr(e error) {
	if e != nil {
		panic(e)
	}
}

