package rle

import ()

type Encoder interface {
	Dump() []byte
	Write([]byte)
}

type RLEncoder struct {
	dict []byte
	len  int
	buf  []byte
}

func NewRLEncoder() *RLEncoder {
	return &RLEncoder{
		dict: make([]byte, 32),
		buf:  make([]byte, 0),
	}
}

func (self *RLEncoder) Dump() []byte {
	return self.buf
}

func (self *RLEncoder) Write(s []byte) {
	if len(s) == 0 {
		return
	}

	var last byte = s[0]
	var run byte = 1

	for i := 1; i < len(s); i++ {
		c := s[i]

		if c == last && run < 255 {
			run++
		} else {
			self.buf = append(self.buf, run, last)
			run = 1
			last = c
		}
	}
}
