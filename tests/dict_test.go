package tests

import (
	"testing"

	"github.com/repong/hardict"
	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	assert.Equal(t, hardict.ExistInvalidWord("封杀"), false)
}
