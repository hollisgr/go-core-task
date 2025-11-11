package main

import (
	"reflect"
	"testing"
)

func TestAdd(t *testing.T) {
	testValue := 123
	myMap := New()
	myMap.Add("test1", testValue)
	value, ok := myMap.Get("test1")
	if !ok || value != testValue {
		t.Errorf("test add value failure")
	}
}

func TestRemove(t *testing.T) {
	myMap := New()
	myMap.Add("test1", 123)
	myMap.Remove("test1")
	ok := myMap.Exist("test1")
	if ok {
		t.Errorf("test remove value failure")
	}
}

func TestCopy(t *testing.T) {
	myMap := New()
	myMap.Add("test1", 123)
	myMap.Add("test2", 321)
	myMap.Add("test3", 221)

	exp := map[string]int{
		"test1": 123,
		"test2": 321,
		"test3": 221,
	}
	copy := myMap.Copy()
	if !reflect.DeepEqual(exp, copy) {
		t.Errorf("expected: %v, got: %v", exp, copy)
	}
}

func TestExist(t *testing.T) {
	myMap := New()
	myMap.Add("test1", 123)
	okFalse := myMap.Exist("test2")
	if okFalse {
		t.Errorf("test exist value failure")
	}
	okTrue := myMap.Exist("test1")
	if !okTrue {
		t.Errorf("test exist value failure")
	}
}

func TestGet(t *testing.T) {
	testValue := 123
	myMap := New()
	myMap.Add("test1", testValue)
	value, ok := myMap.Get("test1")
	if value != testValue || !ok {
		t.Errorf("test get failure")
	}
}
