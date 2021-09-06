package databox

type StringDataBox struct {
	value string
	size int
}

func NewStringDataBox(v string, size ...int) DataBox {
	if len(size) == 1 {
		return &StringDataBox{
			value: v,
			size: size[0],
		}
	}
	return &StringDataBox{
		value: v,
		size: len(v),
	}
}
func (b *StringDataBox) DataType() *DataType {
	return StringType(b.size)
}
func (b *StringDataBox) GetTypeId() TypeId {
	return STRING
}

func (b *StringDataBox) GetValue() interface{} {
	return b.value
}

