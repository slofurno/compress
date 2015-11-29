package rle

import ()

type LZW struct {
	lookup []lzwEntry
}

func NewLZW(alphabet []byte) *LZW {
	lookup := make([]lzwEntry, 0, 4096)
	lookup = append(lookup, lzwEntry{char: 0, prev: 0})
	//	lookup := []*lzwEntry{&lzwEntry{char: 0, prev: 0}}

	for _, a := range alphabet {
		lookup = append(lookup, lzwEntry{prev: 0, char: a})
	}

	return &LZW{lookup: lookup}
}

type lzwEntry struct {
	prev uint16
	char byte
}

func (self *LZW) LLen() int {
	return len(self.lookup)
}

func (self *LZW) rebuild(index uint16) []byte {
	word := []byte{}

	cur := self.lookup[index]
	for cur.char != 0 {
		word = append(word, cur.char)
		cur = self.lookup[cur.prev]
	}

	return word

}

func (self *LZW) Decode(bytes []byte) []byte {
	input := []uint16{}
	for i := 0; i < len(bytes); i += 2 {
		input = append(input, (uint16(bytes[i]) | (uint16(bytes[i+1]) << 8)))
	}

	output := []byte{self.lookup[input[0]].char}

	for i := 1; i < len(input); i++ {
		var word []byte

		if input[i] >= uint16(len(self.lookup)) {

			sc := self.rebuild(input[i-1])
			self.lookup = append(self.lookup, lzwEntry{
				char: sc[len(sc)-1],
				prev: input[i-1],
			})
			//have Sc want cScSc
			//word = append(word, sc[len(sc)-1])
			//word = append(word, sc...)
			//word = append(word, sc...)
			word = self.rebuild(input[i])
		} else {
			word = self.rebuild(input[i])
			self.lookup = append(self.lookup, lzwEntry{
				char: word[len(word)-1],
				prev: input[i-1],
			})
		}

		for j := len(word) - 1; j >= 0; j-- {
			output = append(output, word[j])
		}
		//last output + first letter of current word must be next added to dict by encoder
	}

	return output
}

func (self *LZW) Encode(data []byte) []byte {
	lookup := self.lookup
	output := make([]byte, 0, 2*len(data))
	var last uint16 = 0

	i := 0

	for i < len(data) {
		j := last + 1
		for j < uint16(len(lookup)) &&
			(lookup[j].char != data[i] || lookup[j].prev != last) {
			j++
		}

		if j == uint16(len(lookup)) {
			if last == 0 {
				//last char isnt in our alphabet
				break
			}

			lookup = append(lookup, lzwEntry{char: data[i], prev: last})
			output = append(output, byte(last&255), byte(last>>8))
			last = 0
		} else {
			last = j
			i++
		}
	}

	self.lookup = lookup
	if len(lookup) >= 65535 {
		panic("lookup bigger then max index")
	}

	return output
}
