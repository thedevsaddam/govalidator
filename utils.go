package govalidator

import (
	"fmt"
	"io"
	"net/http"
	"path/filepath"
	"reflect"
	"strings"
)

// containsRequiredField check rules contain any required field
func isContainRequiredField(rules []string) bool {
	for _, rule := range rules {
		if rule == "required" {
			return true
		}
	}
	return false
}

// isRuleExist check if the provided rule name is exist or not
func isRuleExist(rule string) bool {
	if strings.Contains(rule, ":") {
		rule = strings.Split(rule, ":")[0]
	}
	extendedRules := []string{"size", "mime", "ext"}
	for _, r := range extendedRules {
		if r == rule {
			return true
		}
	}
	if _, ok := rulesFuncMap[rule]; ok {
		return true
	}
	return false
}

// toString force data to be string
func toString(v interface{}) string {
	str, ok := v.(string)
	if !ok {
		str = fmt.Sprintf("%#v", v)
	}
	return str
}

// isEmpty check a type is Zero
func isEmpty(x interface{}) bool {
	rt := reflect.TypeOf(x)
	if rt == nil {
		return true
	}
	rv := reflect.ValueOf(x)
	switch rv.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice:
		return rv.Len() == 0
	}
	return reflect.DeepEqual(x, reflect.Zero(rt).Interface())
}

// Sizer interface
type Sizer interface {
	Size() int64
}

// getFileInfo read file from request and return file name, extension, mime and size
func getFileInfo(r *http.Request, field string) (bool, string, string, string, int64, error) {
	file, multipartFileHeader, err := r.FormFile(field)
	if err != nil {
		return false, "", "", "", 0, err
	}
	// Create a buffer to store the header of the file in
	fileHeader := make([]byte, 512)

	// Copy the headers into the FileHeader buffer
	if _, err := file.Read(fileHeader); err != nil {
		if err != io.EOF {
			return false, "", "", "", 0, err
		}
	}

	// set position back to start.
	if _, err := file.Seek(0, 0); err != nil {
		return false, "", "", "", 0, err
	}

	mime := http.DetectContentType(fileHeader)
	if subs := "; charset=utf-8"; strings.Contains(mime, subs) {
		mime = strings.Replace(mime, subs, "", -1)
	}
	if subs := ";charset=utf-8"; strings.Contains(mime, subs) {
		mime = strings.Replace(mime, subs, "", -1)
	}
	if subs := "; charset=UTF-8"; strings.Contains(mime, subs) {
		mime = strings.Replace(mime, subs, "", -1)
	}
	if subs := ";charset=UTF-8"; strings.Contains(mime, subs) {
		mime = strings.Replace(mime, subs, "", -1)
	}
	fExist := false
	if file != nil {
		fExist = true
	}
	return fExist, multipartFileHeader.Filename,
		strings.TrimPrefix(filepath.Ext(multipartFileHeader.Filename), "."),
		strings.TrimSpace(mime),
		file.(Sizer).Size(),
		nil
}
