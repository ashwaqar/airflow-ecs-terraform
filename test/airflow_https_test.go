package test

import (
	"testing"
	"time"

	"github.com/gruntwork-io/terratest/modules/logger"
	"github.com/gruntwork-io/terratest/modules/terraform"
	testStructure "github.com/gruntwork-io/terratest/modules/test-structure"
)

func getTerraformOptionsAirflowHTTPS(t *testing.T) *terraform.Options {
	tempTestFolder := testStructure.CopyTerraformFolderToTemp(t, "..", ".")

	terraformOptions := &terraform.Options{
		TerraformDir:       tempTestFolder,
		Vars:               map[string]interface{}{},
		MaxRetries:         5,
		TimeBetweenRetries: 5 * time.Minute,
		NoColor:            true,
		Logger:             logger.TestingT,
	}

	return terraformOptions
}

func TestAirflowHTTPS(t *testing.T) {
	t.Parallel()
	terraformOptions := getTerraformOptionsAirflowHTTPS(t)

	// terraform destroy => when test completes
	defer terraform.Destroy(t, terraformOptions)
	TerraformInitPlanApply(t, terraformOptions)
}
