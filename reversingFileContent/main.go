package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Please specify a source and a destination file")
		return
	}
	src, errOpeningSrc := os.Open(os.Args[1])
	if errOpeningSrc != nil {
		fmt.Println("Error:", errOpeningSrc)
		return
	}
	defer src.Close()

	// OpenFile allows to open a file with any permissions
	dst, errOpeningDst := os.OpenFile(os.Args[2], os.O_WRONLY|os.O_CREATE, 0644)
	if errOpeningDst != nil {
		fmt.Println("Error:", errOpeningDst)
		return
	}
	defer dst.Close()

	// we are going to the end of the file
	cur, errSeeking := src.Seek(0, io.SeekEnd)
	if errSeeking != nil {
		fmt.Println("Error:", errSeeking)
		return
	}

	// defining a byte buffer
	b := make([]byte, 16)
	var err error
	for step, r, w := int64(16), 0, 0; cur != 0; {
		if cur < step { // ensure cursor is 0 at max
			b, step = b[:cur], cur
		}
		cur = cur - step
		_, err = src.Seek(cur, io.SeekStart) // go backwards
		if err != nil {
			break
		}
		if r, err = src.Read(b); err != nil || r != len(b) {
			if err == nil { // all buffer should be read
				err = fmt.Errorf("read: expected %d bytes, got %d", len(b), r)
			}
			break
		}
		for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
			switch { // Swap (\r\n) so they get back in place
			case b[i] == '\r' && b[i+1] == '\n':
				b[i], b[i+1] = b[i+1], b[i]
			case j != len(b)-1 && b[j-1] == '\r' && b[j] == '\n':
				b[j], b[j-1] = b[j-1], b[j]
			}
			b[i], b[j] = b[j], b[i] // swap bytes
		}
		if w, err = dst.Write(b); err != nil || w != len(b) {
			if err != nil {
				err = fmt.Errorf("write: expected %d bytes, got %d", len(b), w)
			}
		}
	}
	if err != nil && err != io.EOF { // we expect an EOF
		fmt.Println("\n\nError:", err)
	}
}
