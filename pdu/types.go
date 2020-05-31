package pdu

type Name string

const (
	SystemID   Name = "system_id"
	SystemType Name = "system_type"
	Password   Name = "password"
)

const (
	BindTransmitterID ID = 0x00000002
)

type Field interface {
	Len() int
	Raw() interface{}
	String() string
	Bytes() []byte
}

type Variable struct {
	Data []byte
}

func (v *Variable) Len() int {
	return len(v.Bytes())
}

func (v *Variable) Bytes() []byte {
	if len(v.Data) > 0 && v.Data[len(v.Data)-1] == 0x00 {
		return v.Data
	}
	return append(v.Data, 0x00)
}

func (v *Variable) String() string {
	if l := len(v.Data); l > 0 && v.Data[l-1] == 0x00 {
		return string(v.Data[:l-1])
	}
	return string(v.Data)
}

func (v *Variable) Raw() interface{} {
	return v.Data
}
