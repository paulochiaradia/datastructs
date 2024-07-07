package main

import (
	"errors"
	"fmt"
)

type Set struct {
	intergerMap map[int]bool
}

func (set *Set) New() {
	set.intergerMap = make(map[int]bool)
}
func (set *Set) ContainsElement(element int) bool {
	_, exist := set.intergerMap[element]
	return exist
}

func (set *Set) AddElement(element int, b bool) {
	if !set.ContainsElement(element) {
		set.intergerMap[element] = b
	} else {
		fmt.Println("This data already exists")
	}
}

func (set *Set) Delete(element int) {
	delete(set.intergerMap, element)
}

func (set *Set) BothContain(element int, anotherSet *Set) bool {
	countainInSet := set.ContainsElement(element)
	countainInAnotherSet := anotherSet.ContainsElement(element)
	return countainInSet && countainInAnotherSet
}

func (set *Set) Intersect(anotherSet *Set) (*Set, error) {
	intersectSet := &Set{}
	intersectSet.New()
	for data := range set.intergerMap {
		if set.BothContain(data, anotherSet) {
			condition := set.intergerMap[data]
			intersectSet.AddElement(data, condition)
		}
	}
	if len(intersectSet.intergerMap) == 0 {
		return nil, errors.New("intersection is nil")
	} else {
		return intersectSet, nil
	}

}

func (set *Set) Union(anotherSet *Set) (*Set, error) {
	unionSet := &Set{}
	unionSet.New()
	// for data := range set.intergerMap {
	// 	if set.BothContain(data, anotherSet) {
	// 		continue
	// 	}
	// 	condition := set.intergerMap[data]
	// 	unionSet.AddElement(data, condition)
	// }

	// for data := range anotherSet.intergerMap {
	// 	if set.BothContain(data, set) {
	// 		continue
	// 	}
	// 	condition := anotherSet.intergerMap[data]
	// 	unionSet.AddElement(data, condition)
	// }
	var value int
	for value = range set.intergerMap {
		condition := set.intergerMap[value]
		unionSet.AddElement(value, condition)
	}

	for value = range anotherSet.intergerMap {
		condition := anotherSet.intergerMap[value]
		unionSet.AddElement(value, condition)
	}
	if len(unionSet.intergerMap) == 0 {
		return nil, errors.New("intersection is nil")
	} else {
		return unionSet, nil
	}

}

func main() {
	set := &Set{}
	anotherSet := &Set{}

	set.New()
	anotherSet.New()

	set.AddElement(1, true)
	set.AddElement(2, false)
	set.AddElement(3, false)

	anotherSet.AddElement(10, true)
	anotherSet.AddElement(2, false)
	anotherSet.AddElement(3, false)

	intersection, err := set.Intersect(anotherSet)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(intersection)
	}

	union, err := set.Union(anotherSet)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(union)
	}

}
