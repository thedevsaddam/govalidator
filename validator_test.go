package govalidator

import (
	"net/http"
	"net/url"
	"testing"
)

func TestValidator_SetDefaultRequired(t *testing.T) {
	v := New(Options{})
	v.SetDefaultRequired(true)
	if !v.Opts.RequiredDefault {
		t.Error("SetDefaultRequired failed")
	}
}

func TestValidator_Validate(t *testing.T) {
	var URL *url.URL
	URL, _ = url.Parse("http://www.example.com")
	params := url.Values{}
	params.Add("name", "John Doe")
	params.Add("age", "27")
	params.Add("email", "john@mail.com")
	params.Add("zip", "8233")
	URL.RawQuery = params.Encode()
	r, _ := http.NewRequest("GET", URL.String(), nil)
	rulesList := MapData{
		"name":  []string{"required"},
		"age":   []string{"numeric_between:18,60"},
		"email": []string{"email"},
		"zip":   []string{"digits:4"},
	}

	opts := Options{
		Request: r,
		Rules:   rulesList,
	}
	v := New(opts)
	if err := v.Validate(); len(err) > 0 {
		t.Error("Validate failed to validate correct inputs!")
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Validate did not panic")
		}
	}()

	v1 := New(Options{Rules: MapData{}})
	v1.Validate()
}

func Benchmark_Validate(b *testing.B) {
	var URL *url.URL
	URL, _ = url.Parse("http://www.example.com")
	params := url.Values{}
	params.Add("name", "John Doe")
	params.Add("age", "27")
	params.Add("email", "john@mail.com")
	params.Add("zip", "8233")
	URL.RawQuery = params.Encode()
	r, _ := http.NewRequest("GET", URL.String(), nil)
	rulesList := MapData{
		"name":  []string{"required"},
		"age":   []string{"numeric_between:18,60"},
		"email": []string{"email"},
		"zip":   []string{"digits:4"},
	}

	opts := Options{
		Request: r,
		Rules:   rulesList,
	}
	v := New(opts)
	for n := 0; n < b.N; n++ {
		v.Validate()
	}
}
