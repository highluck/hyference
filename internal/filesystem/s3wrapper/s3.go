package s3wrapper

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	errors2 "github.com/hyference/errors"
	"github.com/rs/zerolog/log"
	"io/fs"
	"io/ioutil"
)

type Client struct {
	s3       *s3.S3
	s3Config S3Config
}

func New(cfg S3Config) *Client {
	awsCfg := aws.NewConfig()
	awsCfg.Region = aws.String(cfg.BucketName)
	awsCfg.S3ForcePathStyle = aws.Bool(true)
	sess := session.Must(session.NewSession(awsCfg))
	newS3 := s3.New(sess, &aws.Config{})

	return &Client{
		s3: newS3,
	}
}

func (c *Client) Get(path string) (*s3.GetObjectOutput, error) {
	log.Info().Msgf("bucket : %s, key : %s", c.s3Config.BucketName, path)

	res, err := c.s3.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(c.s3Config.BucketName),
		Key:    aws.String(normalizePath(path)),
	})
	return res, err
}

func (c *Client) GetByAlphaIgnore(path string) (*s3.GetObjectOutput, error) {
	log.Info().Msgf("bucket : %s, key : %s", c.s3Config.BucketName, path)

	res, err := c.s3.GetObject(&s3.GetObjectInput{
		Bucket: aws.String("karrot-search-bucket-prod"),
		Key:    aws.String(normalizePath(path)),
	})
	return res, err
}

func (c *Client) DownloadModel(model string, s3ModelPath string) error {
	op := errors2.GetMethodName()
	s3ModelResponse, err := c.Get(s3ModelPath)
	if err != nil {
		return errors2.Wrapper(err, op)
	}

	modelBytes, err := ioutil.ReadAll(s3ModelResponse.Body)
	if err != nil {
		return errors2.Wrapper(err, op)
	}
	if err := ioutil.WriteFile(model, modelBytes, fs.ModePerm); err != nil {
		return errors2.Wrapper(err, op)
	}
	return nil
}

func normalizePath(path string) string {
	if path == "" {
		return ""
	}

	if path[:1] == "/" {
		path = path[1:]
	}
	return path
}
