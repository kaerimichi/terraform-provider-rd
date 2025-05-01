package rd

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func rubberDuckDataSource() *schema.Resource {
	return &schema.Resource{
		Read: func(d *schema.ResourceData, m interface{}) error {
			id := d.Get("id").(string)
			// Simulate data retrieval
			d.SetId(id)
			d.Set("value", "This is a data source for item: "+id)
			return nil
		},
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"value": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}
