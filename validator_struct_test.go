package govalidator

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

func TestValidator_ValidateStructJSON(t *testing.T) {
	type User struct {
		Name    string `validate:"name|required"`
		Email   string `validate:"email|email|required"`
		Address string `validate:"addresss|between:10,50|required"`
		Age     int    `validate:"age|numeric_between:18,60|required"`
		Zip     string `validate:"zip|digits:5|required"`
	}

	postUser := User{
		Name:    "",
		Email:   "inalid email",
		Address: "invalid",
		Age:     17,
		Zip:     "122",
	}

	var user User

	body, _ := json.Marshal(postUser)
	req, _ := http.NewRequest("POST", "http://www.example.com", bytes.NewReader(body))

	opts := Options{
		Request: req,
		Data:    &user,
	}

	vd := New(opts)
	validationErr := vd.ValidateStructJSON()
	if len(validationErr) != 5 {
		t.Error("ValidateStructJSON failed")
	}

	req2, _ := http.NewRequest("POST", "http://www.example.com", bytes.NewReader([]byte("Invalid json")))

	opts2 := Options{
		Request: req2,
		Data:    &user,
	}

	vd2 := New(opts2)
	validationErr2 := vd2.ValidateStructJSON()
	if len(validationErr2) != 1 {
		t.Error("ValidateStructJSON failed to validate JSON body")
	}

	//check the panic
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("ValidateStructJSON did not panic")
		}
	}()

	opts3 := Options{
		Request: req,
		Data:    user,
	}

	vd3 := New(opts3)
	vd3.ValidateStructJSON()
}

func Benchmark_ValidateStructJSON(b *testing.B) {
	type User struct {
		Name    string `validate:"name|required"`
		Email   string `validate:"email|email|required"`
		Address string `validate:"addresss|between:10,50|required"`
		Age     int    `validate:"age|numeric_between:18,60|required"`
		Zip     string `validate:"zip|digits:5|required"`
	}

	postUser := User{
		Name:    "",
		Email:   "inalid email",
		Address: "invalid",
		Age:     17,
		Zip:     "122",
	}

	var user User

	body, _ := json.Marshal(postUser)
	req, _ := http.NewRequest("POST", "http://www.example.com", bytes.NewReader(body))

	opts := Options{
		Request: req,
		Data:    &user,
	}
	vd := New(opts)
	for n := 0; n < b.N; n++ {
		vd.ValidateStructJSON()
	}
}
