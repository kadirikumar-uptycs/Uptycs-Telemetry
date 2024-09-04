package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/1722101709/Telemetry/pkg/clients"
	"github.com/1722101709/Telemetry/pkg/routes"
)

func main() {

	clients.StartClients()
	clients.InitGCP()
	err := clients.InitAzure()

	if err != nil {
		fmt.Println(err)
	}

	// AWS API Endpoints
	http.HandleFunc("/getAWSCloudTrailEvents", routes.CloudTrailEvents)
	http.HandleFunc("/getAWSS3Buckets", routes.S3Buckets)
	http.HandleFunc("/getAWSEC2Instances", routes.EC2Instances)
	http.HandleFunc("/getAWSIAMPolicies", routes.IAMPolicies)
	http.HandleFunc("/getAWSIAMRoles", routes.IAMRoles)
	http.HandleFunc("/getAWSIAMUsers", routes.IAMUsers)
	http.HandleFunc("/getAWSVPCs", routes.VPCs)

	// GCP API Endpoints
	http.HandleFunc("/getGCPBuckets", routes.StorageBuckets)
	http.HandleFunc("/getGCPCloudFunctions", routes.CloudFunctions)
	http.HandleFunc("/getGCPComputeEngines", routes.ComputeEngines)

	port := 1729

	fmt.Printf("Starting server on %d \n", port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))

	return

}
