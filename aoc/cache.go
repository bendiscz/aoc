package aoc

type Cacheable[T comparable] interface {
	Key() T
}

type CachedFunc[T comparable, IN Cacheable[T], OUT any] struct {
	results map[T]OUT
	fn      func(IN) OUT
}

func NewCachedFunc[T comparable, IN Cacheable[T], OUT any](fn func(IN) OUT) func(IN) OUT {
	c := &CachedFunc[T, IN, OUT]{
		results: make(map[T]OUT),
		fn:      fn,
	}
	return func(in IN) OUT { return c.Call(in) }
}

func (c *CachedFunc[T, IN, OUT]) Call(in IN) OUT {
	key := in.Key()
	if r, ok := c.results[key]; ok {
		return r
	}
	r := c.fn(in)
	c.results[key] = r
	return r
}
