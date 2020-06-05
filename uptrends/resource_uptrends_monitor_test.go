package uptrends

import (
	"fmt"
	"testing"

	"github.com/craigsands/uptrends-sdk-go"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccUptrendsMonitor_basic(t *testing.T) {
	resourceName := "uptrends_monitor.foo"
	rName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckUptrendsMonitorDestroy,
		Steps: []resource.TestStep{
			// Test: Create
			{
				Config: testAccUptrendsMonitorConfig(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUptrendsMonitorExists(resourceName),
				),
			},
			// Test: Update
			{
				Config: testAccUptrendsMonitorConfigUpdated(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUptrendsMonitorExists(resourceName),
				),
			},
		},
	})
}

func testAccCheckUptrendsMonitorExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		providerConf := testAccProvider.Meta().(*ProviderConfiguration)
		client := providerConf.Client
		auth := providerConf.Auth

		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not found: %s", n)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("no monitor ID is set")
		}

		opts := uptrends.MonitorGetMonitorOpts{}
		found, _, err := client.MonitorApi.MonitorGetMonitor(auth, rs.Primary.ID, &opts)
		if err != nil {
			return err
		}

		if found.MonitorGuid != rs.Primary.ID {
			return fmt.Errorf("monitor not found: %v - %v", rs.Primary.ID, found)
		}

		return nil
	}
}

func testAccCheckUptrendsMonitorDestroy(s *terraform.State) error {
	providerConf := testAccProvider.Meta().(*ProviderConfiguration)
	client := providerConf.Client
	auth := providerConf.Auth

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "uptrends_monitor" {
			continue
		}

		opts := uptrends.MonitorGetMonitorOpts{}
		_, _, err := client.MonitorApi.MonitorGetMonitor(auth, rs.Primary.ID, &opts)
		if err == nil {
			return fmt.Errorf("monitor still exists")
		}

	}
	return nil
}

func testAccUptrendsMonitorConfig(name string) string {
	return fmt.Sprintf(`
resource "uptrends_monitor" "foo" {
	name         = "%[1]s"
	monitor_type = "Http"
	url          = "https://example.com"
}
`, name)
}

func testAccUptrendsMonitorConfigUpdated(name string) string {
	return fmt.Sprintf(`
resource "uptrends_monitor" "foo" {
	name         = "%[1]s-updated"
	monitor_type = "Http"
	url          = "https://example-updated.com"
}
`, name)
}
