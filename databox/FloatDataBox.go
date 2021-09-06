package databox

type FloatDataBox struct {
	value float64
}

func NewFloatDataBox(v float64) DataBox {
	return &FloatDataBox{
		value: v,
	}
}
func (b *FloatDataBox) DataType() *DataType {
	return FloatType()
}
func (b *FloatDataBox) GetTypeId() TypeId {
	return FLOAT
}

func (b *FloatDataBox) GetValue() interface{} {
	return b.value
}

