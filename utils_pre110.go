// +build !go1.10

package govalidator

import (
	"io"
	"net/http"
	"path/filepath"
	"strings"
)

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
