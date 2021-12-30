package scptest

import (
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	log.Println("### Provider Start")
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"scptest_resource_group": resourceResourceGroup(),
		},
	}
}
