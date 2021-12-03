package strain

type Ints []int
type Lists [][]int
type Strings []string

func (i Ints) Keep(filter func(int) bool) (res Ints) {
	if i == nil {
		return nil
	}
	for _, e := range i {
		if filter(e) {
			res = append(res, e)
		}
	}
	return res
}

func (i Ints) Discard(filter func(int) bool) (res Ints) {
	if i == nil {
		return nil
	}
	for _, e := range i {
		if !filter(e) {
			res = append(res, e)
		}
	}
	return res
}

func (l Lists) Keep(filter func([]int) bool) (res Lists) {
	if l == nil {
		return nil
	}
	for _, e := range l {
		if filter(e) {
			res = append(res, e)
		}
	}
	return res
}

func (s Strings) Keep(filter func(string) bool) (res Strings) {
	if s == nil {
		return nil
	}
	for _, e := range s {
		if filter(e) {
			res = append(res, e)
		}
	}
	return res
}
