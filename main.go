package main

import (
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"

	"github.com/JamesLKC/terraform-provider-uuid/scptest"
)

func main() {
	/*
		plugin.Serve(&plugin.ServeOpts{
			ProviderFunc: func() terraform.ResourceProvider {
				return Provider()
			},
		})
	*/
	log.Println("### main Start")
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() *schema.Provider {
			return scptest.Provider()
		},
	})
	log.Println("### main End")
}
