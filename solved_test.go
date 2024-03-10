package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolvedacApi(t *testing.T) {
	res, err := fetchUserData("abiriadev")
	assert := assert.New(t)

	assert.Nil(err)
	assert.Equal(res.Handle, "abiriadev")
}
