package csvformat

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
)

// csv 파일을 파싱할 Movie "구조체"를 생성
type Movie struct {
	Title    string
	Director string
	Year     int
}

func ReadCSV(b io.Reader) ([]Movie, error) {
	// csv 리더 객체 생성.
	r := csv.NewReader(b)

	// 리더 객체의 설정을 지정해줌.
	r.Comma = ';'
	r.Comment = '-'
	// 구조체 배열 선언.
	var movies []Movie

	// r.Read 로 해더를 가져옴. 아까 설정한 ; - 설정으로 read 함.
	_, err := r.Read()
	if err != nil && err != io.EOF {
		return nil, err
	}

	// for 무한루프문 = while 처럼 씀.
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		// ParseInt (문자열, base,비트크기) = base와 비트크기로 문자열을 해석하고 정수를 리턴.
		year, err := strconv.ParseInt(record[2], 10, 64)
		if err != nil {
			return nil, err
		}

		m := Movie{record[0], record[1], int(year)}
		// movies는 구조체 배열의 포인터라 갱신 해주게 됨.
		movies = append(movies, m)
	}
	return movies, nil
}

// CSV 에서 문자열을 가져온 후 버퍼로 변환해 읽음.
func AddMoviesFromText() error {
	in := `
- first our headers
movie title;director;year released
- then some data
Guardians of the Galaxy Vol. 2;James Gunn;2017
Star Wars: Episode VIII;Rian Johnson;2017
`
	// 문자열에서 버퍼로 변환
	b := bytes.NewBufferString(in)

	m, err := ReadCSV(b)
	if err != nil {
		return err
	}
	// %#v : a Go-syntax representation of the value
	// %v : 구조체의 인스턴스를 출력함.
	fmt.Printf("%#v\n", m)
	return nil
}
