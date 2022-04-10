package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// WordCount 함수는 파일을 입력받아 각 단어를 키로 하고,
// 단어의 등장 수를 값으로 하는 맵을 반환한다.
func WordCount(f io.Reader) map[string]int {
	result := make(map[string]int)

	scanner := bufio.NewScanner(f)
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
