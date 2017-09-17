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

		var ret common.Wrapper

		var objects []common.Metadata

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

			objects = append(objects, tmp)

		}

		ret.StatusCode = 200
		ret.Headers.ContentType = "application/json"
		tmp, _ := json.Marshal(objects)
		ret.Body = string(tmp)

		return ret, nil
	})
}
