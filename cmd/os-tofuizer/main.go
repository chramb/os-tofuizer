package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/chramb/os-tofiuzer/pkg/schema/flavors"
	img "github.com/chramb/os-tofiuzer/pkg/schema/images"
	"github.com/hashicorp/hcl/v2/hclwrite"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	flvs "github.com/gophercloud/gophercloud/openstack/compute/v2/flavors"
	"github.com/gophercloud/utils/openstack/clientconfig"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "os-tofuizer",
		Usage: "generate terraform configuration of given resources",
		Action: func(ctx *cli.Context) error {
			return nil
		},
		EnableBashCompletion: true,
		Commands: []*cli.Command{
			{
				Name:    "generate",
				Aliases: []string{"g"},
				Usage:   "Generate resource conifg",
				Action: func(cCtx *cli.Context) error {
					generate()
					return nil
				},
			},
			{
				Name:    "json",
				Aliases: []string{"j"},
				Usage:   "Generate resource json",
				Action: func(cCtx *cli.Context) error {
					_json()
					return nil
				},
			},
		},
	}

	app.Run(os.Args)
}

type Resource interface{}

func _generate() {
	opts := new(clientconfig.ClientOpts)
	opts.Cloud = "cluster"

	provider, err := clientconfig.AuthenticatedClient(opts)
	_ = provider
	_ = err
	client, err := openstack.NewImageServiceV2(provider, gophercloud.EndpointOpts{})
	_ = err
	f := hclwrite.NewEmptyFile()
	f.Body().AppendBlock(img.Generate(client, "90b86069-456b-4f26-a9fa-d646a112c9f4"))
	fmt.Printf("%s", f.Bytes())
}

func _json() {
	opts := new(clientconfig.ClientOpts)
	opts.Cloud = "cluster"
	provider, _ := clientconfig.AuthenticatedClient(opts)
	client, _ := openstack.NewComputeV2(provider, gophercloud.EndpointOpts{})
	client.Microversion = "2.55"

	flavor := flavors.FlavorContainer{
		Flavor: flavors.Flavor_v2{
			ResourceType: "openstack_compute_flavor_v2",
			ResourceName: "name",
		}}

	// fmt.Printf("%s\n", flavor.Get(client, "10").Json())
	_ = flavor
	// var i interface{}
	_ = flvs.Get(client, "10")
	js, _ := json.Marshal(flvs.Get(client, "1").Body)
	fmt.Printf("%s\n", js)
}

func generate() {
	opts := new(clientconfig.ClientOpts)
	opts.Cloud = "cluster"
	provider, err := clientconfig.AuthenticatedClient(opts)
	_ = err
	client, err := openstack.NewComputeV2(provider, gophercloud.EndpointOpts{})
	client.Microversion = "2.55"
	_, _ = client, err

	flavor := flavors.FlavorContainer{
		Flavor: flavors.Flavor_v2{
			ResourceType: "openstack_compute_flavor_v2",
			ResourceName: "name",
		}}
	// flavor := flavors.FlavorContainer{}

	f := hclwrite.NewEmptyFile()
	f.Body().AppendBlock(flavor.Get(client, "1").Generate())
	// fmt.Printf("%s\n", flavor.Get(client, "10").Json())
	fmt.Printf("%s\n", f.Bytes())
	// fmt.Printf("%s\n", Body)
}

/*
package main

import (
    "fmt"
    "log"
    "os"

    "github.com/urfave/cli/v2"
)

func main() {
    app := &cli.App{
        EnableBashCompletion: true,
        Commands: []*cli.Command{
            {
                Name:    "add",
                Aliases: []string{"a"},
                Usage:   "add a task to the list",
                Action: func(cCtx *cli.Context) error {
                    fmt.Println("added task: ", cCtx.Args().First())
                    return nil
                },
            },
            {
                Name:    "complete",
                Aliases: []string{"c"},
                Usage:   "complete a task on the list",
                Action: func(cCtx *cli.Context) error {
                    fmt.Println("completed task: ", cCtx.Args().First())
                    return nil
                },
            },
            {
                Name:    "template",
                Aliases: []string{"t"},
                Usage:   "options for task templates",
                Subcommands: []*cli.Command{
                    {
                        Name:  "add",
                        Usage: "add a new template",
                        Action: func(cCtx *cli.Context) error {
                            fmt.Println("new task template: ", cCtx.Args().First())
                            return nil
                        },
                    },
                    {
                        Name:  "remove",
                        Usage: "remove an existing template",
                        Action: func(cCtx *cli.Context) error {
                            fmt.Println("removed task template: ", cCtx.Args().First())
                            return nil
                        },
                    },
                },
            },
        },
    }

    if err := app.Run(os.Args); err != nil {
        log.Fatal(err)
    }
}

*/
