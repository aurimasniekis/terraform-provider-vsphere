// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package vsphere

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccResourceVSphereVmClass_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccVSphereVmClassConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("vsphere_virtual_machine_class.vm_class_1", "id"),
				),
			},
		},
	})
}

func TestAccResourceVSphereVmClass_vgpu(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccVSphereVmClassConfig_vgpu(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("vsphere_virtual_machine_class.vm_class_1", "id"),
				),
			},
		},
	})
}

func testAccVSphereVmClassConfig() string {
	return `
		resource "vsphere_virtual_machine_class" "vm_class_1" {
			name = "test-class-11"
			cpus = 4
			memory = 4096
			memory_reservation = 100
		}
`
}

func testAccVSphereVmClassConfig_vgpu() string {
	return `
		resource "vsphere_virtual_machine_class" "vm_class_1" {
			name = "test-class-11"
			cpus = 4
			memory = 4096
			memory_reservation = 100
			vgpu_devices = [ "mockup-vmiop" ]
		}
`
}
