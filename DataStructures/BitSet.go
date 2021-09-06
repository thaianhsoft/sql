package BitSet

import (
	"errors"
	"fmt"
	"log"
)
func InceedLengthOriginalParams(have interface{}, want interface{}) error {
	str := fmt.Sprintf(" have(%v)\n want(%v)\n", have, want)
	return errors.New(str)
}
type BitSet struct {
	bitset []int
	size int
}
func New(size ...int) *BitSet {
	bs := &BitSet{}
	if len(size) == 1 {
		if size[0] >= 32 {
			vSize := (size[0] % 32)+1
			bs.bitset = make([]int, vSize)
			bs.size = size[0]
			return bs
		} else {
			bs.bitset = make([]int, 1)
			bs.size = size[0]
			return bs
		}
	}
	bs.bitset = make([]int, 0)
	bs.size = 0
	return bs
}

func (b*BitSet) getBit(bit_ith int) int {
	return bit_ith%32
}

func (b*BitSet) getIndex(bit_ith int) int {
	r := bit_ith%32
	index := (bit_ith-r)/32
	return index
}

func (b*BitSet) Set(bit int) error {
	if bit >= b.size {
		return InceedLengthOriginalParams(bit, b.size)
	}
	index := b.getIndex(bit)
	vBit := b.getBit(bit)
	if (b.bitset[index] & (1<<vBit)) == 0 {
		b.bitset[index] |= 1 << vBit
 	}
	return nil
}

func (b *BitSet) Clear(bit int) error {
	if bit >= b.size {
		return InceedLengthOriginalParams(bit, b.size)
	}
	index := b.getIndex(bit)
	vBit := b.getBit(bit)
	if (b.bitset[index] & (1<<vBit)) != 0 {
		b.bitset[index] &= ^(1<<vBit)
	}
	return nil
}

func (b *BitSet) Exist(bit int) bool {
	if bit >= b.size {
		log.Panic(InceedLengthOriginalParams(bit, b.size))
	}
	index := b.getIndex(bit)
	vBit := b.getBit(bit)
	return (b.bitset[index] & (1<<vBit)) == 0
}

func (b *BitSet) Resize(newSize int) {
	if newSize < b.size {
		log.Panic("don't accept resize with smaller size to original bitset, make loss integrity bit is set")
	}
	vSize := (newSize % 32)+1
	change := vSize - len(b.bitset)
	b.size = vSize
	add := make([]int, change)
	b.bitset = append(b.bitset, add...)
}

func (b *BitSet) AddSize (addSize int) {
	add := make([]int, addSize)
	b.bitset = append(b.bitset, add...)
	b.size += addSize
}

func (b *BitSet) GetSize() int {
	return b.size
}

func (b *BitSet) toString() string {
	str := ""
	for _, one := range b.bitset {
		for i := 0; i < 32; i++ {
			if (one & (1<<i)) != 0 {
				str += "1"
			} else {
				str += "0"
			}
		}
	}
	return str
}