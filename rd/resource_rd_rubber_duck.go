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

			id, err := models.CreateRubberDuck(&rubberDuck)
			if err != nil {
				return err
			}

			d.SetId(id)

			return nil
		},
		Read: func(d *schema.ResourceData, m interface{}) error {
			id := d.Id()

			rubberDuck, err := models.GetRubberDuck(id)
			if err != nil {
				return err
			}

			d.Set("color", rubberDuck.Color)
			d.Set("material", rubberDuck.Material)
			d.Set("size", rubberDuck.Size)

			return nil
		},
		Update: func(d *schema.ResourceData, m interface{}) error {
			id := d.Id()

			rubberDuck := models.RubberDuck{
				ID:       id,
				Color:    d.Get("color").(string),
				Material: d.Get("material").(string),
				Size:     d.Get("size").(string),
			}

			err := models.UpdateRubberDuck(id, &rubberDuck)
			if err != nil {
				return err
			}

			return nil
		},
		Delete: func(d *schema.ResourceData, m interface{}) error {
			err := models.DeleteRubberDuck(d.Id())
			if err != nil {
				return err
			}

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
