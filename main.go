package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/kaerimichi/terraform-provider-rd/rd"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: rd.Provider,
	})
}
