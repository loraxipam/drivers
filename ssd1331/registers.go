package ssd1331

// Registers
const (
	DRAWLINE       = 0x21
	DRAWRECT       = 0x22
	FILL           = 0x26
	SETCOLUMN      = 0x15
	SETROW         = 0x75
	CONTRASTA      = 0x81
	CONTRASTB      = 0x82
	CONTRASTC      = 0x83
	MASTERCURRENT  = 0x87
	SETREMAP       = 0xA0
	STARTLINE      = 0xA1
	DISPLAYOFFSET  = 0xA2
	NORMALDISPLAY  = 0xA4
	DISPLAYALLON   = 0xA5
	DISPLAYALLOFF  = 0xA6
	INVERTDISPLAY  = 0xA7
	SETMULTIPLEX   = 0xA8
	SETMASTER      = 0xAD
	DISPLAYOFF     = 0xAE
	DISPLAYON      = 0xAF
	POWERMODE      = 0xB0
	PRECHARGE      = 0xB1
	CLOCKDIV       = 0xB3
	PRECHARGEA     = 0x8A
	PRECHARGEB     = 0x8B
	PRECHARGEC     = 0x8C
	PRECHARGELEVEL = 0xBB
	VCOMH          = 0xBE
)
