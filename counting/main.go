package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	args := os.Args

	if len(args) != 2 {
		fmt.Println("Please specify a path")
		return
	}

	root, errCheckingAbs := filepath.Abs(args[1])
	if errCheckingAbs != nil {
		fmt.Println("Cannot get absolute path: ", errCheckingAbs)
		os.Exit(70)
	}

	fmt.Println("Listing files in ", root)
	var c struct{
		files int
		dirs int
	}

	errWalking := filepath.Walk(root, func(p string, info os.FileInfo, err error) error {
		if info.IsDir() {
			c.dirs++
		} else {
			c.files++
		}
		fmt.Println("-", p)
		return nil
	})

	if errWalking != nil {
		fmt.Println("An error occurred ", errWalking)
		os.Exit(70)
	}

	fmt.Printf("Total: %d files in %d directories", c.files, c.dirs)
}
