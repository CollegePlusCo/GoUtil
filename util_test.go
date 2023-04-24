package util

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMapToStruct(t *testing.T) {
	type X struct {
		Money string `json:"money"`
		B     *bool  `json:"b"`
	}
	test := map[string]any{
		"money": "money",
		"b":     true,
	}
	s, err := MapToStruct[X](test)
	fmt.Println(*s.B)
	assert.Nil(t, err)
}
