package csvformat

import (
	"bytes"
	"encoding/csv"
	"io"
	"os"
)

// Book 구조체 생성
type Book struct {
	Author string
	Title  string
}

// Books : Book 구조체의 배열로 선언.
type Books []Book

// 앞에 있는 관호는 리시버로 interface 의 메소드로 이해.
func (books *Books) ToCSV(w io.Writer) error {
	// csv 타입의 writer로 n 선언.
	n := csv.NewWriter(w)
	// N에 문자열 배열 wirte
	err := n.Write([]string{"Author", "Title"})
	if err != nil {
		return err
	}
	// books 배열 끝까지 book 객체형태로 n에 저장.
	for _, book := range *books {
		err := n.Write([]string{book.Author, book.Title})
		if err != nil {
			return err
		}
	}
	// Flush를 사용하여 모든 버퍼링된 작업이 라이터에 적용되었는지 확인
	n.Flush()
	// n에 에러가 있는지 리턴.
	return n.Error()
}

func WriteCSVOutput() error {
	//책 정보(Books) 선언
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
	// Books객체기 때문에 ToCSV 메소드 사용 가능.
	return b.ToCSV(os.Stdout)
}

func WriteCSVBuffer() (*bytes.Buffer, error) {
	// Books 객체로 초기화
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
	// struct{}는 그냥 변수로 받고, &struct{}는 포인터로 받지만, &struct{}로 선언된 객체를 참조할 때는 자동으로 (*struct).var 형식으로 참조하기 때문에 동작에 다를 게 없다.
	// 차이점은 , &struct{}는 struct의 포인터를 받아오기 때문에 참조 횟수가 늘어나면 struct{}에 비해 약간 더 느리다.
	// 그럼 왜쓰냐? interface에서 선언만 해 둔 메서드를 이후에 오버라이딩할 경우, 해당 함수는 꼭 *struct 타입이어야 한다.

	w := &bytes.Buffer{}
	err := b.ToCSV(w)
	return w, err
}
