package clients

import (
	"context"
	"errors"
	"os"

	"google.golang.org/api/cloudfunctions/v1"
	"google.golang.org/api/compute/v1"
	"google.golang.org/api/option"
	"google.golang.org/api/storage/v1"
)

func InitGCP() {
	gcpAccessKeysPath := "./config/GCP-ACCESS-KEYS.json"

	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", gcpAccessKeysPath)
}

func GetGCPService(service string) (interface{}, error) {
	switch service {
	case "CloudFunction":
		service, err := cloudfunctions.NewService(context.Background(), option.WithScopes(cloudfunctions.CloudPlatformScope))
		if err != nil {
			return nil, errors.New("Error While Creating Cloud Function service")
		}
		return service, nil

	case "compute":
		service, err := compute.NewService(context.Background(), option.WithScopes(compute.ComputeScope))
		if err != nil {
			return nil, errors.New("Error While Creating Compute Service service")
		}
		return service, nil

	case "storage":
		service, err := storage.NewService(context.Background(), option.WithScopes(storage.DevstorageReadOnlyScope))
		if err != nil {
			return nil, errors.New("Error While Creating Storage Service service")
		}
		return service, nil

	default:
		return nil, errors.New("Service Not Allowed")
	}
}
