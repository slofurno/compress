package rle

import (
	"fmt"
	"io/ioutil"
	"testing"
)

/*
func TestRLE(t *testing.T) {

	s, err := ioutil.ReadFile("test/frame.txt")
	if err != nil {
		panic("read error")
	}

	encoder := NewRLEncoder()
	encoder.Write(s)
	//fmt.Println(encoder.Dump())
	fmt.Println("input len", len(s))
	fmt.Println("output len", len(encoder.buf))

}
*/
func TestLZW78(t *testing.T) {

	s, _ := ioutil.ReadFile("test/short.txt")
	fmt.Println("source:", string(s))

	alphabet := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	encoder := NewLZ78(alphabet)
	encoded := encoder.Encode(s)

	decoder := NewLZ78(alphabet)
	decoded := decoder.Decode(encoded)

	fmt.Println("decoded:", string(decoded))

}

func TestLZW78_long(t *testing.T) {
	s, _ := ioutil.ReadFile("test/frame.txt")
	fmt.Println("source:", string(s[2000:4000]))
	alphabet := []byte(".^:-=+*#%@")

	encoder := NewLZ78(alphabet)
	encoded := encoder.Encode(s)
	//fmt.Println("encoded:", encoded)

	decoder := NewLZ78(alphabet)
	decoded := decoder.Decode(encoded)

	fmt.Println("decoded:", string(decoded[2000:4000]))

	fmt.Println(len(s), len(encoded), len(decoded))

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
