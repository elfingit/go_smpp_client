package main

import (
	"github.com/elfingit/go_smpp_client/pdu"
)

type SmppClient struct {
	transport *SmppTransport
}

func (client *SmppClient) bindTransmitter(system_id string, password string) error {
	var header = pdu.Header{}

	header.ID = pdu.BindTransmitterID
	header.Status = 0
	header.Seq = 1

	return header.SerializeTo(client.transport.connection)
}
