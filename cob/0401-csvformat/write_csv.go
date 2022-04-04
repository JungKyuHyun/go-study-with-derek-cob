package csvformat

import (
	"bytes"
	"encoding/csv"
	"io"
	"os"
)

type Book struct {
	Author string
	Title  string
}

type Books []Book

func (books *Books) ToCSV(w io.Writer) error {
	n := csv.NewWriter(w)
	err := n.Write([]string{"Author", "Title"})
	if err != nil {
		return err
	}
	for _, book := range *books {
		err := n.Write([]string{book.Author, book.Title})
		if err != nil {
			return err
		}
	}
	n.Flush()
	return n.Error()
}

// WriteCSVOutput 함수는 여러 책의 정보를 초기화하고 이를 표준 출력에 기록한다
func WriteCSVOutput() error {
	b := Books{
		Book{
			Author: "F Scott Fitzgerald",
			Title:  "The Great Gatsby",
		},
		Book{
			Author: "J D Salinger",
			Title:  "The Catcher in the Rye",
		},
	}
	return b.ToCSV(os.Stdout)
}

// WriteCSVBuffer 함수는 여러 책의 정보가 담긴 버퍼 csv를 반환한다
func WriteCSVBuffer() (*bytes.Buffer, error) {
	b := Books{
		Book{
			Author: "F Scott Fitzgerald",
			Title:  "The Great Gatsby",
		},
		Book{
			Author: "J D Salinger",
			Title:  "The Catcher in the Rye",
		},
	}
	w := &bytes.Buffer{}
	err := b.ToCSV(w)
	return w, err
}
