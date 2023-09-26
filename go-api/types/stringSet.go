package types

type StringSet []string

func (s StringSet) Contains(str string) bool {
	for _, es := range s {
		if es == str {
			return true
		}
	}

	return false
}

func NewStringSet(strs ...string) StringSet {
	s := StringSet{}
	s = append(s, strs...)
	return s
}
