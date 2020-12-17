package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func CopyNOffset(dst io.Writer, src io.ReadSeeker, offset, length int64) (int64, error) {
	relativeToTheStartOfTheSource := io.SeekStart
	if _, err := src.Seek(offset, relativeToTheStartOfTheSource); err != nil {
		return 0, err
	}
	return io.CopyN(dst, src, length)
}

func main() {
	r := strings.NewReader("This is an example of CopyN with offset")
	i := int64(0)
	l := int64(r.Len())
	step := int64(5)
	for ; i < l; i += step {
		_, _ = CopyNOffset(os.Stdout, r, i, step)
		fmt.Println()
	}
}
