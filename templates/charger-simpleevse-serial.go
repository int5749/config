package templates 

import (
	"github.com/evcc-io/config/registry"
)

func init() {
	template := registry.Template{
		Class:  "charger",
		Type:   "simpleevse",
		Name:   "EVSE DIN (Serial port)",
		Sample: `# http://evracing.cz/simple-evse-wallbox
device: /dev/ttyUSB0 # serial RS485 interface`,
	}

	registry.Add(template)
}
