package lib

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserRender(t *testing.T) {
	assert := assert.New(t)

	mockUser := User{
		Handle:      "abiriadev",
		Bio:         "hello, world!",
		SolvedCount: 1234,
		VoteCount:   56,
		Tier:        14,
		Rating:      1351,
		Rank:        1000,
	}

	rendered, err := mockUser.Render()
	assert.Nil(err)

	lines := strings.Split(rendered, "\n")

	assert.Equal("hello, world!", lines[1])

	t.Log(rendered)
}
