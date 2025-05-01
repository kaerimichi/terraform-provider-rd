package rd

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/kaerimichi/terraform-provider-rd/models"
)

func rubberDuckDataSource() *schema.Resource {
	return &schema.Resource{
		Read: func(d *schema.ResourceData, m interface{}) error {
			id := d.Get("id").(string)

			rubberDuck, err := models.GetRubberDuck(id)
			if err != nil {
				return err
			}

			d.SetId(rubberDuck.ID)

			if err := d.Set("color", rubberDuck.Color); err != nil {
				return err
			}

			if err := d.Set("material", rubberDuck.Material); err != nil {
				return err
			}

			if err := d.Set("size", rubberDuck.Size); err != nil {
				return err
			}

			return nil
		},
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"color": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"material": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"size": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}
