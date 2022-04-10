package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// CatchSig 함수는 SIGINT 인터럽트에 대한 리스너를 설정한다.
func CatchSig(ch chan os.Signal, done chan bool) {
	// 신호를 대기하기 위해 블록한다.
	sig := <-ch
	// 신호를 받으면 출력
	fmt.Println("\nsig received:", sig)

	// 모든 신호에 대한 헨들러를 설정한다.
	switch sig {
	case syscall.SIGINT:
		fmt.Println("handling a SIGINT now!")
	case syscall.SIGTERM:
		fmt.Println("handling a SIGTERM in an entirely different way!")
	default:
		fmt.Println("unexpected signal received")
	}
	// 종료
	done <- true
}

func main() {
	// 채널 초기화
	signals := make(chan os.Signal)
	done := make(chan bool)

	// signals 라이브러리에 연결
	// signal.Notify 함수는 신호 알림과 처리를 원하는 신호의 타입을 전달하는데 채널을 필요로 한다.
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	// 이 고루틴에 의해 신호가 잡혔으면 done으로 쓴다.
	go CatchSig(signals, done)

	fmt.Println("Press ctrl-c to terminate...")
	// 누군가 done으로 쓸 때까지 프로그램을 블록한다.

	<-done
	fmt.Println("Done!")

}
