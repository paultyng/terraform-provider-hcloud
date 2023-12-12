package datacenter_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/hetznercloud/terraform-provider-hcloud/internal/datacenter"
	"github.com/hetznercloud/terraform-provider-hcloud/internal/e2etests"
	"github.com/hetznercloud/terraform-provider-hcloud/internal/testtemplate"
)

func TestAccHcloudDataSourceDatacenterTest(t *testing.T) {
	tmplMan := testtemplate.Manager{}

	dcByName := &datacenter.DData{
		DatacenterName: "fsn1-dc14",
	}
	dcByName.SetRName("dc_by_name")
	dcByID := &datacenter.DData{
		DatacenterID: "4",
	}
	dcByID.SetRName("dc_by_id")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 e2etests.PreCheck(t),
		ProtoV6ProviderFactories: e2etests.ProtoV6ProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: tmplMan.Render(t,
					"testdata/d/hcloud_datacenter", dcByName,
					"testdata/d/hcloud_datacenter", dcByID,
				),

				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dcByName.TFID(), "id", "4"),
					resource.TestCheckResourceAttr(dcByName.TFID(), "name", "fsn1-dc14"),

					resource.TestCheckResourceAttr(dcByID.TFID(), "id", "4"),
					resource.TestCheckResourceAttr(dcByID.TFID(), "name", "fsn1-dc14"),
				),
			},
		},
	})
}

func TestAccHcloudDataSourceDatacenter_UpgradePluginFramework(t *testing.T) {
	tmplMan := testtemplate.Manager{}

	dcByName := &datacenter.DData{
		DatacenterName: "fsn1-dc14",
	}
	dcByName.SetRName("dc_by_name")
	dcByID := &datacenter.DData{
		DatacenterID: "4",
	}
	dcByID.SetRName("dc_by_id")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck: e2etests.PreCheck(t),
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"hcloud": {
						VersionConstraint: "1.44.1",
						Source:            "hetznercloud/hcloud",
					},
				},

				Config: tmplMan.Render(t,
					"testdata/d/hcloud_datacenter", dcByName,
					"testdata/d/hcloud_datacenter", dcByID,
					"testdata/r/terraform_data_resource", dcByName,
					"testdata/r/terraform_data_resource", dcByID,
				),
			},
			{
				ProtoV6ProviderFactories: e2etests.ProtoV6ProviderFactories(),

				Config: tmplMan.Render(t,
					"testdata/d/hcloud_datacenter", dcByName,
					"testdata/d/hcloud_datacenter", dcByID,
					"testdata/r/terraform_data_resource", dcByName,
					"testdata/r/terraform_data_resource", dcByID,
				),

				PlanOnly: true,
			},
		},
	})
}

func TestAccHcloudDataSourceDatacentersTest(t *testing.T) {
	tmplMan := testtemplate.Manager{}

	datacentersD := &datacenter.DDataList{}
	datacentersD.SetRName("ds")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 e2etests.PreCheck(t),
		ProtoV6ProviderFactories: e2etests.ProtoV6ProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: tmplMan.Render(t,
					"testdata/d/hcloud_datacenters", datacentersD,
				),

				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(datacentersD.TFID(), "datacenter_ids.0", "2"),
					resource.TestCheckResourceAttr(datacentersD.TFID(), "datacenter_ids.1", "3"),
					resource.TestCheckResourceAttr(datacentersD.TFID(), "datacenter_ids.2", "4"),
					resource.TestCheckResourceAttr(datacentersD.TFID(), "datacenter_ids.3", "5"),
					resource.TestCheckResourceAttr(datacentersD.TFID(), "datacenter_ids.4", "6"),
					resource.TestCheckResourceAttr(datacentersD.TFID(), "names.0", "nbg1-dc3"),
					resource.TestCheckResourceAttr(datacentersD.TFID(), "names.1", "hel1-dc2"),
					resource.TestCheckResourceAttr(datacentersD.TFID(), "names.2", "fsn1-dc14"),
					resource.TestCheckResourceAttr(datacentersD.TFID(), "names.3", "ash-dc1"),
					resource.TestCheckResourceAttr(datacentersD.TFID(), "names.4", "hil-dc1"),
					resource.TestCheckResourceAttr(datacentersD.TFID(), "datacenters.#", "5"),
					resource.TestCheckResourceAttr(datacentersD.TFID(), "datacenters.0.name", "nbg1-dc3"),
					resource.TestCheckResourceAttr(datacentersD.TFID(), "datacenters.1.name", "hel1-dc2"),
					resource.TestCheckResourceAttr(datacentersD.TFID(), "datacenters.2.name", "fsn1-dc14"),
					resource.TestCheckResourceAttr(datacentersD.TFID(), "datacenters.3.name", "ash-dc1"),
					resource.TestCheckResourceAttr(datacentersD.TFID(), "datacenters.4.name", "hil-dc1"),
				),
			},
		},
	})
}

func TestAccHcloudDataSourceDatacenters_UpgradePluginFramework(t *testing.T) {
	tmplMan := testtemplate.Manager{}

	datacentersD := &datacenter.DDataList{}
	datacentersD.SetRName("ds")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck: e2etests.PreCheck(t),
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"hcloud": {
						VersionConstraint: "1.44.1",
						Source:            "hetznercloud/hcloud",
					},
				},

				Config: tmplMan.Render(t,
					"testdata/d/hcloud_datacenters", datacentersD,
					"testdata/r/terraform_data_resource", datacentersD,
				),
			},
			{
				ProtoV6ProviderFactories: e2etests.ProtoV6ProviderFactories(),

				Config: tmplMan.Render(t,
					"testdata/d/hcloud_datacenters", datacentersD,
					"testdata/r/terraform_data_resource", datacentersD,
				),

				PlanOnly: true,
			},
		},
	})
}
