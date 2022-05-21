package types

import "reflect"

type List struct {
	objects []Object
}

func NewList(objects ...Object) *List {
	return &List{objects: objects}
}

func MakeList(o Object) (*List, bool) {
	list, ok := o.(*List)
	return list, ok
}

func (l List) Values() []Object {
	return l.objects
}

func (l *List) Add(o Object) {
	l.objects = append(l.objects, o)
}

func (l List) IsEmpty() bool {
	return len(l.objects) == 0
}

func (l List) Length() int {
	return len(l.objects)
}

func (l List) Get(idx int) Object {
	return l.objects[idx]
}

func IsList(o Object) bool {
	return reflect.TypeOf(o) == reflect.TypeOf(&List{})
}

func MustMakeList(o Object) *List {
	return o.(*List)
}
