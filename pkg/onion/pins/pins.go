package v2plus

import ()

type ExpansionDock struct {
	Ethernet
	USB
	PWM
	I2C
	I2S
	UART
	SPI
	RGBLed
	Reset int
}

type Ethernet struct {
	TxP int
	TxN int
	RxP int
	RxN int
}

type USB struct {
	Plus int
	Neg  int
}

type PWM struct {
	Ch0 int
	Ch1 int
}

type I2C struct {
	SDA int
	SCL int
}

type I2S struct {
	CLK int
	WS  int
	SDO int
	SDI int
}

type UART struct {
	Rx int
	Tx int
}

type SPI struct {
	MISO int
	MOSI int
	CLK  int
	CS1  int
}

type RGBLed struct {
	Red   int
	Green int
	Blue  int
}
