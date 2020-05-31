package pdu

import (
	"encoding/binary"
	"io"
)

type (
	ID uint32
)

const HeaderLen = 16

type Header struct {
	Len    uint32
	ID     ID
	Status uint32
	Seq    uint32
}

func (h *Header) SetLen(length int) {
	length += HeaderLen
	h.Len = uint32(length)
}

func (h *Header) SerializeTo(w io.Writer) error {
	b := make([]byte, HeaderLen)
	binary.BigEndian.PutUint32(b[0:4], h.Len)
	binary.BigEndian.PutUint32(b[4:8], uint32(h.ID))
	binary.BigEndian.PutUint32(b[8:12], uint32(h.Status))
	binary.BigEndian.PutUint32(b[12:16], h.Seq)
	_, err := w.Write(b)
	return err
}
