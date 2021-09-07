package collection

type List interface{
	Iterator
	list()
}

type array struct {
	datas []interface{}
	index int
	size int
}

func Array(values ...interface{}) *array {
	return &array{
		datas: values,
		index: -1,
		size: len(values),
	}
}

func (a *array) HasNext() bool {
	return a.index + 1 < a.size
}
func (a *array) Next() interface{} {
	a.index++
	return a.datas[a.index]
}
