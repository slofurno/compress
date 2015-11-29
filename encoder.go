package rle

type Encoder interface {
	Decode([]byte) []byte
	Encode([]byte) []byte
}
