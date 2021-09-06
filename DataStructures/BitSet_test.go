package BitSet

import (
	"log"
	"testing"
)

func assert(t *testing.T, have interface{}, want interface{}) {
	if have != want {
		log.Panicf("have(%v)\n, want(%v)\n", have, want)
		t.Failed()
	}
}

func TestOperations(t *testing.T) {
	size := 200
	bs := New(size) 
	vSize := (size % 32)+1
	assert(t, vSize, len(bs.bitset))

	bit := 199
	b := bit % 32
	i := (bit-b)/32
	v := bs.bitset[i]
	v |= 1<<b
	bs.Set(bit)
	assert(t, v, bs.bitset[i])
	log.Println(v, bs.bitset[i])
	log.Println("Test Operations is passed")
}

func TestReSize(t *testing.T) {
	size := 200
	bs := New(size)
	new_size := 300
	vNewSize := (new_size%32)+1
	vSize := (size%32)+1

	change := vNewSize - vSize
	bs.Resize(new_size)
	assert(t, change, len(bs.bitset)-vSize)
	log.Println(change, len(bs.bitset))
	log.Println("Test Resize is passed")
}

func TestToString(t *testing.T) {
	size := 200
	bs := New(size)
	bs.Set(199)
	bs.Set(54)
	log.Println(bs.toString())
}