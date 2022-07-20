package main

type Queue interface {
	push()
	pop()
	top() interface{}
	length() int
}

type TSlice struct {
	i    []interface{}
	size int
}

func (a *TSlice) push(data interface{}) {

	a.size++
	a.i = append(a.i, data)

}

func (a *TSlice) pop() {
	a.size--
	a.i = a.i[1:]
}

func (a *TSlice) top() interface{} {
	return a.i[0]
}

func (a *TSlice) length() int {
	return a.size
}
