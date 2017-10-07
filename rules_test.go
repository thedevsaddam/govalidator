package govalidator

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

func Test_AddCustomRule(t *testing.T) {
	AddCustomRule("__x__", func(f string, rule string, message string, v interface{}) error {
		if v.(string) != "xyz" {
			return fmt.Errorf("The %s field must be xyz", f)
		}
		return nil
	})
	if len(rulesFuncMap) <= 0 {
		t.Error("AddCustomRule failed to add new rule")
	}
}

func Test_AddCustomRule_panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("AddCustomRule failed to panic")
		}
	}()
	AddCustomRule("__x__", func(f string, rule string, message string, v interface{}) error {
		if v.(string) != "xyz" {
			return fmt.Errorf("The %s field must be xyz", f)
		}
		return nil
	})
}

func Test_validateExtraRules(t *testing.T) {
	errsBag := url.Values{}
	validateCustomRules("f_field", "__x__", "a", "", errsBag)
	if len(errsBag) != 1 {
		t.Error("validateExtraRules failed")
	}
}

//================================= rules =================================
func Test_Required(t *testing.T) {
	type tRequired struct {
		Str     string  `json:"_str"`
		Int     int     `json:"_int"`
		Int8    int8    `json:"_int8"`
		Int16   int16   `json:"_int16"`
		Int32   int32   `json:"_int32"`
		Int64   int64   `json:"_int64"`
		Uint    uint    `json:"_uint"`
		Uint8   uint8   `json:"_uint8"`
		Uint16  uint16  `json:"_uint16"`
		Uint32  uint32  `json:"_uint32"`
		Uint64  uint64  `json:"_uint64"`
		Uintptr uintptr `json:"_uintptr"`
		Flaot32 float32 `json:"_float32"`
		Flaot64 float64 `json:"_float64"`
	}

	rules := MapData{
		"_str":     []string{"required"},
		"_int":     []string{"required"},
		"_int8":    []string{"required"},
		"_int16":   []string{"required"},
		"_int32":   []string{"required"},
		"_int64":   []string{"required"},
		"_uint":    []string{"required"},
		"_uint8":   []string{"required"},
		"_uint16":  []string{"required"},
		"_uint32":  []string{"required"},
		"_uint64":  []string{"required"},
		"_uintptr": []string{"required"},
		"_float32": []string{"required"},
		"_float64": []string{"required"},
	}

	postRequired := tRequired{}

	var trequired tRequired

	body, _ := json.Marshal(postRequired)
	req, _ := http.NewRequest("POST", "http://www.example.com", bytes.NewReader(body))

	opts := Options{
		Request: req,
		Data:    &trequired,
		Rules:   rules,
	}

	vd := New(opts)
	validationErr := vd.ValidateJSON()
	if len(validationErr) != 14 {
		t.Error("required validation failed!")
	}
}

func Test_Regex(t *testing.T) {
	type tRegex struct {
		Name string `json:"name"`
	}

	postRegex := tRegex{Name: "john"}
	var tregex tRegex

	body, _ := json.Marshal(postRegex)
	req, _ := http.NewRequest("POST", "http://www.example.com", bytes.NewReader(body))

	rules := MapData{
		"name": []string{"regex:^[0-9]+$"},
	}

	opts := Options{
		Request: req,
		Data:    &tregex,
		Rules:   rules,
	}

	vd := New(opts)
	validationErr := vd.ValidateJSON()
	if len(validationErr) != 1 {
		t.Error("regex validation failed!")
	}
}

func Test_Alpha(t *testing.T) {
	type user struct {
		Name string `json:"name"`
	}

	postUser := user{Name: "9080"}
	var userObj user

	body, _ := json.Marshal(postUser)
	req, _ := http.NewRequest("POST", "http://www.example.com", bytes.NewReader(body))

	rules := MapData{
		"name": []string{"alpha"},
	}

	opts := Options{
		Request: req,
		Data:    &userObj,
		Rules:   rules,
	}

	vd := New(opts)
	validationErr := vd.ValidateJSON()
	if len(validationErr) != 1 {
		t.Error("alpha validation failed!")
	}
}

func Test_AlphaDash(t *testing.T) {
	type user struct {
		Name string `json:"name"`
	}

	postUser := user{Name: "9090$"}
	var userObj user

	body, _ := json.Marshal(postUser)
	req, _ := http.NewRequest("POST", "http://www.example.com", bytes.NewReader(body))

	rules := MapData{
		"name": []string{"alpha_dash"},
	}

	opts := Options{
		Request: req,
		Data:    &userObj,
		Rules:   rules,
	}

	vd := New(opts)
	validationErr := vd.ValidateJSON()
	if len(validationErr) != 1 {
		t.Log(validationErr)
		t.Error("alpha_dash validation failed!")
	}
}

func Test_AlphaNumeric(t *testing.T) {
	type user struct {
		Name string `json:"name"`
	}

	postUser := user{Name: "aE*Sb$"}
	var userObj user

	body, _ := json.Marshal(postUser)
	req, _ := http.NewRequest("POST", "http://www.example.com", bytes.NewReader(body))

	rules := MapData{
		"name": []string{"alpha_num"},
	}

	opts := Options{
		Request: req,
		Data:    &userObj,
		Rules:   rules,
	}

	vd := New(opts)
	validationErr := vd.ValidateJSON()
	if len(validationErr) != 1 {
		t.Log(validationErr)
		t.Error("alpha_num validation failed!")
	}
}

func Test_Boolean(t *testing.T) {
	type Bools struct {
		BoolStr     string  `json:"boolStr"`
		BoolInt     int     `json:"boolInt"`
		BoolInt8    int8    `json:"boolInt8"`
		BoolInt16   int16   `json:"boolInt16"`
		BoolInt32   int32   `json:"boolInt32"`
		BoolInt64   int64   `json:"boolInt64"`
		BoolUint    uint    `json:"boolUint"`
		BoolUint8   uint8   `json:"boolUint8"`
		BoolUint16  uint16  `json:"boolUint16"`
		BoolUint32  uint32  `json:"boolUint32"`
		BoolUint64  uint64  `json:"boolUint64"`
		BoolUintptr uintptr `json:"boolUintptr"`
		Bool        bool    `json:"_bool"`
	}

	postBools := Bools{BoolStr: "abc", BoolInt: 90}
	var boolObj Bools

	body, _ := json.Marshal(postBools)
	req, _ := http.NewRequest("POST", "http://www.example.com", bytes.NewReader(body))

	rules := MapData{
		"_bools":      []string{"bool"},
		"boolStr":     []string{"bool"},
		"boolInt":     []string{"bool"},
		"boolInt8":    []string{"bool"},
		"boolInt16":   []string{"bool"},
		"boolInt32":   []string{"bool"},
		"boolInt64":   []string{"bool"},
		"boolUint":    []string{"bool"},
		"boolUint8":   []string{"bool"},
		"boolUint16":  []string{"bool"},
		"boolUint32":  []string{"bool"},
		"boolUint64":  []string{"bool"},
		"boolUintptr": []string{"bool"},
	}

	opts := Options{
		Request: req,
		Data:    &boolObj,
		Rules:   rules,
	}

	vd := New(opts)
	validationErr := vd.ValidateJSON()
	if len(validationErr) != 2 {
		t.Log(validationErr)
		t.Error("bool validation failed!")
	}
}

func Test_Between(t *testing.T) {
	type user struct {
		Str     string  `json:"str"`
		Int     int     `json:"_int"`
		Int8    int8    `json:"_int8"`
		Int16   int16   `json:"_int16"`
		Int32   int32   `json:"_int32"`
		Int64   int64   `json:"_int64"`
		Uint    uint    `json:"_uint"`
		Uint8   uint8   `json:"_uint8"`
		Uint16  uint16  `json:"_uint16"`
		Uint32  uint32  `json:"_uint32"`
		Uint64  uint64  `json:"_uint64"`
		Uintptr uintptr `json:"_uintptr"`
		Float32 float32 `json:"_float32"`
		Float64 float64 `json:"_float64"`
		Slice   []int   `json:"_slice"`
	}

	postUser := user{}
	var userObj user

	body, _ := json.Marshal(postUser)
	req, _ := http.NewRequest("POST", "http://www.example.com", bytes.NewReader(body))

	rules := MapData{
		"str":      []string{"between:3,5"},
		"_int":     []string{"between:3,5"},
		"_int8":    []string{"between:3,5"},
		"_int16":   []string{"between:3,5"},
		"_int32":   []string{"between:3,5"},
		"_int64":   []string{"between:3,5"},
		"_uint":    []string{"between:3,5"},
		"_uint8":   []string{"between:3,5"},
		"_uint16":  []string{"between:3,5"},
		"_uint32":  []string{"between:3,5"},
		"_uint64":  []string{"between:3,5"},
		"_uintptr": []string{"between:3,5"},
		"_float32": []string{"between:3.5,5.9"},
		"_float64": []string{"between:3.3,6.2"},
		"_slice":   []string{"between:3,5"},
	}

	opts := Options{
		Request: req,
		Data:    &userObj,
		Rules:   rules,
	}

	vd := New(opts)
	validationErr := vd.ValidateJSON()
	if len(validationErr) != 15 {
		t.Error("between validation failed!")
	}
}

func Test_CreditCard(t *testing.T) {
	type user struct {
		CreditCard string `json:"credit_card"`
	}

	postUser := user{CreditCard: "87080"}
	var userObj user

	body, _ := json.Marshal(postUser)
	req, _ := http.NewRequest("POST", "http://www.example.com", bytes.NewReader(body))

	rules := MapData{
		"credit_card": []string{"credit_card"},
	}

	opts := Options{
		Request: req,
		Data:    &userObj,
		Rules:   rules,
	}

	vd := New(opts)
	validationErr := vd.ValidateJSON()
	if len(validationErr) != 1 {
		t.Error("Credit card validation failed!")
	}
}

func Test_Coordinate(t *testing.T) {
	type user struct {
		Coordinate string `json:"coordinate"`
	}

	postUser := user{Coordinate: "8080"}
	var userObj user

	body, _ := json.Marshal(postUser)
	req, _ := http.NewRequest("POST", "http://www.example.com", bytes.NewReader(body))

	rules := MapData{
		"coordinate": []string{"coordinate"},
	}

	opts := Options{
		Request: req,
		Data:    &userObj,
		Rules:   rules,
	}

	vd := New(opts)
	validationErr := vd.ValidateJSON()
	if len(validationErr) != 1 {
		t.Error("Coordinate validation failed!")
	}
}

func Test_CSSColor(t *testing.T) {
	type user struct {
		Color string `json:"color"`
	}

	postUser := user{Color: "8080"}
	var userObj user

	body, _ := json.Marshal(postUser)
	req, _ := http.NewRequest("POST", "http://www.example.com", bytes.NewReader(body))

	rules := MapData{
		"color": []string{"css_color"},
	}

	opts := Options{
		Request: req,
		Data:    &userObj,
		Rules:   rules,
	}

	vd := New(opts)
	validationErr := vd.ValidateJSON()
	if len(validationErr) != 1 {
		t.Error("CSS color validation failed!")
	}
}

func Test_Digits(t *testing.T) {
	type user struct {
		Zip   string `json:"zip"`
		Level string `json:"level"`
	}

	postUser := user{Zip: "8322", Level: "10"}
	var userObj user

	body, _ := json.Marshal(postUser)
	req, _ := http.NewRequest("POST", "http://www.example.com", bytes.NewReader(body))

	rules := MapData{
		"zip":   []string{"digits:5"},
		"level": []string{"digits:1"},
	}

	opts := Options{
		Request: req,
		Data:    &userObj,
		Rules:   rules,
	}

	vd := New(opts)
	validationErr := vd.ValidateJSON()
	if len(validationErr) != 2 {
		t.Error("Digits validation failed!")
	}
}
