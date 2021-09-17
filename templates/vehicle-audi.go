package templates 

import (
	"github.com/evcc-io/config/registry"
)

func init() {
	template := registry.Template{
		Class:  "vehicle",
		Type:   "audi",
		Name:   "Audi (eTron, Q55, etc)",
		Sample: `title: eTron # display name for UI
capacity: 14 # kWh
user: # user
password: # password
vin: WAUZZZ... # optional`,
	}

	registry.Add(template)
}
