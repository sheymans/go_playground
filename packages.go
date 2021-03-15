package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func randomStuff() {
	// strings package
	fmt.Println(strings.Contains("stijn", "ij"))
	fmt.Println(strings.Index("stijn", "i"))
	fmt.Println(strings.Join([]string{"a", "b"}, ":"))
	// -1 means replace _all_ occurrences
	fmt.Println(strings.Replace("stijn", "ij", "a",-1))
	fmt.Println(strings.Split("sti,jn", ","))
	fmt.Println(strings.Split("stijn", ""))
	fmt.Println(strings.ToLower("STIJN"))

	// Reading files, long-winded way
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("could not open file test.txt")
		return
	}
	// we defer which closes this as end of function
	defer file.Close()
	//read file by first checking size, creating a byte array of that size and reading the contents in it
	stat, err := file.Stat()
	fmt.Println("file size", stat.Size())
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	str := string(bs)
	fmt.Println(str)

	// Reading files, better way
	bs2, _ := ioutil.ReadFile("test.txt")
	str2 := string(bs2)
	fmt.Println(str2)

	// Create a file
	file3, err3 := os.Create("test2.txt")
	if err3 != nil {
		fmt.Println("creating file went wrong")
		return
	}
	defer file3.Close()
	_, _ = file3.WriteString("we wrote to this")

}
