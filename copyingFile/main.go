package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	args := os.Args

	if len(args) != 3 {
		fmt.Println("Please specify paths")
		return
	}

	if _, errCopying := copyFile(args[1], args[2]); errCopying != nil {
		fmt.Println("Error copying file", errCopying)
		os.Exit(70)
	}
}

func copyFile(from, to string) (int64, error) {
	src, errOpening := os.Open(from)
	if errOpening != nil {
		return 0, errOpening
	}
	defer src.Close()

	dst, errOpeningNewFile := os.OpenFile(to, os.O_WRONLY|os.O_CREATE, 0644)
	if errOpeningNewFile != nil {
		return 0, errOpeningNewFile
	}
	defer dst.Close()

	return io.Copy(dst, src)
}
