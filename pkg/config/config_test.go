package config

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateConfig(t *testing.T) {
	f, err := os.CreateTemp("", "")
	defer os.Remove(f.Name())
	if err != nil {
		panic(err)
	}
	if os.Remove(f.Name()) != nil {
		panic(err)
	}
	cfg, err := ReadOrCreate(f.Name())
	assert.Nil(t, err)
	assert.NotNil(t, cfg)
}

func TestReadLocalConfig(t *testing.T) {
	f, err := os.CreateTemp("", "")
	defer os.Remove(f.Name())
	if err != nil {
		panic(err)
	}
	expectedCfg := &Config{
		GCS: GCS{
			Bucket: "example",
			Creds:  "creds",
		},
	}
	b, err := json.Marshal(expectedCfg)
	if err != nil {
		panic(err)
	}
	if os.WriteFile(f.Name(), b, 0600) != nil {
		panic(err)
	}
	cfg, err := ReadOrCreate(f.Name())
	assert.Nil(t, err)
	assert.Equal(t, cfg.GCS.Bucket, expectedCfg.GCS.Bucket)
}
