package main

import "fmt"

// Byte is a bitmask in byte.
type Byte byte

// Has returns true if this bitmask contains another bitmask.
func (b Byte) Has(bb Byte) bool {
	return (b & bb) != 0
}

func (b *Byte) Set(bb Byte) {
	*b |= bb
}

func (b *Byte) Clear(bb Byte) {
	*b &= ^bb
}

func (b *Byte) Toggle(bb Byte) {
	*b ^= bb
}

func main() {
	var bt Byte
	bt.Set('c')
	fmt.Println(bt.Has('c'))
}

func xor(b []byte) []byte {
	r := make([]byte, len(b))
	for i := range b {
		r[i] = b[i] ^ 'z'
	}
	return r
}
