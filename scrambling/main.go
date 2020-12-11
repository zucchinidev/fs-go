package main

import (
	"fmt"
	"io"
	"math/rand"
	"strings"
	"unicode"
	"unicode/utf8"
)

func NewScrambleWriter(w io.Writer, r *rand.Rand, chance float64) *ScrambleWriter {
	return &ScrambleWriter{w: w, r: r, chance: chance}
}

type ScrambleWriter struct {
	w      io.Writer
	r      *rand.Rand
	chance float64
}

func (s *ScrambleWriter) shambleWrite(runes []rune, noLetterRune rune) (n int, err error) {
	//scramble after first letter
	removeFirstLetter := 1
	for i := removeFirstLetter; i < len(runes)-1; i++ {

		if s.r.Float64() > s.chance {
			continue
		}
		highLimit := len(runes) - 1
		j := s.r.Intn(highLimit) + removeFirstLetter
		runes[i], runes[j] = runes[j], runes[i]
	}

	if noLetterRune != 0 {
		runes = append(runes, noLetterRune)
	}

	var bytes = make([]byte, 10)
	for _, currentRune := range runes {
		writtenWidth := utf8.EncodeRune(bytes, currentRune)
		v, err := s.w.Write(bytes[:writtenWidth])
		if err != nil {
			return n, err
		}
		n += v
	}
	return
}

func (s *ScrambleWriter) Write(b []byte) (n int, err error) {
	var runes = make([]rune, 0, 10)
	runeElem, index, width := rune(0), 0, 0

	for ; index < len(b); index += width {
		// read the rune
		runeElem, width = utf8.DecodeRune(b[index:])

		if unicode.IsLetter(runeElem) {
			runes = append(runes, runeElem)
			continue
		}

		v, err := s.shambleWrite(runes, runeElem)
		if err != nil {
			return n, err
		}
		n += v
		runes = runes[:0] // reset slice
	}
	// when in the last turn there are no more bytes in the segment and it exits the loop
	if len(runes) != 0 {
		v, err := s.shambleWrite(runes, 0)
		if err != nil {
			return n, err
		}
		n += v
	}
	return
}

func main() {
	var s strings.Builder
	w := NewScrambleWriter(&s, rand.New(rand.NewSource(1)), 0.7)
	_, _ = fmt.Fprint(w, "Hello! this is a sample text.\nCan you read it? Yes")
	fmt.Println(s.String())
}
