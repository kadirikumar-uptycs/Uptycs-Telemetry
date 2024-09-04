package gcp_handlers

import (
	"errors"
	"strings"

	"google.golang.org/api/storage/v1"
)

func GetStorageInfo(service *storage.Service, projectId string) (interface{}, error) {

	response, err := service.Buckets.List(projectId).Do()

	if err != nil {
		return nil, errors.New("Error While Retrieving data from the Storage Buckets")
	}

	data := make([]map[string]interface{}, len(response.Items))

	for i, bucket := range response.Items {

		bucketData := make(map[string]interface{})
		bucketData["bucketName"] = bucket.Name
		bucketData["bucketCreatedTime"] = bucket.TimeCreated
		bucketData["bucketRegion"] = bucket.Location
		bucketData["bucketStorageClass"] = bucket.StorageClass

		filesResponse, err := service.Objects.List(bucket.Name).Do()

		if err != nil {
			return nil, errors.New("Error while retrieving files from GCP Storage Bucket")
		}

		var files, folders []string

		for _, file := range filesResponse.Items {
			if strings.HasSuffix(file.Name, "/") {
				folders = append(folders, file.Name)
			} else {
				files = append(files, file.Name)
			}
		}

		bucketMetaData, err := service.Buckets.Get(bucket.Name).Do()

		var tags []map[string]interface{}

		for key, value := range bucketMetaData.Labels {
			tags = append(tags, map[string]interface{}{
				"key":   key,
				"value": value,
			})
		}
		if len(tags) == 0 {
			tags = make([]map[string]interface{}, 0)
		}

		bucketData["bucketFiles"] = files
		bucketData["bucketFolders"] = folders
		bucketData["bucketTags"] = tags

		data[i] = bucketData

	}

	return data, nil
}
