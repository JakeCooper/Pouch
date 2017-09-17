package common

// Location of the time on disk
type Location struct {
	AbsPath string
	RelPath string
}

// Metadata is data about at file, used to sync files between locations
type Metadata struct {
	Name       string `json:"name"`
	ObjectType string `json:"objectType"`
	FilePath   string `json:"filePath"`
	Modified   string `json:"modified"`
}
