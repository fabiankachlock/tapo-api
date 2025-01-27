package response

type DefaultStateType string

const (
	DefaultStateTypeCustom     DefaultStateType = "custom"
	DefaultStateTypeLastStates DefaultStateType = "last_states"
)

type DefaultPowerType string

const (
	DefaultPowerTypeAlwaysOn   DefaultPowerType = "always_on"
	DefaultPowerTypeLastStates DefaultPowerType = "last_states"
)

type DefaultBrightnessState struct {
	Type  DefaultStateType `json:"type"`
	Value uint8            `json:"value"`
}
