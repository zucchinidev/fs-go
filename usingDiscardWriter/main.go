package main

import (
	"io"
	"io/ioutil"
	"strings"
)

func main() {
	r := strings.NewReader("let's read this message\n")
	// discard the context writing to /dev/null, a null device.
	// This means that writing to this variable ignores the data.
	io.Copy(ioutil.Discard, r)
}
