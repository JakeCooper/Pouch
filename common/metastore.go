package common

// MetadataStore is an interface for storing metadata about files
type MetadataStore interface {
	// InitialIndex runs on first config of pouch and generates & stores metadata
	// for all files already in the pouch root
	InitialIndex(config *Configuration) error
	AddFile(path string) error
	Update(path string, meta *Metadata) error
	Remove(path string) error
}
