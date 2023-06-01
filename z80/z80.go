package z80

type CPU struct {
	SPR
}

type SPR struct {
	I uint8
        R uint8
	IX uint16
	IY uint16
	SP uint16
	PC uint16
}
