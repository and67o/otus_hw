package main

import (
	"context"
	"flag"
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

func main() {
	flag.Parse()

	args := os.Args
	if len(args) != 4 {
		log.Fatal("count of variable error")
	}

	address := net.JoinHostPort(os.Args[2], os.Args[3])

	bcg := context.Background()
	ctx, cancel := context.WithCancel(bcg)

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

	go receive(client, cancel)

	go send(client, cancel)

	watchSignal(ctx, cancel)
}

func watchSignal(ctx context.Context, cancel context.CancelFunc) {
	signalsChan := make(chan os.Signal, 1)
	signal.Notify(signalsChan, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-signalsChan:
		cancel()
	case <-ctx.Done():
		close(signalsChan)
	}
}

func receive(client TelnetClient, cancel context.CancelFunc) {
	err := client.Receive()
	if err != nil {
		cancel()
	}
}

func send(client TelnetClient, cancel context.CancelFunc) {
	err := client.Send()
	if err != nil {
		cancel()
	}
}
