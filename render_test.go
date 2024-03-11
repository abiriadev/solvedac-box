package main

import (
	"testing"
	"unicode/utf8"

	"github.com/stretchr/testify/assert"
)

func TestSolvedacRender(t *testing.T) {
	res, err := fetchUserData("cgiosy")
	assert := assert.New(t)

	assert.Nil(err)

	if res, err := NewGistPropFromUser(res); err != nil {
		assert.Nil(err)
	} else {
		rendered, err := res.Render()
		if err != nil {
			assert.Nil(err)
		}

		t.Log(rendered)
		assert.Equal(53, utf8.RuneCountInString(rendered))
	}
}
