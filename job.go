package workerpool

type Job[T any] interface {
	Execute() T
}

type SimpJob[T any] struct {
	id int
	fn inFunc[T]
}

func (j SimpJob[T]) Execute() T {
	res := j.fn(j.id)
	return res
}

// No Channel just provide a ptr to store your value when its ready

type PtrJob[T any] struct {
	id  int
	ptr *T
	fn  inFunc[T]
}

func (j PtrJob[T]) Execute() T {
	*j.ptr = j.fn(j.id)
	return *j.ptr
}
