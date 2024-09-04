package aws_handlers

import (
	"errors"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

func GetS3Buckets(s3Client *s3.S3) ([]map[string]interface{}, error) {

	input := &s3.ListBucketsInput{}

	buckets, err := s3Client.ListBuckets(input)

	if err != nil {
		return nil, errors.New("Error while Fetching AWS S3 Buckets")
	}

	data := make([]map[string]interface{}, len(buckets.Buckets))

	for i, bucket := range buckets.Buckets {
		bucketName := *bucket.Name

		objectsResponse, err := s3Client.ListObjectsV2(&s3.ListObjectsV2Input{
			Bucket: aws.String(bucketName),
		})

		if err != nil {
			continue
		}

		var files, folders []string

		for _, object := range objectsResponse.Contents {
			if strings.HasSuffix(*object.Key, "/") {
				folders = append(folders, *object.Key)
			} else {
				files = append(files, *object.Key)
			}
		}

		// Get tags for the bucket
		tagsResponse, err := s3Client.GetBucketTagging(&s3.GetBucketTaggingInput{
			Bucket: aws.String(bucketName),
		})

		var tags []map[string]string
		if err == nil {
			for _, tagSet := range tagsResponse.TagSet {
				tag := map[string]string{
					"Key":   *tagSet.Key,
					"Value": *tagSet.Value,
				}
				tags = append(tags, tag)
			}
		}

		bucketDetails := map[string]interface{}{
			"bucketName": bucketName,
			"files":      files,
			"folders":    folders,
			"tags":       tags,
		}

		data[i] = bucketDetails
	}

	return data, nil
}
