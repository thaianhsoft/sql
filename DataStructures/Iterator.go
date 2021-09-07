package collection

type Iterator interface {
	HasNext() bool
	Next() interface{}
}