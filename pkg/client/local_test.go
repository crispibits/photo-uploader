package client

import (
	"fmt"
	"os"
	"path"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDoesNotCreateDirectory(t *testing.T) {
	target := path.Join(os.TempDir(), fmt.Sprintf("%d", time.Now().Unix()))
	var client Client
	client, cerr := NewLocalClient(target, false)
	_, err := os.Stat(target)
	assert.True(t, os.IsNotExist(err))
	assert.NotNil(t, cerr)
	assert.NotNil(t, client)
}

func TestCreatesDirectory(t *testing.T) {
	target := path.Join(os.TempDir(), fmt.Sprintf("%d", time.Now().Unix()))
	var client Client
	client, cerr := NewLocalClient(target, true)
	_, err := os.Stat(target)
	assert.False(t, os.IsNotExist(err))
	assert.Nil(t, cerr)
	assert.NotNil(t, client)
}

func TestUploadsFile(t *testing.T) {
	target := path.Join(os.TempDir(), fmt.Sprintf("%d", time.Now().Unix()))
	var client Client
	client, _ = NewLocalClient(target, true)
	err := client.Upload([]byte("hello world"), "testfile.txt")
	assert.Nil(t, err)
	_, cerr := os.Stat(path.Join(target, "testfile.txt"))
	assert.False(t, os.IsNotExist(cerr))
}
