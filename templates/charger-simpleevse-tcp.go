package templates 

import (
	"github.com/evcc-io/config/registry"
)

func init() {
	template := registry.Template{
		Class:  "charger",
		Type:   "simpleevse",
		Name:   "EVSE DIN (Modbus-TCP)",
		Sample: `# http://evracing.cz/simple-evse-wallbox
uri: 192.0.2.2:502 # Modbus TCP converter address`,
	}

	registry.Add(template)
}
