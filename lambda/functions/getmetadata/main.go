package main

import (
	"encoding/json"

	"errors"
	"github.com/JakeCooper/Pouch/common"
	"github.com/apex/go-apex"
	"os"
	"strings"
)

func main() {
	apex.HandleFunc(func(event json.RawMessage, ctx *apex.Context) (interface{}, error) {
		bucketName := os.Getenv("S3Root")

		bucket, err := common.GetS3Bucket(bucketName)
		if err != nil {
			return nil, errors.New("Error while getting S3 Bucket: " + err.Error())
		}

		m, err := bucket.GetBucketContents()
		if err != nil {
			return nil, errors.New("Error while getting the files: " + err.Error())
		}

		var ret common.MetadataResponse
		for k, v := range *m {
			var tmp common.Metadata
			tmp.FilePath = k
			tmp.Modified = v.LastModified
			tmp.ObjectType = "file"

			if len(k) > 0 && k[len(k)-1] == '/' {
				tmp.ObjectType = "folder"
			}
			tmpNames := strings.Split(strings.TrimRight(k, "/"), "/")
			tmp.Name = tmpNames[len(tmpNames)-1]

			ret.Objects = append(ret.Objects, tmp)

		}

		return ret, nil
	})
}
