package api

import (
	"bytes"
	"crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/cookiejar"
	"time"

	"github.com/fabiankachlock/tapo-api/pkg/klap"
)

type ApiClient struct {
	Ip          net.IP
	Email       string
	Password    string
	HandshakeTS time.Time
	url         string
	client      *http.Client
	cipher      *klap.KLAPCipher
	cookieJar   *cookiejar.Jar
}

type ApiClientGeneric[T any] TapoResponse[T]

func NewClient(ip, email, password string) (*ApiClient, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}

	client := &ApiClient{
		Ip:        net.ParseIP(ip),
		Email:     email,
		Password:  password,
		url:       fmt.Sprintf("http://%s/app", ip),
		cookieJar: jar,
		client: &http.Client{
			Jar: jar,
		},
	}

	return client, nil
}

func (d *ApiClient) Login() error {
	hashedUsername := sha1.Sum([]byte(d.Email))
	hashedPassword := sha1.Sum([]byte(d.Password))
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

	d.cipher = klap.NewCipher(localSeed, remoteSeed, authHash[:])
	return nil
}

func (d *ApiClient) RefreshSession() error {
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

func (d *ApiClient) Request(method string, params map[string]interface{}) ([]byte, error) {
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

func (d *ApiClient) handshake1(url string, localSeed []byte, authHash []byte) ([]byte, error) {
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

func (d *ApiClient) handshake2(url string, localSeed, remoteSeed, authHash []byte) error {
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
