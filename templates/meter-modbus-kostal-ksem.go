package templates 

import (
	"github.com/evcc-io/config/registry"
)

func init() {
	template := registry.Template{
		Class:  "meter",
		Type:   "modbus",
		Name:   "Kostal Smart Energy Meter (Grid Meter)",
		Sample: `model: sunspec
uri: 192.0.2.2:502
id: 71 # kostal default sunspec modbus id`,
	}

	registry.Add(template)
}
