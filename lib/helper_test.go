package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTierMap(t *testing.T) {
	assert := assert.New(t)

	assert.Equal("Ruby III", tierMap[28])
}

func TestTierPercentageMap(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(1250, tierPercentageMap[14])
}
