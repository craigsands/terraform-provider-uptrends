package uptrends

import (
	"os"
	"testing"
	
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

var (
	testAccProviders map[string]terraform.ResourceProvider
    testAccProvider *schema.Provider
)

func init() {
    testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"uptrends": testAccProvider,
	}
}

// testAccPreCheck validates the necessary test keys exist in the testing environment
func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("UPTRENDS_USERNAME"); v == "" {
		t.Fatal("UPTRENDS_USERNAME must be set for acceptance tests")
	}

	if v := os.Getenv("UPTRENDS_PASSWORD"); v == "" {
		t.Fatal("UPTRENDS_PASSWORD must be set for acceptance tests")
	}
}
