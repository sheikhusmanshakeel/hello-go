package indexer

import (
	"fmt"
)

type I interface {
		M()
		N()
	}


type W interface {
	func1(float64) string
	func2(int64)string
}

type T struct {
	S string
}

func (t T) func2(int64) string {
	panic("implement me")
}

func (t T) N() {
	panic("implement me")
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t T) M() {
	fmt.Println(t.S)
}

func (t T)  func1(v float64) string{
	if t.S == "Hello Oz"{
		return fmt.Sprintf("%s %f", t.S, v)
	} else{
		return ""
	}
}

func runWeirdInterface() {
	val := T{("Hello Oz")}
	var i1 I = val
	i1.M()

	v := i1.(W)
	v.func1(34)
	var vType = fmt.Sprintf("%T",v)


	println(&v)
	println(&i1)
	println(vType)


	v9 := T{("Hello v9")}
	var w W = v9
	w.func1(34)

	v10 := w.(I)
	v10.M()

}
