package templates 

import (
	"github.com/evcc-io/config/registry"
)

func init() {
	template := registry.Template{
		Class:  "meter",
		Type:   "modbus",
		Name:   "Eastron SDM Modbus Meter via serial port (Grid Meter, PV Meter)",
		Sample: `model: sdm # specific non-sunspec meter
id: 1
# chose either locally attached:
device: /dev/ttyUSB0 # serial port
baudrate: 9600
comset: 8N1`,
	}

	registry.Add(template)
}
