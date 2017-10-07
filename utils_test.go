package govalidator

import (
	"reflect"
	"testing"
)

func Test_isContainRequiredField(t *testing.T) {
	if !isContainRequiredField([]string{"required", "email"}) {
		t.Error("isContainRequiredField failed!")
	}

	if isContainRequiredField([]string{"numeric", "min:5"}) {
		t.Error("isContainRequiredField failed!")
	}
}

func Benchmark_isContainRequiredField(b *testing.B) {
	for n := 0; n < b.N; n++ {
		isContainRequiredField([]string{"required", "email"})
	}
}

type person struct{}

func (person) Details() string {
	return "John Doe"
}

func (person) Age(age string) string {
	return "Age: " + age
}

func Test_isRuleExist(t *testing.T) {
	if !isRuleExist("required") {
		t.Error("isRuleExist failed for valid rule")
	}
	if isRuleExist("not exist") {
		t.Error("isRuleExist failed for invalid rule")
	}
}

func Test_toString(t *testing.T) {
	var Int int
	Int = 100
	str := toString(Int)
	typ := reflect.ValueOf(str).Kind()
	if typ != reflect.String {
		t.Error("toString failed!")
	}
}
