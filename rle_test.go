package rle

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestRLE(t *testing.T) {

	s, err := ioutil.ReadFile("test/frame.txt")
	if err != nil {
		panic("read error")
	}

	encoder := NewRLEncoder()
	encoder.Write(s)
	fmt.Println(encoder.Dump())
	fmt.Println("input len", len(s))
	fmt.Println("output len", len(encoder.buf))

}

func TestLZW78(t *testing.T) {

	s, _ := ioutil.ReadFile("test/short.txt")

	alphabet := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	encoder := NewLZ78(alphabet)
	encoder.Encode(s)

	fmt.Println(encoder.output)

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
