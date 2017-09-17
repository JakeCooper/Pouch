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

type FileContext struct {
	FileName string `json:"filename"`
	RawData  []byte `json:"raw_data"`
}

type Head struct {
	ContentType string `json:"Content-Type"`
}

type Wrap2 struct {
	Body string `json:"body"`
}

type Wrapper struct {
	StatusCode int    `json:"statusCode"`
	Headers    Head   `json:"headers"`
	Body       string `json:"body"`
}

type MetadataResponse struct {
	Objects []Metadata `json:"objects"`
}
