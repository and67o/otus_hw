package main

import (
	"fmt"
	"log"
	_time "time"

	"github.com/beevik/ntp"
)

const host = "0.beevik-ntp.pool.ntp.org"

func main() {
	currentTime := _time.Now()
	time, err := ntp.Time(host)
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Printf("current time: %v\n", currentTime)
	fmt.Printf("exact time: %v\n", time)
}
