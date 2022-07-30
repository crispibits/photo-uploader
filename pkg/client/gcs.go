package client

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"time"

	"cloud.google.com/go/storage"
	"github.com/crispibits/photo-uploader/pkg/config"
	"github.com/crispibits/photo-uploader/pkg/util"
	"google.golang.org/api/option"
)

type GCSClient struct {
	Client
	config    *config.Config
	GcsClient *storage.Client
	Bucket    *storage.BucketHandle
}

type PermissionError struct {
	err      string
	Required []string
	Allowed  []string
}

func (e *PermissionError) Error() string {
	return e.err
}

func NewGCSClient(cfg *config.Config) (Client, error) {
	c := &GCSClient{}
	c.config = cfg
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()
	opt := option.WithCredentialsJSON([]byte(cfg.GCS.Creds))
	gcsClient, err := storage.NewClient(ctx, opt)
	if err != nil {
		return c, err
	}
	c.Bucket = gcsClient.Bucket(cfg.GCS.Bucket)
	var permissions []string
	permissions = append(permissions, "storage.buckets.get")
	permissions = append(permissions, "storage.objects.list")
	permissions = append(permissions, "storage.objects.get")
	permissions = append(permissions, "storage.objects.create")
	permissions = append(permissions, "storage.objects.delete")
	permissions = append(permissions, "storage.objects.update")
	allowed, err := c.Bucket.IAM().TestPermissions(ctx, permissions)
	if err != nil {
		fmt.Println(err)
		return c, err
	}
	if len(permissions) != len(allowed) {
		perr := &PermissionError{}
		perr.err = fmt.Sprintf("Insufficient Permissions, [%d] of [%d]", len(allowed), len(permissions))
		perr.Allowed = allowed
		perr.Required = permissions
		//return errors.New("Insufficient permissions")
		return c, perr
	}
	// Check that we have access to this bucket
	_, err = c.Bucket.Attrs(ctx)
	if err != nil {
		return c, err
	}
	c.GcsClient = gcsClient
	//cfg.Client = client
	return c, err
}

func (gcs *GCSClient) Upload(content []byte, name string) error {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()
	o := gcs.Bucket.Object(name)
	wc := o.NewWriter(ctx)
	wc.ChunkSize = 0 // note retries are not supported for chunk size 0.
	if _, err := io.Copy(wc, bytes.NewReader(content)); err != nil {
		return fmt.Errorf("io.Copy: %v", err)
	}
	// Data can continue to be added to the file until the writer is closed.
	if err := wc.Close(); err != nil {
		return fmt.Errorf("Writer.Close: %v", err)
	}

	// Update the object to set the metadata.
	objectAttrsToUpdate := storage.ObjectAttrsToUpdate{
		Metadata: map[string]string{
			"keyToAddOrUpdate": "value",
		},
	}
	if _, err := o.Update(ctx, objectAttrsToUpdate); err != nil {
		return fmt.Errorf("ObjectHandle(%q).Update: %v", name, err)
	}

	s1, err := getSha1(gcs, name)
	if err != nil {
		return fmt.Errorf("sha1 read: %v", err)
	}
	s2 := util.Sha1(content)

	if s1 != s2 {
		return fmt.Errorf("integrity check failed, %s vs %s", s1, s2)
	}

	fmt.Printf("%v uploaded to %v.\n", name, "wibble")
	return nil
}

func (gcs *GCSClient) Config() *config.Config {
	return gcs.config
}

func getSha1(gcs *GCSClient, name string) (string, error) {
	var backBytes []byte
	var err error
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()
	verify, err := gcs.Bucket.Object(name).NewReader(ctx)
	if err != nil {
		return "", err
	}
	if backBytes, err = ioutil.ReadAll(verify); err != nil {
		return "", fmt.Errorf("verify.Read: %v", err)
	}
	return util.Sha1(backBytes), err
}

func (gcs *GCSClient) Delete(name string) error {
	return nil
}
