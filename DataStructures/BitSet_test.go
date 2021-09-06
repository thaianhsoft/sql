package collection

import (
	"testing"

	"github.com/thaianhsoft/sql/utils"
)

func TestOperations(t *testing.T) {
	size := 200
	bs := New(size)
	bit := 199
	b := bit % 32
	i := (bit - b) / 32
	v := bs.bitset[i]
	v |= 1 << b
	bs.Set(bit)
	utils.Assert(t, v, bs.bitset[i])
}

func TestReSize(t *testing.T) {
	size := 200
	bs := New(size)
	new_size := 300
	vNewSize := (new_size % 32) + 1
	vSize := (size % 32) + 1

	change := vNewSize - vSize
	bs.Resize(new_size)
	utils.Assert(t, change, len(bs.bitset)-vSize)
}

func TestToString(t *testing.T) {
	size := 10
	bs := New(size)
	bit := 9
	bs.Set(bit)

	str := ""
	for i := 0; i < 32; i++ {
		if i == bit {
			str += "1"
		} else {
			str += "0"
		}
	}

	utils.Assert(t, str, bs.toString())
}
