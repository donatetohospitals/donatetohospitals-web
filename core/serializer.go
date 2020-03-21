package core

type Serializer interface {
	Encode(interface{}) ([]byte, error)
	Decode([]byte) (interface{}, error)
}

// TODO(daniel): add serializer definitions here
