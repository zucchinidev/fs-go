package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args

	if len(args) != 2 {
		fmt.Println("Please specify a path")
		return
	}

	b, errReadingFile := ioutil.ReadFile(args[1])
	if errReadingFile != nil {
		fmt.Println("Error:", errReadingFile)
		os.Exit(70)
	}

	fmt.Println(string(b))
}
