package clients

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure/auth"
)

var (
	Authorizer autorest.Authorizer
)

func InitAzure() error {

	data, err := os.ReadFile("config\azure_credentials.json")

	if err != nil {
		return errors.New("Error while reading credentials file")
	}

	var credentials map[string]string

	err = json.Unmarshal(data, &credentials)

	client_id := credentials["client_id"]
	client_secret := credentials["client_secret"]
	tenant_id := credentials["tenant_id"]
	subscription_id := credentials["subscription_id"]

	os.Setenv("AZURE_CLIENT_ID", client_id)
	os.Setenv("AZURE_CLIENT_SECRET", client_secret)
	os.Setenv("AZURE_TENANT_ID", tenant_id)
	os.Setenv("AZURE_SUBSCRIPTION_ID", subscription_id)

	Authorizer, err = auth.NewAuthorizerFromEnvironment()

	if err != nil {
		return errors.New("Error while creating New Authorizer")
	}

	return nil
}
