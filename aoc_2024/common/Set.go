package common

type Set[T comparable] map[T]struct{}

func (s Set[T]) Intersection(comp Set[T]) Set[T] {
    xsections := make(Set[T])

    for val := range s {
        if _, ok := comp[val]; ok { xsections[val] = struct{}{} }
    }

    return xsections
}

func ToSet[T comparable](list []T) Set[T] {
    set := Set[T]{}
    for _, val := range list { set[val] = struct{}{} }
    return set
}

func (s Set[T]) Add(val T) {
    s[val] = struct{}{}
}

func (s Set[T]) Remove(val T) {
    delete(s, val)
}

func (s Set[T]) Contains(val T) bool {
    _, ok := s[val]
    return ok
}
