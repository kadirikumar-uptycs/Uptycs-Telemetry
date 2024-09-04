package aws_handlers

import (
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go/service/iam"
)

func GetIAMPolicies(client *iam.IAM) ([]map[string]interface{}, error) {
	policiesResponse, err := client.ListPolicies(&iam.ListPoliciesInput{})

	if err != nil {
		return nil, errors.New("Error While Fetching IAM Poilicies")
	}

	data := make([]map[string]interface{}, len(policiesResponse.Policies))

	for i, policy := range policiesResponse.Policies {

		policyData := make(map[string]interface{})
		policyData["policyId"] = *policy.PolicyId
		policyData["policyARN"] = *policy.Arn
		policyData["PolicyName"] = *policy.PolicyName
		policyData["policyCreatedTime"] = *policy.CreateDate
		policyData["policyUpdatedDate"] = *policy.UpdateDate

		data[i] = policyData
	}
	return data, nil
}

func GetIAMRoles(client *iam.IAM) ([]map[string]interface{}, error) {
	rolesResponse, err := client.ListRoles(&iam.ListRolesInput{})

	if err != nil {
		return nil, errors.New("Error while Fetching IAM Roles")
	}

	data := make([]map[string]interface{}, len(rolesResponse.Roles))

	for i, role := range rolesResponse.Roles {

		roleData := make(map[string]interface{})

		roleData["RoleId"] = *role.RoleId
		roleData["RoleName"] = *role.RoleName
		roleData["RoleARN"] = *role.Arn
		roleData["RoleCreatedTime"] = *role.CreateDate
		roleData["RoleDescription"] = *role.Description

		data[i] = roleData
	}
	return data, nil
}

func GetIAMUsers(client *iam.IAM) ([]map[string]interface{}, error) {

	usersResponse, err := client.ListUsers(&iam.ListUsersInput{})

	if err != nil {
		return nil, errors.New("Error While Fetching IAM Users")
	}

	data := make([]map[string]interface{}, len(usersResponse.Users))

	for i, user := range usersResponse.Users {
		userData := make(map[string]interface{})

		fmt.Print(*user)

		userData["userId"] = *user.UserId
		userData["userARN"] = *user.Arn
		userData["userName"] = *user.UserName
		userData["userCreatedTime"] = *user.CreateDate
		// userData["userPermissionsBoundary"] = *user.PermissionsBoundary

		data[i] = userData
	}
	return data, nil
}
