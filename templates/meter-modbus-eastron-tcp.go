package templates 

import (
	"github.com/evcc-io/config/registry"
)

func init() {
	template := registry.Template{
		Class:  "meter",
		Type:   "modbus",
		Name:   "Eastron SDM Modbus Meter via TCP (Grid Meter, PV Meter)",
		Sample: `model: sdm # specific non-sunspec meter
id: 1
uri: 192.0.2.2:502
rtu: true # serial modbus rtu (rs485) device connected using simple ethernet adapter`,
	}

	registry.Add(template)
}
