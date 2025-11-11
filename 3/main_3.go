package main

import (
	"fmt"
)

type StringIntMap struct {
	data map[string]int
}

func main() {
	myMap := New()
	myMap.Add("test1", 123)
	myMap.Add("test2", 321)
	myMap.Add("test3", 322)
	fmt.Println("NEW")
	myMap.Print()

	myMap.Add("test4", 212)
	fmt.Println("AFTER ADD")
	myMap.Print()

	myMap.Remove("test1")
	fmt.Println("AFTER REMOVE")
	myMap.Print()

	copy := myMap.Copy()
	fmt.Println("COPY:", copy)

	ok := myMap.Exist("test1")
	fmt.Println(`IS EXIST KEY "test1":`, ok)

	value, getOk := myMap.Get("test4")
	fmt.Printf("GETTING KEY \"test4\", value: %d, success: %v\n", value, getOk)
}

func New() StringIntMap {
	return StringIntMap{
		data: map[string]int{},
	}
}

func (s *StringIntMap) Print() {
	for k, v := range s.data {
		fmt.Printf("key: \"%s\", value: %d\n", k, v)
	}
}

func (s *StringIntMap) Add(key string, value int) {
	s.data[key] = value
}

func (s *StringIntMap) Remove(key string) {
	delete(s.data, key)
}

func (s *StringIntMap) Copy() map[string]int {
	newMap := map[string]int{}

	// maps.Copy(newMap, s.data) можно использовать этот метод, но я написал цикл, т.к. не хотел подключать пакет maps

	for k, v := range s.data {
		newMap[k] = v
	}

	return newMap
}

func (s *StringIntMap) Exist(key string) bool {
	_, ok := s.data[key]
	return ok
}

func (s *StringIntMap) Get(key string) (int, bool) {
	value, ok := s.data[key]
	return value, ok
}
