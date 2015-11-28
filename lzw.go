package rle

import (
	"fmt"
)

type LZEncoder struct {
	head  *lzw
	len   int
	codes []int
}

func NewLZEncoder() *LZEncoder {
	self := &LZEncoder{
		codes: make([]int, 0),
	}
	self.head = self.new_lzw()
	return self
}

func (self *LZEncoder) new_lzw() *lzw {
	self.len++
	return &lzw{
		vals:  map[byte]*lzw{},
		index: self.len,
	}
}

type lzw struct {
	index int
	vals  map[byte]*lzw
}

func (self *LZEncoder) Write(data []byte) {
	cur := self.head

	if len(data) == 1 {
		cur.vals[data[0]] = self.new_lzw()
		return
	}

	for i := 0; i < len(data); i++ {
		if _, ok := cur.vals[data[i]]; !ok {
			fmt.Println("nada")
			cur.vals[data[i]] = self.new_lzw()
			self.codes = append(self.codes, cur.index)
			cur = self.head.vals[data[i]]
		} else {
			cur = cur.vals[data[i]]
		}
	}

	self.codes = append(self.codes, cur.index)
}
