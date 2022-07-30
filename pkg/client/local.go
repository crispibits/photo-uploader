package client

import (
	"errors"
	"os"
	"path"

	"github.com/crispibits/photo-uploader/pkg/util"
)

type LocalClient struct {
	Client
	target string
}

// Creates a new LocalClient to write to the target directory
func NewLocalClient(target string, create bool) (Client, error) {
	c := &LocalClient{target: target}
	_, err := os.Stat(target)
	if os.IsNotExist(err) && !create {
		return c, err
	}
	if os.MkdirAll(target, 0770) != nil {
		return c, err
	}
	return c, nil
}

func (c *LocalClient) Upload(content []byte, name string) error {
	var err error
	path := path.Join(c.target, name)
	if err = os.WriteFile(path, content, 0660); err != nil {
		return err
	}
	b, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	s1 := util.Sha1(content)
	s2 := util.Sha1(b)
	if s1 != s2 {
		return errors.New("integrity check failed")
	}
	return err
}

func (c *LocalClient) Delete(name string) error {
	return nil
}
