package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var timeout time.Duration

func init() {
	flag.DurationVar(&timeout, "timeout", 10*time.Second, "timeout connection")
}

const (
	minArgs = 3
	maxArgs = 4
)

func main() {
	flag.Parse()

	args := os.Args
	if (len(args) > maxArgs) || (len(args) < minArgs) {
		log.Fatal("count of variable error")
	}
	host, port := os.Args[len(os.Args)-2], os.Args[len(os.Args)-1]
	address := net.JoinHostPort(host, port)

	client := NewTelnetClient(address, timeout, os.Stdin, os.Stdout)

	err := client.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := client.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	errorsChan := make(chan error)
	go send(client, errorsChan)
	go receive(client, errorsChan)

	signalsChan := make(chan os.Signal, 1)
	signal.Notify(signalsChan, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	select {
	case <-signalsChan:
		signal.Stop(signalsChan)
		return
	case err = <-errorsChan:
		if err != nil {
			log.Fatal(err)
		}
		_, _ = fmt.Fprintf(os.Stderr, "End\n")
		return
	}
}

func receive(client TelnetClient, errorsChan chan error) {
	errorsChan <- client.Receive()
}

func send(client TelnetClient, errorsChan chan error) {
	errorsChan <- client.Send()
}
