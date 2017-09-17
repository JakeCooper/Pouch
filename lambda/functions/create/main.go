package main

import (
	"encoding/json"
	"os"

	"errors"
	"github.com/JakeCooper/Pouch/common"
	"github.com/apex/go-apex"
	"github.com/goamz/goamz/s3"
)

func main() {
	apex.HandleFunc(func(event json.RawMessage, ctx *apex.Context) (interface{}, error) {
		var m common.FileContext

		if err := json.Unmarshal(event, &m); err != nil {
			return nil, errors.New("Error while marshaling JSON: " + err.Error())
		}

		bucketName := os.Getenv("S3Root")

		bucket, err := common.GetS3Bucket(bucketName)
		if err != nil {
			return nil, errors.New("Error while getting S3 Bucket: " + err.Error())
		}

		err = bucket.Put(m.FileName, m.RawData, "text", s3.BucketOwnerFull, s3.Options{})
		if err != nil {
			return nil, errors.New("Error while putting into bucket:" + err.Error())
		}

		return m, nil
	})
}
