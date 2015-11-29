package rle

import ()

type LZW struct {
	runs map[string]int
}

func NewLZW(alphabet []byte) *LZW {
	runs := make(map[string]int)
	runs[string(10)] = -1
	i := 0
	for i < len(alphabet) {
		runs[string(alphabet[i])] = i + 1
		i++
	}

	return &LZW{runs: runs}
}

type LZWD struct {
	lookup []lzwEntry
}

func NewLZWD(alphabet []byte) *LZWD {
	lookup := make([]lzwEntry, 0, 4096)
	lookup = append(lookup, lzwEntry{char: 0, prev: 0})

	for _, a := range alphabet {
		lookup = append(lookup, lzwEntry{prev: 0, char: a})
	}

	return &LZWD{lookup: lookup}
}

type lzwEntry struct {
	prev uint16
	char byte
}

func (self *LZWD) rebuild(index uint16) []byte {
	word := []byte{}

	cur := self.lookup[index]
	for cur.char != 0 {
		word = append(word, cur.char)
		cur = self.lookup[cur.prev]
	}

	return word

}

func (self *LZWD) Decode(bytes []byte) []byte {
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
	output := make([]byte, 0, 2*len(data))
	words := self.runs
	count := len(words)

	var cur []byte
	last := 0
	var index int
	var ok bool

	i := 0

	for i < len(data) {

		cur = append(cur, data[i])

		if index, ok = words[string(cur)]; !ok {

			if last == 0 {
				break
			}

			words[string(cur)] = count

			output = append(output, byte(last&255), byte(last>>8))
			count++
			last = 0
			cur = nil
		} else {
			last = index
			i++
		}

	}

	self.runs = words
	if len(words) >= 65535 {
		panic("lookup bigger then max index")
	}

	return output
}
