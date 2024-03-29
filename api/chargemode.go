package api

//go:generate enumer -type ChargeMode -trimprefix Mode -text -transform=lower
type ChargeMode int

// Charge modes
const (
	ModeEmpty ChargeMode = iota
	ModeOff
	ModeNow
	ModeMinPV
	ModePV
)
