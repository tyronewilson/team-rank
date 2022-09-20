package util

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestErrorCollector(t *testing.T) {
	c := NewErrorCollector()
	c.HandleError(fmt.Errorf("test"))
	assert.Len(t, c.ReadErrs().(ErrorList), 1)
	c.HandleError(fmt.Errorf("test2"))
	err := c.ReadErrs()
	assert.Len(t, err.(ErrorList), 2)
}

func TestErrorList(t *testing.T) {
	errs := ErrorList{fmt.Errorf("test"), fmt.Errorf("test2")}
	assert.Equal(t, "test\ntest2", errs.Error())
}
