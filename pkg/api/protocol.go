package api

import "fmt"

type Protocol interface {
	Login() error
	RefreshSession() error
	Request(method string, params interface{}) ([]byte, error)
}

type ProtocolOptions struct {
	Email    string
	Password string
	Url      string
}

func NewOptions(ip, email, password string) ProtocolOptions {
	return ProtocolOptions{
		Url:      fmt.Sprintf("http://%s/app", ip),
		Email:    email,
		Password: password,
	}
}
