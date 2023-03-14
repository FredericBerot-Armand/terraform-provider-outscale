package outscale

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccOutscaleOAPISnapshotExportTask_basic(t *testing.T) {
	osuBucketNames := []string{
		acctest.RandomWithPrefix("terraform-export-bucket-"),
		acctest.RandomWithPrefix("terraform-export-bucket-"),
	}
	region := os.Getenv("OUTSCALE_REGION")
	tags := `tags {
		key = "test"
		value = "test"
	}
	tags {
		key = "test-1"
		value = "test-1"
	}`
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccOutscaleOAPISnapshotExportTaskConfig(region, "", osuBucketNames[0]),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOutscaleOAPISnapshotExportTaskExists("outscale_snapshot_export_task.outscale_snapshot_export_task"),
				),
			},
			{
				Config: testAccOutscaleOAPISnapshotExportTaskConfig(region, tags, osuBucketNames[1]),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOutscaleOAPISnapshotExportTaskExists("outscale_snapshot_export_task.outscale_snapshot_export_task"),
				),
			},
		},
	})
}

func testAccCheckOutscaleOAPISnapshotExportTaskExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No image task id is set")
		}

		return nil
	}
}

func testAccOutscaleOAPISnapshotExportTaskConfig(region, tags, osuBucketName string) string {
	return fmt.Sprintf(`
	resource "outscale_volume" "outscale_volume_snap" {
    	subregion_name   = "%[1]sa"
    	size                = 10
	}
	resource "outscale_snapshot" "outscale_snapshot" {
    	volume_id = outscale_volume.outscale_volume_snap.volume_id
	}
	resource "outscale_snapshot_export_task" "outscale_snapshot_export_task" {
		snapshot_id                     = outscale_snapshot.outscale_snapshot.snapshot_id
		osu_export {
			disk_image_format = "qcow2"
			osu_bucket        = "%[3]s"
			osu_prefix        = "new-export"
		}
		%[2]s
	}
	`, region, tags, osuBucketName)
}
