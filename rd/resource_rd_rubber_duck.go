package rd

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/kaerimichi/terraform-provider-rd/models"
)

func rubberDuckResource() *schema.Resource {
	return &schema.Resource{
		Create: func(d *schema.ResourceData, m interface{}) error {
			rubberDuck := models.RubberDuck{
				Color:    d.Get("color").(string),
				Material: d.Get("material").(string),
				Size:     d.Get("size").(string),
			}

			id, err := rubberDuck.CreateRubberDuck(&rubberDuck)
			if err != nil {
				return err
			}

			d.SetId(id)

			return nil
		},
		Read: func(d *schema.ResourceData, m interface{}) error {
			return nil
		},
		Update: func(d *schema.ResourceData, m interface{}) error {
			return nil
		},
		Delete: func(d *schema.ResourceData, m interface{}) error {
			d.SetId("")
			return nil
		},
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"color": {
				Type:     schema.TypeString,
				Required: true,
			},
			"material": {
				Type:     schema.TypeString,
				Required: true,
			},
			"size": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}
