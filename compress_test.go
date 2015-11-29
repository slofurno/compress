package rle

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestRLE_frame(t *testing.T) {

	s, err := ioutil.ReadFile("test/frame.txt")
	if err != nil {
		panic("read error")
	}

	encoder := NewRLE()
	encoded := encoder.Encode(s)
	//fmt.Println(encoder.Dump())
	fmt.Println("input len", len(s))
	fmt.Println("output len", len(encoded))

}

/*
func TestLZW_small(t *testing.T) {

	s, _ := ioutil.ReadFile("test/short.txt")
	fmt.Println("source:", string(s))

	alphabet := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	encoder := NewLZW(alphabet)
	encoded := encoder.Encode(s)

	fmt.Println(encoded)

	decoder := NewLZWD(alphabet)
	decoded := decoder.Decode(encoded)

	fmt.Println("decoded:", string(decoded))

}
*/

func TestLZW_frame(t *testing.T) {
	s, _ := ioutil.ReadFile("test/frame.txt")
	fmt.Println("source:", string(s[2000:4000]))
	alphabet := []byte(".^:-=+*#%@")

	encoder := NewLZW(alphabet)
	encoded := encoder.Encode(s)
	//fmt.Println("encoded:", encoded)

	decoder := NewLZWD(alphabet)
	decoded := decoder.Decode(encoded)

	fmt.Println("decoded:", string(decoded[2000:4000]))

	fmt.Println(len(s), len(encoded), len(decoded))

}

func Test_trie(t *testing.T) {
	tr := &trie{
		val: 24,
	}
	tr.nodes[0] = &trie{}
	fmt.Println(tr.nodes[4])
}

/*
func TestLZW(t *testing.T) {
	s, _ := ioutil.ReadFile("test/frame.txt")
	encoder := NewLZEncoder()
	init := []byte{'.', '^', ':', '-', '=', '+', '*', '#', '%', '@'}

	for _, i := range init {
		encoder.Write([]byte{i})
	}
	encoder.Write(s)
}*/
