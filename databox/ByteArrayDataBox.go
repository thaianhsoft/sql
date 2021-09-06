package databox

type ByteArrayDataBox struct {
	value []byte
}

func NewByteArrayDataBox(b []uint8, n int) DataBox {
	if len(b) == n {
		return &ByteArrayDataBox{
			value: b,
		}
	}
	return nil
}
func (b *ByteArrayDataBox) DataType() *DataType {
	return ByteArrayType(len(b.value))
}
func (b *ByteArrayDataBox) GetTypeId() TypeId {
	return BYTE_ARRAY
}

func (b *ByteArrayDataBox) GetValue() interface{} {
	return b.value
}
