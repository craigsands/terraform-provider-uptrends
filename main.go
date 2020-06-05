package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/craigsands/terraform-provider-uptrends/uptrends"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: uptrends.Provider})
}
