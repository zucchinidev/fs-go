package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

const grr = "G.R.R. Martin"

type book struct {
	Author string `json:"author"`
	Title  string `json:"title"`
	Year   int    `json:"year,omitempty"`
}

func main() {
	dst, err := os.OpenFile("books_json_line_protocol.jsonl", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer dst.Close()

	bookList := []book{
		{Author: grr, Title: "A Game of Thrones", Year: 1996},
		{Author: grr, Title: "A Clash of Kings", Year: 1998},
		{Author: grr, Title: "A Storm of Swords", Year: 2000},
		{Author: grr, Title: "A Feast for Crows", Year: 2005},
		{Author: grr, Title: "A Dance with Dragons", Year: 2011},
		{Author: grr, Title: "The Winds of Winter"},
		{Author: grr, Title: "A Dream of Spring"},
	}
	b := bytes.NewBuffer(make([]byte, 0, 16))
	for _, v := range bookList {
		j, errMarshalling := json.Marshal(v)
		if errMarshalling != nil {
			fmt.Println("Error:", err)
			return
		}
		_, _ = fmt.Fprintf(b, "%s", j)

		b.WriteRune('\n')

		if _, errWriting := b.WriteTo(dst); errWriting != nil { // copies bytes, drains buffer
			fmt.Println("Error:", errWriting)
			return
		}
	}
}
