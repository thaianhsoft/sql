package databox

type BoolDataBox struct {
	value bool
}

func NewBoolDataBox(b bool) DataBox {
	return &BoolDataBox{
		value: b,
	}
}
func (b *BoolDataBox) DataType() *DataType {
	return BoolType()
}
func (b *BoolDataBox) GetTypeId() TypeId {
	return BOOL
}

func (b *BoolDataBox) GetValue() interface{} {
	return b.value
}

