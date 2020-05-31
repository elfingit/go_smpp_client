package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	host, host_exists := os.LookupEnv("HOST")
	port, port_exists := os.LookupEnv("PORT")

	if host_exists && port_exists {
		log.Printf("Host: %s Port: %s", host, port)

		transport := new(SmppTransport)

		err := transport.Open(host, port)

		if err != nil {
			log.Fatal(err)
		}

		protocol := new(SmppClient)
		protocol.transport = transport

		err = protocol.bindTransmitter("login", "pass", "test")

		if err != nil {
			log.Fatal(err)
		}

		transport.Close()
	}
}
