package tapo

type TapoProtocolType string

const (
	TapoProtocolKLAP TapoProtocolType = "klap"
)

type TapoClientOption struct {
	Timeout  int
	Protocol TapoProtocolType
}

var DefaultOptions = TapoClientOption{
	Timeout:  30,
	Protocol: TapoProtocolKLAP,
}

func WithTimeout(timeout int) func(*TapoClientOption) {
	return func(o *TapoClientOption) {
		o.Timeout = timeout
	}
}

func WithProtocol(protocol TapoProtocolType) func(*TapoClientOption) {
	return func(o *TapoClientOption) {
		o.Protocol = protocol
	}
}
