package main

import (
	"fmt"
	"net"
)

const addr = "localhost:8888"

func main() {
	conns := &connections{
		addrs: make(map[string]*net.UDPAddr),
	}

	fmt.Printf("serving on %s\n", addr)

	// udp addr 설정
	addr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		panic(err)
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		panic(err)
	}
	// cleanup
	defer conn.Close()

	go broadcast(conn, conns)

	msg := make([]byte, 1024)
	for {
		// 반복해서 루프를 돌며 메시지 읽어옴.
		_, retAddr, err := conn.ReadFromUDP(msg)
		if err != nil {
			continue
		}

		//수신 메시지 저장.
		conns.mu.Lock()
		conns.addrs[retAddr.String()] = retAddr
		conns.mu.Unlock()
		fmt.Printf("%s connected\n", retAddr)
	}
}