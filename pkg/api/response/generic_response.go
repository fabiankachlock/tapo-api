package response

import (
	"encoding/json"
	"fmt"
)

type GenericResponse struct {
	raw []byte
}

func NewGenericResponse(data []byte) GenericResponse {
	return GenericResponse{data}
}

func (r GenericResponse) Raw() []byte {
	return r.raw
}

func (r GenericResponse) UnmarshalRaw(value any) error {
	return json.Unmarshal(r.raw, value)
}

func (r GenericResponse) UnmarshalResponse(result interface{}) (TapoResponse[any], error) {
	// marshal the data as a TapoResponse capsulating the given result type
	tapoResponse := TapoResponse[any]{Result: result}
	err := json.Unmarshal(r.raw, &tapoResponse)
	if err != nil {
		return TapoResponse[any]{}, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	return tapoResponse, nil
}
