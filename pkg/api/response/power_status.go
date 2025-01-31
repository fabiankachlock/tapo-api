package response

type OvercurrentStatus string

const (
	OvercurrentStatusLifted OvercurrentStatus = "lifted"
	OvercurrentStatusNormal OvercurrentStatus = "normal"
)

type OverheatStatus string

const (
	OverheatStatusCoolDown   OverheatStatus = "cool_down"
	OverheatStatusNormal     OverheatStatus = "normal"
	OverheatStatusOverheated OverheatStatus = "overheated"
)

type PowerProtectionStatus string

const (
	PowerProtectionStatusNormal     PowerProtectionStatus = "normal"
	PowerProtectionStatusOverloaded PowerProtectionStatus = "overloaded"
)
