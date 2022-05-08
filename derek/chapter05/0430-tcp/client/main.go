package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

const addr = "localhost:8888"

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("Enter some text: ")
		data, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("encountered an error reading input: %s\n", err.Error())
			continue
		}
		// 서버 주소로 연결
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			fmt.Printf("encountered an error connecting: %s\n", err.Error())
		}

		// 전송 데이터 생성
		fmt.Fprintf(conn, data)

		// response 가져옴.
		status, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Printf("encountered an error reading response: %s\n", err.Error())
		}
		fmt.Printf("Received back: %s", status)
	
		conn.Close()
	}
}