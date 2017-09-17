package common

import (
	"fmt"
	"io"
	"os"
	"path"

	"github.com/goamz/goamz/s3"
	"github.com/pkg/errors"
)

// NewS3CloudStorage returns an instance of S3CloudStorage
func NewS3CloudStorage(config *Configuration, bucket *s3.Bucket) *S3CloudStorage {
	return &S3CloudStorage{bucket, config}
}

// S3CloudStorage implements cloud storage interface using S3
type S3CloudStorage struct {
	bucket *s3.Bucket
	config *Configuration
}

// Create uploads a new file to remote storage
// path should be relative to POUCHROOT
func (s *S3CloudStorage) Create(fp string) error {
	p := path.Join(s.config.PouchRoot, fp)
	fmt.Println(p)
	stat, err := os.Stat(p)
	if os.IsNotExist(err) {
		fmt.Println("WE FAILIN")
		return err
	}

	f, err := os.Open(p)
	if err != nil {
		return errors.Wrap(err, "could not open file to be uploaded")
	}
	defer f.Close()

	opts := s3.Options{}
	err = s.bucket.PutReader(fp, f, stat.Size(), "application/x-binary", s3.BucketOwnerFull, opts)
	if err != nil {
		return errors.Wrap(err, "could not store file")
	}
	return nil
}

// Get s a file from remote store and writes it into the supplied path
// The path should be relative to POUCHROOT
func (s *S3CloudStorage) Get(fp string) error {
	p := path.Join(s.config.PouchRoot, fp)
	f, err := os.Create(p)
	if err != nil {
		return errors.Wrapf(err, "could not create file %s", p)
	}
	defer f.Close()
	rc, err := s.bucket.GetReader(fp)
	if err != nil {
		return errors.Wrapf(err, "could not create read closer for S3 file", fp)
	}
	defer rc.Close()
	_, err = io.Copy(f, rc)
	if err != nil {
		return errors.Wrap(err, "cannto copy bytes")
	}

	stat, err := os.Stat(p)
	if err != nil {
		return err
	}
	fmt.Printf("file downloaded, %d bytes\n", stat.Size())

	return nil
}

// Update a file that exists remote, path is local path to updated file
func (s *S3CloudStorage) Update(path string) error {
	return s.Create(path)
}

// Delete a file on remote fs
func (s *S3CloudStorage) Delete(path string) error {
	return s.bucket.Del(path)
}

// List all files on remote fs
func (s *S3CloudStorage) List() ([]string, error) {
	return nil, nil
}
