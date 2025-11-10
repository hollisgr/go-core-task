package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"reflect"
	"strings"
)

type task1 struct {
	d data
	r reflect.Type
}

type data struct {
	Decimal int
	Octal   int
	Hex     int
	Float   float32
	String  string
	Bool    bool
	Complex complex64
}

func main() {
	data := CreateData()
	data.Print()

	dataStr := data.ToString()
	fmt.Println(dataStr)

	hashedStr := Hash(dataStr, "go-2024 ")
	fmt.Println(hashedStr)
}

func CreateData() data {
	res := data{
		Decimal: 27,
		Octal:   0154,  // 108
		Hex:     0x289, // 649
		Float:   0.087,
		String:  "hello",
		Bool:    true,
		Complex: 3 + 4i,
	}
	return res
}

func (d *data) Print() {
	// рефлектим типы полей в нашей структуре
	reflectedData := reflect.TypeOf(*d)

	for i := range reflectedData.NumField() {
		field := reflectedData.Field(i)
		fmt.Printf("Field name: %s, type: %s\n", field.Name, field.Type)
	}
}

func (d *data) ToString() string {
	var str string

	// рефлектим значения полей в нашей структуре
	reflectedData := reflect.ValueOf(*d)

	for i := range reflectedData.NumField() {
		value := reflectedData.Field(i).Interface()
		str += fmt.Sprintf("%v ", value)
	}
	res := strings.TrimRight(str, " ")
	return res
}

func Hash(src string, salt string) string {
	// переводим переданные строки в массивы рун
	runes := []rune(src)
	saltRunes := []rune(salt)

	// определяем середину среза и подготавливаем два массива рун, между которыми будем вставлять соль
	mid := len(runes) / 2
	left := runes[:mid]
	right := runes[mid:]

	// соединяем 3 массива рун в один новый массив с солью посередине
	var newRunes []rune
	newRunes = append(newRunes, left...)
	newRunes = append(newRunes, saltRunes...)
	newRunes = append(newRunes, right...)

	// создаем хэш сумму из полученного массива рун, т.к. метод write ожидает массив байтов,
	// (а мы не можем привести массив рун к массиву байтов) приходится промежуточно приводиться к строке
	hasher := sha256.New()
	hasher.Write([]byte(string(newRunes)))
	hashedBytes := hasher.Sum(nil)

	return hex.EncodeToString(hashedBytes)
}
