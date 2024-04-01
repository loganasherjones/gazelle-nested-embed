package main

import (
	"fmt"
	"github.com/loganasherjones/gazelle-nested-embed/broken"
	"github.com/loganasherjones/gazelle-nested-embed/working"
)

func main() {
	fmt.Println("Hello World")
	err := working.PrintFile("file.json")
	if err != nil {
		fmt.Println(err)
		panic("error printing working file")
	}
	err = broken.PrintFile("file.json")
	if err != nil {
		fmt.Println(err)
		panic("error printing broken file")
	}
}
