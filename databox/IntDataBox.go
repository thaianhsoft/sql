package databox

type IntDataBox struct {
	value int
}

func NewIntDataBox(v int) DataBox {
	return &IntDataBox{
		value: v,
	}
}
func (b *IntDataBox) DataType() *DataType {
	return IntType()
}
func (b *IntDataBox) GetTypeId() TypeId {
	return INT
}

func (b *IntDataBox) GetValue() interface{} {
	return b.value
}

