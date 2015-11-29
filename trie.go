package rle

import ()

type trie struct {
	nodes [10]*trie
	val   int
}
