package csvformat

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
)

// Movie 구조체는 파싱된 CSV를 저장하는데 사용된다.
type Movie struct {
	Title    string
	Director string
	Year     int
}

// ReadCSV 함수는 io.Reader에 전달된 CSV를 처리하는 예제를 보여준다.
func ReadCSV(b io.Reader) ([]Movie, error) {
	r := csv.NewReader(b)
	r.Comma = ';'
	r.Comment = '-'
	var moives []Movie

	_, err := r.Read()
	// EOF (End of File) 데이터 소스로부터 더 이상 읽을 수 있는 데이터가 없음을 의미
	if err != nil && err != io.EOF {
		return moives, err
	}

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		// ParseInt(s string, base int, bitSize int) (i int64, err error)
		year, err := strconv.ParseInt(record[2], 10, 64)
		if err != nil {
			return nil, err
		}
		m := Movie{record[0], record[1], int(year)}
		moives = append(moives, m)
	}
	return moives, nil
}

// AddMoviesFromText 함수는 CSV 파서를 사용한다.
func AddMoviesFromText() error {
	// 이 예제는 csv 패키지를 사용해 문자열을 가져온 후 버퍼로 변환해 읽는 예를 보여준다.
	in := `
- first our headers;movie title;directior;year released
- then some data
Guardians of the Galaxy Vol. 2;James Gunn;2017 Star Wars: Episode VIII, Rian Johnson;2017 
`
	b := bytes.NewBufferString(in)
	m, err := ReadCSV(b)
	if err != nil {
		return err
	}
	fmt.Printf("%#vn", m)
	return nil
}
