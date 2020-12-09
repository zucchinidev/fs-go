package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strings"
	"unicode"
	"unicode/utf8"
)

type AngryReader struct {
	r io.Reader
}

func NewAngryReader(r io.Reader) *AngryReader {
	return &AngryReader{r: r}
}

func (a *AngryReader) Read(b []byte) (int, error) {
	n, errReading := a.r.Read(b)
	if errReading != nil {
		return 0, errReading
	}

	index := 0
	runeElem := rune(0)
	widthBytes := 0
	for ; index < n; index += widthBytes {
		// read a rune, unpacks the first UTF-8 encoding in p and returns the rune
		runeElem, widthBytes = utf8.DecodeRune(b[index:])
		// check is the rune is a letter
		if !unicode.IsLetter(runeElem) {
			continue
		}

		// uppercase version of the rune found
		runeElemUpper := unicode.ToUpper(runeElem)
		// EncodeRune writes into p (which must be large enough) the UTF-8 encoding of the rune.
		// It returns the number of bytes written.
		// If the number is not the same is because the rune is other
		bytesEncodeRune := utf8.EncodeRune(b[index:], runeElemUpper)
		if bytesEncodeRune != widthBytes {
			return n, fmt.Errorf(
				"The runes are not the same because they do not have the same size: %c->%c, size mismatch %d->%d",
				runeElem,
				runeElemUpper,
				widthBytes,
				bytesEncodeRune,
			)
		}
	}
	return n, nil
}

func main() {
	reader := strings.NewReader("The power of the streams using the io.Reader interface")
	r := NewAngryReader(reader)
	b, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(b))
}
