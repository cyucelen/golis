package types

type Vector struct {
	objects []Object
}

func NewVector(objects ...Object) *Vector {
	return &Vector{objects: objects}
}

func (v Vector) Values() []Object {
	return v.objects
}

func (v *Vector) Add(o Object) {
	v.objects = append(v.objects, o)
}

func (v Vector) Get(idx int) Object {
	return v.objects[idx]
}

func (v Vector) Length() int {
	return len(v.objects)
}
