package gostructs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDecoder(t *testing.T) {
	decoder, err := NewDecoder(&DecoderConfig{})
	if err != nil {
		t.Error(err)
	}
	if decoder == nil {
		t.Error("decoder should not be nil")
	}
}

func TestDecode(t *testing.T) {
	type (
		Server struct {
			Name string `json:"name"`
			Age  int    `json:"age"`
		}
	)
	srv := Server{Name: "test", Age: 30}

	decoder, _ := NewDecoder(&DecoderConfig{ShouldSnakeCase: true})
	result, _ := decoder.Decode(srv)

	assert.Equal(t, result.Name, "server")
	assert.Equal(t, result.Attributes["name"], "test")
	assert.Equal(t, result.Attributes["age"], 30)
}
