package src

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_it_works(t *testing.T) {
	assert.NotNil(t, NewGame())
}
