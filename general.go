package util

import (
	"encoding/json"
	"regexp"
)

func NewTrue() *bool {
	b := true
	return &b
}

func NewFalse() *bool {
	b := false
	return &b
}

var nonAlphanumericRegex = regexp.MustCompile(`[^a-zA-Z0-9 ]+`)

func ClearString(str string) string {
	return nonAlphanumericRegex.ReplaceAllString(str, "")
}

func MapToStruct[T any](source any) (*T, error) {
	var elem T
	marshStr, marshErr := json.Marshal(source)
	if marshErr != nil {
		return nil, marshErr
	}
	unMarshErr := json.Unmarshal(marshStr, &elem)
	if unMarshErr != nil {
		return nil, unMarshErr
	}
	return &elem, nil
}
