package api

type Protocol interface {
	Login(url, username, password string) error
	RefreshSession(username, password string) error
	ExecuteRequest(request []byte, withToken bool) ([]byte, error)
}
