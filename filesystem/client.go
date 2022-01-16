package filesystem

import "github.com/aws/aws-sdk-go/service/s3"


type ClientInterface interface {
	Get(path string) (*s3.GetObjectOutput, error)
	DownloadModel(model string, s3ModelPath string) error
}
