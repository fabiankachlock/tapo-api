package response

// Decodable is an interface for response types that can be decoded.
// Some response fields are for example base64 encoded and need to be decoded after unmarshaling.
// If a response type implements this interface, the api client will automatically call the Decode method after unmarshaling.
type Decodable interface {
	Decode() error
}
