package templates 

import (
	"github.com/evcc-io/config/registry"
)

func init() {
	template := registry.Template{
		Class:  "vehicle",
		Type:   "enyaq",
		Name:   "SKODA (Enyaq)",
		Sample: `title: Enyaq # display name for UI
capacity: 10 # kWh
user: # user
password: # password
vin: WVWZZZ... # optional`,
	}

	registry.Add(template)
}
