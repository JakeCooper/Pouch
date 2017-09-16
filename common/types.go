package common

// CloudStorage : Interface for CloudStorage
type CloudStorage interface {
	// Create uploads a new file to remote storage
	Create(path string) error
	// Get s a file from remote store and writes it into the supplied path
	Get(path string) error
	// Update a file that exists remote, path is local path to updated file
	Update(path string) error
	// Delete a file on remote fs
	Delete(path string) error
	// List all files on remote fs
	List() ([]string, error)
}

type CreateRequest struct {
	FileName string `json:fileName`
}
