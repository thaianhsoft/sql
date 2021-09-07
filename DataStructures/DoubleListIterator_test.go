package collection

import (
	"log"
	"testing"

	"github.com/thaianhsoft/sql/utils"
)

func TestNewDoubleList(t *testing.T) {
	datas := []interface{}{1, 2, "thaianh"}

	var iter Iterator = DoubleList(datas...)
	for _, v := range datas {
		utils.Assert(t, iter.HasNext(), true)
		utils.Assert(t, iter.Next(), v)
	}
	utils.Assert(t, iter.HasNext(), false)
	log.Printf("%v passed", t.Name())
}

func TestPopAndPushDoubleList(t *testing.T) {
	datas := []interface{}{1, 2, "thaianh"}
	var iter *doubleList = DoubleList(datas...)
	x := iter.Pop()
	utils.Assert(t, x.v, 1) // pop is 1
	utils.Assert(t, iter.Next(), 2) // get iter is next 2

	y := iter.PopBack()
	utils.Assert(t, y.v, "thaianh") // pop is "thaianh"

	utils.Assert(t, iter.Next(), nil) // empty because iter is nil
}