package dto

import "io"

type PutObjectInput struct {
	Bucket string    `json:"bucket"`
	Key    string    `json:"key"`
	LocalFilePath string `json:"localFilePath"`
	Body   io.Reader `json:"body"`
}
