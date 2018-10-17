package alicloud

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccAlicloudCenRouteEntry_importBasic(t *testing.T) {
	resourceName := "alicloud_cen_route_entry.foo"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCenRouteEntryDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCenRouteEntryConfig,
			},

			resource.TestStep{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
