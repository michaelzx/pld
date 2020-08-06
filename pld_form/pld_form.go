package pld_form

import (
	"mime/multipart"
	"strconv"
)

func GetSingleFile(form *multipart.Form, name string) *multipart.FileHeader {
	files := form.File[name]
	if len(files) != 1 {
		return nil
	}
	return files[0]
}

func GetSingleString(form *multipart.Form, name string) string {
	values := form.Value[name]
	if len(values) != 1 {
		return ""
	}
	return values[0]
}

func GetSingleInt64(form *multipart.Form, name string, defaultValue int64) int64 {
	values := form.Value[name]
	if len(values) != 1 {
		return defaultValue
	}
	i64, err := strconv.ParseInt(values[0], 10, 64)
	if err != nil {
		return defaultValue
	}
	return i64
}
