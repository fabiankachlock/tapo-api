package response

import (
	"encoding/json"
	"errors"
	"fmt"
)

const (
	ResponseOk              = 0
	ErrorInvalidRequest     = -1002
	ErrorMalformedRequest   = -1003
	ErrorInvalidPublicKey   = -1010
	ErrorInvalidCredentials = -1501
	ErrorSessionTimeOut     = 9999
)

var (
	ErrNonSuccessfulResponse = errors.New("non successful response")
)

type TapoResponse[T any] struct {
	Result    T   `json:"result"`
	ErrorCode int `json:"error_code"`
}

func (r TapoResponse[T]) IsOk() bool {
	return r.ErrorCode == ResponseOk
}

func (r TapoResponse[T]) HasError() bool {
	return r.ErrorCode != ResponseOk
}

func (r TapoResponse[T]) GetError() error {
	if r.ErrorCode == ResponseOk {
		return nil
	}
	return fmt.Errorf("%w: error code: %d", ErrNonSuccessfulResponse, r.ErrorCode)
}

type TapoMultipleResponse[T any] struct {
	Result TapoMultipleResult[T] `json:"result"`
}

type TapoMultipleResult[T any] struct {
	Responses []TapoResponse[T] `json:"responses"`
}

// UnmarshalResponse unmarshals the response from the Tapo API.
func UnmarshalResponse[T any](data []byte) (TapoResponse[T], error) {
	jsonData := TapoResponse[T]{}
	err := json.Unmarshal(data, &jsonData)
	return jsonData, err
}

type EmptyResponse struct{}
