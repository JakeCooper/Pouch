package common

// Location of the time on disk
type Location struct {
	AbsPath string
	RelPath string
}

// Metadata is data about at file, used to sync files between locations
type Metadata struct {
	Location
	FileName   string `json:"filename"`
	MerkleRoot string `json:"merkle_root"`
	PartSize   int    `json:"part_size"`
	NParts     int    `json:"n_parts"`
	ModifiedAt int64  `json:"modified_at"` // the most recent time in the PartHeaders
	ReceivedAt int64  `json:"received_at"`
}
