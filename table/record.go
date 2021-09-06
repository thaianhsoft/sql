package table

import (
	"reflect"

	"github.com/thaianhsoft/sql/databox"
)

type Record struct {
	datas []databox.DataBox
}

func NewRecord(values ...interface{}) *Record {
	datas := make([]databox.DataBox, 0)
	for _, value := range values {
		r := reflect.TypeOf(value)
		var box databox.DataBox
		if r.Kind() == reflect.Int {
			box = databox.NewIntDataBox(value.(int))
		}
		if r.Kind() == reflect.Float32 || r.Kind() == reflect.Float64 {
			box = databox.NewFloatDataBox(value.(float64))
		}
		if r.Kind() == reflect.String {
			box = databox.NewStringDataBox(value.(string))
		}
		if r.Kind() == reflect.Int64 {
			box = databox.NewLongDataBox(value.(int64))
		}
		if r.Kind() == reflect.Slice {
			if r.Elem().Kind() == reflect.Uint8 {
				l := len(value.([]uint8))
				box = databox.NewByteArrayDataBox(value.([]uint8), l)
			}
		}
		datas = append(datas, box)
	}
	return &Record{
		datas: datas,
	}
}

func (r *Record) GetValues() []databox.DataBox {
	return r.datas
}

func (r *Record) GetValueAt(i int) databox.DataBox {
	if i < len(r.datas) {
		return r.datas[i]
	}
	return nil
}

func (r *Record) Concat(other Record) {
	r.datas = append(r.datas, other.datas...)
}
