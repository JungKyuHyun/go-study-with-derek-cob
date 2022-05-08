package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

const addr = "localhost:8888"

func echoBackCapitalized(conn net.Conn) {
	// 소켓 통신용 리더 세팅.
	reader := bufio.NewReader(conn)

	// grab the first line of data encountered
	data, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("error reading data: %s\n", err.Error())
		return
	}
	// 받은 데이터 출력.
	fmt.Printf("Received: %s", data)
	// 받은 데이터 저장.
	conn.Write([]byte(strings.ToUpper(data)))
	// 통신 끝나면 종료
	conn.Close()
}

func main() {
	// tcp로 리시버 설정
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	defer ln.Close()
	fmt.Printf("listening on: %s\n", addr)
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Printf("encountered an error accepting connection: %s\n", err.Error())
			// 에러시 try again
			continue
		}
		
		go echoBackCapitalized(conn)
	}
}