package filesystem

import (
	"github.com/aws/aws-sdk-go/service/s3"
	"strings"
)

type FileSystemClientTypeMap map[string]FileSystemClientType
type FileSystemClientType string

const S3 = FileSystemClientType("s3")
const UnKnown = FileSystemClientType("unknown")

var typeMap = FileSystemClientTypeMap{
	"s3": S3,
}

type ClientInterface interface {
	Get(path string) (*s3.GetObjectOutput, error)
	DownloadModel(model string, s3ModelPath string) error
}

type ClientDetail struct {
	Bucket string
	Region string
}

func GetFileSystemClientType(types string) FileSystemClientType {
	typeNormal := strings.ToLower(types)
	if v, ok := typeMap[typeNormal]; ok {
		return v
	}
	return UnKnown
}
