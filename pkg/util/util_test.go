package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSha1(t *testing.T) {
	sha := Sha1([]byte(`teststring`))
	assert.Equal(t, sha, "b8473b86d4c2072ca9b08bd28e373e8253e865c4")
}

func TestMd5(t *testing.T) {
	md5 := Md5([]byte(`teststring`))
	assert.Equal(t, md5, "d67c5cbf5b01c9f91932e3b8def5e5f8")
}
