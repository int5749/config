package templates 

import (
	"github.com/evcc-io/config/registry"
)

func init() {
	template := registry.Template{
		Class:  "vehicle",
		Type:   "seat",
		Name:   "SEAT (Cupra, Mii, etc)",
		Sample: `title: Cupra # display name for UI
capacity: 10 # kWh
user: # user
password: # password
vin: WVWZZZ... # optional`,
	}

	registry.Add(template)
}
