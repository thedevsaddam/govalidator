package govalidator

import (
	"reflect"
	"strings"
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

func Test_isRuleExist(t *testing.T) {
	if !isRuleExist("required") {
		t.Error("isRuleExist failed for valid rule")
	}
	if isRuleExist("not exist") {
		t.Error("isRuleExist failed for invalid rule")
	}
	if !isRuleExist("mime") {
		t.Error("extended rules failed")
	}
}

func Test_toString(t *testing.T) {
	Int := 100
	str := toString(Int)
	typ := reflect.ValueOf(str).Kind()
	if typ != reflect.String {
		t.Error("toString failed!")
	}
}

func Test_isEmpty(t *testing.T) {
	var Int int
	var Int8 int
	var Float32 float32
	var Str string
	var Slice []int
	var e interface{}
	list := map[string]interface{}{
		"_int":             Int,
		"_int8":            Int8,
		"_float32":         Float32,
		"_str":             Str,
		"_slice":           Slice,
		"_empty_interface": e,
	}
	for k, v := range list {
		if !isEmpty(v) {
			t.Errorf("%v failed", k)
		}
	}
}

func Test_getFileInfo(t *testing.T) {
	req, err := buildMocFormReq()
	if err != nil {
		t.Error("request failed", err)
	}
	fExist, fn, ext, mime, size, _ := getFileInfo(req, "file")
	if !fExist {
		t.Error("file does not exist")
	}
	if fn != "BENCHMARK.md" {
		t.Error("failed to get file name")
	}
	if ext != "md" {
		t.Error("failed to get file extension")
	}
	if !strings.Contains(mime, "text/plain") {
		t.Log(mime)
		t.Error("failed to get file mime")
	}
	if size <= 0 {
		t.Error("failed to get file size")
	}
}
