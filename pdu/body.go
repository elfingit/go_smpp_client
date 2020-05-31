package pdu

type Body struct {
	h *Header
	f Map
}

func New(n Name, data []byte) Field {
	switch n {
	case
		Password,
		SystemType,
		SystemID:
		if data == nil {
			data = []byte{}
		}
		return &Variable{Data: data}

	default:
		return nil
	}
}
