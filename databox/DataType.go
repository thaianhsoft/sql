package databox

import (
	"bytes"

	"github.com/thaianhsoft/sql/algo"
)

type DataType struct {
	t           TypeId
	sizeInBytes int
}

func BoolType() *DataType {
	return &DataType{
		t:           BOOL,
		sizeInBytes: 1,
	}
}

func IntType() *DataType {
	return &DataType{
		t:           INT,
		sizeInBytes: 4,
	}
}

func FloatType() *DataType {
	return &DataType{
		t:           FLOAT,
		sizeInBytes: 8,
	}
}

func StringType(n int) *DataType {
	return &DataType{
		t:           STRING,
		sizeInBytes: n,
	}
}

func LongType() *DataType {
	return &DataType{
		t:           LONG,
		sizeInBytes: 8,
	}
}

func ByteArrayType(n int) *DataType {
	return &DataType{
		t:           BYTE_ARRAY,
		sizeInBytes: n,
	}
}

func (d *DataType) GetTypeId() TypeId {
	return d.t
}

func (d *DataType) GetSizeInBytes() int {
	return d.sizeInBytes
}

func (d *DataType) ToBytes() []byte {
	buf := bytes.NewBuffer(make([]byte, 8))
	return buf.Bytes()
}

func (d *DataType) HashCode() uint64 {
	return algo.HashCode(d.t, d.sizeInBytes)
}
