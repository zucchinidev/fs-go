package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	r := strings.NewReader("let's read this message\n")
	b := bytes.NewBuffer(nil)
	w := io.MultiWriter(b, os.Stdout)
	io.Copy(w, r)           // prints to the standard output
	fmt.Println(b.String()) // buffer also contains string now
}
