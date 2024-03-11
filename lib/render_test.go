package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolvedacRender(t *testing.T) {
	res, err := NewBoxClient("").FetchUserData("abiriadev")
	assert := assert.New(t)

	assert.Nil(err)

	rendered, err := res.Render()
	assert.Nil(err)

	t.Log(rendered)
}
