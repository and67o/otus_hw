package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"time"
)

type TelnetClient interface {
	Connect() error
	Close() error
	Send() error
	Receive() error
}

const network = "tcp"

func NewTelnetClient(address string, timeout time.Duration, in io.ReadCloser, out io.Writer) TelnetClient {
	return &telnetClient{
		in:         in,
		out:        out,
		address:    address,
		connection: nil,
		timeout:    timeout,
	}
}

type telnetClient struct {
	in         io.ReadCloser
	out        io.Writer
	address    string
	connection net.Conn
	timeout    time.Duration
}

func (t *telnetClient) Connect() error {
	connection, err := net.DialTimeout(network, t.address, t.timeout)
	if err != nil {
		return fmt.Errorf("connection error: %w", err)
	}
	t.connection = connection
	return nil
}

func (t *telnetClient) Close() error {
	if t.connection == nil {
		return nil
	}
	err := t.connection.Close()
	return fmt.Errorf("close connection error: %w", err)
}

func (t *telnetClient) Send() error {
	scanner := bufio.NewScanner(t.in)
	for scanner.Scan() {
		text := scanner.Text()
		_, err := t.connection.Write([]byte(text + "\n"))
		if err != nil {
			return fmt.Errorf("send error: %w", err)
		}
	}
	return nil
}

func (t *telnetClient) Receive() error {
	if t.connection == nil {
		return nil
	}
	_, err := io.Copy(t.out, t.connection)
	return fmt.Errorf("receive error: %w", err)
}
