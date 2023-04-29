package slicelist

type SliceList[T comparable] struct {
	l := []T{}
}

func NewSliceList[T comparable]() *SliceList[T] {
	return &SliceList[T]{}
}

func (l *SliceList[T]) Append(value T) {

}

func (l *SliceList[T]) Prepend(value T) {

}

func (l *SliceList[T]) InsertAt(value T, position int) {

}

func (l *SliceList[T]) Remove(value T) {

}

func (l *SliceList[T]) String() string {
	return ""
}

func (l *SliceList[T]) Search(value T) int {
	return 0
}

func (l *SliceList[T]) Get(position int) {

}

func (l *SliceList[T]) Size() int {
	return 0
}
