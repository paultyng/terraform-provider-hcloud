package floatingip_test

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/terraform-providers/terraform-provider-hcloud/internal/floatingip"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/terraform-providers/terraform-provider-hcloud/internal/testsupport"
	"github.com/terraform-providers/terraform-provider-hcloud/internal/testtemplate"
)

func TestAccHcloudDataSourceFloatingIPTest(t *testing.T) {
	tmplMan := testtemplate.Manager{}

	res := &floatingip.RData{
		Name: "floatingip-ds-test",
		Type: "ipv4",
		Labels: map[string]string{
			"key": strconv.Itoa(acctest.RandInt()),
		},
		HomeLocationName: "nbg1",
	}
	res.SetRName("floatingip-ds-test")
	floatingipByName := &floatingip.DData{
		FloatingIPName: res.TFID() + ".name",
	}
	floatingipByName.SetRName("floatingip_by_name")
	floatingipByID := &floatingip.DData{
		FloatingIPID: res.TFID() + ".id",
	}
	floatingipByID.SetRName("floatingip_by_id")
	floatingipBySel := &floatingip.DData{
		LabelSelector: fmt.Sprintf("key=${%s.labels[\"key\"]}", res.TFID()),
	}
	floatingipBySel.SetRName("floatingip_by_sel")

	resource.Test(t, resource.TestCase{
		PreCheck:     testsupport.AccTestPreCheck(t),
		Providers:    testsupport.AccTestProviders(),
		CheckDestroy: testsupport.CheckResourcesDestroyed(floatingip.ResourceType, floatingip.ByID(t, nil)),
		Steps: []resource.TestStep{
			{
				Config: tmplMan.Render(t,
					"testdata/r/hcloud_floating_ip", res,
					"testdata/d/hcloud_floating_ip", floatingipByName,
					"testdata/d/hcloud_floating_ip", floatingipByID,
					"testdata/d/hcloud_floating_ip", floatingipBySel,
				),

				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(floatingipByName.TFID(),
						"name", fmt.Sprintf("%s--%d", res.Name, tmplMan.RandInt)),
					resource.TestCheckResourceAttr(floatingipByName.TFID(), "home_location", res.HomeLocationName),

					resource.TestCheckResourceAttr(floatingipByID.TFID(),
						"name", fmt.Sprintf("%s--%d", res.Name, tmplMan.RandInt)),
					resource.TestCheckResourceAttr(floatingipByID.TFID(), "home_location", res.HomeLocationName),

					resource.TestCheckResourceAttr(floatingipBySel.TFID(),
						"name", fmt.Sprintf("%s--%d", res.Name, tmplMan.RandInt)),
					resource.TestCheckResourceAttr(floatingipBySel.TFID(), "home_location", res.HomeLocationName),
				),
			},
		},
	})
}
