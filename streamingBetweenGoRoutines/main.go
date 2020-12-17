package main

import (
	"fmt"
	"io"
)

func main() {
	reader, writer := io.Pipe()

	go func(w io.WriteCloser) {
		for _, s := range []string{"communicating streams", " using go routines ", " with the built-in ", " pipe function "} {
			fmt.Printf("-> writing %q\n", s)
			_, _ = fmt.Fprint(w, s)
		}
		_ = w.Close() // we need to close the writer so as to send EOF error when a reader tries to read more data
	}(writer)

	var err error
	for n, b := 0, make([]byte, 100); err == nil; {
		fmt.Println("<- waiting...")
		n, err = reader.Read(b)
		if err == nil {
			fmt.Printf("<- received %q\n", string(b[:n]))
		}
	}

	if err != io.EOF {
		fmt.Println("error:", err)
	}
}
