package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	acc := initial
	for _, v := range s {
		acc = fn(acc, v)
	}
	return acc
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	acc := initial
	for i := len(s) - 1; i >= 0; i-- {
		acc = fn(s[i], acc)
	}
	return acc
}

func (s IntList) Filter(fn func(int) bool) IntList {
	out := IntList{}
	for _, v := range s {
		if fn(v) {
			out = append(out, v)
		}
	}
	return out
}

func (s IntList) Length() int {
	return len(s)
}

func (s IntList) Map(fn func(int) int) IntList {
	out := make(IntList, len(s))
	for i, v := range s {
		out[i] = fn(v)
	}
	return out
}

func (s IntList) Reverse() IntList {
	out := make(IntList, len(s))
	for i := range s {
		out[len(s)-1-i] = s[i]
	}
	return out
}

func (s IntList) Append(lst IntList) IntList {
	return append(s, lst...)
}

func (s IntList) Concat(lists []IntList) IntList {
	out := IntList{}
	for _, lst := range lists {
		out = append(out, lst...)
	}
	return out
}
