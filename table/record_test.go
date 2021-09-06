package table

import (
	"log"
	"testing"

	"github.com/thaianhsoft/sql/utils"
)

func TestNewRecord(t *testing.T) {
	Int := 1
	String := "thaianh"
	Float := 3.2
	ByteArray := []uint8("abc")
	datas := make([]interface{}, 0)
	datas = append(datas, Int, String, Float, ByteArray)
	r := NewRecord(Int, String, Float, ByteArray)
	for i := 0; i < len(r.datas); i++ {
		if i == 3 {
			// byte array
			v := r.datas[i].GetValue().([]uint8)
			for j := 0; j < len(v); j++ {
				utils.Assert(t, datas[i].([]uint8)[j], v[j])
			}
		} else {
			utils.Assert(t, datas[i], r.datas[i].GetValue())
		}
	}
}
