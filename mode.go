package onword

type Mode int

const (
	// ModeNormal indicates normal execution.
	ModeNormal Mode = iota

	// ModeDefName indicates the consumption of a UserDef name.
	ModeDefName

	// ModeDefBody indicates the consumption of a UserDef body.
)
