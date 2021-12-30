package scptest

import (
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceResourceGroup() *schema.Resource {
	log.Println("### resourceResourceGroup Start")
	return &schema.Resource{
		Create: resourceResourceGroupCreate,
		Read:   resourceResourceGroupRead,
		Update: resourceResourceGroupUpdate,
		Delete: resourceResourceGroupDelete,

		Schema: map[string]*schema.Schema{
			"uuid_count": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}
func resourceResourceGroupCreate(d *schema.ResourceData, m interface{}) error {
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
	return nil
}

func resourceResourceGroupRead(d *schema.ResourceData, m interface{}) error {
	log.Println("### resourceResourceGroupRead Start")
	return nil
}

func resourceResourceGroupUpdate(d *schema.ResourceData, m interface{}) error {
	log.Println("### resourceResourceGroupUpdate Start")
	/*
		return resourceResourceGroupRead(d, m)
	*/
	return nil
}

func resourceResourceGroupDelete(d *schema.ResourceData, m interface{}) error {
	log.Println("### resourceResourceGroupDelete Start")
	/*
		d.SetId("")
	*/
	return nil
}
