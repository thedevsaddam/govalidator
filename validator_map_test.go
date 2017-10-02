package govalidator

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

func TestValidator_ValidateMapJSON(t *testing.T) {
	type User struct {
		Name    string `json:"name"`
		Email   string `json:"email"`
		Address string `json:"address"`
		Age     int    `json:"age"`
		Zip     string `json:"zip"`
	}

	postUser := User{
		Name:    "",
		Email:   "inalid email",
		Address: "invalid",
		Age:     17,
		Zip:     "122",
	}

	var user map[string]interface{}

	body, _ := json.Marshal(postUser)
	req, _ := http.NewRequest("POST", "http://www.example.com", bytes.NewReader(body))

	rules := MapData{
		"name":    []string{"required"},
		"email":   []string{"email"},
		"address": []string{"between:10,15"},
		"age":     []string{"numeric_between:18,30"},
		"zip":     []string{"digits:10"},
	}

	opts := Options{
		Request: req,
		Data:    &user,
		Rules:   rules,
	}

	vd := New(opts)
	validationErr := vd.ValidateMapJSON()
	if len(validationErr) != 5 {
		t.Error("ValidateMapJSON failed")
	}

	req2, _ := http.NewRequest("POST", "http://www.example.com", bytes.NewReader([]byte("Invalid json")))

	opts2 := Options{
		Request: req2,
		Data:    &user,
		Rules:   rules,
	}

	vd2 := New(opts2)
	validationErr2 := vd2.ValidateMapJSON()
	if len(validationErr2) != 1 {
		t.Error("ValidateMapJSON failed to validate JSON body")
	}

	//check the panic
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("ValidateJSON did not panic")
		}
	}()

	opts3 := Options{
		Request: req,
		Data:    user,
	}

	vd3 := New(opts3)
	vd3.ValidateMapJSON()
}

func Benchmark_ValidateMapJSON(b *testing.B) {
	type User struct {
		Name    string `json:"name"`
		Email   string `json:"email"`
		Address string `json:"address"`
		Age     int    `json:"age"`
		Zip     string `json:"zip"`
	}

	postUser := User{
		Name:    "",
		Email:   "inalid email",
		Address: "invalid",
		Age:     17,
		Zip:     "122",
	}

	var user map[string]interface{}

	body, _ := json.Marshal(postUser)
	req, _ := http.NewRequest("POST", "http://www.example.com", bytes.NewReader(body))

	rules := MapData{
		"name":    []string{"required"},
		"email":   []string{"email"},
		"address": []string{"between:10,15"},
		"age":     []string{"numeric_between:18,30"},
		"zip":     []string{"digits:10"},
	}

	opts := Options{
		Request: req,
		Data:    &user,
		Rules:   rules,
	}

	vd := New(opts)
	for n := 0; n < b.N; n++ {
		vd.ValidateMapJSON()
	}
}
