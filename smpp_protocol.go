package main

import (
	"io"

	"github.com/elfingit/go_smpp_client/pdu"
)

type SmppClient struct {
	transport *SmppTransport
}

type BindTransmitterPdu struct {
	system_id        string
	password         string
	system_type      string
	iterface_version int
	addr_ton         int
	addr_npi         int
	address_range    string
}

func (pdu *BindTransmitterPdu) Length() int {
	var length int = 0

	length += len(pdu.system_id)
	length += len(pdu.password)
	length += len(pdu.system_type)

	length += 4

	return length
}

func (body *BindTransmitterPdu) SerializeTo(w io.Writer) error {

	b := make([]byte, body.Length())

	system_id := []byte(body.system_id)
	b = append(b, system_id...)

	password := []byte(body.password)
	b = append(b, password...)

	system_type := []byte(body.system_type)
	b = append(b, system_type...)

	iv_b := make([]byte, 1)
	iv_b[0] = byte(body.iterface_version)
	b = append(b, iv_b...)

	addrT_b := make([]byte, 1)
	addrT_b[0] = byte(body.addr_ton)
	b = append(b, addrT_b...)

	addrN_b := make([]byte, 1)
	addrN_b[0] = byte(body.addr_npi)
	b = append(b, addrN_b...)

	addrR_b := []byte(body.address_range)
	b = append(b, addrR_b...)

	_, errW := w.Write(b)

	return errW
}

func (client *SmppClient) bindTransmitter(system_id string, password string, system_type string) error {
	var header = pdu.Header{}

	header.ID = pdu.BindTransmitterID
	header.Status = 0
	header.Seq = 1

	var body = BindTransmitterPdu{}

	body.system_id = system_id
	body.password = password
	body.system_type = system_type
	body.iterface_version = 34

	header.Len = uint32(body.Length())

	err := header.SerializeTo(client.transport.connection)

	if err != nil {
		return err
	}

	return body.SerializeTo(client.transport.connection)
}
