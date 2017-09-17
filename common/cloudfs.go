package common

// WatcherObj : For adding/removing watchers on files
type WatcherObj struct {
	Path   string
	Action string
}

// InitCloudFS returns an S3 interface
func InitCloudFS(config *Configuration) (CloudStorage, error) {
	return nil, nil
}
