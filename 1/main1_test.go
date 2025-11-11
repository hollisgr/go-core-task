package main

import (
	"reflect"
	"testing"
)

func TestCreateData(t *testing.T) {
	expData := data{
		Decimal: 27,
		Octal:   0154,
		Hex:     0x289,
		Float:   0.087,
		String:  "hello",
		Bool:    true,
		Complex: 3 + 4i,
	}

	testData := CreateData()
	if !reflect.DeepEqual(expData, testData) {
		t.Errorf("expected: %v, got: %v\n", expData, testData)
	}
}

func TestToString(t *testing.T) {
	testData := CreateData()
	testDataStr := testData.ToString()
	expStr := "27 108 649 0.087 hello true (3+4i)"
	if testDataStr != expStr {
		t.Errorf("expected: %v, got: %v\n", expStr, testData)
	}
}

func TestHash(t *testing.T) {
	expHash := "5e42fc16043951025a8920c2acf4d4a5a960968bdb43346c028c155affad92e4"
	src := "testtest"
	salt := "salt"
	hashStr := Hash(src, salt)
	if expHash != hashStr {
		t.Errorf("expected: %v, got: %v\n", expHash, hashStr)
	}
}
