package rle

import (
	"fmt"
)

type LZ78 struct {
	lookup []*lz78entry
	output []uint16
}

func NewLZ78(alphabet []byte) *LZ78 {
	lookup := []*lz78entry{&lz78entry{char: 0, prev: 0}}

	for _, a := range alphabet {
		lookup = append(lookup, &lz78entry{prev: 0, char: a})
	}

	return &LZ78{
		output: []uint16{},
		lookup: lookup,
	}
}

type lz78entry struct {
	char byte
	prev uint16
}

func (self *LZ78) rebuild(index uint16) []byte {
	word := []byte{}

	cur := self.lookup[index]
	for cur.char != 0 {
		word = append(word, cur.char)
		cur = self.lookup[cur.prev]
	}

	return word

}

func (self *LZ78) Decode(bytes []byte) []byte {
	input := []uint16{}
	for i := 0; i < len(bytes); i += 2 {
		input = append(input, (uint16(bytes[i]) | (uint16(bytes[i+1]) << 8)))
	}

	output := []byte{self.lookup[input[0]].char}

	for i := 1; i < len(input); i++ {
		var word []byte

		if input[i] >= uint16(len(self.lookup)) {

			sc := self.rebuild(input[i-1])
			fmt.Println("rebuilt : ", string(sc))
			self.lookup = append(self.lookup, &lz78entry{char: sc[len(sc)-1], prev: input[i-1]})
			//have Sc want cScSc
			//word = append(word, sc[len(sc)-1])
			//word = append(word, sc...)
			//word = append(word, sc...)
			word = self.rebuild(input[i])

		} else {

			word = self.rebuild(input[i])
			self.lookup = append(self.lookup, &lz78entry{char: word[len(word)-1], prev: input[i-1]})
		}

		for j := len(word) - 1; j >= 0; j-- {
			output = append(output, word[j])
		}
		//last output + first letter of current word must be next added to dict by encoder
	}

	return output
}

func (self *LZ78) Encode(data []byte) []byte {
	lookup := self.lookup
	output := []uint16{}
	var last uint16 = 0

	i := 0

	for i < len(data) {
		j := last + 1
		for j < uint16(len(lookup)) && !(lookup[j].char == data[i] && lookup[j].prev == last) {
			j++
		}
		//		fmt.Printf("%c", data[i])

		if j == uint16(len(lookup)) {
			if last == 0 {
				fmt.Println("i think were done here")
				break
			}
			lookup = append(lookup, &lz78entry{char: data[i], prev: last})
			output = append(output, last)
			//fmt.Println(",", len(lookup))
			last = 0
		} else {
			last = j
			i++
		}
	}

	self.lookup = lookup

	fmt.Println("lookup size: ", len(lookup))

	if len(lookup) >= 65535 {
		panic("lookup bigger then max index")
	}

	tevs := []byte{}

	for i = 0; i < len(output); i++ {
		tevs = append(tevs, byte(output[i]&255), byte(output[i]>>8))
	}

	return tevs
}
