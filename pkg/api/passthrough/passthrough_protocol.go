package passthrough

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"time"

	"github.com/fabiankachlock/tapo-api/pkg/api"
	"github.com/fabiankachlock/tapo-api/pkg/api/request"
	"github.com/fabiankachlock/tapo-api/pkg/api/response"
)

type PassthroughProtocol struct {
	email     string
	password  string
	url       string
	client    *http.Client
	cipher    *PassthroughCipher
	keyPair   *PassthroughKeyPair
	cookieJar *cookiejar.Jar
	token     string
}

var _ api.Protocol = (*PassthroughProtocol)(nil)

func NewProtocol(opts api.ProtocolOptions) (*PassthroughProtocol, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}

	client := &PassthroughProtocol{
		email:     opts.Email,
		password:  opts.Password,
		url:       opts.Url,
		cookieJar: jar,
		keyPair:   NewPassthroughKeyPair(),
		client: &http.Client{
			Jar: jar,
		},
	}

	return client, nil
}

// Login logs in to the Tapo API.
func (d *PassthroughProtocol) Login() error {
	err := d.handshake(d.url)
	if err != nil {
		return err
	}
	usernameDigest := d.cipher.DigestUsername(d.email)
	username := base64.StdEncoding.EncodeToString(usernameDigest)
	password := base64.StdEncoding.EncodeToString([]byte(d.password))

	resp, err := d.Request(request.RequestLoginDevice, request.LoginDeviceParams{
		Username: username,
		Password: password,
	})
	if err != nil {
		return err
	}

	tokenResult, err := response.UnmarshalResponse[response.TokenResponse](resp)
	if err != nil {
		return err
	}

	if tokenResult.ErrorCode != response.ResponseOk {
		return fmt.Errorf("received error code while logging in")
	}

	d.token = tokenResult.Result.Token

	return nil
}

// RefreshSession refreshes the authentication session of the client.
func (d *PassthroughProtocol) RefreshSession() error {
	// clear cookies
	jar, err := cookiejar.New(nil)
	if err != nil {
		return err
	}
	d.client = &http.Client{
		Jar: jar,
	}

	return d.Login()
}

// Request sends a request to the Tapo API.
func (d *PassthroughProtocol) Request(method string, params interface{}) ([]byte, error) {
	// TODO: https://github.dev/mihai-dinculescu/tapo/blob/main/tapo/src/api/protocol/passthrough_cipher.rs
	return []byte{}, nil
}

func (d *PassthroughProtocol) handshake(url string) error {
	params := request.HandshakeParams{
		Key: string(d.keyPair.GetPublicKey()),
	}
	request := map[string]interface{}{
		"method":           request.RequestHandshake,
		"params":           params,
		"requestTimeMilis": time.Now().UnixMilli(),
		"terminalUUID":     "00-00-00-00-00-00",
	}
	jsonBody, err := json.Marshal(request)
	if err != nil {
		return err
	}

	resp, err := d.client.Post(url, "application/json", bytes.NewReader(jsonBody))
	if err != nil {
		return err
	}

	buf := make([]byte, resp.ContentLength)
	resp.Body.Read(buf)
	responseJson, err := response.UnmarshalResponse[response.HandshakeResponse](buf)
	if responseJson.ErrorCode != response.ResponseOk {
		return fmt.Errorf("received error code while on handshake")
	}

	cipher, err := NewPassthroughCipher([]byte(responseJson.Result.Key), d.keyPair)
	if err != nil {
		return err
	}
	d.cipher = cipher
	return nil
}
