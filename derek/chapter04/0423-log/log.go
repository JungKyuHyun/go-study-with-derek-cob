package log

import (
	"bytes"
	"fmt"
	"log"
)

func Log() {
	// bytes.Buffer 는 로그 기록용
	buf := bytes.Buffer{}

	//  log.New(저장 장소 , 접두사, flag)
	logger := log.New(&buf, "logger: ", log.Lshortfile|log.Ldate)

	logger.Println("test")

	logger.SetPrefix("new logger: ")

	logger.Printf("you can also add args(%v) and use Fatalln to log and crash", true)

	fmt.Println(buf.String())
}