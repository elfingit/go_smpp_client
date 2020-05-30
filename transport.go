package main

import (
	"errors"
	"net"
)

type SmppTransport struct {
	isOpen     bool
	connection net.Conn
}

func (transport *SmppTransport) Open(host string, port string) error {
	address := host + ":" + port

	connection, err := net.Dial("tcp", address)

	if err != nil {
		return err
	}

	transport.connection = connection
	transport.isOpen = true

	return nil
}

func (transport *SmppTransport) Write(b []byte) error {

	if !transport.isOpen {
		return errors.New("Firstly you are must open connection")
	}

	_, err := transport.connection.Write(b)

	if err != nil {
		return err
	}

	return nil
}

func (transport *SmppTransport) Close() error {
	if transport.isOpen {
		err := transport.connection.Close()

		if err != nil {
			return err
		}
	}

	return nil
}
