# PCF8591 I2C AD/DA Converter
This package provides a simple way to connect to the PCF8591 AD/DA converter,
a device which convert analog to digital converter with 4 channel and digital
to analog converter with single channel.

# Hardware Support
Tested on Raspberry Pi with kernel 3.12

# Usage
Import the packages:

	// import the require modules
	import(
		"github.com/trihatmaja/hwio"
		"github.com/trihatmaja/hwio/devices/pcf8591"
	)

Initialise by fetching an i2c module from the driver. You can get instances of devices attached to
the bus.

	// Assume i2c module is enabled. Get the i2c module from the driver. This is an example for Raspberry PI.
	m, e := hwio.GetModule("i2c")

	// Assert that it is an I2C module
	i2c := m.(hwio.I2CModule)

Get the PCF8591 device, so you make requests of it:

	// Get the device on this i2c bus. Usually the address is 0x48 so you could use 0
	converter := pcf8591.NewPCF8591(i2c, 0)

Read the value:

	// Get the value of channel 0 up to 3
	val := converter.ReadChannel(0)

Write the value:

	// Set the value to 0xFF
	converter.WriteByte(0xFF)
