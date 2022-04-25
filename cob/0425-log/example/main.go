package main

import (
	"fmt"

	log "github.com/jungkyuhyun/go-study-with-derek-cob/cob/0425-log"
)

func main() {
	fmt.Println("basic logging and modification of logger:")
	log.Log()
	fmt.Println("logging 'handled' errors:")
	log.FinalDestination()
}
