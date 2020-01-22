package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrue(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("new knocking", "new knocking", "Account owner is new knocking")
}
