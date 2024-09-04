package routes

import (
	"net/http"

	"github.com/1722101709/Telemetry/pkg/clients"
	"github.com/1722101709/Telemetry/pkg/handlers"
	aws_handlers "github.com/1722101709/Telemetry/pkg/handlers/aws"
	gcp_handlers "github.com/1722101709/Telemetry/pkg/handlers/gcp"
	"google.golang.org/api/cloudfunctions/v1"
	"google.golang.org/api/compute/v1"
	"google.golang.org/api/storage/v1"
)

func CloudTrailEvents(w http.ResponseWriter, r *http.Request) {
	client := clients.CloudTrailClient
	data, err := aws_handlers.GetCloudTrailEvents(client)

	if err != nil {
		w = *handlers.ApiResponse(w, http.StatusBadRequest, err)
	}

	w = *handlers.ApiResponse(w, http.StatusOK, data)
}

func S3Buckets(w http.ResponseWriter, r *http.Request) {
	client := clients.S3Client

	data, err := aws_handlers.GetS3Buckets(client)

	if err != nil {
		w = *handlers.ApiResponse(w, http.StatusBadRequest, err)
	}

	w = *handlers.ApiResponse(w, http.StatusOK, data)

}

func EC2Instances(w http.ResponseWriter, r *http.Request) {
	client := clients.EC2Client

	data, err := aws_handlers.GetEC2Instances(client)

	if err != nil {
		w = *handlers.ApiResponse(w, http.StatusBadRequest, err)
	}

	w = *handlers.ApiResponse(w, http.StatusOK, data)

}

func IAMPolicies(w http.ResponseWriter, r *http.Request) {
	client := clients.IAMClient

	data, err := aws_handlers.GetIAMPolicies(client)

	if err != nil {
		w = *handlers.ApiResponse(w, http.StatusBadRequest, err)
	}

	w = *handlers.ApiResponse(w, http.StatusOK, data)
}

func IAMRoles(w http.ResponseWriter, r *http.Request) {
	client := clients.IAMClient

	data, err := aws_handlers.GetIAMRoles(client)

	if err != nil {
		w = *handlers.ApiResponse(w, http.StatusBadRequest, err)
	}

	w = *handlers.ApiResponse(w, http.StatusOK, data)
}

func IAMUsers(w http.ResponseWriter, r *http.Request) {
	client := clients.IAMClient

	data, err := aws_handlers.GetIAMUsers(client)

	if err != nil {
		w = *handlers.ApiResponse(w, http.StatusBadRequest, err)
	}

	w = *handlers.ApiResponse(w, http.StatusOK, data)
}

func VPCs(w http.ResponseWriter, r *http.Request) {
	client := clients.EC2Client

	data, err := aws_handlers.GetVPCs(client)

	if err != nil {
		w = *handlers.ApiResponse(w, http.StatusBadRequest, err)
	}

	w = *handlers.ApiResponse(w, http.StatusOK, data)
}

// GCP Functions

func StorageBuckets(w http.ResponseWriter, r *http.Request) {

	GCP_PROJECT_ID := "extreme-cycling-399615"

	temp, err := clients.GetGCPService("storage")

	if err != nil {
		w = *handlers.ApiResponse(w, http.StatusBadRequest, err)
	}

	storageService := temp.(*storage.Service)

	data, err := gcp_handlers.GetStorageInfo(storageService, GCP_PROJECT_ID)

	if err != nil {
		w = *handlers.ApiResponse(w, http.StatusBadRequest, err)
	}

	w = *handlers.ApiResponse(w, http.StatusOK, data)
}

func CloudFunctions(w http.ResponseWriter, r *http.Request) {

	GCP_PROJECT_ID := "extreme-cycling-399615"

	temp, err := clients.GetGCPService("CloudFunction")

	if err != nil {
		w = *handlers.ApiResponse(w, http.StatusBadRequest, err)
	}

	cloudFunctionService := temp.(*cloudfunctions.Service)

	data, err := gcp_handlers.GetCloudFunctionsInfo(cloudFunctionService, GCP_PROJECT_ID)

	if err != nil {
		w = *handlers.ApiResponse(w, http.StatusBadRequest, err)
	}

	w = *handlers.ApiResponse(w, http.StatusOK, data)
}

func ComputeEngines(w http.ResponseWriter, r *http.Request) {

	GCP_PROJECT_ID := "extreme-cycling-399615"

	temp, err := clients.GetGCPService("compute")

	if err != nil {
		w = *handlers.ApiResponse(w, http.StatusBadRequest, err.Error())
	} else {
		computeService := temp.(*compute.Service)

		data, err := gcp_handlers.GetComputeInfo(computeService, GCP_PROJECT_ID)

		if err != nil {
			w = *handlers.ApiResponse(w, http.StatusBadRequest, err.Error())
			return
		} else {
			w = *handlers.ApiResponse(w, http.StatusOK, data)
		}

	}
}
