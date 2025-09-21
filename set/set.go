package set

type Set[T comparable] map[T]struct{}

func New[T comparable](items ...T) Set[T] {
	set := make(Set[T], len(items))
	for _, item := range items {
		set[item] = struct{}{}
	}
	return set
}
func (s Set[T]) Slice() []T {
	slice := make([]T, 0, len(s))
	for item := range s {
		slice = append(slice, item)
	}
	return slice
}

func (s Set[T]) Raw() map[T]struct{} { return map[T]struct{}(s) }
func (s Set[T]) Add(item T)          { s[item] = struct{}{} }
func (s Set[T]) Has(item T) bool     { _, ok := s[item]; return ok }
