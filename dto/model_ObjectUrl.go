package dto

type ObjectUrlRequest struct{
	BucketName string `json:"bucketName"`
	ObjectKey string `json:"objectKey"`
	Region string `json:"region"`
}
