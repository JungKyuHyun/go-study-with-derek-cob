package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// singint 인터럽트할 리스너 설정.
func CatchSig(ch chan os.Signal, done chan bool) {
	// 신호 대기용 블락
	// 보내기 전용 및 받기 전용 채널은 채널 앞 뒤로 <- 연산자를 붙여서 만듭니다.
	// 보내기 전용 : chan <- 자료형
	// 받기 전용 : <-chan 자료형
	sig := <-ch
	// 받기 전용이기에 리스너.
	fmt.Println("\nsig received:", sig)

	// 신호에 대한 시그널 핸들러 지정.
	switch sig {
	case syscall.SIGINT:
		fmt.Println("handling a SIGINT now!")
	case syscall.SIGTERM:
		fmt.Println("handling a SIGTERM in an entirely different way!")
	default:
		fmt.Println("unexpected signal received")
	}

	// 종료 문구.
	done <- true
}

func main() {
	// 채널 초기화
	signals := make(chan os.Signal)
	done := make(chan bool)

	// signals 라이브러리에 연결.
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	// go 루틴(비동기로) 신호가 잡히면 done으로 쓴다.
	go CatchSig(signals, done)

	//누군가 done을 할때까지 프로그램을 블럭
	fmt.Println("Press ctrl-c to terminate...")
	<-done
	fmt.Println("Done!")
}
