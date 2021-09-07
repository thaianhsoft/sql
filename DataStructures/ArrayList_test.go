package collection

import (
	"log"
	"testing"

	"github.com/thaianhsoft/sql/utils"
)

func TestArrayList(t *testing.T) {
	values := []interface{}{1, 2, "hello", "conkac"}
	var iter Iterator = Array(values...)
	for i := 0; i < len(values); i++ {
		utils.Assert(t, true, iter.HasNext())
		utils.Assert(t, values[i], iter.Next())
	}
	utils.Assert(t, false, iter.HasNext())
	log.Printf("%v passed", t.Name())
}
