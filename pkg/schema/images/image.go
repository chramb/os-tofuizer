package images

import (
	"fmt"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/images"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

type Image_v2 struct {
	ResourceType string `json:"-" hcl:",label"`
	ResourceName string `json:"-" hcl:",label"`

	Name string `json:"Name" hcl:"name"`

	ContainerFormat string `json:"container_format" hcl:"container_format"`
	DiskFormat      string `json:"disk_format" hcl:"disk_format"`
}

func Generate(client *gophercloud.ServiceClient, id string) *hclwrite.Block {
	r := images.Get(client, id)

	res := Image_v2{
		ResourceType: "openstack_images_image_v2",
		ResourceName: "name",
	}
	// TODO: figure out custom extractor: https://github.com/gophercloud/gophercloud/blob/master/docs/FAQ.md#overriding-default-unmarshaljson-method
	err := r.ExtractInto(&res)

	if err != nil {
		fmt.Printf("%s\n", err)
	}
	// f := hclwrite.NewEmptyFile()
	return gohcl.EncodeAsBlock(&res, "resource")
}
