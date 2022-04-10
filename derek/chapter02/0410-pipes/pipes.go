package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 단어 : 빈도수 의 map 만들기.
func WordCount(f io.Reader) map[string]int {
	// map 객체 만들기.
	result := make(map[string]int)

	// f를 읽을 수 있는 scanner 생성.
	scanner := bufio.NewScanner(f)
	// ' '를 기준으로 단어로 분리. 사실상 split(' ')와 같음.
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		result[scanner.Text()]++
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}

	return result
}

func main() {
	fmt.Printf("string: number_of_occurrences\n\n")
	for key, value := range WordCount(os.Stdin) {
		fmt.Printf("%s: %d\n", key, value)
	}
}
