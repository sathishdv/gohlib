package collection

import "fmt"

type Set struct {
	elements map[interface{}]bool
}

func NewSet(elements ...interface{}) *Set {
	if len(elements) > 0 {
		set := &Set{
			elements: make(map[interface{}]bool),
		}
		for _, e := range elements {
			set.Add(e)
		}
		return set
	}
	return &Set{
		elements: make(map[interface{}]bool),
	}
}

func (s *Set) Len() int {
	return len(s.elements)
}

func (s *Set) IsEmpty() bool {
	return s.Len() == 0
}

func (s *Set) Add(e interface{}) bool {
	if _, ok := s.elements[e]; ok {
		return false
	}
	s.elements[e] = true
	return true
}

func (s *Set) Remove(e interface{}) bool {
	if _, ok := s.elements[e]; !ok {
		return false
	}

	delete(s.elements, e)
	return true
}

func (s *Set) Contains(e interface{}) bool {
	return s.elements[e]
}

func (s *Set) Clear() {
	s.elements = make(map[interface{}]bool)
}

func (s *Set) ToSlice() []interface{} {
	elements := make([]interface{}, 0)
	for e := range s.elements {
		elements = append(elements, e)
	}
	return elements
}

func (s *Set) Union(other *Set) *Set {
	union := NewSet()
	for e := range s.elements {
		union.Add(e)
	}
	for e := range other.elements {
		union.Add(e)
	}
	return union
}

func (s *Set) Intersection(other *Set) *Set {
	intersection := NewSet()
	for e := range s.elements {
		if other.Contains(e) {
			intersection.Add(e)
		}
	}
	return intersection
}

func (s *Set) Difference(other *Set) *Set {
	difference := NewSet()
	for e := range s.elements {
		if !other.Contains(e) {
			difference.Add(e)
		}
	}
	return difference
}

func (s *Set) IsSubset(other *Set) bool {
	for e := range s.elements {
		if !other.Contains(e) {
			return false
		}
	}
	return true
}

func (s *Set) IsSuperset(other *Set) bool {
	return other.IsSubset(s)
}

func (s *Set) Equal(other *Set) bool {
	return s.IsSubset(other) && s.IsSuperset(other)
}

func (s *Set) String() string {
	return fmt.Sprintf("%v", s.ToSlice())
}
