package uptrends

import (
	"context"

	"github.com/craigsands/uptrends-sdk-go"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"username": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("UPTRENDS_USERNAME", nil),
			},
			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("UPTRENDS_PASSWORD", nil),
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"uptrends_monitor": resourceUptrendsMonitor(),
		},

		ConfigureFunc: providerConfigure,
	}
}

//ProviderConfiguration contains the initialized API clients to communicate with the Uptrends API
type ProviderConfiguration struct {
	Client *uptrends.APIClient
	Auth   context.Context
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	username := d.Get("username").(string)
	password := d.Get("password").(string)

	config := uptrends.NewConfiguration()
	client := uptrends.NewAPIClient(config)
	auth := context.WithValue(
		context.Background(),
		uptrends.ContextBasicAuth,
		uptrends.BasicAuth{
			UserName: username,
			Password: password,
		},
	)

	return &ProviderConfiguration{
		Client: client,
		Auth:   auth,
	}, nil
}
