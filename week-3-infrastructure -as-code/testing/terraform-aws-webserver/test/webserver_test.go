package test

import (
	"testing"
	"fmt"
	"time"

	http_helper "github.com/gruntwork-io/terratest/modules/http-helper"
	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestTerraformWebserverExample(t *testing.T) {
	// The values to pass into Terraform CLI
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		//Path to where example Terraform code is located
		TerraformDir: "..\\examples\\webserver",

		//Variables to pass to the terraform code using -var
		Vars: map[string]interface{}{
			"region":     "us-west-2",
			"servername": "testwebserver",
		},
	})
	//run a Terraform Init and Apply
	terraform.InitAndApply(t, terraformOptions)

	//run a Terraform Destroy at the end of test
	defer terraform.Destroy(t, terraformOptions)

	publicIp := terraform.Output(t,terraformOptions,"public_ip")

	url := fmt.Sprintf("http://%s:8080",publicIp)

	http_helper.HttpGetWithRetry(t,url,nil,200,"I made a terraform module!", 30, 5*time.Second)
}
