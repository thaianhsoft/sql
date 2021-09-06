package databox

type LongDataBox struct {
	value int64
}

func NewLongDataBox(v int64) DataBox {
	return &LongDataBox{
		value: v,
	}
}
func (b *LongDataBox) DataType() *DataType {
	return LongType()
}
func (b *LongDataBox) GetTypeId() TypeId {
	return LONG
}

func (b *LongDataBox) GetValue() interface{} {
	return b.value
}

