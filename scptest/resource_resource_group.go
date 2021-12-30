package scptest

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceResourceGroup() *schema.Resource {
	log.Println("### resourceResourceGroup Start")
	return &schema.Resource{
		CreateContext: resourceResourceGroupCreate,
		ReadContext:   resourceResourceGroupRead,
		UpdateContext: resourceResourceGroupUpdate,
		DeleteContext: resourceResourceGroupDelete,

		Schema: map[string]*schema.Schema{
			"uuid_count": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}
func resourceResourceGroupCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	/*
		uuid_count := d.Get("uuid_count").(string)
		d.SetId(uuid_count)
		// https://www.uuidtools.com/api/generate/v1/count/uuid_count
		resp, err := http.Get("https://www.uuidtools.com/api/generate/v1/count/" + uuid_count)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		return resourceResourceGroupRead(d, m)
	*/
	log.Println("### resourceResourceGroupCreate Start")
	var diags diag.Diagnostics
	return diags
}

func resourceResourceGroupRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Println("### resourceResourceGroupRead Start")
	var diags diag.Diagnostics
	return diags
}

func resourceResourceGroupUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Println("### resourceResourceGroupUpdate Start")
	/*
		return resourceResourceGroupRead(d, m)
	*/
	var diags diag.Diagnostics
	return diags
}

func resourceResourceGroupDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Println("### resourceResourceGroupDelete Start")
	/*
		d.SetId("")
	*/
	var diags diag.Diagnostics
	return diags
}
