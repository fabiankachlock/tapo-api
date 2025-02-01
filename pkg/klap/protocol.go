package klap

import (
	"bytes"
	"crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"time"
)

type KLAPProtocol struct {
	client      *http.Client
	cookieJar   *cookiejar.Jar
	HandshakeTS time.Time
	url         string
	cipher      *KLAPCipher
}

func NewProtocol(timeout time.Duration) (*KLAPProtocol, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, fmt.Errorf("klap protocol: failed to create cookie jar: %w", err)
	}

	client := &KLAPProtocol{
		cookieJar: jar,
		client: &http.Client{
			Jar:     jar,
			Timeout: timeout,
		},
	}

	return client, nil
}

func (p *KLAPProtocol) handshake(url, username, password string) error {
	hashedUsername := sha1.Sum([]byte(username))
	hashedPassword := sha1.Sum([]byte(password))
	authHash := sha256.Sum256(append(hashedUsername[:], hashedPassword[:]...))

	localSeed := make([]byte, 16)
	_, err := rand.Read(localSeed)
	if err != nil {
		return fmt.Errorf("failed to generate local seed: %w", err)
	}

	remoteSeed, err := p.handshake1(url, localSeed, authHash[:])
	if err != nil {
		return fmt.Errorf("handshake1 failed: %w", err)
	}

	err = p.handshake2(url, localSeed, remoteSeed, authHash[:])
	if err != nil {
		return fmt.Errorf("handshake2 failed: %w", err)
	}

	p.url = url
	p.cipher = NewCipher(localSeed, remoteSeed, authHash[:])
	return nil
}

func (p *KLAPProtocol) handshake1(url string, localSeed []byte, authHash []byte) ([]byte, error) {
	resp, err := p.client.Post(fmt.Sprintf("%s/handshake1", url), "application/x-www-form-urlencoded", bytes.NewReader(localSeed))
	if err != nil {
		return []byte{}, fmt.Errorf("failed to send handshake1 request: %w", err)
	}

	buf, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return []byte{}, fmt.Errorf("failed to read handshake1 response: %w", err)
	}

	remoteSeed := buf[0:16]
	serverHash := buf[16:]
	localHash := sha256.Sum256(append(append(localSeed, remoteSeed...), authHash...))

	if string(localHash[:]) != string(serverHash) {
		return []byte{}, errors.New("handshake1 hashes do not match")
	}
	return remoteSeed, nil
}

func (p *KLAPProtocol) handshake2(url string, localSeed, remoteSeed, authHash []byte) error {
	payload := sha256.Sum256(append(append(remoteSeed, localSeed...), authHash...))
	resp, err := p.client.Post(fmt.Sprintf("%s/handshake2", url), "application/x-www-form-urlencoded", bytes.NewReader(payload[:]))
	if err != nil {
		return fmt.Errorf("failed to send handshake2 request: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return errors.New("handshake2 failed")
	}
	return nil
}

func (d *KLAPProtocol) Login(url, username, password string) error {
	err := d.handshake(url, username, password)
	if err != nil {
		return fmt.Errorf("login failed: %w", err)
	}
	return nil
}

func (d *KLAPProtocol) RefreshSession(username, password string) error {
	// clear cookies
	jar, err := cookiejar.New(nil)
	if err != nil {
		return fmt.Errorf("failed to create new cookie jar: %w", err)
	}
	d.client = &http.Client{
		Jar: jar,
	}

	err = d.handshake(d.url, username, password)
	if err != nil {
		return fmt.Errorf("failed to refresh session: %w", err)
	}
	return nil
}

// Request sends a request to the Tapo API.
func (d *KLAPProtocol) ExecuteRequest(request []byte, withToken bool) ([]byte, error) {
	payload, seq, err := d.cipher.Encrypt(request)
	if err != nil {
		return []byte{}, fmt.Errorf("failed to encrypt request: %w", err)
	}

	resp, err := d.client.Post(fmt.Sprintf("%s/request?seq=%d", d.url, seq), "application/x-www-form-urlencoded", bytes.NewReader(payload))
	if err != nil {
		return []byte{}, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	buf, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, fmt.Errorf("failed to read response: %w", err)
	}

	decrypted, err := d.cipher.Decrypt(buf)
	if err != nil {
		return []byte{}, fmt.Errorf("failed to decrypt response: %w", err)
	}

	return decrypted, nil
}
