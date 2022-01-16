package filesystem

import (
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/hyference/internal/filesystem/s3wrapper"
	"strings"
)

var _ Client = &s3wrapper.Client{}

type FileSystemClientTypeMap map[string]FileSystemClientType
type FileSystemClientType string

const S3 = FileSystemClientType("s3")
const UnKnown = FileSystemClientType("unknown")

var typeMap = FileSystemClientTypeMap{
	"s3": S3,
}

type Client interface {
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

func New(clientType string, detail ClientDetail) Client {
	types := GetFileSystemClientType(clientType)
	switch types {
	case S3:
		s3Config := s3wrapper.S3Config{
			Region:     detail.Region,
			BucketName: detail.Bucket,
		}
		client := s3wrapper.New(s3Config)
		return client
	default:
		return nil
	}
}
