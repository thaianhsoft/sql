package databox

type DataBox interface {
	DataType() *DataType
	GetTypeId() TypeId
	GetValue() interface{}
}

