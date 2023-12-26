package flavors

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/flavors"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

type Flavor_v2 struct {
	ResourceType string `json:"-" hcl:",label"`
	ResourceName string `json:"-" hcl:",label"`

	// Region string `json:"-" hcl:"region"`
	Name        string  `json:"name" hcl:"name"`
	Description *string `json:"description" hcl:"description"`
	Ram         int     `json:"ram" hcl:"ram"`
	FlavorID    string  `json:"id" hcl:"flavor_id"`
	VCPUs       int     `json:"vcpus" hcl:"vcpus"`
	Disk        int     `json:"disk" hcl:"disk"`

	Ephemeral  int     `json:"OS-FLV-EXT-DATA:ephemeral" hcl:"ephemeral"`
	Swap       int     `json:"swap" hcl:"swap"`
	RxTxFactor float64 `json:"rxtx_factor" hcl:"rx_tx_factor"`
	IsPublic   bool    `json:"os-flavor-access:is_public" hcl:"is_public"`

	ExtraSpecs map[string]string `json:"extra_specs" hcl:"extra_specs"`
}

type Flavor struct {
	ID          string  `json:"id"`
	Disk        int     `json:"disk"`
	RAM         int     `json:"ram"`
	Name        string  `json:"name"`
	RxTxFactor  float64 `json:"rxtx_factor"`
	Swap        int     `json:"-"`
	VCPUs       int     `json:"vcpus"`
	IsPublic    bool    `json:"os-flavor-access:is_public"`
	Ephemeral   int     `json:"OS-FLV-EXT-DATA:ephemeral"`
	Description string  `json:"description"`
}

type FlavorContainer struct {
	Flavor       Flavor_v2 `json:"flavor"`
	flavorResult flavors.GetResult
	specsResult  flavors.ListExtraSpecsResult
}

func (f *FlavorContainer) Get(client *gophercloud.ServiceClient, id string) *FlavorContainer {
	f.flavorResult = flavors.Get(client, id)
	f.specsResult = flavors.ListExtraSpecs(client, id)

	return f
}

func (f *FlavorContainer) Generate() *hclwrite.Block {
	err := f.flavorResult.ExtractInto(f)
	if err != nil {
		fmt.Printf("%s\n", err)
	}

	err = f.specsResult.ExtractInto(&f.Flavor)
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	// if *f.Flavor.Description == "" {
	// 	f.Flavor.Description = nil
	// }
	if f.Flavor.Description != nil && *f.Flavor.Description == "" {
		f.Flavor.Description = nil
	}

	return gohcl.EncodeAsBlock(f.Flavor, "resource")
}

func (f *FlavorContainer) Json() interface{} {
	bs, err := json.Marshal(f.flavorResult.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(-1)
	}
	return bs
}
