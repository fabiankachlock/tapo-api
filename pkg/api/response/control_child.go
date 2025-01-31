package response

type ControlChildResponse[T any] struct {
	Response T `json:"responseData"`
}
