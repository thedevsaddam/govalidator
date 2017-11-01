package govalidator

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"testing"
)

// buildMocFormReq prepare a moc form data request with a test file
func buildMocFormReq() (*http.Request, error) {
	fPath := "doc/BENCHMARK.md"
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	file, err := os.Open(fPath)
	if err != nil {
		return nil, err
	}
	part, err := writer.CreateFormFile("file", filepath.Base(fPath))
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)
	file.Close()
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", "www.example.com", body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	return req, nil
}

func Test_validateFiles(t *testing.T) {
	req, err := buildMocFormReq()
	if err != nil {
		t.Error("request failed", err)
	}
	rules := MapData{
		"file": []string{"ext:jpg,pdf", "size:10", "mime:application/pdf", "required"},
	}

	opts := Options{
		Request: req,
		Rules:   rules,
	}

	vd := New(opts)
	validationErr := vd.Validate()
	if len(validationErr) != 1 {
		t.Error("file validation failed!")
	}
}

func Test_validateFiles_message(t *testing.T) {
	req, err := buildMocFormReq()
	if err != nil {
		t.Error("request failed", err)
	}
	rules := MapData{
		"file": []string{"ext:jpg,pdf", "size:10", "mime:application/pdf", "required"},
	}

	msgs := MapData{
		"file": []string{"ext:custom_message"},
	}

	opts := Options{
		Request:  req,
		Rules:    rules,
		Messages: msgs,
	}

	vd := New(opts)
	validationErr := vd.Validate()
	if len(validationErr) != 1 {
		t.Error("file validation failed!")
	}
	if validationErr.Get("file") != "custom_message" {
		t.Log(validationErr)
		t.Error("failed custom message for file validation")
	}
}
