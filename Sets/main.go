package main

import (
	"errors"
	"fmt"
)

type Set struct {
	integerMap map[int]bool
}

func (s *Set) New() {
	s.integerMap = make(map[int]bool)
}

// Contains
func (s *Set) ContainsElement(element int) bool {
	_, exist := s.integerMap[element]
	return exist
}

func (s *Set) ContainsElementInTwoSets(set *Set, element int) bool {
	existOne := s.ContainsElement(element)
	existTwo := set.ContainsElement(element)
	return existOne && existTwo

}

func (s *Set) Add(number int, condition bool) {
	if !s.ContainsElement(number) {
		s.integerMap[number] = condition
	}
}

func (s *Set) Delete(number int) {
	delete(s.integerMap, number)
}

func (s *Set) Union(set *Set) (*Set, error) {
	union := &Set{}
	union.New()

	for element := range s.integerMap {
		condition := s.integerMap[element]
		union.Add(element, condition)
	}

	for element := range set.integerMap {
		condition := set.integerMap[element]
		union.Add(element, condition)
	}
	if len(union.integerMap) == 0 {
		return nil, errors.New("union is empty")
	}
	return union, nil
}

func (s *Set) Intersect(set *Set) (*Set, error) {
	interserct := &Set{}
	interserct.New()

	for element := range s.integerMap {
		if s.ContainsElementInTwoSets(set, element) && s.integerMap[element] == set.integerMap[element] {
			condition := s.integerMap[element]
			interserct.Add(element, condition)
		}
	}

	if len(interserct.integerMap) == 0 {
		return nil, errors.New("intersect is empty")
	}
	return interserct, nil

}

func (s *Set) Difference(set *Set) (*Set, error) {
	difference := &Set{}
	difference.New()
	for element := range s.integerMap {
		if !(s.ContainsElementInTwoSets(set, element)) || s.integerMap[element] != set.integerMap[element] {
			conditon := s.integerMap[element]
			difference.Add(element, conditon)
		}
	}
	if len(difference.integerMap) == 0 {
		return nil, errors.New("difference is empty")
	}
	return difference, nil
}

func main() {
	s := Set{}
	s.New()
	s2 := Set{}
	s2.New()
	s.Add(30000, false)
	s.Add(70, false)
	s.Add(3, true)
	s.Add(80, false)

	s2.Add(6, false)
	s2.Add(70, true)
	s2.Add(3, true)
	s2.Add(86, false)

	fmt.Println(s)
	union, err := s.Union(&s2)
	if err != nil {
		panic(err)
	}
	fmt.Println(union)

	fmt.Println("---------")

	intersect, err := s.Intersect(&s2)
	if err != nil {
		panic(err)
	}
	fmt.Println(intersect)

	fmt.Println("---------")

	difference, err := s2.Difference(&s)
	if err != nil {
		panic(err)
	}
	fmt.Println(difference)
}
