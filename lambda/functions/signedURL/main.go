package main

import (
	"encoding/json"

	"errors"
	"github.com/JakeCooper/Pouch/common"
	"github.com/apex/go-apex"
	"log"
	"os"
	"time"
)

func main() {
	apex.HandleFunc(func(event json.RawMessage, ctx *apex.Context) (interface{}, error) {
		bucketName := os.Getenv("S3Root")

		bucket, err := common.GetS3Bucket(bucketName)
		if err != nil {
			return nil, errors.New("Error while getting S3 Bucket: " + err.Error())
		}

		var wrap common.Wrap2

		err = json.Unmarshal(event, &wrap)
		if err != nil {
			return nil, errors.New("Error while unmarshalling data: " + err.Error())
		}
		var file common.FileContext

		log.Println(wrap)
		err = json.Unmarshal([]byte(wrap.Body), &file)
		if err != nil {
			return nil, errors.New("Error while unmarshalling data: " + err.Error())
		}

		var ret common.Wrapper

		log.Println(file.FileName)

		ret.StatusCode = 200
		ret.Headers.ContentType = "application/json"
		m := map[string]string{
			"event":    string(event),
			"filename": file.FileName,
			"url":      bucket.SignedURL(file.FileName, time.Now().Add(time.Second*60*3))}

		s, err := json.Marshal(m)
		if err != nil {
			return nil, errors.New("Error while unmarshalling data: " + err.Error())
		}
		ret.Body = string(s)

		return ret, nil
	})
}
