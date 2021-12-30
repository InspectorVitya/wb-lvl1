/*
Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
*/
package main

import (
	"fmt"
	"strings"
)

type Set struct {
	collection map[string]struct{}
}

func (s *Set) Collections() map[string]struct{} {
	return s.collection
}

func NewSet(elems ...string) *Set {
	set := &Set{
		make(map[string]struct{}),
	}
	if len(elems) > 0 {
		for _, val := range elems {
			set.Add(val)
		}
	}
	return set
}

// Add
// Если элемент уже присутствует в множестве, возвращается false, иначе true.
func (s *Set) Add(value string) bool {
	if exist := s.Contains(value); !exist {
		s.collection[value] = struct{}{}
		return true
	}
	return false
}

// Contains
// Возвращает true, если множество содержит указанный элемент.
// В противном случае возвращает false.
func (s Set) Contains(value string) bool {
	_, ok := s.collection[value]
	return ok
}

func (s *Set) Size() int {
	return len(s.collection)
}

// Remove
// Удаляет указанный элемент из множества и возвращает true.
// В случае, если элемент нет и он не удален, возвращает false.
func (s *Set) Remove(value string) bool {
	exist := s.Contains(value)
	if !exist {
		return false
	}
	delete(s.collection, value)
	return true
}

// Union
//Возвращает множество, полученное операцией объединения его с указанным.
func (s *Set) Union(set *Set) *Set {
	union := &Set{
		make(map[string]struct{}),
	}

	for val := range s.Collections() {
		union.Add(val)
	}

	for val := range set.Collections() {
		union.Add(val)
	}

	return union
}

// Intersection
// Возвращает множество, полученное операцией пересечения его с указанным.
func (s *Set) Intersection(set *Set) *Set {
	intersection := &Set{
		make(map[string]struct{}),
	}

	for val := range set.Collections() {
		if exist := s.Contains(val); exist {
			intersection.Add(val)
		}
	}

	return intersection
}

func (s *Set) String() string {
	var b strings.Builder
	for val := range s.collection {
		b.WriteString(val)
		b.WriteString(" ")
	}
	return b.String()
}
func main() {
	str := []string{"cat", "cat", "dog", "cat", "tree"}
	set1 := NewSet(str...)
	set2 := NewSet("cat", "wb", "twitch", "tree", "dog")
	fmt.Println(set1)

	set1.Add("golang")
	fmt.Println("Add: ", set1)

	set1.Remove("cat")
	fmt.Println("Remove: ", set1)

	setIntersection := set1.Intersection(set2)
	fmt.Println("Intersection: ", setIntersection)

	setUnion := set1.Union(set2)
	fmt.Println("Union: ", setUnion)
}
