package main

import (
	"encoding/json"

	"github.com/JakeCooper/Pouch/common"
	"github.com/apex/go-apex"
	"github.com/goamz/goamz/s3"
)

func main() {
	apex.HandleFunc(func(event json.RawMessage, ctx *apex.Context) (interface{}, error) {
		var m common.CreateRequest

		if err := json.Unmarshal(event, &m); err != nil {
			return nil, err
		}

		s := "fxUCDZihTKtBDbuv"
		bucket, err := common.GetS3Bucket(s)
		if err != nil {
			return nil, err
		}

		err = bucket.Put("test", []byte(s), "text", s3.BucketOwnerFull, s3.Options{})
		if err != nil {
			return nil, err
		}

		return m, nil
	})
}
