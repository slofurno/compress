package rle

import ()

type RLE struct {
	dict []byte
	buf  []byte
}

func NewRLE() *RLE {
	return &RLE{
		dict: make([]byte, 32),
		buf:  make([]byte, 0),
	}
}

func (self *RLE) Dump() []byte {
	return self.buf
}

func (self *RLE) Encode(data []byte) []byte {
	if len(data) == 0 {
		return self.buf
	}

	var last byte = data[0]
	var run byte = 1

	for i := 1; i < len(data); i++ {
		c := data[i]

		if c == last && run < 255 {
			run++
		} else {
			self.buf = append(self.buf, run, last)
			run = 1
			last = c
		}
	}

	return self.buf
}
