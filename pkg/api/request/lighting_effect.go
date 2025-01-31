package request

type LightingEffectType string

const (
	LightingEffectSequence LightingEffectType = "sequence"
	LightingEffectRandom   LightingEffectType = "random"
	LightingEffectPulse    LightingEffectType = "pulse"
	LightingEffectStatic   LightingEffectType = "static"
)

type LightingEffect struct {
	Brightness         uint8              `json:"brightness"`
	IsCustom           uint8              `json:"custom"`
	DisplayColors      [3]uint8           `json:"display_colors"`
	Enabled            uint8              `json:"enabled"`
	Id                 string             `json:"id"`
	Name               string             `json:"name"`
	Type               LightingEffectType `json:"type"`
	Backgrounds        [][3]uint16        `json:"backgrounds,omitempty"`
	BrightnessRange    []uint8            `json:"brightness_range,omitempty"`
	Direction          uint8              `json:"direction,omitempty"`
	Duration           uint64             `json:"duration,omitempty"`
	ExpansionStrategy  uint8              `json:"expansion_strategy,omitempty"`
	FadeOff            uint16             `json:"fadeoff,omitempty"`
	HueRange           [2]uint16          `json:"hue_range,omitempty"`
	InitStates         [][3]uint16        `json:"init_states,omitempty"`
	RandomSeed         uint64             `json:"random_seed,omitempty"`
	RepeatTimes        uint8              `json:"repeat_times,omitempty"`
	RunTime            uint64             `json:"run_time,omitempty"`
	SaturationRange    [2]uint8           `json:"saturation_range,omitempty"`
	SegmentLength      uint8              `json:"segment_length,omitempty"`
	Segments           []uint8            `json:"segments,omitempty"`
	Sequence           [][3]uint8         `json:"sequence,omitempty"`
	Spread             uint8              `json:"spread,omitempty"`
	Transition         uint16             `json:"transition,omitempty"`
	TransitionRange    [2]uint16          `json:"transition_range,omitempty"`
	TransitionSequence []uint16           `json:"trans_sequence,omitempty"`
}
