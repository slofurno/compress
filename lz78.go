package rle

import (
	"fmt"
)

type LZ78 struct {
	lookup []*lz78entry
	output []int
}

func NewLZ78(alphabet []byte) *LZ78 {
	lookup := []*lz78entry{&lz78entry{char: 0, prev: -1}}

	for _, a := range alphabet {
		lookup = append(lookup, &lz78entry{prev: 0, char: a})
	}

	return &LZ78{
		output: []int{},
		lookup: lookup,
	}
}

type lz78entry struct {
	char byte
	prev int
}

func (self *LZ78) Encode(data []byte) {
	fmt.Println(data)
	lookup := self.lookup
	output := self.output
	last := 0

	i := 0

	for i < len(data) {
		j := last + 1
		for j < len(lookup) && !(lookup[j].char == data[i] && lookup[j].prev == last) {
			j++
		}
		fmt.Printf("%c", data[i])

		if j == len(lookup) {
			lookup = append(lookup, &lz78entry{char: data[i], prev: last})
			output = append(output, last)
			fmt.Println("")
			last = 0
		} else {
			last = j
			i++
		}
	}

	output = append(output, last)
	fmt.Println(last, '#', string(data))

	self.lookup = lookup
	self.output = output
}
