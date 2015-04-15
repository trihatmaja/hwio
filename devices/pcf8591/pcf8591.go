package pcf8591

// reference :
// http://blog.chrysocome.net/2012/12/i2c-analog-to-digital-converter.html

import (
	"github.com/trihatmaja/hwio"
)

const(
	DEFAULT_BASE_ADDRESS = 0x48
	REG_CHAN0 = 0x00
	REG_CHAN1 = 0x01
	REG_CHAN2 = 0x02
	REG_CHAN3 = 0x03
	REG_WRITE = 0x41
)

type PCF8591 struct {
	device hwio.I2CDevice
}

func NewPCF8591(module hwio.I2CModule, address int) (*PCF8591) {
	if address < 8 {
		address += DEFAULT_BASE_ADDRESS
	}
	device := module.GetDevice(address)
	result := &PCF8591{device: device}
	return result
}

func (d *PCF8591) ReadByte(chan int) (byte, error){
	var val byte
	switch chan {
	    case 0:
		val = REG_CHAN0
	    case 1:
		val = REG_CHAN1
	    case 2:
		val = REG_CHAN2
	    case 3:
		val = REG_CHAN3
	}
	return d.device.ReadByte(val)
}

func (d *PCF8591) Read(chan int, bit int)([]byte, error){
	var val []byte
        switch chan {
            case 0:
                val = REG_CHAN0
            case 1:
                val = REG_CHAN1
            case 2:
                val = REG_CHAN2
            case 3:
                val = REG_CHAN3
        }
        return d.device.Read(val, bit)

} 

func (d *PCF8591) WriteByte(val int){
	return d.device.WriteByte(REG_WRITE, val)
}
