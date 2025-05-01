package rd

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"rd_rubber_duck": rubberDuckResource(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"rd_rubber_duck": rubberDuckDataSource(),
		},
	}
}
