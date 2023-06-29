package z80

type CPU struct {
	GPRc
	GPRa
	SPR
}

type GPRc struct {
        A uint8
        F uint8
        B uint8
        C uint8
        D uint8
        E uint8
        H uint8
        L uint8
}

type GPRa struct {
        A_ uint8
        F_ uint8
        B_ uint8
        C_ uint8
        D_ uint8
        E_ uint8
        H_ uint8
        L_ uint8
}

type SPR struct {
	I uint8
        R uint8

	IX uint16
	IY uint16
	SP uint16
	PC uint16
}
