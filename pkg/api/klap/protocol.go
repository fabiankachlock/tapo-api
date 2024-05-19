package klap

import (
	"bytes"
	"crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/http/cookiejar"
	"time"

	"github.com/fabiankachlock/tapo-api/pkg/api"
)

type KLAPProtocol struct {
	email     string
	password  string
	url       string
	client    *http.Client
	cipher    *KLAPCipher
	cookieJar *cookiejar.Jar
}

var _ api.Protocol = (*KLAPProtocol)(nil)

func NewProtocol(opts api.ProtocolOptions) (*KLAPProtocol, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}

	client := &KLAPProtocol{
		email:     opts.Email,
		password:  opts.Password,
		url:       opts.Url,
		cookieJar: jar,
		client: &http.Client{
			Jar: jar,
		},
	}

	return client, nil
}

// Login logs in to the Tapo API.
func (d *KLAPProtocol) Login() error {
	hashedUsername := sha1.Sum([]byte(d.email))
	hashedPassword := sha1.Sum([]byte(d.password))
	authHash := sha256.Sum256(append(hashedUsername[:], hashedPassword[:]...))

	localSeed := make([]byte, 16)
	rand.Read(localSeed)

	remoteSeed, err := d.handshake1(d.url, localSeed, authHash[:])
	if err != nil {
		return err
	}

	err = d.handshake2(d.url, localSeed, remoteSeed, authHash[:])
	if err != nil {
		return err
	}

	d.cipher = NewCipher(localSeed, remoteSeed, authHash[:])
	return nil
}

// RefreshSession refreshes the authentication session of the client.
func (d *KLAPProtocol) RefreshSession() error {
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
func (d *KLAPProtocol) Request(method string, params interface{}) ([]byte, error) {
	request := map[string]interface{}{
		"method":           method,
		"params":           params,
		"requestTimeMilis": time.Now().UnixMilli(),
		"terminalUUID":     "00-00-00-00-00-00",
	}
	requestData, err := json.Marshal(request)
	if err != nil {
		return []byte{}, err
	}

	payload, seq, err := d.cipher.Encrypt(requestData)
	if err != nil {
		return []byte{}, err
	}

	resp, err := d.client.Post(fmt.Sprintf("%s/request?seq=%d", d.url, seq), "application/x-www-form-urlencoded", bytes.NewReader(payload))
	if err != nil {
		return []byte{}, err
	}

	buf := make([]byte, resp.ContentLength)
	resp.Body.Read(buf)
	decrypted, err := d.cipher.Decrypt(seq, buf)
	if err != nil {
		return []byte{}, err
	}

	return decrypted, nil
}

func (d *KLAPProtocol) handshake1(url string, localSeed []byte, authHash []byte) ([]byte, error) {
	resp, err := d.client.Post(fmt.Sprintf("%s/handshake1", url), "application/x-www-form-urlencoded", bytes.NewReader(localSeed))
	if err != nil {
		return []byte{}, err
	}

	buf := make([]byte, resp.ContentLength)
	resp.Body.Read(buf)
	remoteSeed := buf[0:16]
	serverHash := buf[16:]
	localHash := sha256.Sum256(append(append(localSeed, remoteSeed...), authHash...))

	if string(localHash[:]) != string(serverHash) {
		return []byte{}, errors.New("hashes dont match")
	}
	return remoteSeed, nil
}

func (d *KLAPProtocol) handshake2(url string, localSeed, remoteSeed, authHash []byte) error {
	payload := sha256.Sum256(append(append(remoteSeed, localSeed...), authHash...))
	resp, err := d.client.Post(fmt.Sprintf("%s/handshake2", url), "application/x-www-form-urlencoded", bytes.NewReader(payload[:]))
	if err != nil {
		return err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		log.Println(resp.Status)
		return errors.New("handshake 2 failed")
	}
	return nil
}
