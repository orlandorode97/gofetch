package gofetch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetName(t *testing.T) {
	t.Run("get name", func(t *testing.T) {
		linux := NewLinux()
		name, err := linux.GetName()
		assert.Nil(t, err, "should not return an error")
		assert.NotNil(t, name, "name is not nil")
	})
	t.Run("error after getting name", func(t *testing.T) {

	})
}
