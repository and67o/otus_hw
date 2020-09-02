package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"log"
	_time "time"
)

const host = "0.beevik-ntp.pool.ntp.org"
func main() {
	currentTime:=_time.Now()
	time, err := ntp.Time(host)
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Printf("current time: %v\n", currentTime)
	fmt.Printf("exact time: %v\n", time)
}
