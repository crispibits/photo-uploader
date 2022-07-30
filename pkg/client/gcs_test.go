package client

import (
	"os"
	"testing"

	"github.com/crispibits/photo-uploader/pkg/config"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func TestCreateConfig(t *testing.T) {
	creds, err := os.ReadFile("config.yaml")
	assert.Nil(t, err)
	gcsConfig := &config.GCS{}
	assert.Nil(t, yaml.Unmarshal(creds, gcsConfig))
	cfg := &config.Config{GCS: *gcsConfig}
	var client Client
	client, err = NewGCSClient(cfg)
	assert.Nil(t, err)
	assert.NotNil(t, client)
}

func TestUpload(t *testing.T) {
	creds, err := os.ReadFile("config.yaml")
	assert.Nil(t, err)
	gcsConfig := &config.GCS{}
	assert.Nil(t, yaml.Unmarshal(creds, gcsConfig))
	cfg := &config.Config{GCS: *gcsConfig}
	var client Client
	client, err = NewGCSClient(cfg)
	assert.Nil(t, err)
	assert.NotNil(t, client)
	b := []byte("hello world")
	err = client.Upload(b, "foobar/helloworld.txt")
	assert.Nil(t, err)
}
