package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSHA256(t *testing.T) {
	data := SHA256("A random text")
	assert.Equal(t, data, "7331e43726674fab230da59585a1e2e6b7070701f2fa7fd982b8bc42f8265e18")
}
