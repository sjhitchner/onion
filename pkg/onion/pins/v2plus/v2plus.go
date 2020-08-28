package v2plus

import (
	p "github.com/sjhitchner/onion-tutorial/pkg/onion/pins"
)

var ExpansionDock = p.ExpansionDock{
	Ethernet: p.Ethernet{
		TxP: 0,
		TxN: 0,
		RxP: 0,
		RxN: 0,
	},
	USB: p.USB{
		Plus: 0,
		Neg:  0,
	},
	PWM: p.PWM{
		Ch0: 18,
		Ch1: 19,
	},
	I2C: p.I2C{
		SCL: 4,
		SDA: 5,
	},
	I2S: p.I2S{
		CLK: 3,
		WS:  2,
		SDO: 1,
		SDI: 0,
	},
	UART: p.UART{
		Rx: 46,
		Tx: 45,
	},
	SPI: p.SPI{
		MISO: 0,
		MOSI: 0,
		CLK:  0,
		CS1:  0,
	},
	RGBLed: p.RGBLed{
		Red:   17,
		Green: 16,
		Blue:  15,
	},
}
