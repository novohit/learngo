package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Tool struct {
	val int
}

func (tool Tool) String() string {
	return "implement String()"
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	printFileContents(file)
}
func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
func main() {
	var tool Tool
	tool.val = 1

	t := &Tool{}
	t2 := Tool{}
	t3 := new(Tool)
	fmt.Println(tool.String())
	fmt.Println(t.String())
	fmt.Println(t2.String())
	fmt.Println(t3.String())
	printFile("systeminterface/abc.txt")
	fmt.Println("===========")
	s := `abc"d"
	kkkk
	
	hello`
	printFileContents(strings.NewReader(s))
}
