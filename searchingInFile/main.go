package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type queryWriter struct {
	Query []byte
	io.Writer
}

func (q queryWriter) Write(b []byte) (n int, err error) {
	lines := bytes.Split(b, []byte{'\n'})
	l := len(q.Query)

	for _, line := range lines {
		index := bytes.Index(line, q.Query)
		if index == -1 {
			// there is not match in these specific part
			continue
		}

		output := [][]byte{
			line[:index],          // what's before the match
			[]byte("\x1b[31m"),    // start red color
			line[index : index+l], // the match
			[]byte("\x1b[39m"),    // start default color
			line[index+l:],        // whatever is left
		}
		for _, s := range output {
			v, errWriting := q.Writer.Write(s)
			n += v
			if errWriting != nil {
				return 0, errWriting
			}
		}
		_, _ = fmt.Fprintln(q.Writer)
	}
	return n, nil
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Please specify a path and a search string.")
		return
	}
	root, err := filepath.Abs(os.Args[1]) // get absolute path
	if err != nil {
		fmt.Println("Cannot get absolute path:", err)
		return
	}
	query := []byte(strings.Join(os.Args[2:], " "))
	fmt.Printf("Searching for %q in %s...\n", query, root)
	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		fmt.Println(path)
		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer f.Close()

		// TeeReader returns a Reader that writes to w what it reads from r
		// then we will read from the file and will write to the query writer that is connected to the standard output
		reader := io.TeeReader(f, queryWriter{
			Query:  query,
			Writer: os.Stdout,
		})
		_, err = ioutil.ReadAll(reader)
		return err
	})
	if err != nil {
		fmt.Println(err)
	}
}
