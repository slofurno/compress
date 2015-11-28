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

func (self *LZ78) rebuild(index int) []byte {
	word := []byte{}

	cur := self.lookup[index]
	for cur.prev != -1 {
		word = append(word, cur.char)
		cur = self.lookup[cur.prev]
	}

	return word

}

func (self *LZ78) Decode(input []int) []byte {
	output := []byte{self.lookup[input[0]].char}

	for i := 1; i < len(input); i++ {
		var word []byte

		if input[i] >= len(self.lookup) {

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

func (self *LZ78) Encode(data []byte) []int {
	fmt.Println(data)
	lookup := self.lookup
	output := []int{}
	last := 0

	i := 0

	for i < len(data) {
		j := last + 1
		for j < len(lookup) && !(lookup[j].char == data[i] && lookup[j].prev == last) {
			j++
		}
		//		fmt.Printf("%c", data[i])

		if j == len(lookup) {
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
	return output
}
