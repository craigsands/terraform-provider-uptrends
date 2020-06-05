package uptrends

import (
	"log"

	"github.com/craigsands/uptrends-sdk-go"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceUptrendsMonitor() *schema.Resource {
	return &schema.Resource{
		Create: resourceUptrendsMonitorCreate,
		Read:   resourceUptrendsMonitorRead,
		Update: resourceUptrendsMonitorUpdate,
		Delete: resourceUptrendsMonitorDelete,

		Schema: map[string]*schema.Schema{
			"monitor_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"url": {
				Type:     schema.TypeString,
				Required: true,
			},
			"notes": {
				Type: schema.TypeString,
				Optional: true,
				Default: "Managed by Terraform",
			},
		},
	}
}

func buildMonitorStruct(d *schema.ResourceData) *uptrends.Monitor {
	monitor_type := uptrends.MonitorType(d.Get("monitor_type").(string))
	monitor := uptrends.Monitor{
		MonitorType: &monitor_type,
		Name:        d.Get("name").(string),
		Notes: d.Get("notes").(string),
		Url:         d.Get("url").(string),

		// use GetOkExists for boolean values
	}

	return &monitor
}

func readMonitorStruct(monitor *uptrends.Monitor, d *schema.ResourceData) error {
	d.Set("monitor_type", monitor.MonitorType)
	d.Set("name", monitor.Name)
	d.Set("notes", monitor.Notes)
	d.Set("url", monitor.Url)

	return nil
}

func resourceUptrendsMonitorCreate(d *schema.ResourceData, m interface{}) error {
	providerConf := m.(*ProviderConfiguration)
	client := providerConf.Client
	auth := providerConf.Auth

	monitor := buildMonitorStruct(d)

	log.Printf("[INFO] Creating Uptrends %v monitor %s", &monitor.MonitorType, monitor.Name)

	created, _, err := client.MonitorApi.MonitorPostMonitor(auth, *monitor)
	if err != nil {
		log.Fatal(err)
	}

	d.SetId(created.MonitorGuid)
	return resourceUptrendsMonitorRead(d, m)
}

func resourceUptrendsMonitorRead(d *schema.ResourceData, m interface{}) error {
	providerConf := m.(*ProviderConfiguration)
	client := providerConf.Client
	auth := providerConf.Auth

	opts := uptrends.MonitorGetMonitorOpts{}
	monitor, r, err := client.MonitorApi.MonitorGetMonitor(auth, d.Id(), &opts)
	if err != nil {
		if r.StatusCode == 400 {
			log.Printf("[WARN] No monitor found: %s", d.Id())
			d.SetId("")
			return nil
		}

		return err
	}

	return readMonitorStruct(&monitor, d)
}

func resourceUptrendsMonitorUpdate(d *schema.ResourceData, m interface{}) error {
	providerConf := m.(*ProviderConfiguration)
	client := providerConf.Client
	auth := providerConf.Auth

	monitor := buildMonitorStruct(d)

	log.Printf("[INFO] Updating Uptrends %v monitor %s", &monitor.MonitorType, monitor.Name)

	_, err := client.MonitorApi.MonitorPatchMonitor(auth, *monitor, d.Id())
	if err != nil {
		return err
	}
	
	return resourceUptrendsMonitorRead(d, m)
}

func resourceUptrendsMonitorDelete(d *schema.ResourceData, m interface{}) error {
	providerConf := m.(*ProviderConfiguration)
	client := providerConf.Client
	auth := providerConf.Auth

	monitor_type := uptrends.MonitorType(d.Get("monitor_type").(string))
	log.Printf("[INFO] Deleting Uptrends %v monitor %s", &monitor_type, d.Id())

	_, err := client.MonitorApi.MonitorDeleteMonitor(auth, d.Id())
	if err != nil {
		return err
	}

	return nil
}
