package client

import "github.com/crispibits/photo-uploader/pkg/config"

type Client interface {
	Config() *config.Config
	Upload(content []byte, name string) error
	Delete(name string) error
}
