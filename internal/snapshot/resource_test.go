package snapshot_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/hetznercloud/hcloud-go/v2/hcloud"
	"github.com/hetznercloud/terraform-provider-hcloud/internal/server"
	"github.com/hetznercloud/terraform-provider-hcloud/internal/snapshot"
	"github.com/hetznercloud/terraform-provider-hcloud/internal/sshkey"
	"github.com/hetznercloud/terraform-provider-hcloud/internal/teste2e"
	"github.com/hetznercloud/terraform-provider-hcloud/internal/testsupport"
	"github.com/hetznercloud/terraform-provider-hcloud/internal/testtemplate"
)

func TestAccSnapshotResource(t *testing.T) {
	var s hcloud.Image
	tmplMan := testtemplate.Manager{}

	sk := sshkey.NewRData(t, "snapshot-basic")
	resServer := &server.RData{
		Name:  "snapshot-test",
		Type:  teste2e.TestServerType,
		Image: teste2e.TestImage,
		Labels: map[string]string{
			"tf-test": fmt.Sprintf("tf-test-snapshot-%d", tmplMan.RandInt),
		},
		SSHKeys: []string{sk.TFID() + ".id"},
	}
	resServer.SetRName("server-snapshot")
	res := &snapshot.RData{
		Description: "snapshot-basic",
		ServerID:    resServer.TFID() + ".id",
		Labels: map[string]string{
			"tf-test": fmt.Sprintf("tf-test-snapshot-%d", tmplMan.RandInt),
		},
	}
	res.SetRName("snapshot-basic")
	resRenamed := &snapshot.RData{
		Description: "snapshot-basic-changed",
		ServerID:    resServer.TFID() + ".id",
		Labels: map[string]string{
			"tf-test": fmt.Sprintf("tf-test-fip-assignment-%d", tmplMan.RandInt),
		}}
	resRenamed.SetRName("snapshot-basic")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 teste2e.PreCheck(t),
		ProtoV6ProviderFactories: teste2e.ProtoV6ProviderFactories(),
		CheckDestroy:             testsupport.CheckResourcesDestroyed(snapshot.ResourceType, snapshot.ByID(t, &s)),
		Steps: []resource.TestStep{
			{
				// Create a new Snapshot using the required values
				// only.
				Config: tmplMan.Render(t,
					"testdata/r/hcloud_ssh_key", sk,
					"testdata/r/hcloud_server", resServer,
					"testdata/r/hcloud_snapshot", res,
				),
				Check: resource.ComposeTestCheckFunc(
					testsupport.CheckResourceExists(res.TFID(), snapshot.ByID(t, &s)),
					resource.TestCheckResourceAttr(res.TFID(), "description", "snapshot-basic"),
				),
			},
			{
				// Try to import the newly created Snapshot
				ResourceName:      res.TFID(),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Update the Snapshot created in the previous step by
				// setting all optional fields and renaming the Snapshot.
				Config: tmplMan.Render(t,
					"testdata/r/hcloud_ssh_key", sk,
					"testdata/r/hcloud_server", resServer,
					"testdata/r/hcloud_snapshot", resRenamed,
				),
				Check: resource.ComposeAggregateTestCheckFunc(
					testsupport.CheckResourceExists(res.TFID(), snapshot.ByID(t, &s)),
					resource.TestCheckResourceAttr(resRenamed.TFID(), "description", "snapshot-basic-changed"),
				),
			},
		},
	})
}
