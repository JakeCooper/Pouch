package common

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path"

	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

// NewCockroachCloudStorage returns an instance of S3CloudStorage
func NewCockroachCloudStorage(config *Configuration, db *sql.DB) *CockroachCloudStorage {
	return &CockroachCloudStorage{db, config}
}

// CockroachCloudStorage implements cloud storage interface using S3
type CockroachCloudStorage struct {
	db     *sql.DB
	config *Configuration
}

// Create uploads a new file to remote storage
// path should be relative to POUCHROOT
func (s *CockroachCloudStorage) Create(fp string) error {
	p := path.Join(s.config.PouchRoot, fp)
	fmt.Println(p)
	stat, err := os.Stat(p)
	if os.IsNotExist(err) {
		return err
	}

	f, err := os.Open(p)
	if err != nil {
		return errors.Wrap(err, "could not open file to be uploaded")
	}
	defer f.Close()

	create := "INSERT INTO files (path, size) VALUES ($1, $2)"
	s.db.Exec(create, fp, stat.Size())
	return nil
}

// Get s a file from remote store and writes it into the supplied path
// The path should be relative to POUCHROOT
func (s *CockroachCloudStorage) Get(fp string) error {
	p := path.Join(s.config.PouchRoot, fp)
	f, err := os.Create(p)
	if err != nil {
		return errors.Wrapf(err, "could not create file %s", p)
	}
	get := "SELECT from files where path = $1"
	_, err = s.db.Exec(get, fp)
	if err != nil {
		log.Fatal(err)
	}
	_, err = os.Open(fp)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	return nil
}

// Update a file that exists remote, path is local path to updated file
func (s *CockroachCloudStorage) Update(path string) error {
	return s.Create(path)
}

// Delete a file on remote fs
func (s *CockroachCloudStorage) Delete(path string) error {
	//return s.bucket.Del(path)
	return nil
}

// List all files on remote fs
func (s *CockroachCloudStorage) List() ([]string, error) {
	return nil, nil
}
