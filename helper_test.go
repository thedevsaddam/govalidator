package govalidator

import "testing"

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

func Test_callMethodByName(t *testing.T) {
	p := &person{}
	_, err := callMethodByName(p, "Details")
	if err != nil {
		t.Error("Failed to call valid method")
	}

	_, e := callMethodByName(p, "print")
	if e == nil {
		t.Error("Failed to detect invalid method")
	}

	_, er := callMethodByName(p, "Age", "30")
	if er != nil {
		t.Error("Failed to call method with args")
	}
}

func Test_isRuleExist(t *testing.T) {
	if !isRuleExist("required") {
		t.Error("isRuleExist failed for valid rule")
	}
	if isRuleExist("not exist") {
		t.Error("isRuleExist failed for invalid rule")
	}
}
